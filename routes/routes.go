package routes

import (
	"okr-service/hanlders"
	"recipe"
)

var AppRoutes = recipe.Routes{
	recipe.Route{"ObjectiveList", "GET", "/objectives", hanlders.TodoIndex},
	recipe.Route{"findObjective", "GET", "/objectives/{objectiveId}", hanlders.TodoShow},
	recipe.Route{"addKeyResultsToObjective", "POST", "/objectives", hanlders.PostObjective},
	recipe.Route{"replaceKeyREsultsForObjective", "PUT", "/objectives/{objectiveId}/key-kesult", hanlders.TodoShow},
	recipe.Route{"ReplaceObjective", "PUT", "/objectives/{objectiveId}", hanlders.TodoShow},
}
