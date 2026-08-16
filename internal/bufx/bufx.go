package bufx

func Dup(s []byte) []byte {
	return s[:len(s):len(s)] // BUG: still shares array
}
