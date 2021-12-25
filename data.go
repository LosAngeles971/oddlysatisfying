package main

import (
	"fmt"
	"math/rand"
)

var data = map[string]interface{}{
	"files": []string{
		"pippo",
		"start",
		"telnet",
		"ftp",
	},
	"extensions": []string{".exe", ".dat", ".bin", ""},
	"disks": []string{"/dev/sda1", "/dev/sda2", "/dev/sda3", "/dev/sda4", "/dev/sda5",},
	"filesystems": []string{"ext3", "ext4", "vfs", "vfat", "ntfs",},
}

func getFilename() string {
	files := data["files"].([]string)
	extensions := data["extensions"].([]string)
	return fmt.Sprintf("%s%s", files[rand.Intn(len(files))], extensions[rand.Intn(len(extensions))])
}

func getDisk() string {
	disks := data["disks"].([]string)
	return disks[rand.Intn(len(disks))]
}

func getFilesystem() string {
	filesystems := data["filesystems"].([]string)
	return filesystems[rand.Intn(len(filesystems))]
}
