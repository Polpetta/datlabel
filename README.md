Docker-label
===

Docker-label is a micro Go library that aims to make the searching, listing and
filtering of labels in Docker containers and services as easy as possible.
On top of that, it offers the possibility to search on Stacks too.

## Build
The library is a standard Go project that uses modules. In order to build it,
open a shell and type:
```shell script
go build github.com/polpetta/docker-label
```

## Testing
The projects has tests too. In order to run them, you need a local and
 working Docker installation. On top of that, your Docker installation needs
  to be set up in Swarm mode. To wrap up, you need to follow this procedure:
```shell script
docker swarm init # This initialize Docker into Swarm configuration
go test -count=1 ./test/ # Launch the actual tests
docker swarm leave -f # Set it back Docker to classic mode
```

## License

This software is licensed under the MIT license.