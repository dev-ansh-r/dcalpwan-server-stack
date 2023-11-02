package cleanup

// ComputeSetComplement returns the complement of the first parameter set relative to the second parameter set.
func ComputeSetComplement(isSet, localSet map[string]struct{}) (complement map[string]struct{}) {
	complement = make(map[string]struct{})
	for ids := range localSet {
		if _, ok := isSet[ids]; ok {
			continue
		}
		complement[ids] = struct{}{}
	}
	return complement
}
