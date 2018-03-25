package robotname

import (
	"math/rand"
	"strconv"
	"time"
)

var used = make(map[string]bool)

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
	robot.name = randomLetters(2) + randomDigits(3)

	robot.verifyName()
}

func (robot *Robot) verifyName() {
	if used[robot.name] {
		robot.generateName()
	} else {
		used[robot.name] = true
	}
}

func randomLetters(count int) string {
	const base = 26
	const asciiOffset = 65

	return generate(count, base, func(random int) string {
		return string(random + asciiOffset)
	})
}

func randomDigits(count int) string {
	const base = 10

	return generate(count, base, func(random int) string {
		return strconv.Itoa(random)
	})
}

func reseed() {
	rand.Seed(time.Now().UnixNano())
}

func generate(count, base int, mod func(random int) string) (result string) {
	reseed()

	for index := 0; index < count; index++ {
		result += mod(rand.Intn(base))
	}

	return
}
