package validation

func (v validation) Email(email string) bool {
	return v.emailRegex.MatchString(email)
}
