package twelve

import (
	"fmt"
)

const verseCount = 12

var components = [verseCount][3]string{
	{"first", "a", "Partridge in a Pear Tree"},
	{"second", "two", "Turtle Doves"},
	{"third", "three", "French Hens"},
	{"fourth", "four", "Calling Birds"},
	{"fifth", "five", "Gold Rings"},
	{"sixth", "six", "Geese-a-Laying"},
	{"seventh", "seven", "Swans-a-Swimming"},
	{"eighth", "eight", "Maids-a-Milking"},
	{"ninth", "nine", "Ladies Dancing"},
	{"tenth", "ten", "Lords-a-Leaping"},
	{"eleventh", "eleven", "Pipers Piping"},
	{"twelfth", "twelve", "Drummers Drumming"},
}

func Song() (song string) {
	for number := 1; number <= verseCount; number++ {
		song += Verse(number) + "\n"
	}
	return
}

func Verse(number int) (verse string) {
	verse = fmt.Sprintf(
		"On the %s day of Christmas my true love gave to me, %s.",
		day(number-1), accumulatedGifts(number-1),
	)

	return
}

func day(index int) string {
	return components[index][0]
}

func accumulatedGifts(index int) (gifts string) {
	gifts = components[index][1] + " " + components[index][2]

	if index == 0 {
		return
	}

	seperator := ", "
	if index == 1 {
		seperator = ", and "
	}

	index--

	gifts += seperator + accumulatedGifts(index)

	return
}
