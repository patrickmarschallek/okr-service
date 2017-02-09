package structures

import (
	"errors"
	"time"

	valid "github.com/asaskevich/govalidator"
)

// ObjectiveTable is the table name.
const ObjectiveTable = "objectives"

// ObjectiveColumns counatins a slice with the column names.
var ObjectiveColumns = []string{
	"title",
	"description",
	"grade",
	"start_date",
	"end_date",
}

var (
	ErrorInvalidTitle       = errors.New("title shouldn't be empty")
	ErrorInvalidDescription = errors.New("description shouldn't be empty")
	ErrorInvalidGrade       = errors.New("grade should be between 0 and 1.")
)

// Objectives a list of Objective
type Objectives []Objective

// Objective data structure for an obvective
type Objective struct {
	ID          uint64      `json:"id" sql:"id"`
	CreatedAt   time.Time   `json:"-" sql:"created_at"`
	UpdatedAt   time.Time   `json:"-" sql:"update_at"`
	DeletedAt   *time.Time  `json:"-" sql:"deleted_at"`
	Title       *string     `json:"title" sql:"title"`
	Description *string     `json:"description" sql:"description"`
	Grade       *Grade      `json:"keyResults" sql:"grade"`
	StartDate   *time.Time  `json:"startDate" sql:"start_date"`
	EndDate     *time.Time  `json:"endDate" sql:"end_date"`
	KeyResults  *KeyResults `json:"keyResults" sql:"-"`
}

// Remaining give the remaining time back
func (o Objective) Remaining() time.Duration {
	return o.EndDate.Sub(*o.StartDate)
}

// Valid validates Objective.
func (o Objective) Valid() error {
	// validate the title is not empty or missing
	if o.Title == nil || valid.IsNull(*o.Title) {
		return ErrorInvalidTitle
	}
	// validate the description is not empty or missing
	if o.Description == nil || valid.IsNull(*o.Description) {
		return ErrorInvalidDescription
	}
	if o.Grade != nil && (*o.Grade > 1 || *o.Grade < 0) {
		return ErrorInvalidGrade
	}
	return nil
}

// Add adds an Objective to the objectives.
func (o Objectives) Add(objective *Objective) Objectives {
	list := make([]Objective, len(o))
	list = append(list, *objective)
	return list
}
