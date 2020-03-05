package hash

import (
	"math/rand"
	"strconv"
	"testing"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func genString(n uint64) string {
	s := make([]byte, 0, n)
	for i := uint64(0); i < n; i++ {
		s = append(s, letters[rand.Intn(len(letters))])
	}
	return string(s)
}

func testCollisions(t *testing.T, hashFunc func(string) uint64) {
	tests := []uint64{
		0,
		10,
		100,
		1000,
		10000,
		100000,
	}
	for _, n := range tests {
		t.Run(strconv.FormatInt(int64(n), 10), func(t *testing.T) {
			var collisions uint64
			m := make(map[uint64]struct{})
			for i := uint64(0); i < n; i++ {
				h := hashFunc(genString(64))
				if _, found := m[h]; found {
					collisions++
				} else {
					m[h] = struct{}{}
				}
			}
			t.Logf("Found %d collisions", collisions)
		})
	}
}

func TestDummyCollisions(t *testing.T) {
	testCollisions(t, Dummy)
}

func TestRot13(t *testing.T) {
	testCollisions(t, Rot13)
}
