package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	neturl "net/url"
	"time"
)

func main() {

}
func queryWithHedgedRequestsWithContext(urls []string) string {
	ch := make(chan string, len(urls))
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	for _, url := range urls {
		go func(u string, c chan string) {
			c <- executeQueryWithContext(u, ctx)
		}(url, ch)

		select {
		case r := <-ch:
			cancel()
			return r
		case <-time.After(21 * time.Millisecond):
		}
	}

	return <-ch
}

func executeQueryWithContext(url string, ctx context.Context) string {
	start := time.Now()
	parsedURL, _ := neturl.Parse(url)
	req := &http.Request{URL: parsedURL}
	req = req.WithContext(ctx)

	response, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println(err.Error())
		return err.Error()
	}

	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Printf("Request time: %d ms from url%s\n", time.Since(start).Nanoseconds()/time.Millisecond.Nanoseconds(), url)
	return fmt.Sprintf("%s from %s", body, url)
}
