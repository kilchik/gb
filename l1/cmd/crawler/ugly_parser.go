package main

import (
	"bytes"
	"context"
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"time"
)

// парсим страницу
func parse(ctx context.Context, url string) (*html.Node, error) {
	select {
	case <-ctx.Done():
		return nil, nil
	default:
		client := http.Client{Timeout: 1 * time.Second}
		request, err := http.NewRequest("GET", url, bytes.NewReader([]byte{}))
		if err != nil {
			return nil, fmt.Errorf("can't send request")
		}

		r, err := client.Do(request)
		if err != nil {
			return nil, fmt.Errorf("can't get page")
		}

		b, err := html.Parse(r.Body)
		if err != nil {
			return nil, fmt.Errorf("can't parse page")
		}
		return b, err
	}
}

// ищем заголовок на странице
func pageTitle(ctx context.Context, n *html.Node) string {
	select {
	case <-ctx.Done():
		return ""
	default:
		var title string
		if n.Type == html.ElementNode && n.Data == "title" {
			return n.FirstChild.Data
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			title = pageTitle(ctx, c)
			if title != "" {
				break
			}
		}
		return title
	}
}

// ищем все ссылки на страницы. Используем мапку чтобы избежать дубликатов
func pageLinks(ctx context.Context, links map[string]struct{}, n *html.Node) map[string]struct{} {
	select {
	case <-ctx.Done():
		return nil
	default:
		if links == nil {
			links = make(map[string]struct{})
		}

		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}

				// костылик для простоты
				if _, ok := links[a.Val]; !ok && len(a.Val) > 2 && a.Val[:2] == "//" {
					links["http://"+a.Val[2:]] = struct{}{}
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			links = pageLinks(ctx, links, c)
		}
		return links
	}
}
