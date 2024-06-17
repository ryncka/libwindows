package crt

var (
	calloc  = msvcrt.NewProc("calloc")
	free    = msvcrt.NewProc("free")
	malloc  = msvcrt.NewProc("malloc")
	realloc = msvcrt.NewProc("realloc")
)

func Malloc(size uint32) (uintptr, error) {
	memory, _, err := malloc.Call(uintptr(size))
	if err != ok {
		return 0, err
	}
	return memory, nil
}

func Free(ptr uintptr) error {
	if _, _, err := free.Call(ptr); err != ok {
		return err
	}
	return nil
}

func Calloc(nitems, size uint32) (uintptr, error) {
	memory, _, err := calloc.Call(uintptr(nitems), uintptr(size))
	if err != ok {
		return 0, err
	}
	return memory, nil
}

func Realloc(ptr uintptr, size uint32) (uintptr, error) {
	memory, _, err := realloc.Call(uintptr(ptr), uintptr(size))
	if err != ok {
		return 0, err
	}
	return memory, nil
}
