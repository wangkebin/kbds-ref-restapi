package dal

import (
	"github.com/wangkebin/kbds-ref-restapi/models"

	log "go.uber.org/zap"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func MatchFiles(s string, db *gorm.DB, l *log.Logger) (*models.Files, error) {
	f := make(models.Files, 0)
	res := db.Debug().Where("name like ? group by name,size having count(*) > 1", "%"+s+"%").Find(&f)
	if res.Error != nil {
		return nil, res.Error
	}
	return &f, nil
}

func SaveFiles(f *models.Files, db *gorm.DB, l *log.Logger) error {
	res := db.Debug().Clauses(clause.OnConflict{
		Columns:   []clause.Column{{}},
		UpdateAll: true,
	}).CreateInBatches(f, len(*f))

	if res.Error != nil {
		return res.Error
	}
	return nil
}
