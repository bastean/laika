package createEmptyData

import "github.com/bastean/laika/pkg/context/shared/domain/aggregate"

type CreateEmptyData struct{}

func (create *CreateEmptyData) Run() *aggregate.Data {
	return aggregate.Create()
}

func NewCreateEmptyData() *CreateEmptyData {
	return new(CreateEmptyData)
}
