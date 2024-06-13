package dal

import (
	"github.com/wangkebin/kbds-ref-restapi/models"

	log "go.uber.org/zap"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func SearchFilesByPartName(s string, page int, pagesize int, db *gorm.DB, l *log.Logger) (*models.Files, error) {
	f := make(models.Files, 0)
	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pagesize
	res := db.Debug().Select("*, count(name) as cnt").Offset(offset).Limit(pagesize).
		Where("name like ? group by name,size having cnt > 1", "%"+s+"%").Find(&f)
	if res.Error != nil {
		return nil, res.Error
	}
	return &f, nil
}

func GetDupFilesByName(finfos models.Files, db *gorm.DB, l *log.Logger) (*models.Files, error) {
	f := make(models.Files, 0)
	q := make([][]interface{}, 0)
	for _, finfo := range finfos {
		q = append(q, []interface{}{finfo.Name, finfo.Size})
	}
	res := db.Debug().Where("(name, size) in ?", q).Find(&f)
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
