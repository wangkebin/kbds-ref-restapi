package models

import (
	"strings"

	resource "kbds-ref-restapi/gen/models"
)

type File struct {
	Id   int64  `gorm:"column:id;primarykey;autoincrement"`
	Name string `gorm:"column:name"`
	Loc  string `gorm:"column:loc"`
	Ext  string `gorm:"column:ext"`
	Size int64  `gorm:"column:size"`
}

func (f *File) FromResource(src *resource.File) {
	f.Id = src.ID
	f.Name = strings.TrimSpace(*src.Name)
	f.Loc = strings.TrimSpace(src.Loc)
	f.Ext = strings.TrimSpace(src.Ext)
	f.Size = src.Size

}

func (f *File) ToResource() *resource.File {
	var src = new(resource.File)
	src.ID = f.Id
	src.Name = &f.Name
	src.Loc = f.Loc
	src.Ext = f.Ext
	src.Size = f.Size

	return src
}

func (f *File) TableName() string {
	return "file_info"
}
