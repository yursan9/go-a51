package a51

const (
	rxMask bit = 0x07FFFF // 19 bits
	ryMask bit = 0x3FFFFF // 22 bits
	rzMask bit = 0x7FFFFF // 23 bits

	rxMid bit = 0x000100 // bit 8
	ryMid bit = 0x000400 // bit 10
	rzMid bit = 0x000400 // bit 10

	rxTaps bit = 0x072000 // bits 18, 17, 16, 13
	ryTaps bit = 0x300000 // bits 21, 20
	rzTaps bit = 0x700080 // bits 22, 21, 20, 7

	rxOut bit = 0x040000 // bit 18
	ryOut bit = 0x200000 // bit 21
	rzOut bit = 0x400000 // bit 22
)

// register struct simulate three register
// used to generate keystream
type register struct {
	x bit
	y bit
	z bit
}

// NewRegister return new Register struct
func NewRegister() *register {
	return &register{
		0,
		0,
		0,
	}
}

// majority look at the middle bit of register
// and return the majority bit. E.g (1,0,1)=>1 and (0,0,1)=>0
func (r *register) majority() bit {
	var sum bit
	sum = parity(r.x&rxMid) + parity(r.y&ryMid) + parity(r.z&rzMid)
	if sum >= 2 {
		return 1
	}
	return 0
}

// clock shift two or three register based on majority bit
func (r *register) clock() {
	m := r.majority()

	if (r.x&rxMid != 0) == (m > 0) {
		r.x = clockOne(r.x, rxMask, rxTaps)
	}

	if (r.y&ryMid != 0) == (m > 0) {
		r.y = clockOne(r.y, ryMask, ryTaps)
	}

	if (r.z&rzMid != 0) == (m > 0) {
		r.z = clockOne(r.z, rzMask, rzTaps)
	}
}

// clockAll shift all register ignoring the majority bit of middle bits
func (r *register) clockAll() {
	r.x = clockOne(r.x, rxMask, rxTaps)
	r.y = clockOne(r.y, ryMask, ryTaps)
	r.z = clockOne(r.z, rzMask, rzTaps)
}

// insertBit insert bit into all register
func (r *register) insertBit(x bit) {
	r.x ^= x
	r.y ^= x
	r.z ^= x
}

// getBit generate output bit from 3 end bits of register, XORed
func (r *register) getBit() bit {
	return parity(r.x&rxOut) ^ parity(r.y&ryOut) ^ parity(r.z&rzOut)
}
