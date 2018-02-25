package gigasecond

import (
	"os"
	"testing"
	"time"
)

const (
	fmtD  = "2006-01-02"
	fmtDT = "2006-01-02T15:04:05"
)

func TestAddGigasecond(t *testing.T) {
	for _, tc := range addCases {
		in := parse(tc.in, t)
		want := parse(tc.want, t)
		got := AddGigasecond(in)
		if !got.Equal(want) {
			t.Fatalf(`FAIL: %s
AddGigasecond(%s)
   = %s
want %s`, tc.description, in, got, want)
		}
		t.Log("PASS:", tc.description)
	}
	t.Log("Tested", len(addCases), "cases.")
}

func parse(s string, t *testing.T) time.Time {
	tt, err := time.Parse(fmtDT, s)
	if err != nil {
		tt, err = time.Parse(fmtD, s)
	}
	if err != nil {
		_, statErr := os.Stat("example_gen.go")
		if statErr == nil || os.Getenv("TRAVIS_GO_VERSION") > "" {
			t.Fatal(err)
		} else {
			t.Log(err)
			t.Skip("(This is not your fault, and is unexpected.  " +
				"Please file an issue at https://github.com/exercism/go.)")
		}
	}
	return tt
}

func BenchmarkAddGigasecond(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AddGigasecond(time.Time{})
	}
}
