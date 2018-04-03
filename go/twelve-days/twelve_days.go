package twelve

import (
	"fmt"
)

const verseCount = 12

type verse struct {
	day  string
	gift string
}

func (v *verse) gifts(index int) (gifts string) {
	if index == 0 {
		return v.gift
	}

	gifts = v.gift

	seperator := ", "
	for index--; index >= 0; index-- {
		if index == 0 {
			seperator = ", and "
		}
		gifts += seperator + verses[index].gift
	}

	return
}

var verses = [verseCount]*verse{
	{"first", "a Partridge in a Pear Tree"},
	{"second", "two Turtle Doves"},
	{"third", "three French Hens"},
	{"fourth", "four Calling Birds"},
	{"fifth", "five Gold Rings"},
	{"sixth", "six Geese-a-Laying"},
	{"seventh", "seven Swans-a-Swimming"},
	{"eighth", "eight Maids-a-Milking"},
	{"ninth", "nine Ladies Dancing"},
	{"tenth", "ten Lords-a-Leaping"},
	{"eleventh", "eleven Pipers Piping"},
	{"twelfth", "twelve Drummers Drumming"},
}

// Song returns all twelve verses of the song.
func Song() (song string) {
	for number := 1; number <= verseCount; number++ {
		song += Verse(number) + "\n"
	}

	return
}

// Verse return the nth verse of the song.
func Verse(number int) (verse string) {
	index := number - 1
	verse = fmt.Sprintf(
		"On the %s day of Christmas my true love gave to me, %s.",
		verses[index].day, verses[index].gifts(index),
	)

	return
}
