package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/idexter/dip-visa-app/pkg/visa"
)

// FileApplicationsStorage implements file Visa Applications storage.
type FileApplicationsStorage struct {
	Database string
}

// NewFileApplicationsStorage creates new FileApplicationsStorage.
func NewFileApplicationsStorage() *FileApplicationsStorage {
	return &FileApplicationsStorage{Database: defaultApplicationsDB}
}

// GetVisaApplication gets visa application by application id.
func (as *FileApplicationsStorage) GetVisaApplication(id int) (*visa.Application, error) {

	var apps []StoredApplication
	b, err := ioutil.ReadFile(as.Database)
	if err != nil {
		return nil, fmt.Errorf("couldn't read applications database %w", err)
	}

	if err := json.Unmarshal(b, &apps); err != nil {
		return nil, fmt.Errorf("couldn't unmarshal applications database %w", err)
	}

	for _, v := range apps {
		if v.ID == id {
			return &visa.Application{
				ID:        v.ID,
				Name:      v.Name,
				Arrival:   v.Arrival,
				Departure: v.Departure,
				Money:     v.Money,
			}, nil
		}
	}

	return nil, errors.New("application was not found")
}
