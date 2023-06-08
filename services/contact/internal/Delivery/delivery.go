package Delivery

import (
	"awesomeProject3/services/contact/internal/repository"
	"awesomeProject3/services/contact/internal/useCase"
)

type Delivery struct {
	useCase    useCase.UseCase
	repository repository.Repository
}

func NewDelivery(useCase useCase.UseCase, repository repository.Repository) *Delivery {
	return &Delivery{
		useCase:    useCase,
		repository: repository,
	}
}
