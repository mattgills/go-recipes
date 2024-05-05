package recipes

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"recipe-manager/db"
)

type Recipe struct {
	Name string `json:"name"`
}

type RecipeDbRow struct {
	ID int `json:"id"`
	Recipe
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Recipes Handler")
	db := db.GetDb()
	recipes := make([]RecipeDbRow, 0)

	rows, err := db.QueryContext(context.Background(), `SELECT * FROM recipes;`)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("Could not query DB"))
	}

	defer rows.Close()

	for rows.Next() {

		var recipe RecipeDbRow

		if err := rows.Scan(
			&recipe.ID, &recipe.Name,
		); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte("Could not put data into slice"))
		}
		recipes = append(recipes, recipe)
	}

	recipesJson, _ := json.Marshal(recipes)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(recipesJson)
}
