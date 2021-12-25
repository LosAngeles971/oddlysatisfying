package main

import (
	"fmt"
	"math/rand"
	"time"
)

func download() {
	log := New(WithMode(mode_logrus))
	filenames := data["filenames"].([]string)
	extensions := data["extensions"].([]string)
	program := fmt.Sprintf("%s%s", filenames[rand.Intn(len(filenames))], extensions[rand.Intn(len(extensions))])
	size := rand.Intn(1000000)
	log.Infof("attempt to download file [%s] of %v bytes", program, size)
	log.Debug("opening connection...")
	log.Debug("downloading...")
	c := 0;
	for i :=0; i < 10; i++ {
		c += int(size / 10)
		time.Sleep(time.Duration(rand.Intn(5))*time.Second)
		log.Infof("downloaded %v of %v", c, size)
	}
	log.Infof("file [%s] of %v bytes successfully downloaded", program, size)
}

func adminport(n int) {
	broken := rand.Int31n(100)
	log := New()
	if broken >49 {
		log.Debugf("GigabitEthernet2/%v is up, line protocol is up", n)
		log.Debugf("  %v minute input rate %v bits/sec, %v packets/sec", rand.Intn(10), rand.Intn(1000000), rand.Intn(50))
		log.Debugf("  %v minute output rate %v bits/sec, %v packets/sec", rand.Intn(10), rand.Intn(1000000), rand.Intn(50))
	} else {
		log.Debugf("GigabitEthernet2/%v is up, line protocol is down", n)
		log.Debugf("  %v minute input rate 0 bits/sec, 0 packets/sec", rand.Intn(10))
		log.Debugf("  %v minute output rate 0 bits/sec, 0 packets/sec", rand.Intn(10))
		log.Debugf("GigabitEthernet2/%v is administratively down, line protocol is down", n)
		log.Debugf("  %v minute input rate 0 bits/sec, 0 packets/sec", rand.Intn(10))
		log.Debugf("  %v minute output rate 0 bits/sec, 0 packets/sec", rand.Intn(10))
		log.Debugf("GigabitEthernet2/%v is administratively up, line protocol is up", n)
		log.Debugf("  %v minute input rate 0 bits/sec, 0 packets/sec", rand.Intn(10))
		log.Debugf("  %v minute output rate 0 bits/sec, 0 packets/sec", rand.Intn(10))
		log.Debugf("GigabitEthernet2/%v is up, line protocol is up", n)
		log.Debugf("  %v minute input rate %v bits/sec, %v packets/sec", rand.Intn(10), rand.Intn(1000000), rand.Intn(50))
		log.Debugf("  %v minute output rate %v bits/sec, %v packets/sec", rand.Intn(10), rand.Intn(1000000), rand.Intn(50))
	}
}
func adminports() {
	for i := 1; i < 9; i++ {
		adminport(i)
	}	
}

