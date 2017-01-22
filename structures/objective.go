package structures

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

const ObjectiveTable = "objectives"

// Objectives a list of Objective
type Objectives []Objective

// Objective data structure for an obvective
type Objective struct {
	Description
	Grade
	KeyResults *KeyResults `json:"keyResults" sql:"-"`
	StartDate  time.Time   `json:"startDate"`
	EndDate    time.Time   `json:"endDate"`
}

// Remaining give the remaining time back
func (o Objective) Remaining() time.Duration {
	return o.EndDate.Sub(o.StartDate)
}

func (o Objective) Validate(req *http.Request) bool {
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&o)
	if err != nil {
		log.Fatal(err)
		return false
	}
	defer req.Body.Close()
	return true
}
