package storage

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/idexter/dip-visa-app/pkg/visa"
)

// FileVisasStorage implements Visas database based on files.
type FileVisasStorage struct {
	Database string
}

// NewFileVisasStorage creates new instance for FileVisasStorage.
func NewFileVisasStorage() *FileVisasStorage {
	return &FileVisasStorage{Database: defaultVisasDB}
}

// GetPreviousVisas gets previous Visas for provided applicant.
func (vs *FileVisasStorage) GetPreviousVisas(applicantName string) ([]visa.Visa, error) {

	var visas map[string][]StoredVisa
	b, err := ioutil.ReadFile(vs.Database)
	if err != nil {
		return nil, fmt.Errorf("couldn't read visas database %w", err)
	}

	if err := json.Unmarshal(b, &visas); err != nil {
		return nil, fmt.Errorf("couldn't unmarshal visas database %w", err)
	}

	var ret = make([]visa.Visa, len(visas))

	if v, ok := visas[applicantName]; ok {
		for _, current := range v {
			ret = append(ret, visa.Visa{
				From:      current.From,
				To:        current.To,
				Arrival:   current.Departure,
				Departure: current.Arrival,
			})
		}
	}

	return ret, nil
}
