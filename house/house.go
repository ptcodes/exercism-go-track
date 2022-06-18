package house

import (
	"strings"
)

var characters = []string{
	"the house that Jack built.",
	"the malt\nthat lay in",
	"the rat\nthat ate",
	"the cat\nthat killed",
	"the dog\nthat worried",
	"the cow with the crumpled horn\nthat tossed",
	"the maiden all forlorn\nthat milked",
	"the man all tattered and torn\nthat kissed",
	"the priest all shaven and shorn\nthat married",
	"the rooster that crowed in the morn\nthat woke",
	"the farmer sowing his corn\nthat kept",
	"the horse and the hound and the horn\nthat belonged to",
}

func Verse(v int) string {
	var lines []string
	for i := v - 1; i >= 0; i-- {
		lines = append(lines, characters[i])
	}
	return "This is " + strings.Join(lines, " ")
}

func Song() string {
	var lines []string
	for i := range characters {
		lines = append(lines, Verse(i+1))
	}
	return strings.Join(lines, "\n\n")

}
