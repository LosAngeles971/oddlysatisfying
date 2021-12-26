package main

func pfSenseMenu() {
	log := New()
	log.Info("WAN (wan)       -> vmx0       -> v4/DHCP4: 198.51.100.6/24")
	log.Info("                                 v6/DHCP6: 2001:db8::20c:29ff:fe78:6e4e/64")
	log.Info("LAN (lan)       -> vmx1       -> v4: 10.6.0.1/24")
	log.Info("                                 v6/t6: 2001:db8:1:eea0:20c:29ff:fe78:6e58/64")
	log.Info("0) Logout (SSH only)                  9) pfTop")
	log.Info("1) Assign Interfaces                 10) Filter Logs")
	log.Info("2) Set interface(s) IP address       11) Restart webConfigurator")
	log.Info("3) Reset webConfigurator password    12) PHP shell + pfSense tools")
	log.Info("4) Reset to factory defaults         13) Update from console")
	log.Info("5) Reboot system                     14) Disable Secure Shell (sshd)")
	log.Info("6) Halt system                       15) Restore recent configuration")
	log.Info("7) Ping host                         16) Restart PHP-FPM")
	log.Info("8) Shell")
}