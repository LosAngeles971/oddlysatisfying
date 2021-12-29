package main

func pfSenseMenu() *Log {
    log := New()
    log.add("==============================================================================")
    log.add("                            pfSense Menu                                      ")
    log.add("==============================================================================")
    log.add("WAN (wan)       -> vmx0       -> v4/DHCP4: 198.51.100.6/24")
    log.add("                                 v6/DHCP6: 2001:db8::20c:29ff:fe78:6e4e/64")
    log.add("LAN (lan)       -> vmx1       -> v4: 10.6.0.1/24")
    log.add("                                 v6/t6: 2001:db8:1:eea0:20c:29ff:fe78:6e58/64")
    log.add("0) Logout (SSH only)                  9) pfTop")
    log.add("1) Assign Interfaces                 10) Filter Logs")
    log.add("2) Set interface(s) IP address       11) Restart webConfigurator")
    log.add("3) Reset webConfigurator password    12) PHP shell + pfSense tools")
    log.add("4) Reset to factory defaults         13) Update from console")
    log.add("5) Reboot system                     14) Disable Secure Shell (sshd)")
    log.add("6) Halt system                       15) Restore recent configuration")
    log.add("7) Ping host                         16) Restart PHP-FPM")
    log.add("8) Shell")
    return log
}