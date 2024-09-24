package controller

import (
	"context"
	"errors"
	"fmt"

	"sync"
	"time"

	"github.com/wangkebin/kbds-ref-restapi/dal"
	"github.com/wangkebin/kbds-ref-restapi/fileop"
	"github.com/wangkebin/kbds-ref-restapi/models"

	log "go.uber.org/zap"
)

func GetFiles(ctx context.Context, s string, page int, pagesize int, l *log.Logger) (*models.Files, error) {
	start := time.Now()

	db, err := dal.Connect(l)
	if err != nil {
		l.Sugar().Errorf(err.Error())
		return nil, err
	}

	res, err := dal.SearchFilesByPartName(s, page, pagesize, db, l)
	if err != nil {
		l.Sugar().Errorf(err.Error())
		return nil, err
	}
	l.Info(fmt.Sprintf("GetFiles processing time: %v", time.Since(start)))
	return res, nil
}

func GetDups(ctx context.Context, finfos models.Files, l *log.Logger) (*models.Files, error) {
	start := time.Now()

	db, err := dal.Connect(l)
	if err != nil {
		l.Sugar().Errorf(err.Error())
		return nil, err
	}

	res, err := dal.GetDupFilesByName(finfos, db, l)
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

func DeleteFile(ctx context.Context, fileid int64, l *log.Logger) (string, error) {
	start := time.Now()
	//var wg sync.WaitGroup
	db, dberr := dal.Connect(l)
	if dberr != nil {
		l.Sugar().Errorf(dberr.Error())
		return "", dberr
	}

	f, err := dal.GetFile(fileid, db, l)
	if err != nil {
		l.Sugar().Errorf(err.Error())
		return "file to delete does not exist", err
	}

	//wg.Add(1)
	ch := make(chan error)
	go func() {
		//defer wg.Done()
		ops := fileop.FileOS{}
		err = ops.Delete(f, l)
		if err != nil {
			l.Sugar().Errorf(err.Error())
			ch <- err
		}
		ch <- nil

	}()

	note, err := dal.DeleteFile(fileid, db, l)
	if err != nil {
		l.Sugar().Errorf(err.Error())
		return "failed to delete file record in db", err
	}

	//wg.Wait()

	err = <-ch
	if err != nil {
		l.Sugar().Errorf(err.Error())
		return "failed to delete file", err
	}

	l.Info(fmt.Sprintf("delete file processing time: %v", time.Since(start)))
	return note, nil
}

func DeleteFiles(ctx context.Context, files *[]models.File, l *log.Logger) (string, error) {
	start := time.Now()
	var wg sync.WaitGroup
	db, dberr := dal.Connect(l)
	if dberr != nil {
		l.Sugar().Errorf(dberr.Error())
		return "", dberr
	}

	//wg.Add(1)
	ch := make(chan error)
	errorlist := make([]error, 0)
	wg.Add(len(*files))
	for _, f := range *files {
		go func(f models.File) {
			defer wg.Done()
			ops := fileop.FileOS{}
			err := ops.Delete(&f, l)
			if err != nil {
				l.Sugar().Errorf(err.Error())
				ch <- err
			}
		}(f)
	}
	for er := range ch {
		errorlist = append(errorlist, er)
	}
	ids := make([]int64, 0)
	for _, s := range *files {
		ids = append(ids, s.Id)
	}
	note, err := dal.DeleteFiles(ids, db, l)
	if err != nil {
		l.Sugar().Errorf(err.Error())
		return "failed to delete file record in db", err
	}

	wg.Wait()
	errs := errors.Join(errorlist...)
	if errs != nil {
		l.Sugar().Errorf(errs.Error())
		return "failed to delete files", errs
	}

	l.Info(fmt.Sprintf("delete file processing time: %v", time.Since(start)))
	return note, nil
}
