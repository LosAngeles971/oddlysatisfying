package main

func standard() {
	log := New()
	log.Info("used connections closed")
	log.Info("used memory freed")
	log.Info("used temporary files removed")
	log.Info("******************************************")
	log.Info("*          BUILD SUCCESSFULL             *")
	log.Info("******************************************")
}