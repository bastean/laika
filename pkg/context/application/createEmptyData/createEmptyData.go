package createEmptyData

import "github.com/bastean/laika/pkg/context/domain/aggregate"

type CreateEmptyData struct{}

func (create *CreateEmptyData) Run() *aggregate.Laika {
	laika := new(aggregate.Laika)

	laika.Sniffed = make(map[string][]*aggregate.Data)

	return laika
}

func NewCreateEmptyData() *CreateEmptyData {
	return new(CreateEmptyData)
}
