package models

import (
	"golang.org/x/exp/constraints"
)

type Identifiable[T constraints.Integer] interface {
	SetId(id T)
	GetId() T
}
