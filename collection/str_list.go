// /////////////////////////////////////////////////////////////////////////////
// StringList

package collection

// /////////////////////////////////////////////////////////////////////////////
// StringList

// StringList is a list of string (slice)
type StringList []string

// Remove removes the string from StringList
func (sl *StringList) Remove(elem string) {
	widx := 0
	cpsl := *sl
	for idx, _elem := range cpsl {
		if _elem == elem {
			// ignore this elem by doing nothing
		} else {
			if idx != widx {
				cpsl[widx] = _elem
			}
			widx += 1
		}
	}

	*sl = cpsl[:widx]
}

// Append add the string to the end of StringList
func (sl *StringList) Append(elem string) {
	*sl = append(*sl, elem)
}

// Find get the index of string in StringList, returns -1 if not found
func (sl *StringList) Find(s string) int {
	for idx, elem := range *sl {
		if elem == s {
			return idx
		}
	}
	return -1
}
