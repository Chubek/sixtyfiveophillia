package logical

import (
	"sixtyfiveophillia/memory"
)

func LogicalOR(a, b memory.MemoryBlock) (res memory.Octet) {
	var result memory.Octet

	for i := 0; i <= 7; i++ {
		result[i] = a.Container[i] | b.Container[i]
	}

	return result
}

func LogicalAND(a, b memory.MemoryBlock) (res memory.Octet) {
	var result memory.Octet

	for i := 0; i <= 7; i++ {
		result[i] = a.Container[i] & b.Container[i]
	}

	return result
}

func LogicalXOR(a, b memory.MemoryBlock) (res memory.Octet) {
	var result memory.Octet

	for i := 0; i <= 7; i++ {
		result[i] = a.Container[i] ^ b.Container[i]
	}

	return result
}
