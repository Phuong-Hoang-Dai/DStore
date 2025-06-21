package utils

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func ParseCustomDuration(s string) (time.Duration, error) {
	duration, err := strconv.Atoi(s[0 : len(s)-1])
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	var durationTime time.Duration
	switch suffix := s[len(s)-1]; suffix {
	case 'h':
		durationTime = time.Duration(duration) * time.Hour
	case 'd':
		durationTime = time.Duration(duration) * time.Hour * 24
	default:
		return 0, fmt.Errorf("utils/ParseCustomDuration: Incorrect format. Input: %v", s)
	}
	return durationTime, nil
}

func SpaceToDash(s string) string {
	s = strings.ReplaceAll(s, " ", "-")
	return s
}
