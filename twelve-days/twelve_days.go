package twelve

import (
	"fmt"
	"strings"
)

var n = []string{
	"and a",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
	"ten",
	"eleven",
	"twelve",
}

var nth = []string{
	"first",
	"second",
	"third",
	"fourth",
	"fifth",
	"sixth",
	"seventh",
	"eighth",
	"ninth",
	"tenth",
	"eleventh",
	"twelfth",
}

var gifts = []string{
	"Partridge in a Pear Tree",
	"Turtle Doves",
	"French Hens",
	"Calling Birds",
	"Gold Rings",
	"Geese-a-Laying",
	"Swans-a-Swimming",
	"Maids-a-Milking",
	"Ladies Dancing",
	"Lords-a-Leaping",
	"Pipers Piping",
	"Drummers Drumming",
}

func tail(max int) string {
	var lines []string
	for i := max - 1; i >= 0; i-- {
		prefix := n[i]
		if max == 1 {
			prefix = "a"
		}
		lines = append(lines, fmt.Sprintf("%s %s", prefix, gifts[i]))
	}
	return strings.Join(lines, ", ")
}

func Verse(i int) string {
	return fmt.Sprintf("On the %s day of Christmas my true love gave to me: %s.", nth[i-1], tail(i))
}

func Song() string {
	var lines []string
	for i := range gifts {
		lines = append(lines, Verse(i+1))
	}
	return strings.Join(lines, "\n")
}
