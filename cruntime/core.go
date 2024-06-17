package crt

import "syscall"

const ok = syscall.Errno(0)

var (
	msvcrt = syscall.NewLazyDLL("msvcrt.dll")
)
