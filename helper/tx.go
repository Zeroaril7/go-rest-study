package helper

import "database/sql"

func CommitOrRollBack(tx *sql.Tx) {
	err := recover()
	if err != nil {
		errRollBack := tx.Rollback()
		PanicIfError(errRollBack)
		panic(err)
	}

	errorCommit := tx.Commit()
	PanicIfError(errorCommit)
}
