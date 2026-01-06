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
	time.Sleep(time.Second)
	test_SIGUSER2()
	time.Sleep(time.Second)
	test_SIGHUB()
	time.Sleep(time.Second)
}

func test_SIGUSER1() {
	fmt.Println("Sending syscall.SIGUSR1")
	syscall.Kill(syscall.Getpid(), syscall.SIGUSR1)
}
func test_SIGUSER2() {
	fmt.Println("Sending syscall.SIGUSR2")
	syscall.Kill(syscall.Getpid(), syscall.SIGUSR2)
}
func test_SIGHUB() {
	fmt.Println("Sending syscall.SIGHUB")
	syscall.Kill(syscall.Getpid(), syscall.SIGHUP)
}
