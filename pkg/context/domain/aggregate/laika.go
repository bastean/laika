package aggregate

type Data struct {
	Source  string
	Content string
	Found   map[string][]string
}

type Laika struct {
	Sniffed map[string][]*Data
}

func Create() *Laika {
	laika := new(Laika)

	laika.Sniffed = make(map[string][]*Data)

	return laika
}
