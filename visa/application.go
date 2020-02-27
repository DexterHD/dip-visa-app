package visa

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

func GetVisaApplication(id int, filename string) (*Application, error) {

	var apps []Application
	b, err := ioutil.ReadFile(filename)
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
