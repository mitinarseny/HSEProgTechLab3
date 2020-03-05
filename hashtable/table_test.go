package hashtable

import (
	"strconv"
	"testing"

	"github.com/mitinarseny/HSEProgTechLab3/hash"
	"github.com/stretchr/testify/assert"
)

func TestTable(t *testing.T) {
	tests := [][]struct {
		key   string
		value interface{}
	}{
		{},
		{
			{key: "a", value: 1,},
		},
		{
			{key: "a", value: 1,},
			{key: "b", value: 2,},
		},
		{
			{key: "a", value: 1,},
			{key: "b", value: 2,},
			{key: "a", value: 3,},
			{key: "a", value: 1,},
			{key: "c", value: 5,},
		},
	}
	for _, tt := range tests {
		t.Run(strconv.Itoa(len(tt)), func(t *testing.T) {
			table := New(uint64(len(tt)), hash.Rot13)
			for _, l := range tt {
				table.Add(l.key, l.value)
			}
			for _, l := range tt {
				_, found := table.Remove(l.key)
				assert.True(t, found)
			}
			for _, l := range tt {
				_, found := table.Get(l.key)
				assert.False(t, found)
			}
		})
	}
}