package a51

// KeySetup setup register with given key and frame
func KeySetup(reg *register, key []uint64, frame uint64) {
	for i := bit(0); i < 64; i++ {
		reg.clockAll()
		keybit := (bit(key[i/8]) >> (i & 7)) & 1
		reg.insertBit(keybit)
	}

	for i := bit(0); i < 22; i++ {
		reg.clockAll()
		framebit := (bit(frame) >> i) & 1
		reg.insertBit(framebit)
	}

	for i := 0; i < 100; i++ {
		reg.clock()
	}
}

// GenKeystream generate 114 bits keystream for
// encryption and decryption
func GenKeystream(reg *register) ([]byte, []byte) {
	enc := make([]byte, 15)
	dec := make([]byte, 15)

	for i := bit(0); i < 114; i++ {
		reg.clock()
		enc[i/8] |= byte(reg.getBit() << (7 - (i & 7)))
	}

	for i := bit(0); i < 114; i++ {
		reg.clock()
		dec[i/8] |= byte(reg.getBit() << (7 - (i & 7)))
	}

	return enc, dec
}
