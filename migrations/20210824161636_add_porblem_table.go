package migrations

import (
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upAddPorblemTable, downAddPorblemTable)
}

func upAddPorblemTable(tx *sql.Tx) error {
	_, err := tx.Exec(
		"CREATE TABLE problem (" +
			"id int not null," +
			"user_id int," +
			"message text," +
			"PRIMARY KEY(id)"+
		");",
		)
	return err
}

func downAddPorblemTable(tx *sql.Tx) error {
	_, err := tx.Exec("CREATE TABLE problem;")
	return err
}
