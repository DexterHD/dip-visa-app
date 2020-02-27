package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"time"
)

type Application struct {
	ID        int
	Name      string
	Arrival   time.Time
	Departure time.Time
	Money     float64
}

type Visa struct {
	From      time.Time
	To        time.Time
	Arrival   time.Time
	Departure time.Time
}

var ApplicationsDB = "data/applications.json"
var VisasDB = "data/visas.json"

func GetVisaApplication(id int) (*Application, error) {

	var apps []Application
	b, err := ioutil.ReadFile(ApplicationsDB)
	if err != nil {
		return nil, fmt.Errorf("couldn't read applications database %w", err)
	}

	if err := json.Unmarshal(b, &apps); err != nil {
		return nil, fmt.Errorf("couldn't unmarshal applications database %w", err)
	}

	for _, v := range apps {
		if v.ID == id {
			return &v, nil
		}
	}

	return nil, errors.New("application was not found")
}

func GetPreviousVisas(name string) ([]Visa, error) {

	var visas map[string][]Visa
	b, err := ioutil.ReadFile(VisasDB)
	if err != nil {
		return nil, fmt.Errorf("couldn't read visas database %w", err)
	}

	if err := json.Unmarshal(b, &visas); err != nil {
		return nil, fmt.Errorf("couldn't unmarshal visas database %w", err)
	}

	if v, ok := visas[name]; ok {
		return v, nil
	}

	return []Visa{}, nil
}
