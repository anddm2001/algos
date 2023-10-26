package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gocolly/colly"
	"golang.org/x/net/html"
)

var ctx = context.Background()

type PageData struct {
	Title   string
	H1      []string
	H2_H6   []string
	Content []byte
}

var tagsToDelete = map[string]bool{
	"h1": true,
	"h2": true,
	"h3": true,
	"h4": true,
	"h5": true,
	"h6": true,
}

func main() {
	rand.Seed(time.Now().UnixNano())

	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0, // используем базу данных по умолчанию
	})

	c := colly.NewCollector(
		colly.UserAgent("YourCustomUserAgent"),
	)

	pageData := &PageData{}

	c.OnResponse(func(r *colly.Response) {
		if r.StatusCode != 200 {
			log.Fatalf("Остановка работы парсера: получен HTTP-статус %v", r.StatusCode)
		}

		if xRobotsTag := r.Headers.Get("X-Robots-Tag"); xRobotsTag != "" {
			if strings.Contains(strings.ToLower(xRobotsTag), "noindex") {
				log.Fatalf("Скрапинг страницы запрещен по правилам X-Robots-Tag.")
			}
		}

		if len(r.Body) > 50*1024 {
			log.Fatalf("Предупреждение: размер HTML документа больше 50 КБ. Прекращаем анализ.")
		}

		// Создаем уникальный ключ на основе URL и случайной строки
		key := fmt.Sprintf("%s_%s", r.Ctx.Get("url"), randStringRunes(8))

		// Сохраняем содержимое тега <body> в Redis
		doc, err := html.Parse(bytes.NewReader(r.Body))
		if err != nil {
			log.Fatalf("Ошибка при разборе HTML: %v", err)
			return
		}
		var f func(*html.Node)
		f = func(n *html.Node) {
			if n.Type == html.ElementNode && n.Data == "body" {
				err := rdb.Set(ctx, key, renderNode(n), 0).Err()
				if err != nil {
					log.Fatalf("Не удалось сохранить данные в Redis: %v", err)
				}
			}
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				f(c)
			}
		}
		f(doc)
	})

	c.OnHTML("title", func(e *colly.HTMLElement) {
		pageData.Title = e.Text
	})

	c.OnHTML("h1", func(e *colly.HTMLElement) {
		pageData.H1 = append(pageData.H1, e.Text)
	})

	// Извлекаем содержимое тегов h2-h6
	for i := 2; i <= 6; i++ {
		c.OnHTML(fmt.Sprintf("h%d", i), func(e *colly.HTMLElement) {
			fmt.Printf("H%d: %s\n", i, e.Text)
			pageData.H2_H6 = append(pageData.H2_H6, e.Text)
		})
	}

	// Извлекаем остальное содержимое тега body, исключая h1-h6
	c.OnHTML("body", func(e *colly.HTMLElement) {
		e.ForEach("p, span, div", func(_ int, el *colly.HTMLElement) {
			fmt.Println("Other content:", el.Text)
		})
	})

	c.OnHTML("body", func(e *colly.HTMLElement) {
		// Получаем все содержимое body, затем удаляем из него содержимое тегов h1-h6
		bodyContent, _ := e.DOM.Html()
		pageData.Content = []byte(removeHeadings(bodyContent))
	})

	c.OnScraped(func(r *colly.Response) {
		// После завершения скрапинга, выводим собранные данные
		fmt.Printf("Собранные данные с веб-страницы: %+v\n", pageData)
	})

	// Устанавливаем обработчик ошибок
	c.OnError(func(r *colly.Response, err error) {
		log.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	err := c.Visit("http://example.com")
	if err != nil {
		log.Fatal(err)
	}

	defer rdb.Close()
}

// removeHeadings удаляет содержимое тегов h1-h6 из HTML-контента
func removeHeadings(content string) string {
	doc, err := html.Parse(strings.NewReader(content))
	if err != nil {
		log.Fatal(err)
	}

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && tagsToDelete[n.Data] {
			n.Parent.RemoveChild(n)
		} else {
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				f(c)
			}
		}
	}

	f(doc)

	var b strings.Builder
	html.Render(&b, doc)
	return b.String()
}

func renderNode(n *html.Node) string {
	var buf bytes.Buffer
	w := io.Writer(&buf)
	html.Render(w, n)
	return buf.String()
}

// randStringRunes генерирует случайную строку заданной длины
func randStringRunes(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
