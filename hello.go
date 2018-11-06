package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"time"
	"log"
)

type Recipe struct{
	Name string `json:"name"`
	Description string `json:"description"`
	CreatedAt time.Time `json:"created_at"`

}

var recipeStore = make(map[string]Recipe)

var id int 

//GetRecipeHandler - GET - /api/recipes
func GetRecipeHandler (w http.ResponseWriter, r *http.Request){
	var recipes []Recipe
	for _, v := range recipeStore{
		recipes = append(recipes, v)
	}
}
func main() {
	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/api/recipes", GetRecipeHandler).Methods("GET")
	/*r.HandleFunc("/api/recipes", PostRecipeHandler).Methods("POST")
	r.HandleFunc("/api/recipes/{id}", PutRecipeHandler).Methods("PUT")
	r.HandleFunc("/api/recipes", DeleteRecipeHandler).Methods("DELETE")*/

	server := &http.Server{
		Addr: ":8080",
		Handler: r,
		ReadTimeout: 10,
		WriteTimeout: 10,
	}
	log.Println("listening http://localhost:8080...")
	server.ListenAndServe()
}