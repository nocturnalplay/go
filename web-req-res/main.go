package main

import (
	"net/http"
	"fmt"
	"io"
)

func main()  {
	fmt.Println("welcome to the web server")
	res,err := http.Get("http://localhost")

	if err != nil {
		panic(err)
	}
	body,err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
	defer res.Body.Close()
}