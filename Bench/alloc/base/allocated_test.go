package allocated

import (
	"testing"
)

func BenchmarkAllocated(b *testing.B) {
	msg := []byte("Karbala")
	for i := 0; i < b.N; i++ {
		for j := 0; j < 52; j++ {
			writeMessage(msg)
		}
	}
}
