package allocated

import (
	"bytes"
)

func writeMessage(msg []byte) {
	b := new(bytes.Buffer)
	b.Write(msg)
}
