package hash

func Rot13(s string) uint64 {
	var h uint64
	for _, c := range s {
		h += uint64(c)
		h -= (h << 13) | (h >> 19)
	}
	return h
}

func Dummy(s string) uint64 {
	var h uint64
	for _, c := range s {
		h += uint64(c)
	}
	return h
}