package aggregate

type Sniffed struct {
	Source          string
	Emails          []string
	PhoneNumbers    []string
	SocialMediaUrls []string
}

type Data map[string][]*Sniffed

func Create() Data {
	return Data{}
}
