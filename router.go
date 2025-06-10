package main

import (
	"fmt"
	"net/http"
)

func (s *server) newRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", getIndex)
	mux.HandleFunc("/quit", quit)

	mux.HandleFunc("GET /recipes", getRecipes)
	mux.HandleFunc("GET /calculator", getCalculator)

	return mux
}

func quit(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Upgrade error:", err)
		return
	}
	defer conn.Close()

	fmt.Println("WebSocket client connected")

	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("WebSocket disconnected. Exiting.")
			close(done)
			return
		}
	}
}

func getIndex(w http.ResponseWriter, r *http.Request) {
	parseTemplate("index").Execute(w, nil)
}

func getRecipes(w http.ResponseWriter, r *http.Request) {
	parseTemplate("recipes").Execute(w, nil)
}

func getCalculator(w http.ResponseWriter, r *http.Request) {
	parseTemplate("calculator").Execute(w, nil)
}
