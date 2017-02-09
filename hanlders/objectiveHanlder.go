package hanlders

import (
	"net/http"
	"okr-service/repository"
	"okr-service/structures"
	"recipe"

	"strconv"

	"github.com/gorilla/mux"
	"github.com/uber-go/zap"
)

func FindObjective(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	var uriParams = mux.Vars(r)
	var ID, err = strconv.Atoi(uriParams["objectiveID"])
	if err != nil {
		return nil, recipe.NewErrorResponse("ID isn't a number value", http.StatusBadRequest, err)
	}
	objectives, err := repository.FindOneObjective(ID)

	recipe.Logger.Debug("create observation",
		zap.Object("result", objectives),
	)
	return objectives, err
}

func ListObjectives(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	objectives, err := repository.FindAllObjectives()

	recipe.Logger.Debug("create observation",
		zap.Object("result", objectives),
	)
	return objectives, err
}

// PostObjective handles incoming objective.
func PostObjective(w http.ResponseWriter, req *http.Request) (interface{}, error) {
	recipe.Logger.Debug("start handling",
		zap.Object("length", req.ContentLength),
	)
	var objective structures.Objective

	// validation step
	if err := recipe.JSONDecodeAndValidate(req, &objective); err != nil {
		recipe.Logger.Error("decode/validation error",
			zap.Error(err),
		)
		return nil, recipe.NewErrorResponse("validation error", http.StatusBadRequest, err)
	}

	result, err := repository.SaveObjective(objective)
	if err != nil {
		return nil, recipe.NewErrorResponse("database error", http.StatusInternalServerError, err)
	}

	recipe.Logger.Debug("create observation",
		zap.Object("result", result),
	)
	return objective, nil
}

// ReplaceObjective handles incoming objective.
func ReplaceObjective(w http.ResponseWriter, req *http.Request) (interface{}, error) {
	recipe.Logger.Debug("start handling",
		zap.Object("length", req.ContentLength),
	)
	var objective *structures.Objective

	// validation step
	if err := recipe.JSONDecodeAndValidate(req, objective); err != nil {
		return nil, recipe.NewErrorResponse("validation error", http.StatusBadRequest, err)
	}

	result, err := repository.UpdateObjective(*objective)
	if err != nil {
		return nil, recipe.NewErrorResponse("database error", http.StatusInternalServerError, err)
	}

	recipe.Logger.Debug("put observation",
		zap.Object("result", result),
	)
	return result, nil
}
