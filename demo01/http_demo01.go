package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	client := &http.Client{}
	url := "http://www.baidu.com"
	request, err := http.NewRequest("GET", url, nil)

	if err != nil {
		panic(err)
	}

	respone, _ := client.Do(request)
	stdout := os.Stdout
	_, err = io.Copy(stdout, respone.Body)
	status := respone.StatusCode
	fmt.Println(status)

}
