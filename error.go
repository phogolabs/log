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
func FieldsOfError(err error) Fields {
	fields := Fields{}

	if cerr, ok := err.(ErrorCauser); ok {
		fields["error"] = cerr.Cause().Error()
		fields["source"] = cerr.StackTrace()
		return fields
	}

	chain, ok := err.(gerrors.Chain)
	if !ok {
		chain = gerrors.WrapSkipFrames(err, "", 2)
	}

	var (
		msg   string
		types = make([]string, 0, len(chain))
	)

	for index, e := range chain {
		if index == 0 {
			msg = e.Err.Error()
			fields["source"] = e.Source
		}

		if e.Prefix != "" {
			msg = fmt.Sprintf("%s: %s", e.Prefix, msg)
		}

		for _, tag := range e.Tags {
			fields[tag.Key] = tag.Value
		}

		types = append(types, e.Types...)
	}

	fields["error"] = msg

	if len(types) > 0 {
		fields["types"] = strings.Join(types, ",")
	}

	return fields
}
