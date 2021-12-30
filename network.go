package main

import (
	"math/rand"
	"time"
)

func download() *Log {
	log := New()
	program := getFilename()
	size := rand.Intn(1000000)
	log.addf("attempt to download file [%s] of %v bytes", program, size)
	log.add("opening connection...")
	log.add("downloading...")
	c := 0;
	for i :=0; i < 10; i++ {
		c += int(size / 10)
		time.Sleep(time.Duration(rand.Intn(5))*time.Second)
		log.addf("downloaded %v of %v", c, size)
	}
	log.addf("file [%s] of %v bytes successfully downloaded", program, size)
	return log
}

func adminport(log *Log, n int) {
	broken := rand.Int31n(100)
	if broken >49 {
		log.addf("GigabitEthernet2/%v is up, line protocol is up", n)
		log.addf("  %v minute input rate %v bits/sec, %v packets/sec", rand.Intn(10), rand.Intn(1000000), rand.Intn(50))
		log.addf("  %v minute output rate %v bits/sec, %v packets/sec", rand.Intn(10), rand.Intn(1000000), rand.Intn(50))
	} else {
		log.addf("GigabitEthernet2/%v is up, line protocol is down", n)
		log.addf("  %v minute input rate 0 bits/sec, 0 packets/sec", rand.Intn(10))
		log.addf("  %v minute output rate 0 bits/sec, 0 packets/sec", rand.Intn(10))
		log.addf("GigabitEthernet2/%v is administratively down, line protocol is down", n)
		log.addf("  %v minute input rate 0 bits/sec, 0 packets/sec", rand.Intn(10))
		log.addf("  %v minute output rate 0 bits/sec, 0 packets/sec", rand.Intn(10))
		log.addf("GigabitEthernet2/%v is administratively up, line protocol is up", n)
		log.addf("  %v minute input rate 0 bits/sec, 0 packets/sec", rand.Intn(10))
		log.addf("  %v minute output rate 0 bits/sec, 0 packets/sec", rand.Intn(10))
		log.addf("GigabitEthernet2/%v is up, line protocol is up", n)
		log.addf("  %v minute input rate %v bits/sec, %v packets/sec", rand.Intn(10), rand.Intn(1000000), rand.Intn(50))
		log.addf("  %v minute output rate %v bits/sec, %v packets/sec", rand.Intn(10), rand.Intn(1000000), rand.Intn(50))
	}
}

func adminports() *Log {
	log := New()
	for i := 1; i < 9; i++ {
		adminport(log, i)
	}	
	return log
}

func ciscoboot() *Log {
	log := New()
	log.addf("Launching IOS image at 0x80008000...")
	log.addf("6 Ethernet interfaces")
	log.addf("2 Gigabit Ethernet interface")
	log.addf("")
	log.addf("999K bytes of NVRAM.")
	log.addf("")
	log.addf("8192K bytes of Flash internal SIMM (Sector size 256K).Installed image archive")
	log.addf("")
	log.addf("Press RETURN to get started!")
	log.addf("* %s : %%LINEPROTO-5-UPDOWN: Line protocol on Interface GigabitEthernet2/0, changed state to up",
		time.Now().Format(time.StampMilli))
	log.addf("* %s : %%LINEPROTO-5-UPDOWN: Line protocol on Interface GigabitEthernet2/1, changed state to up",
		time.Now().Format(time.StampMilli))
	log.addf("* %s : %%LINEPROTO-5-UPDOWN: Line protocol on Interface Ethernet3/0, changed state to up",
		time.Now().Format(time.StampMilli))
	log.addf("* %s : %%LINEPROTO-5-UPDOWN: Line protocol on Interface Ethernet3/1, changed state to up",
		time.Now().Format(time.StampMilli))
	log.addf("* %s : %%LINEPROTO-5-UPDOWN: Line protocol on Interface Ethernet3/2, changed state to up",
		time.Now().Format(time.StampMilli))
	log.addf("* %s : %%LINEPROTO-5-UPDOWN: Line protocol on Interface Ethernet3/3, changed state to up",
		time.Now().Format(time.StampMilli))
	log.addf("* %s : %%LINEPROTO-5-UPDOWN: Line protocol on Interface Ethernet3/4, changed state to up",
		time.Now().Format(time.StampMilli))
	log.addf("* %s : %%LINEPROTO-5-UPDOWN: Line protocol on Interface Ethernet3/5, changed state to up",
		time.Now().Format(time.StampMilli))
	log.addf("* %s : %%LINEPROTO-5-UPDOWN: Line protocol on Interface Loopback0, changed state to up",
		time.Now().Format(time.StampMilli))
	log.addf("* %s : %%LINEPROTO-5-UPDOWN: Line protocol on Interface Loopback1, changed state to up",
		time.Now().Format(time.StampMilli))
	log.addf("* %s : %%SYS-5-RESTART: System restarted --",
		time.Now().Format(time.StampMilli))
	log.addf("* %s : %%SNMP-5-COLDSTART: SNMP agent on host R3 is undergoing a cold start",
		time.Now().Format(time.StampMilli))
	log.addf("* %s : %%OSPF-5-ADJCHG: Process 1, Nbr 1.1.1.1 on Ethernet3/1 from LOADING to FULL, Loading Done",
		time.Now().Format(time.StampMilli))
	log.addf("* %s : %%OSPF-5-ADJCHG: Process 1, Nbr 4.4.4.4 on Ethernet3/3 from LOADING to FULL, Loading Done",
		time.Now().Format(time.StampMilli))
	log.addf("* %s : %%BGP-5-ADJCHANGE: neighbor 1.1.1.1 Up",
		time.Now().Format(time.StampMilli))
	log.addf("* %s : %%BGP-5-ADJCHANGE: neighbor 4.4.4.4 Up",
		time.Now().Format(time.StampMilli))
	return log
}

func curlDownload() *Log {
    log := New()
    program := getFilename()
    url := getUrl()
    size := rand.Intn(9999)
    avgspeed := rand.Intn(9999)
    log.addf("\n\ncurl %s/download/%s", url, program)
    log.add("   % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current")
    log.add("                                  Dload  Upload   Total   Spent    Left  Speed")
    log.addf(" 100 %4dk  100 %4dk    0     0  %4dk      0 --:--:-- --:--:-- --:--:-- %4dk\n\n",size,size,avgspeed,avgspeed+123)
	return log
}