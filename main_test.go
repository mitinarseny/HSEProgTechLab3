package main

import (
	"math/rand"
	"strconv"
	"testing"

	"github.com/mitinarseny/HSEProgTechLab3/hash"
	"github.com/mitinarseny/HSEProgTechLab3/hashtable"
)

func BenchmarkHashTableRot13(b *testing.B) {
	benchmarkHashTable(b, hash.Rot13)
}

func BenchmarkHashTableDummy(b *testing.B) {
	benchmarkHashTable(b, hash.Dummy)
}

func benchmarkHashTable(b *testing.B, hashFunc func(string) uint64) {
	tests := []uint64{
		10,
		50,
		100,
		500,
		1000,
		5000,
		10000,
	}
	for _, n := range tests {
		b.Run(strconv.FormatInt(int64(n), 10), func(b *testing.B) {
			t := hashtable.New(n, hashFunc)
			for i := uint64(0); i < n; i++ {
				t.Add(genString(64), i)
			}
			b.ResetTimer()
			key := strconv.Itoa(int(n / 2))
			for i := 0; i < b.N; i++ {
				t.Get(key)
			}
		})
	}
}

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func genString(n uint64) string {
	s := make([]byte, 0, n)
	for i := uint64(0); i < n; i++ {
		s = append(s, letters[rand.Intn(len(letters))])
	}
	return string(s)
}
