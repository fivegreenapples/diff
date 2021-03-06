// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package diff

// InterfacePatch is a description of how to convert one input interface{} slice to a
// second interface{} slice. It is simply a slice of changeT structs.
type InterfacePatch []InterfaceChange

// A InterfaceChange describes a single instance of a difference acting at a particular
// offset into the original interface{} slice. It details the offset, the new items to be
// added at this offset, and the number of items that should be skipped in the original
// slice.
type InterfaceChange struct {
	Offset int
	Add    []interface{}
	Skip   int
}

// MakeInterfacePatch calculates the InterfacePatch from a to b
func MakeInterfacePatch(a, b []interface{}) InterfacePatch {
	firstDiff, lcs := InterfaceLCS(a, b)

	indexA := firstDiff
	indexB := firstDiff
	indexCommon := firstDiff
	patch := InterfacePatch{}

	for indexA < len(a) || indexB < len(b) {

		for indexA < len(a) && indexB < len(b) && a[indexA] == b[indexB] {
			indexA++
			indexB++
			indexCommon++
		}

		newChange := InterfaceChange{
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

// ApplyInterfacePatch creates a new slice given an input slice, a, and a InterfacePatch p.
func ApplyInterfacePatch(a []interface{}, p InterfacePatch) []interface{} {

	result := []interface{}{}

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
