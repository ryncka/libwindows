package crt

var (
	memcpy  = msvcrt.NewProc("memcpy")
	memset  = msvcrt.NewProc("memset")
	strcpy  = msvcrt.NewProc("strcpy")
	strncpy = msvcrt.NewProc("strncpy")
	strlen  = msvcrt.NewProc("strlen")
)

func Memcpy(dest, src uintptr, n uint32) error {
	_, _, err := memcpy.Call(dest, src, uintptr(n))
	if err != ok {
		return err
	}
	return nil
}

func Memset(str uintptr, c int32, n uint32) error {
	_, _, err := memset.Call(str, uintptr(c), uintptr(n))
	if err != ok {
		return err
	}
	return nil
}

func Strcpy(dest, src uintptr) error {
	_, _, err := strcpy.Call(dest, src)
	if err != ok {
		return err
	}
	return nil
}

func Strncpy(dest, src uintptr, n uint32) error {
	_, _, err := strncpy.Call(dest, src, uintptr(n))
	if err != ok {
		return err
	}
	return nil
}

func Strlen(str uintptr) (uint32, error) {
	result, _, err := strlen.Call(str)
	if err != ok {
		return 0, err
	}
	return uint32(result), nil
}
