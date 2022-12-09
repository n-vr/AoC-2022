// Package shiftregister implements a shift register with a maximum amount of values.
package shiftregister

// ShiftRegister is a shift register with maximum amount of values.
type ShiftRegister struct {
	data    []rune
	maxSize int
}

// Create a new [shiftRegister] with maximum size.
func New(maxSize int) *ShiftRegister {
	return &ShiftRegister{
		data:    make([]rune, 0, maxSize),
		maxSize: maxSize,
	}
}

// Shift a new value in. If there are already the maximum of values inside,
// the last value added will shift out.
func (s *ShiftRegister) ShiftIn(in rune) {
	if len(s.data) == s.maxSize {
		// Remove 1st element.
		s.data = s.data[1:]
	}

	s.data = append(s.data, in)
}

// Return true if all values in the [shiftRegister] are unique.
func (s *ShiftRegister) ValuesUnique() bool {
	var commons uint64

	for _, item := range s.data {
		if commons&(1<<(item-'a')) != 0 {
			return false
		}
		commons |= (1 << (item - 'a'))
	}

	return true
}

// Returns true if the [shiftRegister] is full and all values are unique.
func (s *ShiftRegister) FullAndValuesUnique() bool {
	if len(s.data) < s.maxSize {
		return false
	}
	return s.ValuesUnique()
}
