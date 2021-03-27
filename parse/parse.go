package parse

import (
	"protoc-gen-rest/model"

	pgs "github.com/lyft/protoc-gen-star"
)

type Parser interface {
	GetTemplateInfo(f pgs.File) model.TemplateData
}
