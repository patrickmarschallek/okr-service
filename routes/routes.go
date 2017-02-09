package routes

import (
	"okr-service/hanlders"
	"recipe"
)

var AppRoutes = recipe.Routes{

	recipe.Route{"ObjectiveList", "GET", "/objectives", hanlders.ListObjectives},
	recipe.Route{"findObjective", "GET", "/objectives/{objectiveID}", hanlders.FindObjective},
	recipe.Route{"addKeyResultsToObjective", "POST", "/objectives", hanlders.PostObjective},
	recipe.Route{"ReplaceObjective", "PUT", "/objectives/{objectiveId}", hanlders.ReplaceObjective},

	recipe.Route{"replaceKeyResultsForObjective", "PUT", "/objectives/{objectiveId}/key-kesults", recipe.IndexHandler},
}
