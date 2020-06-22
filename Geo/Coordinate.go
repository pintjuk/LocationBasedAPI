package Geo

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type Coordinate struct {
	// Float32 should be good enouh, it can represent 7 decimal places,
	// witch is good enough for specifying geolocations with centimeter accuracy
	Longitude float32
	Latitude  float32
}

func (c Coordinate) ToString() string {
	return fmt.Sprintf("(%f,%f)", c.Longitude, c.Latitude)
}

/*parse from string with format: (lat, long)*/
func ParseString(s string) (error, Coordinate) {
	if len(s) < 2 {
		return errors.New("failed to parse"), Coordinate{}
	}
	withoutSpaces := strings.TrimFunc(s, func(r rune) bool {
		return unicode.IsSpace(r)
	})
	if withoutSpaces[0] != '(' {
		return errors.New("failed to parse"), Coordinate{}
	}
	if withoutSpaces[len(withoutSpaces)-1] != ')' {
		return errors.New("failed to parse"), Coordinate{}
	}
	withoutBraces := strings.TrimFunc(withoutSpaces, func(r rune) bool {
		return r == '(' || r == ')'
	})
	parts := strings.Split(withoutBraces, ",")
	if len(parts) != 2 {
		return errors.New("failed to parse"), Coordinate{}
	}
	lat, err := strconv.ParseFloat(strings.Trim(parts[0], " "), 32)
	if err != nil {
		return err, Coordinate{}
	}
	long, err2 := strconv.ParseFloat(strings.Trim(parts[1], " "), 32)
	if err2 != nil {
		return err2, Coordinate{}
	}
	return nil, Coordinate{float32(lat), float32(long)}
}
