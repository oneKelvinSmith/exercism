package robotname

import (
	"math"
	"math/rand"
	"strconv"
	"time"
)

// Robot represents a robotic entity.
type Robot struct {
	name string
}

// Name returns the identifier for a Robot.
func (robot *Robot) Name() string {
	if robot.isNamed() {
		return robot.name
	}

	robot.generateName()

	return robot.name
}

// Reset clears the Robot's identifier so that a new one can be generated.
func (robot *Robot) Reset() {
	robot.name = ""
}

func (robot *Robot) isNamed() bool {
	return robot.name != ""
}

func (robot *Robot) generateName() {
	const letterCount = 2
	const digitCount = 3

	rand.Seed(time.Now().UnixNano())

	robot.name = randomLetters(letterCount) + randomDigits(digitCount)
}

func randomLetters(size int) string {
	const asciiRange = 25
	const asciiOffset = 65

	letters := make([]rune, size)
	for index := 0; index < size; index++ {
		letters[index] = rune(rand.Intn(asciiRange) + asciiOffset)
	}

	return string(letters)
}

func randomDigits(size int) string {
	const base = 10

	digits := 0
	for index := 0; index < size; index++ {
		digits += int(math.Pow(base, float64(index))) * rand.Intn(base-1)
	}

	return strconv.Itoa(digits)
}
