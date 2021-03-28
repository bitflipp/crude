package crude

import (
	"database/sql"
)

func WithTransaction(db *sql.DB, f func(tx *sql.Tx) error) (err error) {
	var tx *sql.Tx
	tx, err = db.Begin()
	if err != nil {
		return
	}
	defer func() {
		if err != nil {
			err = tx.Rollback()
		}
	}()
	err = f(tx)
	if err != nil {
		return
	}
	err = tx.Commit()
	return
}
