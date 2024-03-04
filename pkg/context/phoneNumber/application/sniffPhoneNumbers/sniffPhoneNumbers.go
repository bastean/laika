package sniffEmails

import (
	"github.com/bastean/laika/pkg/context/shared/domain/aggregate"
)

type SniffPhoneNumbers struct{}

func (sniff *SniffPhoneNumbers) Run(data *aggregate.Data) {
	// TODO feat(context): add sniff phone numbers
}

func NewSniffPhoneNumbers() *SniffPhoneNumbers {
	return new(SniffPhoneNumbers)
}
