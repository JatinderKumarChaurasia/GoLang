package main

import "net/http"

func main()  {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("Hello Go "))
	})
	_ = http.ListenAndServe(":8000",nil)
}
