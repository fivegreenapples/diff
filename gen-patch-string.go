// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package diff

// StringPatch is a description of how to convert one input string slice to a
// second string slice. It is simply a slice of changeT structs.
type StringPatch []StringChange

// A StringChange describes a single instance of a difference acting at a particular
// offset into the original string slice. It details the offset, the new items to be
// added at this offset, and the number of items that should be skipped in the original
// slice.
type StringChange struct {
	Offset int
	Add    []string
	Skip   int
}

// MakeStringPatch calculates the StringPatch from a to b
func MakeStringPatch(a, b []string) StringPatch {
	firstDiff, lcs := StringLCS(a, b)

	indexA := firstDiff
	indexB := firstDiff
	indexCommon := firstDiff
	patch := StringPatch{}

	for indexA < len(a) || indexB < len(b) {

		for indexA < len(a) && indexB < len(b) && a[indexA] == b[indexB] {
			indexA++
			indexB++
			indexCommon++
		}

		newChange := StringChange{
			Offset: indexA,
		}

		for indexA < len(a) && (indexCommon >= len(lcs) || a[indexA] != lcs[indexCommon]) {
			newChange.Skip++
			indexA++
		}
		for indexB < len(b) && (indexCommon >= len(lcs) || b[indexB] != lcs[indexCommon]) {
			newChange.Add = append(newChange.Add, b[indexB])
			indexB++
		}

		patch = append(patch, newChange)
	}
	return patch
}

// ApplyStringPatch creates a new slice given an input slice, a, and a StringPatch p.
func ApplyStringPatch(a []string, p StringPatch) []string {

	result := []string{}

	indexA := 0
	indexP := 0

	for indexP < len(p) || indexA < len(a) {
		if indexP < len(p) {
			currentChange := p[indexP]
			if currentChange.Offset == indexA {
				// Act on current change as we're at the right line number
				result = append(result, currentChange.Add...)
				indexA += currentChange.Skip
				indexP++
				continue
			} else if indexA >= len(a) {
				// this protects us from a duff patch where items in the patch
				// reference beyond the end of the original.
				break
			}
		}

		if indexA < len(a) {
			result = append(result, a[indexA])
			indexA++
		}
	}

	return result
}
