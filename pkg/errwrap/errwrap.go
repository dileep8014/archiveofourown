package errwrap

import (
	"fmt"
	"github.com/shyptr/archiveofourown/pkg/errcode"
)

func Wrap(errp *error, format string, args ...interface{}) {
	if *errp != nil {
		if _, ok := (*errp).(errcode.Error); ok {
			return
		}
		*errp = fmt.Errorf("%s: %w", fmt.Sprintf(format, args...), *errp)
	}
}

func Add(errp *error, format string, args ...interface{}) {
	if *errp != nil {
		if _, ok := (*errp).(errcode.Error); ok {
			return
		}
		*errp = fmt.Errorf("%s: %v", fmt.Sprintf(format, args...), *errp)
	}
}
