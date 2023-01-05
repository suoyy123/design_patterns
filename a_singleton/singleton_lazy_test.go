package a_singleton

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//go test .
func TestGetLazySingleton(t *testing.T) {
	assert.Equal(t, GetLazySingleton(), GetLazySingleton())
}

//go test -bench . -benchmem
func BenchmarkGetLazySingleton(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if GetLazySingleton() != GetLazySingleton() {
				b.Error("test fail")
			}
		}
	})
}