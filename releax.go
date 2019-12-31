package releax

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/itsmanjeet/releax/errcode"
	"github.com/itsmanjeet/releax/msg"
)

// ReleaxRelease struct
type ReleaxRelease struct {
	Version float32 `json:"version"`
	Release string  `json:"release"`
	Build   int     `json:"build"`
}

const (
	releaseInfoDir = "/usr/releax/release.json"
)

// GetInfo of current release
func GetInfo() (*ReleaxRelease, error) {
	releaseData, err := ioutil.ReadFile(releaseInfoDir)
	if err != nil {
		return nil, err
	}

	var release ReleaxRelease
	err = json.Unmarshal(releaseData, &release)
	if err != nil {
		return nil, err
	}
	return &release, nil
}

func NeedSU() {
	if os.Geteuid() != 0 {
		msg.Error("need root permission")
		os.Exit(errcode.NeedRootPermission)
	}
}
