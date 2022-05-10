// @file main.go
// @brief
// @author tenfyzhong
// @email tenfyzhong@qq.com
// @created 2017-06-26 17:54:34
package main

import (
	"flag"
	"fmt"
	"os"
	"syscall"
	"time"
	"unsafe"
)

const (
	// IpcCreate create if key is nonexistent
	IpcCreate = 00001000
)

var mode = flag.Int("mode", 0, "0:write 1:read")

func main() {
	flag.Parse()
	shmid, _, err := syscall.Syscall(syscall.SYS_SHMGET, 2, 4, IpcCreate|0600)
	if err != 0 {
		fmt.Printf("syscall error, err: %v\n", err)
		os.Exit(-1)
	}
	fmt.Printf("shmid: %v\n", shmid)

	shmaddr, _, err := syscall.Syscall(syscall.SYS_SHMAT, shmid, 0, 0)
	if err != 0 {
		fmt.Printf("syscall error, err: %v\n", err)
		os.Exit(-2)
	}
	fmt.Printf("shmaddr: %v\n", shmaddr)

	defer syscall.Syscall(syscall.SYS_SHMDT, shmaddr, 0, 0)

	if *mode == 0 {
		fmt.Println("write mode")
		i := 0
		for {
			fmt.Printf("%d\n", i)
			*(*int)(unsafe.Pointer(uintptr(shmaddr))) = i
			i++
			time.Sleep(1 * time.Second)
		}
	} else {
		fmt.Println("read mode")
		for {
			fmt.Println(*(*int)(unsafe.Pointer(uintptr(shmaddr))))
			time.Sleep(1 * time.Second)
		}
	}
}
