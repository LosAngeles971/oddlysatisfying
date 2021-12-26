package main

func bootMsDos() {
	log := New()
	log.Info("Starting MS-DOS...")
	log.Info("HIMEM is testing extended memory...done.")
	log.Info("A:\\")
}

func c64() {
	log := New(WithDelay(500, true))
	log.Info("    **** COMMODORE 64 BASIC X2 ****")
	log.Info(" 64K RAM SYSTEM 38911 BASIC BYTES FREE")
	log.Info("READY")
	log.Info("LOAD\"*\",8,1")
	log.Info("SEARCHING FOR *")
	log.Info("LOADING FROM 0801 (2049) TO (7786)")
	log.Info("READY")
}