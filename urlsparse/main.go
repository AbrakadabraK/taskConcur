package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
)

type Urls struct {
	urlsChannel            chan string
	answerToRequestChannel chan string
	successCounter         int
	counterMutext          sync.Mutex
}

func New(urls []string) *Urls {
	u := &Urls{
		urlsChannel:            make(chan string),
		answerToRequestChannel: make(chan string),
		counterMutext:          sync.Mutex{},
	}
	go func() {
		for _, url := range urls {
			u.urlsChannel <- url
		}
		close(u.urlsChannel)
	}()
	return u
}

var successCounter int
var counterMutex sync.Mutex // Защита для счётчика успешных запросов

func main() {
	exampleUrls := []string{
		"http://ozon.ru",
		"https://ozon.ru",
		"http://google.com",
		"http://somesite.com",
		"http://non-existent.domain.tld",
		"https://ya.ru",
		"http://ya.ru",
		"http://ёёёё",
	}

	wg := &sync.WaitGroup{}
	urls := New(exampleUrls)

	doneChan := make(chan struct{})
	ctxC, _ := context.WithCancel(context.Background())
	// Запуск горутин для обработки запросов
	go func() {
		for u := range urls.urlsChannel {
			wg.Add(1)
			go func(url string) {
				defer wg.Done()
				status := urls.requestUrls(ctxC, url, doneChan)
				if status != "" { // Если запрос был выполнен и статус не пустой
					urls.answerToRequestChannel <- status
				}
			}(u)
		}
		wg.Wait() // Ждём завершения всех горутин
		close(urls.answerToRequestChannel)
	}()

	// Горутина для печати результатов
	go func() {
		for result := range urls.answerToRequestChannel {
			fmt.Println(result)
		}
		close(doneChan)
	}()
	<-doneChan
	fmt.Println("END")
}

func (u *Urls) requestUrls(ctx context.Context, urlStatus string, doneC chan struct{}) string {
	var res string

	select {
	case <-doneC:
		return res
	default:
	}

	// Выполняем запрос с использованием контекста с тайм-аутом
	req, err := http.NewRequestWithContext(ctx, "GET", urlStatus, nil)
	if err != nil {
		res = fmt.Sprintf("Address url : %s , not ok !\n", urlStatus)
	}

	r, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Sprintf("Address url : %s , not ok !\n", urlStatus)
	}
	defer r.Body.Close()

	// Если статус не 200 OK, выводим сообщение
	if r.StatusCode != http.StatusOK {
		return fmt.Sprintf("Address url : %s , not ok! Status: %d\n", urlStatus, r.StatusCode)
	}
	// Если статус 200 OK, увеличиваем счетчик
	counterMutex.Lock()
	successCounter++
	res = fmt.Sprintf("URL GOOD: %s, Status: %d OK", urlStatus, r.StatusCode)
	counterMutex.Unlock()

	// Проверяем, достигли ли мы 2 успешных ответа, и если да, отменяем остальные запросы
	if successCounter == 2 {
		u.answerToRequestChannel <- res
		doneC <- struct{}{}
		fmt.Println("Received 2 successful responses, cancelling remaining requests...")
		return ""
	}

	return res
}
