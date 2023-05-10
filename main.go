package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Println(req.Method)
		_, _ = res.Write([]byte("Hello world!!!"))
		//_, err := fmt.Fprintf(res, "Hello world!")
		//if err != nil {
		//	fmt.Println("error:", err)
		//}
	})

	_ = http.ListenAndServe(":8080", nil)
	fmt.Println("server is running")
}
