package rabbits

import (
	"zveropolis/animals"
	"zveropolis/repository"
)

type Rabbits struct {
	repo *repository.Repository
}

func NewRabbits(repo *repository.Repository) (*Rabbits, error) {
	return &Rabbits{
		repo: repo,
	}, nil
}

func (f *Rabbits) GetAll() ([]animals.Rabbit, error) {
	list, err := f.repo.GetRabbitsFromPostgres()
	if err != nil {
		return nil, err
	}

	if len(list) == 0 {
		list, err = f.repo.GetRabbitsFromMysql()
	}

	return list, err
}
