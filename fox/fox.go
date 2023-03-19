package fox

import (
	"zveropolis/animals"
	"zveropolis/repository"
)

type Fox struct {
	repo *repository.Repository
}

func NewFox(repo *repository.Repository) (*Fox, error) {
	return &Fox{
		repo: repo,
	}, nil
}

func (f *Fox) GetAll() ([]animals.Fox, error) {
	list, err := f.repo.GetFoxFromMysql()
	if err != nil {
		return nil, err
	}

	if len(list) == 0 {
		list, err = f.repo.GetFoxFromPostgres()
	}

	return list, err
}
