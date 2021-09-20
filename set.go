package main

type HashsetMap map[string]struct{}

type Hashset struct {
	set HashsetMap
}

// return false if value is already in map
func (h *Hashset) Add(value string) bool {
	if _, ok := h.set[value]; ok {
		return false
	}

	h.set[value] = struct{}{}
	return true
}

func New(values ...string) *Hashset {
	hashset := Hashset{make(HashsetMap)}

	for _, value := range values {
		hashset.set[value] = struct{}{}
	}

	return &hashset
}
