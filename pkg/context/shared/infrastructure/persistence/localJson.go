package persistence

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/bastean/laika/pkg/context/shared/domain/aggregate"
	"github.com/bastean/laika/pkg/context/shared/domain/repository"
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

func (localJson *LocalJson) Save(data *aggregate.Data) error {
	dataEncoded, err := json.Marshal(data)

	if err != nil {
		return err
	}

	file, err := localJson.create()

	if err != nil {
		return err
	}

	_, err = file.Write(dataEncoded)

	if err != nil {
		return err
	}

	return nil
}

func (localJson *LocalJson) Read() (*aggregate.Data, error) {
	dataDecoded, err := localJson.load()

	if err != nil {
		return nil, err
	}

	data := new(aggregate.Data)

	if len(dataDecoded) == 0 {
		data.Sniffed = make(map[string][]*aggregate.Sniffed)

		return data, nil
	}

	err = json.Unmarshal(dataDecoded, &data)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func NewLocalJson(path, filename string) repository.Repository {
	return &LocalJson{
		path,
		filename,
	}
}
