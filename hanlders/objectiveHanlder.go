package hanlders

import (
	"encoding/json"
	"net/http"
	"okr-service/repository"
	"okr-service/structures"
	"recipe"
	"time"

	"github.com/uber-go/zap"

	"github.com/gorilla/mux"
)

func TodoIndex(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	objective := structures.Objective{
		Description: structures.Description{
			Title: "be a big player",
			Desc:  "",
		},
		StartDate: time.Now(),
		EndDate:   time.Now().AddDate(0, 0, 30),
		KeyResults: &structures.KeyResults{
			structures.KeyResult{
				Description: structures.Description{
					Title: "deliver first service in beginning of 2017",
					Desc:  "",
				},
			},
		},
	}
	return objective, nil
}

func PostObjective(w http.ResponseWriter, req *http.Request) (interface{}, error) {
	recipe.Logger.Debug("start handling",
		zap.Object("length", req.ContentLength),
	)
	// validation step
	if req.ContentLength == 0 {
		return nil, recipe.NewErrorResponse("empty payload", http.StatusUnprocessableEntity)
	}

	// json to structure
	decoder := json.NewDecoder(req.Body)
	var objective *structures.Objective

	err := decoder.Decode(&objective)
	if err != nil {
		return nil, recipe.NewErrorResponse("empty payload", http.StatusUnprocessableEntity, err)
	}
	defer req.Body.Close()

	result, err := repository.Save(*objective)
	if err != nil {
		return nil, recipe.NewErrorResponse("database error", http.StatusInternalServerError, err)
	}

	recipe.Logger.Debug("ROWS",
		zap.Object("rows", result),
	)

	return objective, nil
}

func TodoShow(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	objectiveId := vars["objectiveId"]
	return objectiveId, nil
}
