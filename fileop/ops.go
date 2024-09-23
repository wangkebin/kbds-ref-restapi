package fileop

import (
	models "github.com/wangkebin/kbds-ref-restapi/models"
	log "go.uber.org/zap"
)

type FileOp interface {
	Delete(file *models.File, l *log.Logger) error
}
