package releax

import "os"

func Exist(f string) bool {
	if _, err := os.Stat(f); err != nil {
		return false
	}

	return true
}
