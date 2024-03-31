package controller

import (
	"context"
	"fmt"
	"kbds-ref-restapi/dal"
	"kbds-ref-restapi/models"
	"time"

	log "go.uber.org/zap"
)

func GetFiles(ctx context.Context, s string, l *log.Logger) (*models.Files, error) {
	start := time.Now()

	db, err := dal.Connect(l)
	if err != nil {
		l.Sugar().Errorf(err.Error())
		return nil, err
	}

	res, err := dal.MatchFiles(s, db, l)
	if err != nil {
		l.Sugar().Errorf(err.Error())
		return nil, err
	}
	l.Info(fmt.Sprintf("GetFiles processing time: %v", time.Since(start)))
	return res, nil
}

func PostFiles(ctx context.Context, f *models.Files, l *log.Logger) error {
	start := time.Now()

	db, err := dal.Connect(l)
	if err != nil {
		l.Sugar().Errorf(err.Error())
		return err
	}

	err = dal.SaveFiles(f, db, l)
	if err != nil {
		l.Sugar().Errorf(err.Error())
		return err
	}
	l.Info(fmt.Sprintf("PostFiles processing time: %v", time.Since(start)))
	return nil
}
