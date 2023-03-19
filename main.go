package main

import (
	"fmt"
	"net/http"
	"os"
	"zveropolis/movie"

	"zveropolis/controllers"
	"zveropolis/fox"
	"zveropolis/rabbits"
	"zveropolis/repository"
)

func main() {
	repo, _ := repository.NewRepository()
	fmt.Println("init repository")

	rabbitAnimal, _ := rabbits.NewRabbits(repo)
	fmt.Println("init rabbits success")

	foxAnimal, _ := fox.NewFox(repo)
	fmt.Println("init fox success")

	movies, _ := movie.NewMovie(foxAnimal, rabbitAnimal)
	fmt.Println("init fox success")

	controller := controllers.NewController(foxAnimal, rabbitAnimal, movies)
	http.HandleFunc("/get-fox", controller.GetFox)
	http.HandleFunc("/get-rabbits", controller.GetRabbits)
	http.HandleFunc("/start-movie", controller.MovieStartNewJob)

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprintf(w, "health\n"); err != nil {
			fmt.Printf("err health handler %v", err)
			os.Exit(1)
		}
	})

	fmt.Println("start dogs app, port: 8090")
	if err := http.ListenAndServe(":8090", nil); err != nil {
		fmt.Printf("start http server error %v", err)
		os.Exit(1)
	}
}
