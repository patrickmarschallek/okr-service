package repository

import (
	"database/sql"
	"okr-service/structures"
	"recipe"
	"recipe/config"

	sq "github.com/Masterminds/squirrel"
	"github.com/kisielk/sqlstruct"
	"github.com/uber-go/zap"
)

func FindOne(id int) (sql.Rows, error) {
	dBCon := recipe.GetDBConn(&config.Settings)
	defer dBCon.Close()

	rows, err := sq.
		Select("*").
		From("obvectives").
		Where("id = ?", id).
		RunWith(dBCon).
		Query()
	recipe.Logger.Debug("find Objective",
		zap.Object("obejctive", rows),
		zap.Int("objectiveId", id),
	)
	return *rows, err
}

func Save(o structures.Objective) (sql.Result, error) {
	dBCon := recipe.GetDBConn(&config.Settings)
	defer dBCon.Close()

	sql, args, err := sq.
		Insert(structures.ObjectiveTable).
		Columns(sqlstruct.Columns(o)).
		Values(o).
		ToSql()

	results, err := dBCon.Exec(sql, args)
	recipe.Logger.Debug("saved Objective",
		zap.String("SQL", sql),
		zap.Object("args", args),
		zap.Error(err),
	)

	return results, err
}
