package movie

import (
	"zveropolis/fox"
	"zveropolis/rabbits"
)

type Movie struct {
	rabbitAnimal *rabbits.Rabbits
	foxAnimal    *fox.Fox
}

func NewMovie(fox *fox.Fox, rabbits *rabbits.Rabbits) (*Movie, error) {
	return &Movie{
		rabbitAnimal: rabbits,
		foxAnimal:    fox,
	}, nil
}

func (m *Movie) StartNewJob() error {
	return nil
}
