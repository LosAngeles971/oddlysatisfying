package main

import (
	"fmt"
	"math/rand"
)

func download() {
	log := New()
	filenames := data["filenames"].([]string)
	extensions := data["extensions"].([]string)
	program := fmt.Sprintf("%s%s", filenames[rand.Intn(len(filenames))], extensions[rand.Intn(len(extensions))])
	size := rand.Intn(1000000)
	log.Infof("attempt to download file [%s] of %v bytes", program, size)
	log.Debug("opening connection...")
	log.Debug("downloading...")
	log.Infof("file [%s] of %v bytes successfully downloaded", program, size)
}

func mvs() {
	log := New()
	log.Info("STC18213 00000090 $HASP100 BPXAS ON STCINRDR")
	log.Info("STC18213 00000090 $HASP373 BPXAS STARTED")
	log.Info("STC18213 80000010 IEF403I BPXAS - STARTED - TIME=13.36.36 - ASID=001F - SC53")
	log.Info("STC16316 00000291 IST663I IPS SRQ REQUEST FROM ISTAPNCP FAILED, SENSE=08570002")
    log.Info(" 	111 00000291 IST664I REAL OLU=USIBMSC.S52TOS48 REAL DLU=USIBMSC.S48TO")
    log.Info(" 	111 00000291 IST889I SID = ED0385CAAEEAAF28")
    log.Info(" 	111 00000291 IST264I REQUIRED RESOURCE S48TOS52 NOT ACTIVE")
    log.Info("  111 00000291 IST314I END")
	log.Info("STC16352 00000291 IST663I IPS SRQ REQUEST FROM ISTAPNCP FAILED, SENSE=087D0001")
    log.Info(" 	883 00000291 IST664I REAL OLU=USIBMSC.S52TOS48 ALIAS DLU=USIBMSC.S48TO")
    log.Info(" 	883 00000291 IST889I SID = ED0385CAAEEAAF28")
    log.Info(" 	883 00000291 IST314I END")
	log.Info("STC28215 00000291 IST663I IPS SRQ REQUEST TO ISTAPNCP FAILED, SENSE=08570002 86")
	log.Info(" 	864 00000291 IST664I REAL OLU=USIBMSC.S52TOS48 ALIAS DLU=USIBMSC.S48TO")
    log.Info(" 	864 00000291 IST889I SID = ED0385CAAEEAAF28")
    log.Info(" 	864 00000291 IST264I REQUIRED RESOURCE S48TOS52 NOT ACTIVE")
    log.Info(" 	864 00000291 IST891I USIBMSC.SC48M GENERATED FAILURE NOTIFICATION")
    log.Info(" 	864 00000291 IST314I END")
}

func log4j() {
	log := New()
	log.Info("setting log4j level...")
	log.Debug("searching for environment variables...")
	log.Debug("searching for configuration files...")
	log.Error("ATTENTION: log4j hacked!")
	log.Error("ATTENTION: log4j under attack!")
	log.Error("starting recovery mode...")
	log.Error("recovery mode started")
	log.Info("system cleaned")
	log.Info("log4j disabled")
}