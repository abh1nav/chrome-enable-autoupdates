# chrome-enable-autoupdates

Golang version of [hjuutilainen/adminscripts/chrome-enable-autoupdates](https://github.com/hjuutilainen/adminscripts/blob/master/chrome-enable-autoupdates.py).
  
This binary enables system wide automatic updates for Google Chrome. It should work for most recent versions of Chrome. No configuration needed
as this is originally intended as a munki postinstall script. This is built as a binary (instead of a Python script) since MacOS 12.3 removes python as part of the core OS.

Download the latest release from the [releases](https://github.com/abh1nav/chrome-enable-autoupdates/releases) page and run it.