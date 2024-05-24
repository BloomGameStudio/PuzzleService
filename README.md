# PuzzleService

The Bloom & LostLight PuzzleService.

- [PuzzleService](#puzzleservice)
  - [Architecture](#architecture)
  - [Installation](#installation)
    - [**Native Bare Metel**](#native-bare-metel)
  - [Usage](#usage)
    - [**Using Docker \&\& DockerCompose**](#using-docker--dockercompose)
    - [**Docker-Compose**](#docker-compose)
    - [**Start the App and listen on port 1323**](#start-the-app-and-listen-on-port-1323)
    - [**Docker**](#docker)
  - [Contributors Guide](#contributors-guide)
  - [Live reloading with Air, installation and usage](#live-reloading-with-air-installation-and-usage)
  - [Air Installation](#air-installation)
    - [Via `go install` (Recommended)](#via-go-install-recommended)
    - [Via install.sh](#via-installsh)
    - [Via goblin.run](#via-goblinrun)
    - [Air Docker/Podman](#air-dockerpodman)
      - [Docker/Podman .${SHELL}rc](#dockerpodman-shellrc)
  - [Air Usage](#air-usage)
    - [Air Runtime arguments](#air-runtime-arguments)
    - [Docker-compose](#docker-compose-1)
    - [Air Debug](#air-debug)
  - [Installation and Usage for Docker users who don't want to use air image](#installation-and-usage-for-docker-users-who-dont-want-to-use-air-image)
  - [API Endpoints](#api-endpoints)
    - [Rest Base Endpoints](#rest-base-endpoints)
    - [Rest Endpoints](#rest-endpoints)
        - [VerifyPuzzle](#verifypuzzle)

## Architecture

The API architecture is designed using "clean architecture" principles. Please ensure your contributions adhere to these principles.

Read more about it here: https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html

## Installation

### **Native Bare Metel**

    // Install Deps/Modules
    go mod download

## Usage

    // Run the server on port 1323
    go run .

### **Using Docker && DockerCompose**
> Note: Docker and Docker Compose are currently not implemented

---

### **Docker-Compose**

Chose your docker compose cli
Depending on what version you have or how you installed docker compose.

The examples will use the more wider used `docker-compose`

For more Information read: https://stackoverflow.com/questions/66514436/difference-between-docker-compose-and-docker-compose

The more wider used `docker-compose`.

        docker-compose <command>

The newer `docker compose`.

        docker compose <command>

### **Start the App and listen on port 1323**

Note: Depending on your system and context you may have to configure your image & container versions
view the official Docker Compose documentation on how Docker determines what and how it runs images & containers and how <docker-compose up> behaves.

TLDR: Docker and by extension Docker Compose will chose the latest container and if that does not exist the latest image to run your application.

**Run the latest version that was build from the branch main.**

Note: If you have build a later version or somehow else have a later version on your system a version that docker thinks is later than what was build from main it will most likely use that.
Which resulst in you not runing the version from main and not runing the intended version.
This will automatically be resolved for you if a new push to main happens.

        docker-compose up
        // CTRL + C to stop

**Run & Build the current state of the currently checkout out branch.**

Note: This will build a image and run and build a container which probably is a later version than the prebuild image built from the main branch.

        docker-compose up --build
        // CTRL + C to stop

        // If there are caching issues or some other problems or you want to be 100% sure that you run and have build the latest version of the current branch you can run:
        docker-compose up --build --force-recreate
        // This will recreate everything and might take longer.

---

### **Docker**

1.  Create a The docker volume for the database

        docker volume create puzzleservicevolume

2.  Run The Container

        // From Github Container Registry via Image
        // You can Replace the tag <main> at the end with whatever tag you want

                docker run --rm -p 1323:1323 -v puzzleservicevolume:/database ghcr.io/bloomgamestudio/puzzleservice:main

        // Build it yourself locally with build tag/name then run it

                docker build -t puzzleservice .

                docker run -p 1323:1323 -v puzzleservicevolume:/database puzzleservice


        // Build it yourself locally without tag/name then run it

                docker build .

                docker run -p 1323:1323 -v puzzleservicevolume:/database <Containername>

---

## Contributors Guide

1. Assign a Issue\*(Github Issue) to yourself and or clearly indicate to at least Balu if not everyone else involved that that you are now working on this task. If you lack permissions to assign a issue to yourself contact Balu or any other person with the needed permissions to assign the issue to you.
2. Continuesly Push your work to Github so that other people can follow the progress passivly. Ask for Help if you are stuck Notify Balu if you cant complete it. Ghosting is 0/10.
3. Test everything and make sure everything works as intended
4. Open a Pull Request. Ask Balu if you lack permissions.
5. Address comments on the PR if there are any.
6. Shake hands firmly.

Contributors shall work on delivering a complete contribution from start to end.

Contributing half finished and untested things is not ideal.

---

## Live reloading with Air, installation and usage

[Official Air documentation](https://github.com/cosmtrek/air)

## Air Installation

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

### Air Docker/Podman

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

## Air Usage

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

### Air Runtime arguments

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

### Air Debug

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

---

## API Endpoints

Sending partial Data is Accepted and works on most Endpoints.
Indications will be made if partial Data is not supported for certain objects or Endpoints.

Sending Incorrect or Malformed JSON will always result in failure for the entire request.

### Rest Base Endpoints

Local Base Endpoint with Default Config:

- http://127.0.0.1:3000

Staging Base Endpoint:

- http://staging.puzzle.bloomstudio.gg
>*Note: This is not yet live

Full Example:

- http://127.0.0.1:1323/verify

### Rest Endpoints

##### VerifyPuzzle

`POST /verify`

Verifies a puzzle.

Expects a puzzle and a solution in that order.

**Headers:** None

**Request Body:**

Expects one of the two:
>*Note: two is currently unimplemented.


1.
    | Name     | Type   | Mandatory | Info              |
    |----------|--------|-----------|-------------------|
    | PuzzleID | INT    | YES       | Has to be unique. |
    | Solution | STRING | YES       |                   |


1. A JSON serilized Puzzle [publicModel](./publicModels/puzzle.go) or a [model](./models/puzzle.go) and a solution [solution](./publicModels/puzzle.go) or a [solution](./models/puzzle.go) object in the body.
    >*Note: Models are placeholders atm.


**Request Body Example With all required Fields:**

```json
{
  "PuzzleID": "33b7e1f3-6f8e-40b9-97dc-c54d9162vb05",
  "Solution": "Solution string or byte array",
}
```

**Response:**

```json
200 OK
```
