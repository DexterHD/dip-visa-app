package visa

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"time"
)

type Visa struct {
	From      time.Time
	To        time.Time
	Arrival   time.Time
	Departure time.Time
}

func GetPreviousVisas(name string, filename string) ([]Visa, error) {

	var visas map[string][]Visa
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("couldn't read visas database %w", err)
	}

	if err := json.Unmarshal(b, &visas); err != nil {
		return nil, fmt.Errorf("couldn't unmarshal visas database %w", err)
	}

	if v, ok := visas[name]; ok {
		return v, nil
	}

	return nil, errors.New("Visas not found")
}
