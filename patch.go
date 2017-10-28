package diff

//go:generate genny -in=$GOFILE -out=gen-patch-string.go gen "GenericType=string PatchT=StringPatch ChangeT=StringChange GenericTwo=String"
//go:generate genny -in=$GOFILE -out=gen-patch-int.go gen "GenericType=int PatchT=IntPatch ChangeT=IntChange GenericTwo=Int"
//go:generate genny -in=$GOFILE -out=gen-patch-bool.go gen "GenericType=bool PatchT=BoolPatch ChangeT=BoolChange GenericTwo=Bool"
//go:generate genny -in=$GOFILE -out=gen-patch-interface.go gen "GenericType=interface{} PatchT=InterfacePatch ChangeT=InterfaceChange GenericTwo=Interface"

// PatchT is a description of how to convert one input GenericType slice to a
// second GenericType slice. It is simply a slice of changeT structs.
type PatchT []ChangeT

// A ChangeT describes a single instance of a difference acting at a particular
// offset into the original GenericType slice. It details the offset, the new items to be
// added at this offset, and the number of items that should be skipped in the original
// slice.
type ChangeT struct {
	Offset int
	Add    []GenericType
	Skip   int
}

// MakeGenericTwoPatch calculates the PatchT from a to b
func MakeGenericTwoPatch(a, b []GenericType) PatchT {
	firstDiff, lcs := GenericTwoLCS(a, b)

	indexA := firstDiff
	indexB := firstDiff
	indexCommon := firstDiff
	patch := PatchT{}

	for indexA < len(a) || indexB < len(b) {

		for indexA < len(a) && indexB < len(b) && a[indexA] == b[indexB] {
			indexA++
			indexB++
			indexCommon++
		}

		newChange := ChangeT{
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

// ApplyGenericTwoPatch creates a new slice given an input slice, a, and a PatchT p.
func ApplyGenericTwoPatch(a []GenericType, p PatchT) []GenericType {

	result := []GenericType{}

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
