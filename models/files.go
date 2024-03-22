package models

import (
	resource "github.com/wangkebin/kbds-ref-restapi/gen/models"
)

type Files []File

func (c *Files) FromResource(src *resource.Files) {
	for _, s := range *src {
		f := new(File)
		f.FromResource(s)
		*c = append(*c, *f)
	}
}

func (c *Files) TableName() string {
	return "dirs"
}
