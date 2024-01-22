package test

import (
	"fmt"
	"io"
	"net/http"
)

func ClientGet() {
	response, err := http.Get("https://practicum.yandex.ru.")
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	fmt.Println("Header")
	fmt.Println("================================")
	for key, value := range response.Header {
		fmt.Println(key, value)
	}
	fmt.Println("================================\n")

	fmt.Println("Body")
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	if len(body) > 512 {
		body = body[:512]
	}
	fmt.Print(string(body))
	fmt.Println("\n================================\n")

}

func Redirect() {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println(req.URL)
			return nil
		},
	}
	response, err := client.Get("http://ya.ru")

	_, err = io.Copy(io.Discard, response.Body)

	if err != nil {
		fmt.Println(err)
	}
}
