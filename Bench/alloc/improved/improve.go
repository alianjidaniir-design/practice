package allocate

import (
	"bytes"
	"io"
)

func WriteMessageBuffer(msg []byte, b bytes.Buffer) {
	b.Write(msg)
}
func WriteMessageBufferPinter(msg []byte, b *bytes.Buffer) {
	b.Write(msg)
}
func WriteMessageBufferWriter(msg []byte, b io.Writer) {
	_, err := b.Write(msg)
	if err != nil {
		return
	}
}
