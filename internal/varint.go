package internal

import (
	"encoding/binary"
	"math/bits"
)

const MaxVarintSize = binary.MaxVarintLen64

func VarintSize(v uint64) int {
	return 1 + (bits.Len64(v)-1)/7
}

func VarintSizeInBytes(b []byte) int {
	for i, bt := range b {
		if bt < 0x80 {
			if i >= MaxVarintSize || i == MaxVarintSize-1 && bt > 1 {
				return -(i + 1) // v overflow
			}
			return i + 1
		}
	}
	return 0 // b is too small
}

func VarintSizeInLastBytes(b []byte) int {
	if len(b) == 0 {
		return 0 // b is too small
	}

	lb := b[len(b)-1]
	if lb > 0x7f {
		return 0 // b is too small
	}

	i := len(b) - 2
	for ; i >= 0; i-- {
		if b[i] < 0x80 {
			break
		}
	}
	n := len(b) - i - 1
	if n > MaxVarintSize || n == MaxVarintSize && lb > 1 {
		return -n // v overflow
	}
	return n
}

func PutVarint(b []byte, v uint64) int {
	return binary.PutUvarint(b, v)
}

func ReadVarint(b []byte) (uint64, int) {
	return binary.Uvarint(b)
}

func ReadVarintLast(b []byte) (uint64, int) {
	if len(b) == 0 {
		return 0, 0 // b is too small
	}

	lb := b[len(b)-1]
	if lb > 0x7f {
		return 0, 0 // b is too small
	}

	var x = uint64(lb)
	i := len(b) - 2
	for ; i >= 0; i-- {
		if b[i] < 0x80 {
			break
		}
		x = x<<7 | uint64(b[i]&0x7f)
	}
	n := len(b) - i - 1
	if n > MaxVarintSize || n == MaxVarintSize && lb > 1 {
		return 0, -n // v overflow
	}
	return x, n
}
