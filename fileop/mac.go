package fileop

import (
	"os"

	models "github.com/wangkebin/kbds-ref-restapi/models"
	log "go.uber.org/zap"
)

type FileOS struct {
}

func (m *FileOS) Delete(file *models.File, l *log.Logger) error {
	f, err := os.Open(file.Loc)
	if err != nil {
		l.Sugar().Errorf(err.Error())
		return err
	}
	// Windows does not allow deleting open file
	err = f.Close()
	if err != nil {
		l.Sugar().Errorf(err.Error())
		return err
	}
	err = os.Remove(file.Loc)
	if err != nil {
		l.Sugar().Errorf(err.Error())
		return err
	}
	l.Sugar().Infof("file %s deleted", file.Loc)
	return nil
}
