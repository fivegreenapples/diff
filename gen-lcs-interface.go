// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package diff

// InterfaceLCS returns the longest common subsequence of two given interface{} slices,
// along with the offset at which the two slices start to differ. This offset
// makes patch creation a bit more efficient.
//
// More informaton about this lcs algorithm can be found on wikipedia:
// https://en.wikipedia.org/wiki/Longest_common_subsequence_problem
func InterfaceLCS(A, B []interface{}) (int, []interface{}) {
	if len(A) == 0 || len(B) == 0 {
		return 0, []interface{}{}
	}

	startIndex := findStartIndexInterface(A, B)
	endIndexA, endIndexB := findEndIndicesInterface(A, B, startIndex)

	// When A is a prefix of B
	// 123
	// 123456
	if startIndex >= endIndexA {
		return startIndex, A
	}

	// When B is a prefix of A
	// 123456
	// 123
	if startIndex >= endIndexB {
		return startIndex, B
	}

	partialLCS := doLCSInterface(A[startIndex:endIndexA], B[startIndex:endIndexB])

	lcs := make([]interface{}, startIndex)
	copy(lcs, A[:startIndex])
	lcs = append(lcs, partialLCS...)
	lcs = append(lcs, A[endIndexA:]...)
	return startIndex, lcs
}

func findStartIndexInterface(A, B []interface{}) int {
	start, i, j := 0, 0, 0
	for i < len(A) && j < len(B) && A[i] == B[j] {
		start = i + 1
		i++
		j++
	}
	return start
}

func findEndIndicesInterface(A, B []interface{}, startIndex int) (int, int) {
	endA, endB := len(A), len(B)
	for endA > startIndex && endB > startIndex && endA > 0 && endB > 0 && A[endA-1] == B[endB-1] {
		endA--
		endB--
	}
	return endA, endB
}

func doLCSInterface(A, B []interface{}) []interface{} {

	partialsLength := len(B) + 1
	partials := make([][]interface{}, partialsLength)

	for _, elA := range A {
		newPartials := make([][]interface{}, partialsLength)
		for iB, elB := range B {
			if elA == elB {
				newPartials[iB+1] = partials[iB]
				newPartials[iB+1] = append(newPartials[iB+1], elA)
			} else {
				if len(partials[iB+1]) > len(newPartials[iB]) {
					newPartials[iB+1] = partials[iB+1]
				} else {
					newPartials[iB+1] = newPartials[iB]
				}
			}
		}
		partials = newPartials
	}
	return partials[partialsLength-1]
}