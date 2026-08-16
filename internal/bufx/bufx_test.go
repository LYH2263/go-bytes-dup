package bufx

import "testing"

func TestDupIndependent(t *testing.T) {
	src := []byte("hello")
	d := Dup(src)
	d[0] = 'H'
	if src[0] != 'h' {
		t.Fatalf("src mutated: %q", src)
	}
}
