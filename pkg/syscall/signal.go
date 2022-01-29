package syscall

import (
	"fmt"
	"syscall"
)

func Do() {
	fmt.Printf("before kill: %v, %v \n", syscall.Getpid(), syscall.SIGINT)
	syscall.Kill(syscall.Getpid(), syscall.SIGINT)
	fmt.Printf("after kill")
}
