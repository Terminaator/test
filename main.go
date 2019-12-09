package main

import (
	"fmt"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "test")
}

func setupRoutes()  {
	http.HandleFunc("/", homePage)
}

func main() {
	fmt.Println("go app")
	setupRoutes()
	http.ListenAndServe(":3000",nil)
}
