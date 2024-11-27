package migrations

import (
	"os"
	"path"
	"path/filepath"
	"reflect"
	"testing"

	"gofr.dev/pkg/gofr"
	"gofr.dev/pkg/gofr/config"
	"gofr.dev/pkg/gofr/migration"
)

func TestAll(t *testing.T) {
	testdb, _ := filepath.Abs(path.Join(t.TempDir(), "./testing.db"))
	t.Logf("Test database: %s", testdb)
	os.Remove(testdb)
	defer os.Remove(testdb)

	t.Setenv("DB_NAME", testdb)
	t.Setenv("DB_DIALECT", "sqlite")

	app := gofr.New()

	// Add migrations to run
	app.Migrate(All(app.Config))
	type args struct {
		cfg config.Config
	}
	tests := []struct {
		name string
		args args
		want map[int64]migration.Migrate
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := All(tt.args.cfg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("All() = %v, want %v", got, tt.want)
			}
		})
	}
}
