package models

import (
	resource "kbds-ref-restapi/gen/models"
)

type Files []File

func (c *Files) FromResource(src *resource.Files) {
	for _, s := range *src {
		f := new(File)
		f.FromResource(s)
		*c = append(*c, *f)
	}
}

func (c *Files) ToResource() *resource.Files {
	res := new(resource.Files)
	for _, s := range *c {
		*res = append(*res, s.ToResource())
	}
	return res
}

func (c *Files) TableName() string {
	return "file_info"
}
