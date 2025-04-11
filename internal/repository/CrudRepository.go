package repository

import (
	"golang.org/x/exp/constraints"
)

type CrudRepository[T any, I constraints.Integer] interface {
	Create(object *T) (bool, error)

	Read(id I) (*T, error)

	ReadAll() ([]*T, error)

	Update(object *T) (bool, error)

	Delete(id I) (*T, error)
}
