package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"zveropolis/fox"
	"zveropolis/movie"
	"zveropolis/rabbits"
)

type AnimalController struct {
	fox     *fox.Fox
	rabbits *rabbits.Rabbits
	movies  *movie.Movie
}

var (
	systemError = "system error"
)

func NewController(fox *fox.Fox, rabbits *rabbits.Rabbits, movies *movie.Movie) *AnimalController {
	return &AnimalController{
		fox:     fox,
		rabbits: rabbits,
		movies:  movies,
	}
}

func (d *AnimalController) GetFox(w http.ResponseWriter, r *http.Request) {
	foxList, _ := d.fox.GetAll()
	listB, err := json.Marshal(foxList)
	if err != nil {
		fmt.Fprintf(w, systemError)
		return
	}
	fmt.Fprintf(w, string(listB))
}

func (d *AnimalController) GetRabbits(w http.ResponseWriter, r *http.Request) {
	rbList, _ := d.rabbits.GetAll()
	listB, err := json.Marshal(rbList)
	if err != nil {
		fmt.Fprintf(w, systemError)
		return
	}
	fmt.Fprintf(w, string(listB))
}

func (d *AnimalController) MovieStartNewJob(w http.ResponseWriter, r *http.Request) {
	go func() {
		if err := d.movies.StartNewJob(); err != nil {
			fmt.Println("error start job", err)
		}
	}()
	fmt.Fprintf(w, "job started")
}
