package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	c := http.Client{Timeout: time.Duration(1) * time.Second}

	myJson := bytes.NewBuffer([]byte(`{"name":"Maximilien"}`))

	resp, err := c.Post("https://www.google.com", "application/json", myJson)
	if err != nil {
		fmt.Errorf("Error %s", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("Body : %s", body)
}
