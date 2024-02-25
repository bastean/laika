package createEmptyData

import "github.com/bastean/laika/pkg/context/domain/aggregate"

type CreateEmptyData struct{}

func (create *CreateEmptyData) Run() *aggregate.Laika {
	return aggregate.Create()
}

func NewCreateEmptyData() *CreateEmptyData {
	return new(CreateEmptyData)
}
