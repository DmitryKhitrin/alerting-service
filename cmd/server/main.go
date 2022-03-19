package main

import (
	"fmt"
	"io"
	"net/http"
)

// HelloWorld — обработчик запроса.
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	// читаем Body
	b, err := io.ReadAll(r.Body)
	// обрабатываем ошибку
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	fmt.Println(b)
}

func main() {
	// маршрутизация запросов обработчику
	http.HandleFunc("/", HelloWorld)
	// запуск сервера с адресом localhost, порт 8080
	http.ListenAndServe(":8080", nil)
}
