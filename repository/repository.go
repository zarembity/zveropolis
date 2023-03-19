package repository

import "zveropolis/animals"

//Repository объект для работы с хранилищами
type Repository struct {
}

func NewRepository() (*Repository, error) {
	return &Repository{}, nil
}

func (r *Repository) GetRabbitsFromPostgres() ([]animals.Rabbit, error) {
	return []animals.Rabbit{
		{
			Name: "Роджер",
			Age:  12,
		},
		{
			Name: "Баксбани",
			Age:  12,
		},
	}, nil
}

func (r *Repository) GetRabbitsFromMysql() ([]animals.Rabbit, error) {
	return []animals.Rabbit{}, nil
}

func (r *Repository) GetFoxFromPostgres() ([]animals.Fox, error) {
	return []animals.Fox{}, nil
}

func (r *Repository) GetFoxFromMysql() ([]animals.Fox, error) {
	return []animals.Fox{
		{
			Name: "Ник",
			Age:  12,
		},
		{
			Name: "Баксбани",
			Age:  12,
		},
	}, nil
}
