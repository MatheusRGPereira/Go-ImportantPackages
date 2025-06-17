package main

import "net/http"

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", homeHandler)
	mux.Handle("/blog", &Blog{})

	http.ListenAndServe(":8080", mux)

}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world!"))
}

type Blog struct {
	title string
}

func (b *Blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Darkness my old friend!"))
}
