package diff

import "github.com/cheekybits/genny/generic"

//go:generate genny -in=$GOFILE -out=gen-lcs-string.go gen "GenericType=string GenericTwo=String"
//go:generate genny -in=$GOFILE -out=gen-lcs-int.go gen "GenericType=int GenericTwo=Int"
//go:generate genny -in=$GOFILE -out=gen-lcs-bool.go gen "GenericType=bool GenericTwo=Bool"
//go:generate genny -in=$GOFILE -out=gen-lcs-interface.go gen "GenericType=interface{} GenericTwo=Interface"

// GenericType defines a place holder type for the actions in this package.
type GenericType generic.Type

// GenericTwo defines a place holder type for the actions in this package.
type GenericTwo generic.Type

// GenericTwoLCS returns the longest common subsequence of two given GenericType slices,
// along with the offset at which the two slices start to differ. This offset
// makes patch creation a bit more efficient.
//
// More informaton about this lcs algorithm can be found on wikipedia:
// https://en.wikipedia.org/wiki/Longest_common_subsequence_problem
func GenericTwoLCS(A, B []GenericType) (int, []GenericType) {
	if len(A) == 0 || len(B) == 0 {
		return 0, []GenericType{}
	}

	startIndex := findStartIndexGenericTwo(A, B)
	endIndexA, endIndexB := findEndIndicesGenericTwo(A, B, startIndex)

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

	partialLCS := doLCSGenericTwo(A[startIndex:endIndexA], B[startIndex:endIndexB])

	lcs := make([]GenericType, startIndex)
	copy(lcs, A[:startIndex])
	lcs = append(lcs, partialLCS...)
	lcs = append(lcs, A[endIndexA:]...)
	return startIndex, lcs
}

func findStartIndexGenericTwo(A, B []GenericType) int {
	start, i, j := 0, 0, 0
	for i < len(A) && j < len(B) && A[i] == B[j] {
		start = i + 1
		i++
		j++
	}
	return start
}

func findEndIndicesGenericTwo(A, B []GenericType, startIndex int) (int, int) {
	endA, endB := len(A), len(B)
	for endA > startIndex && endB > startIndex && endA > 0 && endB > 0 && A[endA-1] == B[endB-1] {
		endA--
		endB--
	}
	return endA, endB
}

func doLCSGenericTwo(A, B []GenericType) []GenericType {

	partialsLength := len(B) + 1
	partials := make([][]GenericType, partialsLength)

	for _, elA := range A {
		newPartials := make([][]GenericType, partialsLength)
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
