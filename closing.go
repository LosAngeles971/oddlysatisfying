package main

func standard() {
	log := New(WithDelay(3))
	log.Info("all connections closed")
	log.Info("memory freed")
	log.Info("temporary files removed")
	log.Info("build successfull")
}