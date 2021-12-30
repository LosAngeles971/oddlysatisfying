package main

func bootMsDos() *Log {
	log := New()
	log.add("Starting MS-DOS...")
	log.add("HIMEM is testing extended memory...done.")
	log.add("A:\\")
	return log
}

func c64() *Log {
	log := New()
	log.add("    **** COMMODORE 64 BASIC X2 ****")
	log.add(" 64K RAM SYSTEM 38911 BASIC BYTES FREE")
	log.add("READY")
	log.add("LOAD\"*\",8,1")
	log.add("SEARCHING FOR *")
	log.add("LOADING FROM 0801 (2049) TO (7786)")
	log.add("READY")
	return log
}