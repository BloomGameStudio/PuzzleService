# PuzzleService

The Bloom & LostLight PuzzleService.

- [PuzzleService](#puzzleservice)
  - [Installation](#installation)
    - [**Native Bare Metel**](#native-bare-metel)
  - [Usage](#usage)
  - [Live reloading with Air, installation and usage](#live-reloading-with-air-installation-and-usage)
  - [Installation](#installation)
    - [Via `go install` (Recommended)](#via-go-install-recommended)
    - [Via install.sh](#via-installsh)
    - [Via goblin.run](#via-goblinrun)
    - [Docker/Podman](#dockerpodman)
      - [Docker/Podman .${SHELL}rc](#dockerpodman-shellrc)
  - [Usage](#usage)
    - [Runtime arguments](#runtime-arguments)
    - [Docker-compose](#docker-compose)
    - [Debug](#debug)
  - [Installation and Usage for Docker users who don't want to use air image](#installation-and-usage-for-docker-users-who-dont-want-to-use-air-image)

## Installation

### **Native Bare Metel**

    // Install Deps/Modules
    go mod download

## Usage

    // Run the server on port 1323
    go run .

## Live reloading with Air, installation and usage

[Official Air documentation](https://github.com/cosmtrek/air)

## Installation

### Via `go install` (Recommended)

With go 1.22 or higher:

```bash
go install github.com/cosmtrek/air@latest
```

### Via install.sh

```bash
# binary will be $(go env GOPATH)/bin/air
curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

# or install it into ./bin/
curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s

air -v
```

### Via [goblin.run](https://goblin.run)

```sh
# binary will be /usr/local/bin/air
curl -sSfL https://goblin.run/github.com/cosmtrek/air | sh

# to put to a custom path
curl -sSfL https://goblin.run/github.com/cosmtrek/air | PREFIX=/tmp sh
```

### Docker/Podman

Please pull this docker image [cosmtrek/air](https://hub.docker.com/r/cosmtrek/air).

```bash
docker/podman run -it --rm \
    -w "<PROJECT>" \
    -e "air_wd=<PROJECT>" \
    -v $(pwd):<PROJECT> \
    -p <PORT>:<APP SERVER PORT> \
    cosmtrek/air
    -c <CONF>
```

#### Docker/Podman .${SHELL}rc

if you want to use air continuously like a normal app, you can create a function in your ${SHELL}rc (bash,zsh,etc...)

```bash
air() {
  podman/docker run -it --rm \
    -w "$PWD" -v "$PWD":"$PWD" \
    -p "$AIR_PORT":"$AIR_PORT" \
    docker.io/cosmtrek/air "$@"
}
```

`<PROJECT>` is your project path in container, eg: /go/example
if you want to enter the container, Please add --entrypoint=bash.

<details>
  <summary>For example</summary>

One of my project runs in docker:

```bash
docker run -it --rm \
  -w "/go/src/github.com/cosmtrek/hub" \
  -v $(pwd):/go/src/github.com/cosmtrek/hub \
  -p 9090:9090 \
  cosmtrek/air
```

Another example:

```bash
cd /go/src/github.com/cosmtrek/hub
AIR_PORT=8080 air -c "config.toml"
```

this will replace `$PWD` with the current directory, `$AIR_PORT` is the port where to publish and `$@` is to accept arguments of the application itself for example -c

</details>

## Usage

For less typing, you could add `alias air='~/.air'` to your `.bashrc` or `.zshrc`.

First enter into your project

```bash
cd /path/to/your_project
```

The simplest usage is run

```bash
# firstly find `.air.toml` in current directory, if not found, use defaults
air -c .air.toml
```

You can initialize the `.air.toml` configuration file to the current directory with the default settings running the following command.

```bash
air init
```

After this, you can just run the `air` command without additional arguments and it will use the `.air.toml` file for configuration.

```bash
air
```

For modifying the configuration refer to the [air_example.toml](air_example.toml) file.

### Runtime arguments

You can pass arguments for running the built binary by adding them after the air command.

```bash
# Will run ./tmp/main bench
air bench

# Will run ./tmp/main server --port 8080
air server --port 8080
```

You can separate the arguments passed for the air command and the built binary with `--` argument.

```bash
# Will run ./tmp/main -h
air -- -h

# Will run air with custom config and pass -h argument to the built binary
air -c .air.toml -- -h
```

### Docker-compose

```yaml
services:
  my-project-with-air:
    image: cosmtrek/air
    # working_dir value has to be the same of mapped volume
    working_dir: /project-package
    ports:
      - <any>:<any>
    environment:
      - ENV_A=${ENV_A}
      - ENV_B=${ENV_B}
      - ENV_C=${ENV_C}
    volumes:
      - ./project-relative-path/:/project-package/
```

### Debug

`air -d` prints all logs.

## Installation and Usage for Docker users who don't want to use air image

`Dockerfile`

```Dockerfile
# Choose whatever you want, version >= 1.16
FROM golang:1.22-alpine

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

COPY go.mod go.sum ./
RUN go mod download

CMD ["air", "-c", ".air.toml"]
```

`docker-compose.yaml`

```yaml
version: "3.8"
services:
  web:
    build:
      context: .
      # Correct the path to your Dockerfile
      dockerfile: Dockerfile
    ports:
      - 8080:3000
    # Important to bind/mount your codebase dir to /app dir for live reload
    volumes:
      - ./:/app
```
