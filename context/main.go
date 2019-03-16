package main

import (
	"context"
	"sync"
	"time"
	"fmt"
	"net/http"
	"io/ioutil"
	"crypto/md5"
)

func main() {
	wg := &sync.WaitGroup{}
	urls :=  []string{"https://www.baidu.com/", "https://www.zhihu.com/"}
	ctx, cancel := context.WithCancel(context.Background())

	for _, url := range urls {
		subctx := context.WithValue(ctx, "url", url)
		wg.Add(1)
		go reqURL(subctx, wg)
	}

	go func() {
		time.Sleep(time.Second*3)
		cancel()
	}()

	wg.Wait()
	fmt.Println("exit the main goroutine")
}

func reqURL(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	url, _ := ctx.Value("url").(string)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("stop getting url: ", url)
			return
		default:
			r, err := http.Get(url)
			if err == nil && r.StatusCode == http.StatusOK {
				body, _ := ioutil.ReadAll(r.Body)
				subctx := context.WithValue(ctx, "resp", fmt.Sprintf("%s%x", url, md5.Sum(body)))
				wg.Add(1)
				go showResp(subctx, wg)
			}
			r.Body.Close()
			time.Sleep(time.Second*1)
		}
	}
}

func showResp(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <- ctx.Done():
			fmt.Println("stop showing response")
			return
		default:
			fmt.Println("printing response: ", ctx.Value("resp"))
			time.Sleep(time.Second *1)
		}
	}
}