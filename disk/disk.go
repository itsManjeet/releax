package disk

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
)

// Disk with
type Disk struct {
	Name           string `json:"name"`
	Label          string `json:"label"`
	UUID           string `json:"uuid"`
	Path           string `json:"path"`
	Size           string `json:"size"`
	MountPoint     string `json:"mountpoint"`
	Type           string `json:"type"`
	FileSystemType string `json:"fstype"`
	Children       []Disk `json:"children"`
}

// BlockDevices struct
type BlockDevices struct {
	Disks []Disk `json:"blockdevices"`
}

// LoadDisk to load all disks with uuid, Label, path
func LoadDisk() (BlockDevices, error) {
	cmdout, err := exec.Command("lsblk", "-J", "-O").Output()
	if err != nil {
		return BlockDevices{}, err
	}
	var blockdevices BlockDevices
	if err = json.Unmarshal(cmdout, &blockdevices); err != nil {
		return BlockDevices{}, err
	}

	return blockdevices, nil
}

// Reload BlockDevice data
func (disk *BlockDevices) Reload() error {
	dsk, err := LoadDisk()
	if err != nil {
		return err
	}
	disk = &dsk
	return nil
}

// Oflabel to get disk of label
func (disk BlockDevices) Oflabel(label string) (Disk, error) {
	for _, dsk := range disk.Disks {
		if dsk.Label == label {
			return dsk, nil
		}
		for _, part := range dsk.Children {
			if part.Label == label {
				return part, nil
			}
		}
	}

	return Disk{}, fmt.Errorf("unable to find disk with label %s", label)
}

// OfPath to get disk of Path
func (disk BlockDevices) OfPath(path string) (Disk, error) {
	for _, dsk := range disk.Disks {
		if dsk.Path == path {
			return dsk, nil
		}
		for _, part := range dsk.Children {
			if part.Path == path {
				return part, nil
			}
		}
	}

	return Disk{}, fmt.Errorf("unable to find disk with label %s", path)
}

// CheckMountAt if mountpoint is already mounted
func (disk BlockDevices) CheckMountAt(path string) bool {
	for _, dsk := range disk.Disks {
		if dsk.MountPoint == "" {
			continue
		}
		if dsk.MountPoint == path {
			return true
		}
	}
	return false
}

// MountTo mount disk to path
func (disk Disk) MountTo(mountpoint string) error {
	if _, err := os.Stat(mountpoint); err != nil {
		os.MkdirAll(mountpoint, 0755)
	}
	cmd, err := exec.Command("mount", disk.Path, mountpoint).Output()
	if err != nil {
		return fmt.Errorf("cmd: %v, error: %v", string(cmd[:]), err.Error())
	}
	return nil
}

// Unmount Disk
func (disk Disk) Unmount() error {
	cmd, err := exec.Command("mount", disk.Path).Output()
	if err != nil {
		return fmt.Errorf("cmd: %v, error: %v", string(cmd[:]), err.Error())
	}
	return nil
}

// SetLabel of ext4 disk
func (disk Disk) SetLabel(label string) error {
	if disk.FileSystemType != "ext4" {
		return fmt.Errorf("cant label disk of partition type other then ext4, current is %s", disk.FileSystemType)
	}
	cmd, err := exec.Command("e2label", disk.Path, label).Output()
	if err != nil {
		return fmt.Errorf("cmd: %v, error: %v", string(cmd[:]), err.Error())
	}
	return nil
}

// Format disk parition
func (disk Disk) Format() error {
	cmd, err := exec.Command("mkfs.ext4", "-F", disk.Path).Output()
	if err != nil {
		return fmt.Errorf("cmd: %v, error: %v", string(cmd[:]), err.Error())
	}
	return nil
}
