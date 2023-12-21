package worker

import (
	"database/sql"
	"time"

	"go.uber.org/zap"
)

type Worker struct {
	Interval time.Duration
	DB       *sql.DB
	Logger   *zap.Logger
}

func (w Worker) Do() {
	w.Logger.Info("database garbage collector started ...")

	for {
		if _, err := w.DB.Query("delete from users where deleted_at is not null"); err != nil {
			w.Logger.Error("failed to run delete query", zap.Error(err))
		}

		time.Sleep(w.Interval)
	}
}
