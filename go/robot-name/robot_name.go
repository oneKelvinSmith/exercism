package robotname

import (
	"math/rand"
	"strconv"
	"time"
)

// Robot represents a robotic entity.
type Robot struct {
	name    string
	isNamed bool
}

// Name returns the identifier for a Robot.
func (robot *Robot) Name() string {
	if robot.isNamed {
		return robot.name
	}

	robot.name = generateName()
	robot.isNamed = true

	return robot.name
}

// Reset clears the Robot's identifier so that a new one can be generated.
func (robot *Robot) Reset() {
	robot.isNamed = false
}

func generateName() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return randomLetter(r) +
		randomLetter(r) +
		randomDigit(r) +
		randomDigit(r) +
		randomDigit(r)
}

func randomLetter(generator *rand.Rand) string {
	const letterMax = 25
	const letterOffset = 65

	return string(generator.Intn(letterMax) + letterOffset)
}

func randomDigit(generator *rand.Rand) string {
	const numberMax = 9

	return strconv.Itoa(generator.Intn(numberMax))
}
