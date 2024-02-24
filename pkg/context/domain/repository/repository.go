package repository

import (
	"github.com/bastean/laika/pkg/context/domain/aggregate"
)

type Repository interface {
	Save(laika *aggregate.Laika)
	Read() *aggregate.Laika
}
