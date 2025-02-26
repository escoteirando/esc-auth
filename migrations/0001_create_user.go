package migrations

import "gofr.dev/pkg/gofr/migration"

const createTableUsersSQLite = `CREATE TABLE users (
	id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	username TEXT NOT NULL UNIQUE,
	password TEXT NOT NULL,
	person_id INTEGER,
	"role" INTEGER NOT NULL
);`

const createTableUsersPG = `` //TODO: Implementar tabela pg.users

func createTableUsers(dialect string) migration.Migrate {
	return migration.Migrate{
		UP: func(d migration.Datasource) error {
			if _, err := d.SQL.Exec(getSQL(dialect, createTableUsersSQLite, createTableUsersPG)); err != nil {
				d.Logger.Errorf("failed on creating users table - %v", err)
				return err
			}

			return nil
		},
	}
}
