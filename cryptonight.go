package cryptonight

/*
#cgo CFLAGS: -maes
#cgo LDFLAGS:
#include "hash-ops.h"
*/
import "C"
import (
	"encoding/binary"
	"unsafe"
)

// Setting the following environment variable before running 'make
// geth' was required to avoid an aes-related compile error:
// CGO_CFLAGS_ALLOW=-maes
//
// Some other possibly useful CFLAGS:
// -I. -Ofast -fuse-linker-plugin -funroll-loops -fvariable-expansion-in-unroller -ftree-loop-if-convert-stores -fmerge-all-constants -fbranch-target-load-optimize2 -fsched2-use-superblocks -falign-loops=16 -falign-functions=16 -falign-jumps=16 -falign-labels=16 -Wno-pointer-sign -Wno-pointer-to-int-cast -maes -march=native -Wl,--stack,10485760

// Direct wrapper around cryptonight's cn_slow_hash. You should
// probably not use this function, and instead use
// HashVariant{1,2,4}ForEthereumHeader below.
//
// Takes input hash material with length at least 43 bytes, and
// returns the 32 byte hash. The number 43 is due to an invariant in
// the C implementation of cn_slow_hash. If you pass less than 43
// bytes, the C code will halt with a runtime error and a reasonably
// descriptive message.
//
// Variant must be 1, 2, or 4. If variant is less then 4, then set
// block_height to zero. In public discourse, a lot of folks use
// alternative names for the different versions of Cryptonight:
// variant 1 is aka Cryptonight v7, variant 2 is aka Cryptonight v8,
// and variant 4 is aka CryptonightR.
func hashCryptonight(input []byte, variant int, block_height uint64) []byte {
	result := make([]byte, 32)
	input_ptr := unsafe.Pointer(&input[0])
	output_ptr := unsafe.Pointer(&result[0])
	C.cn_slow_hash(input_ptr, C.size_t(len(input)), (*C.char)(output_ptr), (C.int)(variant), 0 /*prehashed*/, (C.uint64_t)(block_height))
	return result
}

// Uses the above hash function in a specific way for given ethereum
// header and ethereum nonce (which is 8 bytes, vs more typical 4
// bytes in other protocols).
//
// Monero and other projects that use the cryptonight algorithm have
// historically interpreted the resulting hash bytes as if they are
// little-endian before comparing to the target difficulty, whereas
// Ethereum has always interpreted hash bytes as big endian (mostly
// due to the use of the native golang big integer class). Since we
// want our Ethereum code's usage of cryptonight hashes to agree with
// existing implementations in CPU and GPU mining software, we reverse
// the `result` bytes before returning them in this function (but keep
// `digest` untouched). That way, existing code elsewhere in the
// Ethereum codebase (which interprets everything as big endian) will
// end up agreeing with non-Ethereum implementations without any
// changes.
func hashCryptonightForEthereumHeader(block_header_hash []byte, nonce uint64, variant int, block_height uint64) ([]byte, []byte) {
	// Note: this blob format intentionally looks hacky. We're trying
	// to match the length and some of the byte offsets that monero
	// uses, e.g. its major/minor versions and nonce, so that existing
	// monero-like mining software implementations (both cpu and gpu)
	// remain compatible with fewer changes.
	blob := make([]byte, 76)
	for i := 0; i < len(blob); i++ {
		// Initialize to 0x77 for all bytes.
		blob[i] = 119
	}
	var blen int = 0
	// Major version. It's not really necessary to set this (that is
	// until main net goes live, at which point we'll need to keep it
	// consistent), but just in case there's existing mining software
	// out there which makes use of it, we keep the value in sync with
	// Monero's major version (the major version which corresponds to a
	// particular variant of Cryptonight). You can see a list of Monero
	// major versions (hard forks) in Monero repository's
	// src/cryptonote_core/blockchain.cpp file.
	if variant == 1 {
		blob[blen] = 7
	} else if variant == 2 {
		blob[blen] = 8
	} else {
		blob[blen] = 10
	}
	blen++
	// And minor version. Pretty sure no one uses this anywhere so we
	// don't bother setting it.
	blob[blen] = 0
	blen++
	// 5 byte timestamp
	for i := 0; i < 5; i++ {
		// Initialize to zeroes for timestamp.
		blob[blen] = 0
		blen++
	}
	copy(blob[blen:], block_header_hash)
	blen += 32
	binary.LittleEndian.PutUint64(blob[blen:], nonce)
	digest := hashCryptonight(blob, variant, block_height)

	// Interpret hash result as little endian.
	result := make([]byte, len(digest))
	for i, b := range digest {
		result[len(digest) - i - 1] = b
	}

	return digest, result
}

// Similar to hashimoto, HashVariant{1,2,4}ForEthereumHeader accepts
// 32 byte block header hash and 8 byte nonce, then returns 32 byte
// digest and 32 byte result. Variant 4 (aka CryptonightR) also needs
// to know the block height.
func HashVariant1ForEthereumHeader(block_header_hash []byte, nonce uint64) ([]byte, []byte) {
	return hashCryptonightForEthereumHeader(block_header_hash, nonce, 1, 0 /*block_height*/)
}

func HashVariant2ForEthereumHeader(block_header_hash []byte, nonce uint64) ([]byte, []byte) {
	return hashCryptonightForEthereumHeader(block_header_hash, nonce, 2, 0 /*block_height*/)
}

func HashVariant4ForEthereumHeader(block_header_hash []byte, nonce uint64, block_height uint64) ([]byte, []byte) {
	return hashCryptonightForEthereumHeader(block_header_hash, nonce, 4, block_height)
}
