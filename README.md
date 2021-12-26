# Useless but oddly satisfying program to build nothing!

Hi all!

This project focuses on the importance of positive emotions and relaxing moments for stressed developers and sysadmins.

When you are waiting for building your app or for installing a complex software, and you feel stress, bad or weird mood, then you may use the _oddlysatisfying_ program.

The _oddlysatisfying_ program is a completely useless script for building nothig! But having oddly satisfying log messages till to an always positive ending.

The _oddlysatisfying_ program randomly choose vintage and modern operations, typing nice logging on your screen and at the end, you always receive a positive "BUILD SUCCESSFUL" final message.

For the best effect, you may have two monitors and use the far one for your boring and stressful activities add the close one to run the _oddlysatisfying_ program (maximing the terminal!).

And you may contribute to this project providing ideas, nice log examples and code, in order to keep _oddlysatisfying_ program completely useless but oddly satysfying and relaxing.

To download the superb and useless _oddlysatisfying_ program click on the release tags and then choose the last (or wathever you want) release.
Chosen the release, you may download the right binary for your machine.

If you're lazy, then just launch the _oddlysatisfying_ program.
Otherwise, you may launch  _oddlysatisfying_ program specifying a number, the latter will be used to randomly generate the number of programs to be executed.

Anyway remember! Even if you see attempts to download something, recovery actions, boot of programs, **the superb _oddlysatisfying_ program never really do something**.
The superb _oddlysatisfying_ program is and will be always completely useless and harmless.

## How to contribute

The program is really simple:

* _main.go_ -> includes the list of all available _logging func_ and the _func main_, the latter randomly calls the _logging func_
* _console.go_ -> provides the helper to write log on screen
* _data.go_ -> provides some data to make the logging more realistic (a list of possible file names, a list of possible disk devices, etc.)
* _closing.go_ -> provides the list of possible _func_ to (positively) ends the execution
* _mainframe.go_ -> provides _logging func_ related to the world of mainframe computers
* _linux.go_ -> provides _logging func_ related to the world of linux computers
* _network.go_ -> provides _logging func_ related to the world of network devices and actions
* _pc.go_ -> provides _logging func_ related to the world of personal computers

If you would add a new _logging func_, you just need to chose an existing file or creating a new one, and populate it with new code.

Last but not least, when you add a new _logging func_, you need to update the _vtp_ map into main.go: 

```golang
var vtp = map[int]uselessFunc{
	0: download,
	1: mvs,
	2: bootVax,
	3: c64,
	4: adminports,
	5: fsck,
	6: bootMsDos,
	7: stateLinux,
	8: qradarAutoupdate,
}
```

## Conduct guidelines

* if possible, you should make your _logging func_ dynamic (not static), it means generating at the time of execution the name of servers, files, dates, etc
* _logging func_ may also include failures, recovery actions, loss of data
* if you modify the way the program ends, keep in mind that the execution must always end with the calling of one of the possible _closing func_
* all the possible _closing func_ must always end with a positive, relaxing, peacefull and successfull final message
* the code cannot contain vulgar, aggressive or discriminatory language

## Few words more...

At the date, the superb _oddlysatisfying_ program supports the following architectures:

* Linux/amd64
* Windows/amd64
* Darwin/amd64

But you may ask for additional architecture and (maybe) you'll be satisfied.

Enjoy!
Peace and love, @LosAngeles971 (aka @Bluepulsar971).


## P.s. an example of execution

