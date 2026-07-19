package main

// Proquint encoding per Daniel Wilkerson, "A Proposal for Proquints:
// Readable, Spellable, and Pronounceable" (https://arxiv.org/html/0901.4016).
//
// A proquint is a sequence of consonant-vowel quintuplets where each quint
// encodes 16 bits: four bits each from three consonants and two bits each
// from two vowels, in the pattern C V C V C.
//
//	consonants (16): b d f g h j k l m n p r s t v z
//	vowels      (4): a i o u
//
// A 16-bit unsigned value n encodes to:
//
//	c1 = CONS[(n >> 12) & 0xF]
//	v1 = VOWS[(n >> 10) & 0x3]
//	c2 = CONS[(n >>  6) & 0xF]
//	v2 = VOWS[(n >>  4) & 0x3]
//	c3 = CONS[ n        & 0xF]
//
// proquint-1 is one quint (5 chars, 16-bit space, 65,536 values).
// proquint-2 is two quints joined by '-' (10 chars + hyphen, 32-bit space,
// 4.29 billion values).
//
// This repo uses proquint as the handle encoding because it is short,
// pronounceable, and trivially derivable from any 16- or 32-bit hash output
// without a curated wordlist or central registry.
//
// New coordination artifact IDs stay short and collision-checkable without
// relying on central integer or timestamp allocation.
const (
	proquintCons = "bdfghjklmnprstvz"
	proquintVows = "aiou"
)

// uint16ToProquint encodes one 16-bit value as a single CVCVC proquint. The
// bit layout follows Wilkerson's original paper exactly.
func uint16ToProquint(n uint16) string {
	buf := []byte{
		proquintCons[(n>>12)&0x0f],
		proquintVows[(n>>10)&0x03],
		proquintCons[(n>>6)&0x0f],
		proquintVows[(n>>4)&0x03],
		proquintCons[n&0x0f],
	}
	return string(buf)
}

// proquint1FromBytes converts the first two bytes into one proquint. The
// caller is responsible for passing at least two bytes; mint() always passes
// SHA-256 output, so that precondition is controlled locally.
func proquint1FromBytes(b []byte) string {
	n := uint16(b[0])<<8 | uint16(b[1])
	return uint16ToProquint(n)
}

// proquint2FromBytes converts the first four bytes into two hyphen-separated
// proquints. It is available for future corpus-size growth beyond the
// proquint-1 comfort range.
func proquint2FromBytes(b []byte) string {
	a := uint16(b[0])<<8 | uint16(b[1])
	c := uint16(b[2])<<8 | uint16(b[3])
	return uint16ToProquint(a) + "-" + uint16ToProquint(c)
}
