// /////////////////////////////////////////////////////////////////////////////
// Uint16Set

package collection

// /////////////////////////////////////////////////////////////////////////////
// Uint16Set

// Uint16Set is a set of int
type Uint16Set map[uint16]struct{}

// Contains checks if Stringset contains the string
func (is Uint16Set) Contains(elem uint16) bool {
	_, ok := is[elem]
	return ok
}

// Add adds the string to Uint16Set
func (is Uint16Set) Add(elem uint16) {
	is[elem] = struct{}{}
}

// Remove removes the string from Uint16Set
func (is Uint16Set) Remove(elem uint16) {
	delete(is, elem)
}

// ToList convert Uint16Set to int slice
func (is Uint16Set) ToList() []uint16 {
	keys := make([]uint16, 0, len(is))
	for s := range is {
		keys = append(keys, s)
	}
	return keys
}
