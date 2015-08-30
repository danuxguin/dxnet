package common

const size = 64

type bits uint64

// BitSet is a set of bits that can be set, cleared and queried.
type BitSet []bits

func NewBitSet(n ...uint64) BitSet {

	var arr []bits
	for _, a := range n {
		arr = append(arr, bits(a))
	}

	return BitSet(arr)
}

// Set ensures that the given bit is set in the BitSet.
func (s *BitSet) Set(i uint) {
	if len(*s) < int(i/size+1) {
		r := make([]bits, i/size+1)
		copy(r, *s)
		*s = r
	}
	(*s)[i/size] |= 1 << (i % size)
}

// Clear ensures that the given bit is cleared (not set) in the BitSet.
func (s *BitSet) Clear(i uint) {
	if len(*s) >= int(i/size+1) {
		(*s)[i/size] &^= 1 << (i % size)
	}
}

// IsSet returns true if the given bit is set, false if it is cleared.
func (s *BitSet) IsSet(i uint) bool {
	return (*s)[i/size]&(1<<(i%size)) != 0
}
