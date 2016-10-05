package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("connected from %s\n", r.RemoteAddr)
	fmt.Fprintf(w, "connected from %s", r.RemoteAddr)
}

func main() {
	http.HandleFunc("/", handler) // ハンドラを登録してウェブページを表示させる
	http.ListenAndServe(":18080", nil)
}
