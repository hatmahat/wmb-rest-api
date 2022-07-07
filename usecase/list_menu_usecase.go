package usecase

import (
	"go-api-with-gin2/model"
	"go-api-with-gin2/repo"
)

type ListMenuUseCase interface {
	Retrive() ([]model.Menu, error)
}

type listMenuUseCase struct {
	repo repo.MenuRepo
}

func (c *listMenuUseCase) Retrive() ([]model.Menu, error) {
	return c.repo.FindAll()
}

func NewListMenuUseCase(repo repo.MenuRepo) ListMenuUseCase {
	return &listMenuUseCase{
		repo: repo,
	}
}
