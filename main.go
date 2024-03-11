package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	var path string
	if len(os.Args) == 1 {
		fmt.Println("Error: specify the directory of the website")
		os.Exit(1)
	}
	path = os.Args[1]

	if _, err := os.Stat(path); err != nil {
		fmt.Println("Error: specify a site directory")
		os.Exit(2)
	}

	if _, err := os.Stat(fmt.Sprintf("%v/index.html", path)); err != nil {
		fmt.Println("Error: no index.html file found in path")
		os.Exit(3)
	}

	fmt.Println("====VALID WEBSITE DIRECTORY SPECIFIED====\n\n")

	port := os.Getenv("PORT")
	if port == "" {
		fmt.Println("Error: please specify a port in your environment variables")
	}

	fileServe := http.FileServer(http.Dir(path))
	http.Handle("/", fileServe)

	fmt.Println("=====STARTING SERVER=====")

	http.HandleFunc("/docker", (func(writer http.ResponseWriter, _ *http.Request) {
		containered := os.Getenv("IN_CONTAINER")
		if containered == "TRUE" {
			fmt.Fprintf(writer, "<html><body><h1>I AM A CONTAINER!!</h1></body></html>")
		} else {
			fmt.Fprintf(writer, "<html><body><h1>dockerize me or perish</h1></body></html>")
		}
	}))

	err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
	if err != nil {
		fmt.Println("CANNOT RUN SERVER")
		return
	}

	fmt.Println("Listening on port:", port)
}
