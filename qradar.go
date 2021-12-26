package main

import (
	"time"
)

func qradarAutoupdate() {
	log := New()
	log.Info("AUTOUPDATE - [DEVEL] Signature verification passed!")
	log.Info("AUTOUPDATE - [INFO] Version: 9.12")
	log.Info("AUTOUPDATE - [INFO] Supplied Version: 9.12")
	log.Info("AUTOUPDATE - [INFO] Manifest type is DAU")
	log.Info("AUTOUPDATE - [DEVEL] Myver has already been initialized")
	log.Info("AUTOUPDATE - [DEVEL] Checking \"QRadar\" against \"QRadar\"")
	log.Infof("AUTOUPDATE - [INFO] Installing DAU with serial 1640462743 from %s at 12:00.  Previous DAU was 1640376370 from 12/24/2021 at 21:06", time.Now().Format("12/24/2021"))
	log.Info("AUTOUPDATE - [INFO] Local version of last successful update 7.4.2")
	log.Info("AUTOUPDATE - [INFO] Local version of current software 7.4.2")
	log.Info("AUTOUPDATE - [INFO] Obsoleting DAU file Cmd:  /bin/rm -f /store/configservices/deployed/globalconfig//QVM-Autoupdate-SQL-All-742-2021.12.24.x86_64.rpm")
	log.Info("AUTOUPDATE - [INFO] Obsoleting DAU file Cmd:  /bin/rm -f /store/configservices/deployed/GLOBALSET//QVM-Autoupdate-SQL-All-742-2021.12.24.x86_64.rpm")
	log.Info("AUTOUPDATE - [INFO] Obsoleting DAU file Cmd:  /bin/rm -f /store/configservices/deployed/LOCALSET//QVM-Autoupdate-SQL-All-742-2021.12.24.x86_64.rpm")
	log.Info("AUTOUPDATE - [INFO] Obsoleting DAU file Cmd:  /bin/rm -f /store/configservices/deployed/globalconfig//QVM-Autoupdate-Tools-All-742-2021.12.24.x86_64.rpm")
	log.Info("AUTOUPDATE - [INFO] Obsoleting DAU file Cmd:  /bin/rm -f /store/configservices/deployed/GLOBALSET//QVM-Autoupdate-Tools-All-742-2021.12.24.x86_64.rpm")
	log.Info("AUTOUPDATE - [INFO] Obsoleting DAU file Cmd:  /bin/rm -f /store/configservices/deployed/LOCALSET//QVM-Autoupdate-Tools-All-742-2021.12.24.x86_64.rpm")
	log.Info("AUTOUPDATE - [INFO] Obsoleting DAU file Cmd:  /bin/rm -f /store/configservices/staging/globalconfig//QVM-Autoupdate-Tools-All-742-2021.12.24.x86_64.rpm")
	log.Info("AUTOUPDATE - [DEBUG] sqlite.latest.db - old sha512 6028f2b59e2a666e7f2385095bebd4831002c9ce9dc515d1c85c22541cf093be9d9c675efe9dc4c91ffd5a6c40f5e912d045a38becb4a8eef928a75b05b92c8d eq new sha512 f1d510d973223cf1f0bdefef3da828d13ce637a6937fd95e5f138f13f9554d7c962304ab109bcf31cc5375cf07abc0f648c6ba444e126753dca089f5d5efc5fa")
	log.Info("AUTOUPDATE - [INFO] sqlite.latest.db.gz needs to be updated.")
	log.Info("AUTOUPDATE - [DEVEL] QRadar version is \"7.4.2\"")
	log.Info("AUTOUPDATE - [DEBUG] remotenet.conf - old sha512 4ff5c6f08f729637fbdc8673887aaee5a85f3faf32e684a023670ec4558ff720363d09f7b4cce7d4ea0a2976cb6e8355d2428a52732705d25ef69ae8de4841de eq new sha512 b3ef9671abdf55a73f7920abafd20ce99280a8ed661dfd5f6ab8f5535fd2f0c89bc15307e082e628ad1006c676ad384519ebc06bdd9503de1397e296c708575d")
	log.Info("AUTOUPDATE - [INFO] remotenet.conf.gz needs to be updated.")
	log.Info("AUTOUPDATE - [DEBUG] xforce_feed.txt - old sha512 0ecc949f89be631753979cd8f905018224a2941fd16f2401ab7b78ce388d95c702d99bc26b4b75113d1d8a58e69fd51a5a08115cec8132ff37cc526887575c5a eq new sha512 4f25be1e9c290807abe3371175912ee21636bf79df5bfe4189072bf2132ea2c2e0afe59bfa858079f06817f0d46bed907c74fb427fe966695e7e2a7c8d4871fe")
	log.Info("AUTOUPDATE - [INFO] xforce_feed.txt.gz needs to be updated.")
	log.Info("AUTOUPDATE - [DEBUG] remotenet.conf - old sha512 4ff5c6f08f729637fbdc8673887aaee5a85f3faf32e684a023670ec4558ff720363d09f7b4cce7d4ea0a2976cb6e8355d2428a52732705d25ef69ae8de4841de eq new sha512 52c1b68de8ba794fd2266adfdb20a7dae29a2f8bf0abbb7132c53014f98d2c0c0c25625321b485d2471838cb4e9c192f7d72b0849ce20236561ea619358aaafc")
	log.Info("AUTOUPDATE - [INFO] remotenet.conf.gz needs to be updated.")
	log.Info("AUTOUPDATE - [DEBUG] xforce_feed.txt - old sha512 0ecc949f89be631753979cd8f905018224a2941fd16f2401ab7b78ce388d95c702d99bc26b4b75113d1d8a58e69fd51a5a08115cec8132ff37cc526887575c5a eq new sha512 d98c4fc2d28586043aff6ec8a808eb5e673208a131d9aa30b4d90a8cbe8ea466313f1a49fbdfe43719b24c4e8981cb183bc3a9fca921d8881773874c7d950a3d")
	log.Info("AUTOUPDATE - [INFO] xforce_feed.txt.gz needs to be updated.")
	log.Info("AUTOUPDATE - [DEBUG] supportability-tools.rpm - old sha512 b2f21cfc82d9c4ea104809f264f0cd3f0b464fb561cf800cbaebc8790030247333eae8b507e0f63add471c9a76b8f67c42ae6b1e3d3bfced8118f0d385470dc4 eq new sha512 05848d0186651e16585be1f1184209fe6a24cabd5092159fd8381dc084a70479c6bece0e38f11b924b5758961f15add4c7b9d8dcee63439e7aa781393936bbc9")
	log.Info("AUTOUPDATE - [INFO] supportability-tools.rpm.gz needs to be updated.")
	log.Info("AUTOUPDATE - [DEBUG] supportability-data.rpm - old sha512 193bc3eacac3c19a286b95e4baa76b6693c8731f990ac782c77a7a03dbf6e26063d59c319161a2fa6b998585435ec96dbe8fe25d1edb8610b18ec33d06f37111 eq new sha512 b1743301fa24aab83a4facb9b9c5f6d2216a178e8e5e3edf2ed98cbf114334e06c31b3932163a")
	log.Info("AUTOUPDATE - [INFO] supportability-data.rpm.gz needs to be updated.")
	log.Info("AUTOUPDATE - [INFO] Executed command returned code: 0 ")
}


