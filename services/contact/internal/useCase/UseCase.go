package useCase

import (
	"awesomeProject3/services/contact/internal/domain"
	"context"
)

type UseCase interface {
	CreateContact(ctx context.Context, contact domain.Contact) (domain.Contact, error)
	ReadContact(ctx context.Context, contact domain.Contact) (domain.Contact, error)
	UpdateContact(ctx context.Context, contact domain.Contact) (updated domain.Contact, err error)
	DeleteContact(ctx context.Context, contact domain.Contact) (id int, err error)
	CreateGroup(ctx context.Context, g domain.Group) (domain.Group, error)
	GetGroup(ctx context.Context, g domain.Group) (domain.Group, error)
	AddContactToGroup(ctx context.Context, c domain.Contact, g domain.Group) (domain.Contact, domain.Group, error)
}
