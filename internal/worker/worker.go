package worker

import (
	"database/sql"
	"time"

	"github.com/go-logr/logr"
)

type Worker struct {
	Interval time.Duration
	DB       *sql.DB
	Logger   logr.Logger
}

func (w Worker) Do() {
	w.Logger.Info("database garbage collector started ...")

	for {
		if _, err := w.DB.Query("delete from users where deleted_at is not null"); err != nil {
			w.Logger.Error(err, "failed to run delete query")
		}

		time.Sleep(w.Interval)
	}
}
