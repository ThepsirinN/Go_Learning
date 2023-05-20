package greeting

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "barko"
	want := regexp.MustCompile(name)
	msg, err := Hello("barko")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("barko") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
