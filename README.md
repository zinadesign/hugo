![Hugo](https://raw.githubusercontent.com/spf13/hugo/master/docs/static/img/hugo-logo.png)

A Fast and Flexible Static Site Generator built with love

This is a fork. Official repository located at [github](https://github.com/zinadesign/hugo)

### Build and Install the Binaries from Source (Advanced Install)


Add Hugo and its package dependencies to your go `src` directory.

    go get -v github.com/zinadesign/hugo && go get -u -v github.com/kardianos/govendor && cd $GOPATH/src/github.com/zinadesign/hugo && go install

Once the `get` completes, you should find your new `hugo` (or `hugo.exe`) executable sitting inside `$GOPATH/bin/`.

To update Hugo’s dependencies, use `go get` with the `-u` option.

    go get -u -v github.com/zinadesign/hugo
