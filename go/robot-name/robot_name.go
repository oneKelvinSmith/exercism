package robotname

import (
	"fmt"
	"math/rand"
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
	robot.name = fmt.Sprintf("%c%c%03d", letter(), letter(), digits())

	if used[robot.name] {
		robot.generateName()
	} else {
		used[robot.name] = true
	}
}

func letter() int {
	const alphabetSize = 26
	const asciiOffset = 65

	reseed()

	return rand.Intn(alphabetSize) + asciiOffset
}

func digits() int {
	const upperLimit = 1000

	reseed()

	return rand.Intn(upperLimit)
}

func reseed() {
	rand.Seed(time.Now().UnixNano())
}
