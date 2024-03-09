package repository

import (
	"github.com/bastean/laika/pkg/context/shared/domain/aggregate"
)

type Repository interface {
	Save(data aggregate.Data) error
	Read() (aggregate.Data, error)
}
