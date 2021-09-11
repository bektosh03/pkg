package validation

import "regexp"

type Validation interface {
	Email(email string) error
}

func New() Validation {
	return validation{
		emailRegex: regexp.MustCompile(
			"^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$",
		),
	}
}

type validation struct {
	emailRegex *regexp.Regexp
}
