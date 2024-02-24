package aggregate

type Data struct {
	Source  string
	Content string
	Found   map[string][]string
}

type Laika struct {
	Sniffed map[string][]*Data
}
