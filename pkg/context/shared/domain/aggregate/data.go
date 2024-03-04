package aggregate

type Sniffed struct {
	Source  string
	Content string
	Found   map[string][]string
}

type Data struct {
	Sniffed map[string][]*Sniffed
}

func Create() *Data {
	return &Data{Sniffed: make(map[string][]*Sniffed)}
}
