Datlabel
===

[![Build Status](https://travis-ci.org/Polpetta/datlabel.svg?branch=master)](https://travis-ci.org/Polpetta/datlabel)

Datlabel is a micro Go library that aims to make the searching, listing and
filtering of labels in Docker containers and services as easy as possible.
On top of that, it offers the possibility to search on Stacks too.

## Build
The library is a standard Go project that uses modules. In order to build it,
open a shell and type:
```shell script
go build github.com/polpetta/datlabel
```

## Testing
The project has tests too. In order to run them, you need a local and
working Docker installation. On top of that, your Docker installation needs
to be set up in Swarm mode. To wrap up, you need to follow this procedure:
```shell script
docker swarm init # This initialize Docker into Swarm configuration
go test -tags=unit -count=1 ./... # Launch unit tests
go test -tags=intergation -count=1 ./... # Launch integration tests
docker swarm leave -f # Set it back Docker to classic mode
```

Alternatively, you can execute all the tests in one shot with:
```shell script
go test -tags="unit integration" -count=1 ./...
```

## License

This software is licensed under the MIT license.
