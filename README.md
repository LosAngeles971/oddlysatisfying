# The _oddly satisfying_ program!

Hi all!

This project focuses on the importance of positive emotions and relaxing moments for stressed developers, sysadmins and whoever has to do with computers.

If you are struggling for building your app or configuring a complex software, or your mega-excel is not working anymore, whatever activity you are doing on your computer, if you feel stress or weird mood, then you may use the _oddly satisfying_ program.

The _oddlysatisfying_ program is a completely useless script for building nothig! But having oddly satisfying log messages till to an always positive ending.

This program randomly choose vintage and modern operations, typing nice logging on your screen and eventually an always positive "BUILD SUCCESSFUL" closing message.

For the best effect, you may have two monitors and use the far one for your boring and stressful activities, and the close one to run the _oddly satisfying_ program (maximizing the window!).

And you may contribute to the project providing ideas, nice log examples and code, in order to keep _oddly satisfying_ program completely useless but relaxing, and beyond that this project is an archive of historical logs too (a bootstrap of a VAX system for example).

## The CLI version

Downloading the superb and useless _oddly satisfying_ program is easy, click on the release tags and choose the last (or wathever you want) release.
Chosen the release, you may download the right binary for your machine.

If you're lazy just run the _oddly satisfying_ program.

```sh
./oddlysatisfying
```

Otherwise you may specify the number of logging operations you desire.

```sh
./oddlysatisfying 10
```

Anyway remember: even if you see attempts of downloading, recovery actions, boot of programs, **the superb _oddly satisfying_ program never really do something**, it is and it will be always completely useless and harmless.

## The Web version

The superb and useless _oddly satisfying_ program is available on Web too, the best for lazy and stressed people.
Just point your browser to https://web.bluepulsar971.it/oddlysatisfying and enjoy.

Please keep in mind that the web server is just a Raspberry PI, so be nice, patient ane peaceful with him.

## Your Web server

If you want to run _oddly satisfying_ as web server, just do that:

```sh
mkdir oddlysatisfying
cd oddlysatisfying
# copy the binary here
# copy the html directory here
./oddlysatisfying server
```

## How to contribute

The program is really simple and you may use it to make practice with Golang too.

Here the files: 

* _main.go_ -> includes the list of all available _logging func_ and the _main func_
* _console.go_ -> provides the helper to write log on screen
* _data.go_ -> provides some data to make the logging more realistic (a list of possible file names, a list of possible disk devices, etc.)
* _closing.go_ -> provides the list of possible _func_ to (positively) ends the execution

Beyong above files, there are many files including the logging functions (each file per topic), for example: 

* _mainframe.go_ -> provides _logging func_ related to the world of mainframe computers
* _linux.go_ -> provides _logging func_ related to the world of linux computers
* _network.go_ -> provides _logging func_ related to the world of network devices and actions
* _pc.go_ -> provides _logging func_ related to the world of personal computers

You may add new _logging func_ changing an existing file or creating a new one (if you are introducing a new topic).

Last but not least, when you add a new _logging func_, you need to update the _vtp_ map into main.go.

## Conduct guidelines

* if possible, you should make your _logging func_ dynamic (not static), it means generating at the time of execution the name of servers, files, dates, etc
* _logging func_ may also include failures, recovery actions, loss of data
* if you add new _closing_ functions, keep in mind that all _closing_ functions must always end with a positive, relaxing, peacefull and successfull final message
* the code cannot contain vulgar, aggressive or discriminatory language
* **please, you may comment your logging functions, to describe their purpose and historical value**

## Few more words...

At the date, the superb _oddlysatisfying_ program supports the following architectures:

* Linux/amd64 (tested)
* Windows/amd64 (tested)
* Darwin/amd64 (not tested: if you positively or negatively tried this platform, please open an issue to say it)
* Linux/arm64
* Linux/arm

You may ask for additional architecture and (maybe, at best effort) you'll be satisfied.

Enjoy!
Peace and love.

@LosAngeles971 (aka @Bluepulsar971) - https://www.linkedin.com/in/bluepulsar971/

## Credits and references

http://aleclownes.com/2017/02/01/crt-display.html
https://dev.to/ekeijl/retro-crt-terminal-screen-in-css-js-4afh
https://safi.me.uk/typewriterjs/