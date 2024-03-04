package sniffEmails

import (
	"github.com/bastean/laika/pkg/context/shared/domain/aggregate"
)

type SniffSocialMediaUrls struct{}

func (sniff *SniffSocialMediaUrls) Run(data *aggregate.Data) {
	// TODO feat(context): add sniff social media urls
}

func NewSniffSocialMediaUrls() *SniffSocialMediaUrls {
	return new(SniffSocialMediaUrls)
}
