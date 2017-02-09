package repository

import (
	"database/sql"
	"net/http"
	"okr-service/structures"
	"recipe"
	"recipe/config"

	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/kisielk/sqlstruct"
	"github.com/uber-go/zap"
)

// FindAllObjectives will return all observations.
func FindAllObjectives() (*structures.Objectives, error) {
	var objectives structures.Objectives
	dBCon := recipe.GetDBConn(&config.Settings)
	defer dBCon.Close()

	rows, err := sq.
		Select("*").
		From(structures.ObjectiveTable).
		RunWith(dBCon).
		Query()
	if err != nil {
		return nil, recipe.NewErrorResponse("database error", http.StatusInternalServerError, err)
	}
	defer rows.Close()
	for rows.Next() {
		var o structures.Objective
		err = sqlstruct.Scan(&o, rows)
		if err != nil {
			return nil, recipe.NewErrorResponse("database error", http.StatusInternalServerError, err)
		}
		objectives = objectives.Add(&o)
	}

	recipe.Logger.Debug("find Objective",
		zap.Object("obejctive", objectives),
	)
	return &objectives, err
}

// FindOneObjective finds an objective by the given id.
func FindOneObjective(id int) (*structures.Objective, error) {
	dBCon := recipe.GetDBConn(&config.Settings)
	defer dBCon.Close()

	rows, err := sq.
		Select("*").
		From("obvectives").
		Where("id = ?", id).
		RunWith(dBCon).
		Query()

	if err != nil {
		return nil, recipe.NewErrorResponse("database error", http.StatusInternalServerError, err)
	}
	defer rows.Close()

	if rows.Next() {
		var o *structures.Objective
		err = sqlstruct.Scan(o, rows)
		if err != nil {
			return nil, recipe.NewErrorResponse("database error", http.StatusInternalServerError, err)
		}
		return o, nil
	}

	return nil, recipe.NewErrorResponse(
		"Not Found",
		http.StatusNotFound,
		fmt.Errorf("no objectives found for id %d", id),
	)

}

// SaveObjective stores an objective to the database.
func SaveObjective(o structures.Objective) (sql.Result, error) {
	dBCon := recipe.GetDBConn(&config.Settings)
	defer dBCon.Close()

	results, err := sq.
		Insert(structures.ObjectiveTable).
		Columns(structures.ObjectiveColumns...).
		Values(
			o.Title,
			o.Description,
			o.Grade,
			o.StartDate,
			o.EndDate,
		).
		RunWith(dBCon).
		Exec()

	recipe.Logger.Debug("saved Objective",
		zap.Object("results", results),
		zap.Error(err),
	)

	return results, err
}

// UpdateObjective updates or creates an objective.
func UpdateObjective(o structures.Objective) (sql.Result, error) {
	dBCon := recipe.GetDBConn(&config.Settings)
	defer dBCon.Close()

	results, err := sq.
		Update(structures.ObjectiveTable).
		SetMap(map[string]interface{}{
			"title":       o.Title,
			"description": o.Description,
			"grade":       o.Grade,
			"startDate":   o.StartDate,
			"endDate":     o.EndDate,
		}).
		RunWith(dBCon).
		Exec()

	recipe.Logger.Debug("saved Objective",
		zap.Object("results", results),
		zap.Error(err),
	)

	return results, err
}
