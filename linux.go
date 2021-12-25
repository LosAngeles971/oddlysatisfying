package main

import (
	"math/rand"
	"time"
)

func fsck() {
	dd := getDisk()
	ff := getFilesystem()
	files := rand.Intn(99999999)
	blocks := rand.Intn(999999999)
	log := New()
	log.Infof("Log of fsck -C -a -V -t %s %s", ff, dd)
	log.Info(time.Now())
	log.Infof("[/sbin/fsck.ext4 (1) -- %s] fsck.%s -a -C0 %s", dd, ff, dd)
	log.Infof("%s: clean, %v/%v files, %v/%v blocks", dd, rand.Intn(files), files, rand.Intn(blocks), blocks)
}