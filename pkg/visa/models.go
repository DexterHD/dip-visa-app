package visa

import "time"

// MaxTimeToStay described maximum time in hours from arrival to departure.
const MaxTimeToStay = 24 * 90

// Application described Visa Application domain model.
type Application struct {
	ID        int
	Name      string
	Arrival   time.Time
	Departure time.Time
	Money     float64
}

// Visa described Visa domain model.
type Visa struct {
	From      time.Time
	To        time.Time
	Arrival   time.Time
	Departure time.Time
}

// Report described Visa Application Report domain model.
// In other words this structure described Visa Application Result.
type Report struct {
	ID            int64
	ApplicationID int
	Applicant     string
	Accepted      bool
}
