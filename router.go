package main

import (
	"net/http"
	"strconv"
	"strings"
)

func (s *server) newRouter() *http.ServeMux {
	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("frontend/assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("GET /quit", quitHandler)

	mux.HandleFunc("GET /recipes", s.recipesHandler)
	mux.HandleFunc("GET /recipes/", s.recipesHandler)

	mux.HandleFunc("GET /ingredients", s.ingredientsHandler)
	mux.HandleFunc("POST /ingredients", s.createIngredientHandler)
	mux.HandleFunc("POST /ingredients/delete", s.deleteIngredientHandler)

	return mux
}

func quitHandler(w http.ResponseWriter, r *http.Request) {
	close(done)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	parseTemplate("index").Execute(w, nil)
}

func (s *server) recipesHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.Trim(r.URL.Path, "/")
	parts := strings.Split(path, "/")

	if len(parts) == 1 && parts[0] == "recipes" {
		s.allRecipesHandler(w, r)
	} else if len(parts) == 2 && parts[0] == "recipes" {
		s.singleRecipeHandler(w, r, parts[1])
	} else {
		http.NotFound(w, r)
	}
}

func (s *server) singleRecipeHandler(w http.ResponseWriter, r *http.Request, id string) {
	parseTemplate("create_recipe").Execute(w, nil)
}

func (s *server) allRecipesHandler(w http.ResponseWriter, r *http.Request) {
	type data struct {
		Recipes   []Recipe
		NoOfPages int
	}

	pageStr := r.URL.Query().Get("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	noOfPages := s.repo.getRecipesPagination()

	recipes, err := s.repo.getAllRecipes(page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	parseTemplate("recipes").Execute(w, data{
		Recipes:   recipes,
		NoOfPages: noOfPages,
	})
}

func (s *server) ingredientsHandler(w http.ResponseWriter, r *http.Request) {
	data, err := s.repo.getAllIngredients()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	parseTemplate("ingredients").Execute(w, data)
}

func (s *server) createIngredientHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	description := r.FormValue("description")

	i := Ingredient{
		Name:        name,
		Description: description,
	}

	if err = s.repo.createIngredient(i); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	refererURL := r.Header.Get("Referer")
	http.Redirect(w, r, refererURL, http.StatusSeeOther)
}

func (s *server) deleteIngredientHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	id := r.FormValue("id")

	if err = s.repo.deleteIngredient(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	refererURL := r.Header.Get("Referer")
	http.Redirect(w, r, refererURL, http.StatusSeeOther)
}
