package utils

func IgnoreErr(f func() error) {
	_ = f()
}
