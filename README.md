# Useless but oddly satisfying program to build nothing!

Hi all!

This project focuses on the importance of positive emotions and relaxing moments for stressed developers and sysadmins.

When you are waiting for building your app or for installing a complex software, and you feel stress, bad or weird mood, then you may use the _oddlysatisfying_ program.

The _oddlysatisfying_ program is a completely useless script for building nothig! But having oddly satisfying log messages till to an always positive ending.

The _oddlysatisfying_ program randomly choose vintage and modern operations, typing nice logging on your screen adn at the end, you always receive a positive "BUILD SUCCESSFUL" final message.

For the best effect, you may have two monitor and use the far one for your boring and stressful activity amd the close one to run the _oddlysatisfying_ program.

And you may contribute to this project providing ideas, nice log examples and code, in order to keep _oddlysatisfying_ program completely useless but oddly satysfying and relaxing.

To download the superb and useless _oddlysatisfying_ program click on the release tags and then choose the last (or wathever you want) release.
Chosen the release, you may download the right binary for your machine.

If you're lazy, then just launch the _oddlysatisfying_ program.
Otherwise, you may launch  _oddlysatisfying_ program specifying a number, the latter will be used to randomly generate the number of programs to be executed.

Anyway remember! Even if you see attempt to download something, recovery actions, boot of programs, **the superb _oddlysatisfying_ program never really do something**.
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

If you would add a new _logging func_, you just need to chose an existing file or creating a new one, and populate it with new code.

## Conduct guidelines

* if possible, you should make your _logging func_ dynamic not static, it means generating at the time name of servers, files, dates, etc.
* _logging func_ may include failure, recovery actions, loss of data
* if you modify the way the program ends, keep in mind that the execution must always end with the calling of one of the possible _closing func_
* all the possible _closing func_ must always end with a positive, relaxing, peacefull and successfull final message
* the code cannot contain vulgar, aggressive or discriminatory language

## Few words more...

At the date, the superb _oddlysatisfying_ program supports the following architectures:

* Linux/amd64
* Windows/amd64
* Darwin/amd64

But you asks and (maybe) you'll be satisfied.


Enjoy!
Peace and love, LosAngeles971.