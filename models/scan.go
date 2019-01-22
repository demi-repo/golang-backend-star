package models

import (
	"time"

	"github.com/EFForg/starttls-backend/checker"
)

const CurrentVersion = 1

// Scan stores the result of a scan of a domain
type Scan struct {
	Domain    string               `json:"domain"`    // Input domain
	Data      checker.DomainResult `json:"scandata"`  // Scan results from starttls-checker
	Timestamp time.Time            `json:"timestamp"` // Time at which this scan was conducted
	Version   uint32               `json:"version"`   // Version counter
}
