package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/DexterHD/dip-visa-app/pkg/visa"
)

// StoredApplication describes Visa Application stored in Database.
type StoredApplication struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Arrival   time.Time `json:"arrival"`
	Departure time.Time `json:"departure"`
	Money     float64   `json:"money"`
}

const defaultApplicationsDB = "data/applications.json"
const defaultVisasDB = "data/visas.json"

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

// StoredVisa describes Visa stored in database.
type StoredVisa struct {
	From      time.Time `json:"from"`
	To        time.Time `json:"to"`
	Arrival   time.Time `json:"arrival"`
	Departure time.Time `json:"departure"`
}

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