```bash
C:\Users\losangeles971\myhome\black\oddlysatisfying>oddlysatisfying.exe 15
here we go!  [13]




WAN (wan)       -> vmx0       -> v4/DHCP4: 198.51.100.6/24
                                 v6/DHCP6: 2001:db8::20c:29ff:fe78:6e4e/64
LAN (lan)       -> vmx1       -> v4: 10.6.0.1/24
                                 v6/t6: 2001:db8:1:eea0:20c:29ff:fe78:6e58/64
0) Logout (SSH only)                  9) pfTop
1) Assign Interfaces                 10) Filter Logs
2) Set interface(s) IP address       11) Restart webConfigurator
3) Reset webConfigurator password    12) PHP shell + pfSense tools
4) Reset to factory defaults         13) Update from console
5) Reboot system                     14) Disable Secure Shell (sshd)
6) Halt system                       15) Restore recent configuration
7) Ping host                         16) Restart PHP-FPM
8) Shell
Starting MS-DOS...
HIMEM is testing extended memory...done.
A:\
Log of fsck -C -a -V -t ext4 /dev/sda3
2021-12-26 14:17:03.3263846 +0100 CET m=+0.004581101
[/sbin/fsck.ext4 (1) -- /dev/sda3] fsck.ext4 -a -C0 /dev/sda3
/dev/sda3: clean, 9570864/12268556 files, 221852713/343121015 blocks
time="2021-12-26T14:17:03+01:00" level=info msg="attempt to download file [telnet.bin] of 133243 bytes"
time="2021-12-26T14:17:05+01:00" level=info msg="downloaded 13324 of 133243"
time="2021-12-26T14:17:05+01:00" level=info msg="downloaded 26648 of 133243"
time="2021-12-26T14:17:06+01:00" level=info msg="downloaded 39972 of 133243"
time="2021-12-26T14:17:09+01:00" level=info msg="downloaded 53296 of 133243"
time="2021-12-26T14:17:11+01:00" level=info msg="downloaded 66620 of 133243"
time="2021-12-26T14:17:13+01:00" level=info msg="downloaded 79944 of 133243"
time="2021-12-26T14:17:14+01:00" level=info msg="downloaded 93268 of 133243"
time="2021-12-26T14:17:14+01:00" level=info msg="downloaded 106592 of 133243"
time="2021-12-26T14:17:16+01:00" level=info msg="downloaded 119916 of 133243"
time="2021-12-26T14:17:18+01:00" level=info msg="downloaded 133240 of 133243"
time="2021-12-26T14:17:18+01:00" level=info msg="file [telnet.bin] of 133243 bytes successfully downloaded"
Log of fsck -C -a -V -t ext3 /dev/sda2
2021-12-26 14:17:18.4600727 +0100 CET m=+15.138269201
[/sbin/fsck.ext4 (1) -- /dev/sda2] fsck.ext3 -a -C0 /dev/sda2
/dev/sda2: clean, 1859981/20919089 files, 516034364/919973670 blocks
        Linux 5.15.8-100.fc34.x86_64 (sebalinux)    2021-12-26 14:17:18.4618942 +0100 CET m=+15.140090701  _x86_64_    (12 CPU)


1417   runq-sz  plist-sz   ldavg-1   ldavg-5  ldavg-15   blocked
1417        11      2458      3.03      2.91      2.86         0
1417         2      2457      4.01      3.17      2.98         0
1417         1      2459      2.34      2.73      2.85         0
1417         1      2455      1.90      2.30      2.55         0
1417         1      2460      2.37      2.48      2.59         0
Average:            3      2458      2.73      2.72      2.77         0
STC18213 00000090 $HASP100 BPXAS ON STCINRDR
STC18213 00000090 $HASP373 BPXAS STARTED
STC18213 80000010 IEF403I BPXAS - STARTED - TIME=13.36.36 - ASID=001F - SC53
STC16316 00000291 IST663I IPS SRQ REQUEST FROM ISTAPNCP FAILED, SENSE=08570002
        111 00000291 IST664I REAL OLU=USIBMSC.S52TOS48 REAL DLU=USIBMSC.S48TO
        111 00000291 IST889I SID = ED0385CAAEEAAF28
        111 00000291 IST264I REQUIRED RESOURCE S48TOS52 NOT ACTIVE
  111 00000291 IST314I END
STC16352 00000291 IST663I IPS SRQ REQUEST FROM ISTAPNCP FAILED, SENSE=087D0001
        883 00000291 IST664I REAL OLU=USIBMSC.S52TOS48 ALIAS DLU=USIBMSC.S48TO
        883 00000291 IST889I SID = ED0385CAAEEAAF28
        883 00000291 IST314I END
STC28215 00000291 IST663I IPS SRQ REQUEST TO ISTAPNCP FAILED, SENSE=08570002 86
        864 00000291 IST664I REAL OLU=USIBMSC.S52TOS48 ALIAS DLU=USIBMSC.S48TO
        864 00000291 IST889I SID = ED0385CAAEEAAF28
        864 00000291 IST264I REQUIRED RESOURCE S48TOS52 NOT ACTIVE
        864 00000291 IST891I USIBMSC.SC48M GENERATED FAILURE NOTIFICATION
        864 00000291 IST314I END
STC18213 00000090 $HASP100 BPXAS ON STCINRDR
STC18213 00000090 $HASP373 BPXAS STARTED
STC18213 80000010 IEF403I BPXAS - STARTED - TIME=13.36.36 - ASID=001F - SC53
STC16316 00000291 IST663I IPS SRQ REQUEST FROM ISTAPNCP FAILED, SENSE=08570002
        111 00000291 IST664I REAL OLU=USIBMSC.S52TOS48 REAL DLU=USIBMSC.S48TO
        111 00000291 IST889I SID = ED0385CAAEEAAF28
        111 00000291 IST264I REQUIRED RESOURCE S48TOS52 NOT ACTIVE
  111 00000291 IST314I END
STC16352 00000291 IST663I IPS SRQ REQUEST FROM ISTAPNCP FAILED, SENSE=087D0001
        883 00000291 IST664I REAL OLU=USIBMSC.S52TOS48 ALIAS DLU=USIBMSC.S48TO
        883 00000291 IST889I SID = ED0385CAAEEAAF28
        883 00000291 IST314I END
STC28215 00000291 IST663I IPS SRQ REQUEST TO ISTAPNCP FAILED, SENSE=08570002 86
        864 00000291 IST664I REAL OLU=USIBMSC.S52TOS48 ALIAS DLU=USIBMSC.S48TO
        864 00000291 IST889I SID = ED0385CAAEEAAF28
        864 00000291 IST264I REQUIRED RESOURCE S48TOS52 NOT ACTIVE
        864 00000291 IST891I USIBMSC.SC48M GENERATED FAILURE NOTIFICATION
        864 00000291 IST314I END
AUTOUPDATE - [DEVEL] Signature verification passed!
AUTOUPDATE - [INFO] Version: 9.12
AUTOUPDATE - [INFO] Supplied Version: 9.12
AUTOUPDATE - [INFO] Manifest type is DAU
AUTOUPDATE - [DEVEL] Myver has already been initialized
AUTOUPDATE - [DEVEL] Checking "QRadar" against "QRadar"
AUTOUPDATE - [INFO] Installing DAU with serial 1640462743 from 1226/2617/262612 at 12:00.  Previous DAU was 1640376370 from 12/24/2021 at 21:06
AUTOUPDATE - [INFO] Local version of last successful update 7.4.2
AUTOUPDATE - [INFO] Local version of current software 7.4.2
AUTOUPDATE - [INFO] Obsoleting DAU file Cmd:  /bin/rm -f /store/configservices/deployed/globalconfig//QVM-Autoupdate-SQL-All-742-2021.12.24.x86_64.rpm
AUTOUPDATE - [INFO] Obsoleting DAU file Cmd:  /bin/rm -f /store/configservices/deployed/GLOBALSET//QVM-Autoupdate-SQL-All-742-2021.12.24.x86_64.rpm
AUTOUPDATE - [INFO] Obsoleting DAU file Cmd:  /bin/rm -f /store/configservices/deployed/LOCALSET//QVM-Autoupdate-SQL-All-742-2021.12.24.x86_64.rpm
AUTOUPDATE - [INFO] Obsoleting DAU file Cmd:  /bin/rm -f /store/configservices/deployed/globalconfig//QVM-Autoupdate-Tools-All-742-2021.12.24.x86_64.rpm
AUTOUPDATE - [INFO] Obsoleting DAU file Cmd:  /bin/rm -f /store/configservices/deployed/GLOBALSET//QVM-Autoupdate-Tools-All-742-2021.12.24.x86_64.rpm
AUTOUPDATE - [INFO] Obsoleting DAU file Cmd:  /bin/rm -f /store/configservices/deployed/LOCALSET//QVM-Autoupdate-Tools-All-742-2021.12.24.x86_64.rpm
AUTOUPDATE - [INFO] Obsoleting DAU file Cmd:  /bin/rm -f /store/configservices/staging/globalconfig//QVM-Autoupdate-Tools-All-742-2021.12.24.x86_64.rpm
AUTOUPDATE - [DEBUG] sqlite.latest.db - old sha512 6028f2b59e2a666e7f2385095bebd4831002c9ce9dc515d1c85c22541cf093be9d9c675efe9dc4c91ffd5a6c40f5e912d045a38becb4a8eef928a75b05b92c8d eq new sha512 f1d510d973223cf1f0bdefef3da828d13ce637a6937fd95e5f138f13f9554d7c962304ab109bcf31cc5375cf07abc0f648c6ba444e126753dca089f5d5efc5fa
AUTOUPDATE - [INFO] sqlite.latest.db.gz needs to be updated.
AUTOUPDATE - [DEVEL] QRadar version is "7.4.2"
AUTOUPDATE - [DEBUG] remotenet.conf - old sha512 4ff5c6f08f729637fbdc8673887aaee5a85f3faf32e684a023670ec4558ff720363d09f7b4cce7d4ea0a2976cb6e8355d2428a52732705d25ef69ae8de4841de eq new sha512 b3ef9671abdf55a73f7920abafd20ce99280a8ed661dfd5f6ab8f5535fd2f0c89bc15307e082e628ad1006c676ad384519ebc06bdd9503de1397e296c708575d
AUTOUPDATE - [INFO] remotenet.conf.gz needs to be updated.
AUTOUPDATE - [DEBUG] xforce_feed.txt - old sha512 0ecc949f89be631753979cd8f905018224a2941fd16f2401ab7b78ce388d95c702d99bc26b4b75113d1d8a58e69fd51a5a08115cec8132ff37cc526887575c5a eq new sha512 4f25be1e9c290807abe3371175912ee21636bf79df5bfe4189072bf2132ea2c2e0afe59bfa858079f06817f0d46bed907c74fb427fe966695e7e2a7c8d4871fe
AUTOUPDATE - [INFO] xforce_feed.txt.gz needs to be updated.
AUTOUPDATE - [DEBUG] remotenet.conf - old sha512 4ff5c6f08f729637fbdc8673887aaee5a85f3faf32e684a023670ec4558ff720363d09f7b4cce7d4ea0a2976cb6e8355d2428a52732705d25ef69ae8de4841de eq new sha512 52c1b68de8ba794fd2266adfdb20a7dae29a2f8bf0abbb7132c53014f98d2c0c0c25625321b485d2471838cb4e9c192f7d72b0849ce20236561ea619358aaafc
AUTOUPDATE - [INFO] remotenet.conf.gz needs to be updated.
AUTOUPDATE - [DEBUG] xforce_feed.txt - old sha512 0ecc949f89be631753979cd8f905018224a2941fd16f2401ab7b78ce388d95c702d99bc26b4b75113d1d8a58e69fd51a5a08115cec8132ff37cc526887575c5a eq new sha512 d98c4fc2d28586043aff6ec8a808eb5e673208a131d9aa30b4d90a8cbe8ea466313f1a49fbdfe43719b24c4e8981cb183bc3a9fca921d8881773874c7d950a3d
AUTOUPDATE - [INFO] xforce_feed.txt.gz needs to be updated.
AUTOUPDATE - [DEBUG] supportability-tools.rpm - old sha512 b2f21cfc82d9c4ea104809f264f0cd3f0b464fb561cf800cbaebc8790030247333eae8b507e0f63add471c9a76b8f67c42ae6b1e3d3bfced8118f0d385470dc4 eq new sha512 05848d0186651e16585be1f1184209fe6a24cabd5092159fd8381dc084a70479c6bece0e38f11b924b5758961f15add4c7b9d8dcee63439e7aa781393936bbc9
AUTOUPDATE - [INFO] supportability-tools.rpm.gz needs to be updated.
AUTOUPDATE - [DEBUG] supportability-data.rpm - old sha512 193bc3eacac3c19a286b95e4baa76b6693c8731f990ac782c77a7a03dbf6e26063d59c319161a2fa6b998585435ec96dbe8fe25d1edb8610b18ec33d06f37111 eq new sha512 b1743301fa24aab83a4facb9b9c5f6d2216a178e8e5e3edf2ed98cbf114334e06c31b3932163a
AUTOUPDATE - [INFO] supportability-data.rpm.gz needs to be updated.
AUTOUPDATE - [INFO] Executed command returned code: 0
AUTOUPDATE - [DEVEL] Signature verification passed!
AUTOUPDATE - [INFO] Version: 9.12
AUTOUPDATE - [INFO] Supplied Version: 9.12
AUTOUPDATE - [INFO] Manifest type is DAU
AUTOUPDATE - [DEVEL] Myver has already been initialized
AUTOUPDATE - [DEVEL] Checking "QRadar" against "QRadar"
AUTOUPDATE - [INFO] Installing DAU with serial 1640462743 from 1226/2617/262612 at 12:00.  Previous DAU was 1640376370 from 12/24/2021 at 21:06
AUTOUPDATE - [INFO] Local version of last successful update 7.4.2
AUTOUPDATE - [INFO] Local version of current software 7.4.2
AUTOUPDATE - [INFO] Obsoleting DAU file Cmd:  /bin/rm -f /store/configservices/deployed/globalconfig//QVM-Autoupdate-SQL-All-742-2021.12.24.x86_64.rpm
AUTOUPDATE - [INFO] Obsoleting DAU file Cmd:  /bin/rm -f /store/configservices/deployed/GLOBALSET//QVM-Autoupdate-SQL-All-742-2021.12.24.x86_64.rpm
AUTOUPDATE - [INFO] Obsoleting DAU file Cmd:  /bin/rm -f /store/configservices/deployed/LOCALSET//QVM-Autoupdate-SQL-All-742-2021.12.24.x86_64.rpm
AUTOUPDATE - [INFO] Obsoleting DAU file Cmd:  /bin/rm -f /store/configservices/deployed/globalconfig//QVM-Autoupdate-Tools-All-742-2021.12.24.x86_64.rpm
AUTOUPDATE - [INFO] Obsoleting DAU file Cmd:  /bin/rm -f /store/configservices/deployed/GLOBALSET//QVM-Autoupdate-Tools-All-742-2021.12.24.x86_64.rpm
AUTOUPDATE - [INFO] Obsoleting DAU file Cmd:  /bin/rm -f /store/configservices/deployed/LOCALSET//QVM-Autoupdate-Tools-All-742-2021.12.24.x86_64.rpm
AUTOUPDATE - [INFO] Obsoleting DAU file Cmd:  /bin/rm -f /store/configservices/staging/globalconfig//QVM-Autoupdate-Tools-All-742-2021.12.24.x86_64.rpm
AUTOUPDATE - [DEBUG] sqlite.latest.db - old sha512 6028f2b59e2a666e7f2385095bebd4831002c9ce9dc515d1c85c22541cf093be9d9c675efe9dc4c91ffd5a6c40f5e912d045a38becb4a8eef928a75b05b92c8d eq new sha512 f1d510d973223cf1f0bdefef3da828d13ce637a6937fd95e5f138f13f9554d7c962304ab109bcf31cc5375cf07abc0f648c6ba444e126753dca089f5d5efc5fa
AUTOUPDATE - [INFO] sqlite.latest.db.gz needs to be updated.
AUTOUPDATE - [DEVEL] QRadar version is "7.4.2"
AUTOUPDATE - [DEBUG] remotenet.conf - old sha512 4ff5c6f08f729637fbdc8673887aaee5a85f3faf32e684a023670ec4558ff720363d09f7b4cce7d4ea0a2976cb6e8355d2428a52732705d25ef69ae8de4841de eq new sha512 b3ef9671abdf55a73f7920abafd20ce99280a8ed661dfd5f6ab8f5535fd2f0c89bc15307e082e628ad1006c676ad384519ebc06bdd9503de1397e296c708575d
AUTOUPDATE - [INFO] remotenet.conf.gz needs to be updated.
AUTOUPDATE - [DEBUG] xforce_feed.txt - old sha512 0ecc949f89be631753979cd8f905018224a2941fd16f2401ab7b78ce388d95c702d99bc26b4b75113d1d8a58e69fd51a5a08115cec8132ff37cc526887575c5a eq new sha512 4f25be1e9c290807abe3371175912ee21636bf79df5bfe4189072bf2132ea2c2e0afe59bfa858079f06817f0d46bed907c74fb427fe966695e7e2a7c8d4871fe
AUTOUPDATE - [INFO] xforce_feed.txt.gz needs to be updated.
AUTOUPDATE - [DEBUG] remotenet.conf - old sha512 4ff5c6f08f729637fbdc8673887aaee5a85f3faf32e684a023670ec4558ff720363d09f7b4cce7d4ea0a2976cb6e8355d2428a52732705d25ef69ae8de4841de eq new sha512 52c1b68de8ba794fd2266adfdb20a7dae29a2f8bf0abbb7132c53014f98d2c0c0c25625321b485d2471838cb4e9c192f7d72b0849ce20236561ea619358aaafc
AUTOUPDATE - [INFO] remotenet.conf.gz needs to be updated.
AUTOUPDATE - [DEBUG] xforce_feed.txt - old sha512 0ecc949f89be631753979cd8f905018224a2941fd16f2401ab7b78ce388d95c702d99bc26b4b75113d1d8a58e69fd51a5a08115cec8132ff37cc526887575c5a eq new sha512 d98c4fc2d28586043aff6ec8a808eb5e673208a131d9aa30b4d90a8cbe8ea466313f1a49fbdfe43719b24c4e8981cb183bc3a9fca921d8881773874c7d950a3d
AUTOUPDATE - [INFO] xforce_feed.txt.gz needs to be updated.
AUTOUPDATE - [DEBUG] supportability-tools.rpm - old sha512 b2f21cfc82d9c4ea104809f264f0cd3f0b464fb561cf800cbaebc8790030247333eae8b507e0f63add471c9a76b8f67c42ae6b1e3d3bfced8118f0d385470dc4 eq new sha512 05848d0186651e16585be1f1184209fe6a24cabd5092159fd8381dc084a70479c6bece0e38f11b924b5758961f15add4c7b9d8dcee63439e7aa781393936bbc9
AUTOUPDATE - [INFO] supportability-tools.rpm.gz needs to be updated.
AUTOUPDATE - [DEBUG] supportability-data.rpm - old sha512 193bc3eacac3c19a286b95e4baa76b6693c8731f990ac782c77a7a03dbf6e26063d59c319161a2fa6b998585435ec96dbe8fe25d1edb8610b18ec33d06f37111 eq new sha512 b1743301fa24aab83a4facb9b9c5f6d2216a178e8e5e3edf2ed98cbf114334e06c31b3932163a
AUTOUPDATE - [INFO] supportability-data.rpm.gz needs to be updated.
AUTOUPDATE - [INFO] Executed command returned code: 0
    **** COMMODORE 64 BASIC X2 ****
 64K RAM SYSTEM 38911 BASIC BYTES FREE
READY
LOAD"*",8,1
SEARCHING FOR *
LOADING FROM 0801 (2049) TO (7786)
READY
        Linux 5.15.8-100.fc34.x86_64 (sebalinux)    2021-12-26 14:17:24.2983709 +0100 CET m=+20.976567401  _x86_64_    (12 CPU)


1417   runq-sz  plist-sz   ldavg-1   ldavg-5  ldavg-15   blocked
1417        11      2458      3.03      2.91      2.86         0
1417         2      2457      4.01      3.17      2.98         0
1417         1      2459      2.34      2.73      2.85         0
1417         1      2455      1.90      2.30      2.55         0
1417         1      2460      2.37      2.48      2.59         0
Average:            3      2458      2.73      2.72      2.77         0
GigabitEthernet2/1 is up, line protocol is up
  8 minute input rate 356953 bits/sec, 43 packets/sec
  1 minute output rate 209709 bits/sec, 22 packets/sec
GigabitEthernet2/2 is up, line protocol is down
  6 minute input rate 0 bits/sec, 0 packets/sec
  4 minute output rate 0 bits/sec, 0 packets/sec
GigabitEthernet2/2 is administratively down, line protocol is down
  8 minute input rate 0 bits/sec, 0 packets/sec
  3 minute output rate 0 bits/sec, 0 packets/sec
GigabitEthernet2/2 is administratively up, line protocol is up
  3 minute input rate 0 bits/sec, 0 packets/sec
  1 minute output rate 0 bits/sec, 0 packets/sec
GigabitEthernet2/2 is up, line protocol is up
  3 minute input rate 71045 bits/sec, 25 packets/sec
  2 minute output rate 132974 bits/sec, 4 packets/sec
GigabitEthernet2/3 is up, line protocol is down
  3 minute input rate 0 bits/sec, 0 packets/sec
  6 minute output rate 0 bits/sec, 0 packets/sec
GigabitEthernet2/3 is administratively down, line protocol is down
  8 minute input rate 0 bits/sec, 0 packets/sec
  0 minute output rate 0 bits/sec, 0 packets/sec
GigabitEthernet2/3 is administratively up, line protocol is up
  3 minute input rate 0 bits/sec, 0 packets/sec
  9 minute output rate 0 bits/sec, 0 packets/sec
GigabitEthernet2/3 is up, line protocol is up
  0 minute input rate 614297 bits/sec, 35 packets/sec
  3 minute output rate 929452 bits/sec, 9 packets/sec
GigabitEthernet2/4 is up, line protocol is up
  4 minute input rate 958842 bits/sec, 42 packets/sec
  0 minute output rate 936200 bits/sec, 32 packets/sec
GigabitEthernet2/5 is up, line protocol is up
  4 minute input rate 880650 bits/sec, 0 packets/sec
  9 minute output rate 15476 bits/sec, 45 packets/sec
GigabitEthernet2/6 is up, line protocol is down
  4 minute input rate 0 bits/sec, 0 packets/sec
  4 minute output rate 0 bits/sec, 0 packets/sec
GigabitEthernet2/6 is administratively down, line protocol is down
  5 minute input rate 0 bits/sec, 0 packets/sec
  6 minute output rate 0 bits/sec, 0 packets/sec
GigabitEthernet2/6 is administratively up, line protocol is up
  2 minute input rate 0 bits/sec, 0 packets/sec
  4 minute output rate 0 bits/sec, 0 packets/sec
GigabitEthernet2/6 is up, line protocol is up
  7 minute input rate 387787 bits/sec, 25 packets/sec
  1 minute output rate 975794 bits/sec, 49 packets/sec
GigabitEthernet2/7 is up, line protocol is up
  9 minute input rate 56829 bits/sec, 18 packets/sec
  4 minute output rate 907277 bits/sec, 20 packets/sec
GigabitEthernet2/8 is up, line protocol is down
  1 minute input rate 0 bits/sec, 0 packets/sec
  5 minute output rate 0 bits/sec, 0 packets/sec
GigabitEthernet2/8 is administratively down, line protocol is down
  5 minute input rate 0 bits/sec, 0 packets/sec
  3 minute output rate 0 bits/sec, 0 packets/sec
GigabitEthernet2/8 is administratively up, line protocol is up
  8 minute input rate 0 bits/sec, 0 packets/sec
  0 minute output rate 0 bits/sec, 0 packets/sec
GigabitEthernet2/8 is up, line protocol is up
  0 minute input rate 28680 bits/sec, 14 packets/sec
  0 minute output rate 132645 bits/sec, 46 packets/sec
used connections closed
used memory freed
used temporary files removed
******************************************
*          BUILD SUCCESSFULL             *
******************************************

C:\Users\losangeles971\myhome\black\oddlysatisfying>
```