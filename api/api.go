package api

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"recipe-manager/api/ingredients"
	"recipe-manager/api/recipes"

	_ "github.com/joho/godotenv/autoload"
)

func Init() {
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	http.HandleFunc("/recipes", recipes.GetAll)
	http.HandleFunc("/ingredients", ingredients.GetAll)

	fmt.Printf("Starting Server on %v\n", host+":"+port)
	log.Fatal(http.ListenAndServe(host+":"+port, nil))
}
