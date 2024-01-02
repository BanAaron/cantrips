package cantrips

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var (
	title = cases.Title(language.English)
	upper = cases.Upper(language.English)
	lower = cases.Lower(language.English)
)

func ToTitle(input string) string {
	return title.String(input)
}

func ToUpper(input string) string {
	return upper.String(input)
}

func ToLower(input string) string {
	return lower.String(input)
}
