package diff

import (
	"strings"
	"testing"
)

func TestOne(t *testing.T) {

	a := "ACDFH"
	b := "012ABCFH"

	aa := strings.Split(a, "")
	bb := strings.Split(b, "")

	delta := MakeStringPatch(aa, bb)

	newB := strings.Join(ApplyStringPatch(aa, delta), "")

	if b != newB {
		t.Errorf("Failed to apply delta correctly. Got %s expected %s", newB, b)
	}

}
