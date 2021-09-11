package validation

import "github.com/bektosh03/pkg/errs"

func (v validation) Email(email string) error {
	matches := v.emailRegex.MatchString(email)
	if !matches {
		return errs.ErrBadEmail
	}
	return nil
}
