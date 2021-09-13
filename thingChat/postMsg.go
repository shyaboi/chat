package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func thing(inputy string) {
	url := "http://localhost:3000/msg"
	fmt.Println("URL:>", url)

	var jsonStr = []byte(inputy)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Chat Log:", string(body))
}
