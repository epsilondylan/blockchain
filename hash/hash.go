package hash

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"sync/atomic"
	"math/rand"
	"time"
)

// Size The size of a SHA256 checksum in bytes.
const Size = 32

var (
	stop          int32
	tmpDifficulty int
)

// HashwithDifficulty ...
func HashwithDifficulty(data []byte, d int) (result [Size]byte, nonce int64) {
	rand.Seed(time.Now().UnixNano())
	tmpDifficulty = d
	for nonce = int64(rand.Intn(10000)); ; nonce++ {
		time.Sleep(5*time.Microsecond)
		if atomic.LoadInt32(&stop) == 1 {
			return result, 0
		}
		str := strconv.FormatInt(nonce, 10)
		b := append(data, []byte(str)...)
		result = sha256.Sum256(b)
		if difficulty(result, d) {
			return result, nonce
		}
	}
	return
}

func difficulty(hash [Size]byte, d int) bool {
	dn := d / 2
	sn := d % 2
	for i := 0; i < dn; i++ {
		if hash[i] != 0x00 {
			return false
		}
	}
	if sn != 0 {
		if hash[dn*2+1] > 0x0f {
			return false
		}
	}
	return true
}

// StopHash ...
func StopHash() bool {
	return atomic.CompareAndSwapInt32(&stop, 0, 1)
}

// StartHash ...
func StartHash() bool {
	return atomic.CompareAndSwapInt32(&stop, 1, 0)
}

// Verification to test if the data's hash is equal to a string
func Verification(data []byte, hash string) bool {
	var new [Size]byte
	new = sha256.Sum256(data)
	if !difficulty(new, tmpDifficulty) {
		return false
	}
	return hash == fmt.Sprintf("%x", new)
}
