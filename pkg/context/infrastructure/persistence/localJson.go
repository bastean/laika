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

func (localJson *LocalJson) create() (*os.File, error) {
	return os.Create(filepath.Join(localJson.Path, filepath.Base(localJson.Filename)+".json"))
}

func (localJson *LocalJson) load() ([]byte, error) {
	data, err := os.ReadFile(filepath.Join(localJson.Path, filepath.Base(localJson.Filename)+".json"))

	if os.IsNotExist(err) {
		localJson.create()
	} else if err != nil {
		return nil, err
	}

	return data, nil
}

func (localJson *LocalJson) Save(laika *aggregate.Laika) error {
	data, err := json.Marshal(laika)

	if err != nil {
		return err
	}

	file, err := localJson.create()

	if err != nil {
		return err
	}

	_, err = file.Write(data)

	if err != nil {
		return err
	}

	return nil
}

func (localJson *LocalJson) Read() (*aggregate.Laika, error) {
	data, err := localJson.load()

	if err != nil {
		return nil, err
	}

	laika := new(aggregate.Laika)

	if len(data) == 0 {
		laika.Sniffed = make(map[string][]*aggregate.Data)

		return laika, nil
	}

	err = json.Unmarshal(data, &laika)

	if err != nil {
		return nil, err
	}

	return laika, nil
}

func NewLocalJson(path, filename string) *LocalJson {
	return &LocalJson{
		path,
		filename,
	}
}
