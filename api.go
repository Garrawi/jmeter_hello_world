package main

import (
	"fmt"
	"encoding/json"
	"math/rand"
        "net/http"
	"strconv"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	element_count_str := r.URL.Query().Get("size")
	element_count := rand.Intn(2000);
	if len(element_count_str) != 0 {
		element_count, _ = strconv.Atoi(element_count_str);
	}

	randomElements := make(map[string]int)
	for i := 0; i < element_count; i++ {
		randomElements[fmt.Sprintf("key%v", i)] = rand.Intn(200);
	}

        a, _ := json.Marshal(randomElements)
        w.Write(a)
        return
}

func main() {
        http.HandleFunc("/random", handler)
	fmt.Printf("Listening on 8080")
        http.ListenAndServe(":8080", nil)
}
