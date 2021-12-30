/*
This file provides all available closing functions.
At the end of one execution, the oddly satisfying program always ends randomly chossing a closing function.

NOTE: all closing functions must provide a positive log.
*/
package main

// standard is the first and more simple closing function, this log is purely fictional
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