package main

func standard() *Log {
	log := New()
	log.addf("used connections closed")
	log.addf("used memory freed")
	log.addf("used temporary files removed")
	log.addf("******************************************")
	log.addf("*          BUILD SUCCESSFULL             *")
	log.addf("******************************************")
	return log
}