package main

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"sync"
	"time"
)

type crawlResult struct {
	err error
	msg string
}

type crawler struct {
	sync.Mutex
	rwMutex sync.RWMutex
	visited  map[string]string
	maxDepth int
}

func newCrawler(maxDepth int) *crawler {
	return &crawler{
		visited:  make(map[string]string),
		maxDepth: maxDepth,
	}
}

func (c *crawler) dive() {
	c.rwMutex.RLock()
	defer c.rwMutex.RUnlock()
	c.maxDepth += 2
}

// рекурсивно сканируем страницы
func (c *crawler) run(ctx context.Context, url string, results chan<- crawlResult, depth int) {
	// просто для того, чтобы успевать следить за выводом программы, можно убрать :)
	time.Sleep(2 * time.Second)

	ctxParse, cancel := context.WithTimeout(ctx, 3 * time.Second)
	defer cancel()

	// проверяем что контекст исполнения актуален
	select {
	case <-ctx.Done():
		return

	default:
		// проверка глубины
		if depth >= c.maxDepth {
			return
		}

		page, err := parse(ctxParse, url)
		if err != nil {
			// ошибку отправляем в канал, а не обрабатываем на месте
			results <- crawlResult{
				err: errors.Wrapf(err, "parse page %s", url),
			}
			return
		}

		title := pageTitle(ctxParse, page)
		links := pageLinks(ctxParse, nil, page)

		// блокировка требуется, т.к. мы модифицируем мапку в несколько горутин
		c.Lock()
		c.visited[url] = title
		c.Unlock()

		// отправляем результат в канал, не обрабатывая на месте
		results <- crawlResult{
			err: nil,
			msg: fmt.Sprintf("%s -> %s\n", url, title),
		}

		// рекурсивно ищем ссылки
		for link := range links {
			// если ссылка не найдена, то запускаем анализ по новой ссылке
			if c.checkVisited(link) {
				continue
			}

			go c.run(ctx, link, results, depth+1)
		}
	}
}

func (c *crawler) checkVisited(url string) bool {
	c.Lock()
	defer c.Unlock()

	_, ok := c.visited[url]
	return ok
}
