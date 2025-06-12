package main

import (
	"net/http"
)

func (s *server) newRouter() *http.ServeMux {
	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("frontend/assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("GET /quit", quitHandler)

	mux.HandleFunc("GET /recipes", recipesHandler)

	mux.HandleFunc("GET /ingredients", s.ingredientsHandler)
	mux.HandleFunc("POST /ingredients", s.createIngredientHandler)

	return mux
}

func quitHandler(w http.ResponseWriter, r *http.Request) {
	close(done)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	parseTemplate("index").Execute(w, nil)
}

func recipesHandler(w http.ResponseWriter, r *http.Request) {
	parseTemplate("recipes").Execute(w, nil)
}

func (s *server) ingredientsHandler(w http.ResponseWriter, r *http.Request) {
	data, _ := s.repo.getAllIngredients()
	parseTemplate("ingredients").Execute(w, data)
}

func (s *server) createIngredientHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	i := ingredient{
		Name:        r.FormValue("name"),
		Description: r.FormValue("description"),
	}

	if err = s.repo.createIngredient(i); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/ingredients", http.StatusOK)
}
