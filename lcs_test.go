package diff

import "testing"
import "strings"

type testData struct {
	name           string
	a              string
	b              string
	expectedOffset int
	expectedLCS    string
}

func testGeneric(t *testing.T, name string, inputA, inputB string, expectedOffset int, expectedLCS string) {

	a := strings.Split(inputA, "")
	b := strings.Split(inputB, "")
	expected := strings.Split(expectedLCS, "")

	offset, lcs := StringLCS(a, b)

	if offset != expectedOffset {
		t.Errorf("%s: offset expected to be %d but was %d", name, expectedOffset, offset)
	}

	comparableLCS := lcs
	if len(expected) < len(lcs) {
		comparableLCS = lcs[0:len(expected)]
	}

	for i, val := range comparableLCS {
		if val != expected[i] {
			t.Errorf("%s: value at index %d of LCS does not match expected. Got %s expected %s", name, i, val, expected[i])
		}
	}
	if len(lcs) != len(expectedLCS) {
		t.Errorf("%s: LCS was different length to expected. Got %d items, expected %d", name, len(lcs), len(expectedLCS))
	}

}

func TestAll(t *testing.T) {

	inputs := []testData{
		testData{"Empty A", "", "ABCD", 0, ""},
		testData{"Empty B", "ABCD", "", 0, ""},
		testData{"A == B", "ABCD", "ABCD", 4, "ABCD"},
		testData{"A is prefix of B", "ABCD", "ABCDEFGH", 4, "ABCD"},
		testData{"B is prefix of A", "ABCDEFGH", "ABCD", 4, "ABCD"},
		testData{"A is suffix of B", "EFGH", "ABCDEFGH", 0, "EFGH"},
		testData{"B is suffix of A", "ABCDEFGH", "EFGH", 0, "EFGH"},
		testData{"A is infix of B", "CDEF", "ABCDEFGH", 0, "CDEF"},
		testData{"B is infix of A", "ABCDEFGH", "CDEF", 0, "CDEF"},
		testData{"A is subset of B", "ABEFGKLM", "ABCDEFGHIJKLMNOP", 2, "ABEFGKLM"},
		testData{"Other mix", "BCDEF0123456ABC234", "ABC0123ABC0123", 0, "BC0123ABC23"},
	}

	for _, test := range inputs {
		testGeneric(t, test.name, test.a, test.b, test.expectedOffset, test.expectedLCS)
	}

}
