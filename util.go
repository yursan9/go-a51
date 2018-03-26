package a51

// parity use to search the parity bit of bits
func parity(x bit) bit {
	x ^= x >> 32
	x ^= x >> 16
	x ^= x >> 8
	x ^= x >> 4
	x ^= x >> 2
	x ^= x >> 1

	return x & 1
}

// clockOne shift one register
func clockOne(reg, mask, taps bit) bit {
	t := reg & taps
	reg = (reg << 1) & mask
	reg |= parity(t)
	return reg
}
