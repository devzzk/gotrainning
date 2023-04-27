package demo

import "testing"

func TestOnceClose_Close(t *testing.T) {
	o := &OnceClose{}
	o.Close()
	o.Close()
}
