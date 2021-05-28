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
	mutexVisited  sync.Mutex
	mutexMaxDepth sync.RWMutex
	visited       map[string]string
	maxDepth      int
}

func newCrawler(maxDepth int) *crawler {
	return &crawler{
		visited:  make(map[string]string),
		maxDepth: maxDepth,
	}
}

func (c *crawler) increaseMaxDepth(maxDepth int) {
	c.mutexMaxDepth.Lock()
	c.maxDepth += maxDepth
	c.mutexMaxDepth.Unlock()
}

// рекурсивно сканируем страницы
func (c *crawler) run(ctx context.Context, url string, results chan<- crawlResult, depth int) {
	// просто для того, чтобы успевать следить за выводом программы, можно убрать :)
	time.Sleep(2 * time.Second)

	// проверяем что контекст исполнения актуален
	select {
	case <-ctx.Done():
		return

	default:
		c.mutexMaxDepth.RLock()
		// проверка глубины
		if depth >= c.maxDepth {
			return
		}
		c.mutexMaxDepth.RUnlock()

		page, err := parse(url)
		if err != nil {
			// ошибку отправляем в канал, а не обрабатываем на месте
			results <- crawlResult{
				err: errors.Wrapf(err, "parse page %s", url),
			}
			return
		}

		title := pageTitle(page)
		links := pageLinks(nil, page)

		// блокировка требуется, т.к. мы модифицируем мапку в несколько горутин
		c.mutexVisited.Lock()
		c.visited[url] = title
		c.mutexVisited.Unlock()

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
	c.mutexVisited.Lock()
	defer c.mutexVisited.Unlock()

	_, ok := c.visited[url]
	return ok
}
