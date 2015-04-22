# webit
Simple command to load a static file server pointing at a directory. 

This is heavily inspired by @SimbCo's [httpster project](https://github.com/SimbCo/httpster).

Installation is done via `go get`, which is installed along with Go on most platforms.

    go get github.com/cfchris/webit && go install github.com/cfchris/webit

Then from any directory where you want to have an http service running, just run

    $GOBIN/webit

That will start up a web server on port 3333 and let you serve any static content you wish. If you want to change the port or directory that the server runs from pass in the -p or -d options

    $GOBIN/webit -p 8080 -d /home/someuser/somedir
