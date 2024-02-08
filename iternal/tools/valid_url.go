package tools

import (
	"errors"
	parseUrl "net/url"
)

func ValidUrl(str string) error {
	_, err := parseUrl.ParseRequestURI(str)
	if err != nil {
		return errors.New("не валидный урл " + str)
	}
	return nil
}
