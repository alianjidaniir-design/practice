package allocate

import (
	"bytes"
	"testing"
)

func BenchmarkWBuf(b *testing.B) {
	msg := []byte("hello world")
	buffer := bytes.Buffer{}
	for i := 0; i < b.N; i++ {
		for k := 0; k < 50; k++ {
			WriteMessageBuffer(msg, buffer)
		}
	}
}
func BenchmarkWBufPointerNoReset(b *testing.B) {
	msg := []byte("hello world")
	buffer := bytes.Buffer{}
	for i := 0; i < b.N; i++ {
		for k := 0; k < 50; k++ {
			WriteMessageBuffer(msg, buffer)
		}

	}
}
func BenchmarkWBufPointerReset(b *testing.B) {
	msg := []byte("hello world")
	buffer := bytes.Buffer{}
	for i := 0; i < b.N; i++ {
		for k := 0; k < 50; k++ {
			WriteMessageBuffer(msg, buffer)
			buffer.Reset()
		}
	}

}

func BenchmarkWBufWriterReset(b *testing.B) {
	msg := []byte("hello world")
	buffer := bytes.Buffer{}
	for i := 0; i < b.N; i++ {
		for k := 0; k < 50; k++ {
			WriteMessageBuffer(msg, buffer)
			buffer.Reset()
		}
	}
}
