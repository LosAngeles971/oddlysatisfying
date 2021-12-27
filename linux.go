package main

import (
	"fmt"
	"math/rand"
	"os"
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

func stateLinux() {
	log := New()
	hours, minutes, _ := time.Now().Clock()
	thisHost, err := os.Hostname()
	if err != nil {
		thisHost = "hostname"
	}
	log.Info("\n")
	log.Infof("Linux 5.15.8-100.fc34.x86_64 (%s)    %s    _x86_64_    (12 CPU)\n\n", thisHost, time.Now())
	log.Infof("%s   runq-sz  plist-sz   ldavg-1   ldavg-5  ldavg-15   blocked", fmt.Sprintf("%02d:%02d", hours, minutes))
	log.Infof("%s        11      2458      3.03      2.91      2.86         0", fmt.Sprintf("%02d:%02d", hours, minutes))
	log.Infof("%s         2      2457      4.01      3.17      2.98         0", fmt.Sprintf("%02d:%02d", hours, minutes))
	log.Infof("%s         1      2459      2.34      2.73      2.85         0", fmt.Sprintf("%02d:%02d", hours, minutes))
	log.Infof("%s         1      2455      1.90      2.30      2.55         0", fmt.Sprintf("%02d:%02d", hours, minutes))
	log.Infof("%s         1      2460      2.37      2.48      2.59         0", fmt.Sprintf("%02d:%02d", hours, minutes))
	log.Info("Average:            3      2458      2.73      2.72      2.77         0\n")
}