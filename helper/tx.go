package helper

import "database/sql"

// fix_this_shit !!!

func CommitOrRollback(tx *sql.Tx) {
	err := recover()
	if err != nil {
		errorRollback := tx.Rollback()
		Panic(errorRollback)
		panic(err)
	} else {
		errorCommit := tx.Commit()
		Panic(errorCommit)
	}
}
