package log

import (
	"fmt"
	"strings"

	gerrors "github.com/go-playground/errors"
	perrors "github.com/pkg/errors"
)

// ErrorCauser returns the actual error
type ErrorCauser interface {
	Cause() error
	StackTrace() perrors.StackTrace
}

// FieldsOfError returns the fields for given error
func FieldsOfError(err error) []Field {
	fields := []Field{}

	if cerr, ok := err.(ErrorCauser); ok {
		fields = append(fields, F("error", cerr.Cause().Error()))
		fields = append(fields, F("source", cerr.StackTrace()))
		return fields
	}

	chain, ok := err.(gerrors.Chain)
	if !ok {
		chain = gerrors.WrapSkipFrames(err, "", 2)
	}

	var (
		msg   string
		types = make([]string, 0, len(chain))
		tags  = make([]Field, 0, len(chain))
	)

	for index, e := range chain {
		if index == 0 {
			msg = e.Err.Error()
			fields = append(fields, F("source", e.Source))
		}

		if e.Prefix != "" {
			msg = fmt.Sprintf("%s: %s", e.Prefix, msg)
		}

		for _, tag := range e.Tags {
			tags = append(tags, F(tag.Key, tag.Value))
		}

		types = append(types, e.Types...)
	}

	fields = append(fields, F("error", msg))
	fields = append(fields, tags...)

	if len(types) > 0 {
		fields = append(fields, F("types", strings.Join(types, ",")))
	}

	return fields
}
