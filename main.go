package main

import (
	"fmt"
	"net/http"

	"github.com/martonikon/go-block-explorer/explorer"
)

func main() {
	http.HandleFunc("/blocks", explorer.GetBlocks)

	fmt.Println("Explorer running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
