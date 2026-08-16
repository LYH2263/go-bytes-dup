package bufx

func Dup(s []byte) []byte {
	out := make([]byte, len(s))
	copy(out, s)
	return out
}
