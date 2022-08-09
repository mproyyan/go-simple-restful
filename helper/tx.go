package helper

import (
	"database/sql"
	"fmt"
)

func CommitOrRollback(tx *sql.Tx) {
	err := recover()
	if err != nil {
		errorRollback := tx.Rollback()
		CheckErr(errorRollback)
		fmt.Println(err)
	} else {
		errorCommit := tx.Commit()
		CheckErr(errorCommit)
	}
}
