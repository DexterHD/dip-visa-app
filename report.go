package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func SaveApplicationReport(vs Violations) (int, error) {

	data, err := json.Marshal(vs)
	if err != nil {
		return 0, fmt.Errorf("could not marshall violations, reason %w", err)
	}

	err = ioutil.WriteFile(fmt.Sprintf("reports/violations-%d.json", vs.ApplicationID), data, os.FileMode(0777))
	if err != nil {
		return 0, fmt.Errorf("could not write violations, reason %w", err)
	}

	return vs.ApplicationID, nil
}

func PrintApplicationReport(reportId int) error {

	b, err := ioutil.ReadFile(fmt.Sprintf("reports/violations-%d.json", reportId))
	if err != nil {
		return fmt.Errorf("could not read violations, reason %w", err)
	}

	vs := Violations{}
	err = json.Unmarshal(b, &vs)
	if err != nil {
		return fmt.Errorf("could not unmarshall, reason %w", err)
	}

	fmt.Printf("%v\n", vs)
	// os.Remove(fmt.Sprintf("data/violations-%d.json", reportId))
	return nil
}
