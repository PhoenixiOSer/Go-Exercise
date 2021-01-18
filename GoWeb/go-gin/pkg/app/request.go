package app

import (
	"github.com/EDDYCJY/go-gin-example/pkg/logging"
	"github.com/astaxie/beego/validation"
)

func MarkErrors(errors []*validation.Error) {
	for _,err := range errors {
		logging.Info(err)
	}
	return
}

