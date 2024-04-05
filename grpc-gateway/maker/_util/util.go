package _util

import (
	"strings"
	"time"
	"unicode"
)

func Time() string {
	currentTime := time.Now()
	currentTimeString := currentTime.Format("2006-01-02 15:04:05")
	return currentTimeString
}

func CapitalizeFirstChar(str string) string {
	if len(str) == 0 {
		return str // Return the original string if it's empty
	}
	firstChar := []rune(str)[0]
	upperFirstChar := unicode.ToUpper(firstChar)
	return string(upperFirstChar) + str[1:]
}

func RemoveLastChar(input string) string {
	if len(input) > 0 {
		return input[:len(input)-1]
	}
	return input
}

func GetFirstChar(input string) string {
	if len(input) > 0 {
		return string(input[0])
	}
	return ""
}

func ToKebabCase(input string) string {
	var builder strings.Builder
	for i, r := range input {
		if unicode.IsUpper(r) {
			if i > 0 {
				builder.WriteRune('-')
			}
			builder.WriteRune(unicode.ToLower(r))
		} else {
			builder.WriteRune(r)
		}
	}
	return builder.String()
}

func KebabCase(input string) string {
	segments := strings.Split(input, "/")
	var output string
	for _, segment := range segments {
		fc := GetFirstChar(segment)
		if fc != ":" && fc != "{" {
			output += ToKebabCase(segment) + "/"
		} else {
			output += segment + "/"
		}
	}
	output = RemoveLastChar(output)
	return output
}
