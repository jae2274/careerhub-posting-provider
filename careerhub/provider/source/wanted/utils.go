package wanted

import (
	"regexp"
	"strconv"

	"github.com/jae2274/goutils/enum"
)

type CareerTypeValues struct{}

type CareerType = enum.Enum[CareerTypeValues]

const (
	MIN = CareerType("MIN")
	MAX = CareerType("MAX")
)

func (CareerTypeValues) Values() []string {
	return []string{string(MIN), string(MAX)}
}

var (
	minRegex = []*regexp.Regexp{
		regexp.MustCompile(`경력\s*\d+년\s*이상`),
		regexp.MustCompile(`\d+년\s*이상`),
	}
	maxRegex = []*regexp.Regexp{
		regexp.MustCompile(`경력\s*\d+년\s*이하`),
		regexp.MustCompile(`\d+년\s*이하`),
	}
)

// TODO: 테스트코드 작성
func Career(stringVal string, careerType CareerType) (*int32, error) {
	var regexs []*regexp.Regexp
	switch careerType {
	case MIN:
		regexs = minRegex
	case MAX:
		regexs = maxRegex
	}

	for _, regex := range regexs {
		match := regex.FindString(stringVal)
		if match != "" {
			intMatch := regexp.MustCompile(`\d+`).FindString(match)
			if intMatch != "" {
				num, error := strconv.Atoi(intMatch)
				if error != nil {
					return nil, error
				}
				num32 := int32(num)
				return &num32, nil
			}
		}
	}

	return nil, nil
}
