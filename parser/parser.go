package parser

import (
	"errors"
	"regexp"
)

// Parse parses an HTTP request
func Parse(input string) (req Request, err error) {
	reqRegex, _ := regexp.Compile(`(GET|POST)\s+(.+)\s+HTTP\/1\.1`)

	if matches := reqRegex.MatchString(input); !matches {
		return Request{}, errors.New("invalid request")
	}

	slices := reqRegex.FindAllStringSubmatch(input, -1)
	method := slices[0][1]
	path := slices[0][2]

	return Request{Method: method, Path: path}, nil
}
