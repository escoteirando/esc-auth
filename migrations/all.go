package migrations

import (
	"gofr.dev/pkg/gofr/config"
	"gofr.dev/pkg/gofr/migration"
)

func getSQL(dialect string, sqliteSQL string, pgSQL string) (sql string) {
	if dialect == "sqlite" {
		sql = sqliteSQL
	} else {
		sql = pgSQL
	}
	if len(sql) == 0 {
		panic("missing migration for dialect " + dialect)
	}
	return sql
}

func All(cfg config.Config) map[int64]migration.Migrate {
	dialect := cfg.GetOrDefault("DB_DIALECT", "sqlite")
	return map[int64]migration.Migrate{
		20240226153000: createTableUsers(dialect),
	}
}
