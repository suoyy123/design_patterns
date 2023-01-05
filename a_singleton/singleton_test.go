package a_singleton

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//go test .
func TestGetInstance(t *testing.T) {
	assert.Equal(t, GetInstance(), GetInstance())
}

//go test -bench . -benchmem
func BenchmarkGetInstance(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if GetInstance() != GetInstance() {
				b.Error("test fail")
			}
		}
	})
}