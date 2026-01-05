package mySignal

import (
	"fmt"
	"syscall"
	"testing"
	"time"
)

func TestAll(t *testing.T) {
	go Listener()
	time.Sleep(time.Second)
	test_SIGUSER1()
}
