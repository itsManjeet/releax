package disk

import (
	"fmt"
	"testing"
)

func TestDisk(t *testing.T) {
	disk, err := LoadDisk()
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println(disk)
}
