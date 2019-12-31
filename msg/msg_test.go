package msg

import "testing"

func TestPrintErr(t *testing.T) {
	Error("this is error")
	Notice("this is notice")
}
