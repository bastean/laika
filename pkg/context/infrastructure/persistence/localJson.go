package persistence

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/bastean/laika/pkg/context/domain/aggregate"
)

type LocalJson struct {
	Path     string
	Filename string
}

func (localJson *LocalJson) create() *os.File {
	file, err := os.Create(filepath.Join(localJson.Path, filepath.Base(localJson.Filename)+".json"))

	if err != nil {
		panic(err)
	}

	return file
}

func (localJson *LocalJson) load() []byte {
	data, err := os.ReadFile(filepath.Join(localJson.Path, filepath.Base(localJson.Filename)+".json"))

	if os.IsNotExist(err) {
		localJson.create()
	} else if err != nil {
		panic(err)
	}

	return data
}

func (localJson *LocalJson) Save(laika *aggregate.Laika) {
	data, err := json.Marshal(laika)

	if err != nil {
		panic(err)
	}

	file := localJson.create()

	_, err = file.Write(data)

	if err != nil {
		panic(err)
	}
}

func (localJson *LocalJson) Read() *aggregate.Laika {
	data := localJson.load()

	laika := new(aggregate.Laika)

	if len(data) == 0 {
		laika.Sniffed = make(map[string][]*aggregate.Data)

		return laika
	}

	err := json.Unmarshal(data, &laika)

	if err != nil {
		panic(err)
	}

	return laika
}

func NewLocalJson(path, filename string) *LocalJson {
	return &LocalJson{
		path,
		filename,
	}
}
