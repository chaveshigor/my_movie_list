package validations

import (
	"errors"
	"fmt"
	"regexp"
	"unicode/utf8"
)

type Rules struct {
	Presence   bool
	MinimumLen int
	Regex      string
}

func Validate(errs *[]error, name string, value any, validations Rules) {
	// string validations
	strValue, ok := value.(string)
	if ok {
		if validations.Presence {
			err := presence(name, strValue)
			appendErr(errs, err)
		}

		if validations.MinimumLen > 0 {
			err := stringMinLen(name, strValue, validations.MinimumLen)
			appendErr(errs, err)
		}

		if len(validations.Regex) > 0 {
			err := regexMatch(name, strValue, validations.Regex)
			appendErr(errs, err)
		}
	}
}

func presence(name, value string) error {
	if len(value) == 0 {
		errorMessage := fmt.Sprintf("%s can't be blank", name)
		return errors.New(errorMessage)
	}

	return nil
}

func stringMinLen(name string, value string, minimumLen int) error {
	if utf8.RuneCountInString(value) < minimumLen {
		errorMessage := fmt.Sprintf("%s need to have at least %d characters", name, minimumLen)
		return errors.New(errorMessage)
	}

	return nil
}

func regexMatch(name string, value string, regex string) error {
	re, err := regexp.Compile(regex)
	if err != nil {
		errorMessage := fmt.Sprintf("invalid regex %s", regex)
		return errors.New(errorMessage)
	}

	if !re.MatchString(value) {
		errorMessage := fmt.Sprintf("%s has an invalid format", name)
		return errors.New(errorMessage)
	}

	return nil
}

func appendErr(errs *[]error, err error) {
	if err != nil {
		*errs = append(*errs, err)
	}
}
