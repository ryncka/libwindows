package crt

var (
	wcscpy = msvcrt.NewProc("wcscpy")
	wcslen = msvcrt.NewProc("wcslen")
)

func Wcscpy(dest, src uintptr) error {
	if _, _, err := wcscpy.Call(dest, src); err != ok {
		return err
	}
	return nil
}

func Wcslen(str uintptr) (uint32, error) {
	result, _, err := wcslen.Call(str)
	if err != ok {
		return 0, err
	}
	return uint32(result), nil
}
