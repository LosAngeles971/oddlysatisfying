package main

import (
	"math/rand"
	"time"
)

func download() {
	log := New(WithMode(mode_logrus))
	program := getFilename()
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

func ciscoboot() {
	log := New()
	log.Debugf("Launching IOS image at 0x80008000...")
	log.Debugf("6 Ethernet interfaces")
	log.Debugf("2 Gigabit Ethernet interface")
	log.Debugf("")
	log.Debugf("999K bytes of NVRAM.")
	log.Debugf("")
	log.Debugf("8192K bytes of Flash internal SIMM (Sector size 256K).Installed image archive")
	log.Debugf("")
	log.Debugf("Press RETURN to get started!")
	log.Debugf("* %s : %%LINEPROTO-5-UPDOWN: Line protocol on Interface GigabitEthernet2/0, changed state to up",
		time.Now().Format(time.StampMilli))
	log.Debugf("* %s : %%LINEPROTO-5-UPDOWN: Line protocol on Interface GigabitEthernet2/1, changed state to up",
		time.Now().Format(time.StampMilli))
	log.Debugf("* %s : %%LINEPROTO-5-UPDOWN: Line protocol on Interface Ethernet3/0, changed state to up",
		time.Now().Format(time.StampMilli))
	log.Debugf("* %s : %%LINEPROTO-5-UPDOWN: Line protocol on Interface Ethernet3/1, changed state to up",
		time.Now().Format(time.StampMilli))
	log.Debugf("* %s : %%LINEPROTO-5-UPDOWN: Line protocol on Interface Ethernet3/2, changed state to up",
		time.Now().Format(time.StampMilli))
	log.Debugf("* %s : %%LINEPROTO-5-UPDOWN: Line protocol on Interface Ethernet3/3, changed state to up",
		time.Now().Format(time.StampMilli))
	log.Debugf("* %s : %%LINEPROTO-5-UPDOWN: Line protocol on Interface Ethernet3/4, changed state to up",
		time.Now().Format(time.StampMilli))
	log.Debugf("* %s : %%LINEPROTO-5-UPDOWN: Line protocol on Interface Ethernet3/5, changed state to up",
		time.Now().Format(time.StampMilli))
	log.Debugf("* %s : %%LINEPROTO-5-UPDOWN: Line protocol on Interface Loopback0, changed state to up",
		time.Now().Format(time.StampMilli))
	log.Debugf("* %s : %%LINEPROTO-5-UPDOWN: Line protocol on Interface Loopback1, changed state to up",
		time.Now().Format(time.StampMilli))
	log.Debugf("* %s : %%SYS-5-RESTART: System restarted --",
		time.Now().Format(time.StampMilli))
	log.Debugf("* %s : %%SNMP-5-COLDSTART: SNMP agent on host R3 is undergoing a cold start",
		time.Now().Format(time.StampMilli))
	log.Debugf("* %s : %%OSPF-5-ADJCHG: Process 1, Nbr 1.1.1.1 on Ethernet3/1 from LOADING to FULL, Loading Done",
		time.Now().Format(time.StampMilli))
	log.Debugf("* %s : %%OSPF-5-ADJCHG: Process 1, Nbr 4.4.4.4 on Ethernet3/3 from LOADING to FULL, Loading Done",
		time.Now().Format(time.StampMilli))
	log.Debugf("* %s : %%BGP-5-ADJCHANGE: neighbor 1.1.1.1 Up",
		time.Now().Format(time.StampMilli))
	log.Debugf("* %s : %%BGP-5-ADJCHANGE: neighbor 4.4.4.4 Up",
		time.Now().Format(time.StampMilli))
}
