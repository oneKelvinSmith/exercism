package robotname

import (
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

func randomLetters(size int) (letters string) {
	const base = 26
	const asciiOffset = 65

	letters = ""
	for index := 0; index < size; index++ {
		letters += string(rand.Intn(base) + asciiOffset)
	}

	return
}

func randomDigits(size int) (digits string) {
	const base = 10

	for index := 0; index < size; index++ {
		digits += strconv.Itoa(rand.Intn(base))
	}

	return
}
