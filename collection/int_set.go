// /////////////////////////////////////////////////////////////////////////////
// IntSet

package collection

// /////////////////////////////////////////////////////////////////////////////
// IntSet

// IntSet is a set of int
type IntSet map[int]struct{}

// Contains checks if Stringset contains the string
func (is IntSet) Contains(elem int) bool {
	_, ok := is[elem]
	return ok
}

// Add adds the string to IntSet
func (is IntSet) Add(elem int) {
	is[elem] = struct{}{}
}

// Remove removes the string from IntSet
func (is IntSet) Remove(elem int) {
	delete(is, elem)
}

// ToList convert IntSet to int slice
func (is IntSet) ToList() []int {
	keys := make([]int, 0, len(is))
	for s := range is {
		keys = append(keys, s)
	}
	return keys
}
