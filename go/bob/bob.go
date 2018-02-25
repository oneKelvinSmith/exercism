package bob

import (
	"fmt"
	"strings"
)

// Hey greets with an appropriate remark.
func Hey(remark string) string {
	fmt.Printf("Remark: %v\n", remark)

	shouting := func() bool {
		return strings.ToUpper(remark) == remark
	}

	questioning := func() bool {
		return strings.HasSuffix(remark, "?")
	}

	switch {
	case shouting() && questioning():
		return "Calm down, I know what I'm doing!"
	case shouting():
		return "Whoa, chill out!"
	case questioning():
		return "Sure."
	default:
		return "Whatever."
	}
}
