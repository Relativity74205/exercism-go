package twelve

import (
	"fmt"
	"strings"
)

var items = []string{"Drummers Drumming", "Pipers Piping", "Lords-a-Leaping", "Ladies Dancing", "Maids-a-Milking", "Swans-a-Swimming", "Geese-a-Laying", "Gold Rings", "Calling Birds", "French Hens", "Turtle Doves", "Partridge in a Pear Tree."}

func num2Word(n int) string {
	switch n {
	case 1:
		return "a"
	case 2:
		return "two"
	case 3:
		return "three"
	case 4:
		return "four"
	case 5:
		return "five"
	case 6:
		return "six"
	case 7:
		return "seven"
	case 8:
		return "eight"
	case 9:
		return "nine"
	case 10:
		return "ten"
	case 11:
		return "eleven"
	case 12:
		return "twelve"
	default:
		return ""
	}
}

func num2Day(n int) string {
	switch n {
	case 1:
		return "first"
	case 2:
		return "second"
	case 3:
		return "third"
	case 4:
		return "fourth"
	case 5:
		return "fifth"
	case 6:
		return "sixth"
	case 7:
		return "seventh"
	case 8:
		return "eighth"
	case 9:
		return "ninth"
	case 10:
		return "tenth"
	case 11:
		return "eleventh"
	case 12:
		return "twelfth"
	default:
		return ""
	}
}

func Verse(dayNumber int) string {
	var itemsWithCount []string
	for i := len(items) - dayNumber; i < len(items); i++ {
		count := len(items) - i
		itemsWithCount = append(itemsWithCount, fmt.Sprintf("%s %s", num2Word(count), items[i]))
	}
	enumeration := ""
	if dayNumber == 1 {
		enumeration = itemsWithCount[0]
	} else {
		enumeration = fmt.Sprintf("%s, and %s", strings.Join(itemsWithCount[:dayNumber-1], ", "), itemsWithCount[dayNumber-1])
	}

	return fmt.Sprintf("On the %s day of Christmas my true love gave to me: %s", num2Day(dayNumber), enumeration)
}

func Song() string {
	var verses []string
	for dayNumber := 1; dayNumber <= 12; dayNumber++ {
		verses = append(verses, Verse(dayNumber))
	}
	return strings.Join(verses, "\n")
}
