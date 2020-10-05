// Package storage implements persistent storages for different types of business models.
// It also uses different structures to store data in filesystem.
// It's not required but implemented to demonstrate layer separation.
package storage

import "time"

const defaultApplicationsDB = "data/applications.json"
const defaultVisasDB = "data/visas.json"

// StoredApplication describes Visa Application stored in Database.
type StoredApplication struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Arrival   time.Time `json:"arrival"`
	Departure time.Time `json:"departure"`
	Money     float64   `json:"money"`
}

// StoredVisa describes Visa stored in database.
type StoredVisa struct {
	From      time.Time `json:"from"`
	To        time.Time `json:"to"`
	Arrival   time.Time `json:"arrival"`
	Departure time.Time `json:"departure"`
}
