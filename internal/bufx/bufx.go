package bufx

func Dup(s []byte) []byte {
	d := make([]byte, len(s))
	copy(d, s)
	return d
}
