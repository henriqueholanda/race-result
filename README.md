# Race Result

This project has the objective of read a race file (`race_log.txt`) and generate a final classification
ordered by position and generate some metrics about that.

## Getting started

### Requirements

To run this application you must have installed this following tools:

* [Docker](https://docs.docker.com/engine/installation/)
* [Docker Compose](https://docs.docker.com/compose/install/)
* [Git](https://git-scm.com/)
* [Golang](https://golang.org/doc/install)*

>Note: *Golang is required only to generate the coverage of application.

### Executing the application

To execute this application you just need to follow the steps bellow:

1 - Clone project repository:
```bash
$ git clone https://github.com/henriqueholanda/race-result.git
```

2 - Building the application
```bash
$ make build
```
> Note: This command may spend some time to complete, mainly when you run for the first
time, because it will download all Docker images that project needs from [Docker Store](https://store.docker.com).

3 - Run application
```bash
$ make run
```
> Note: It will show a log of execution that you can follow the running steps

### Project output

To see the project output you need to open the file `race_result.csv` in a editor of your preference.

### Run tests

To run test you just need to run the following command:

```bash
$ make test
```
> Note: You must need to build project before it.

### Generating coverage

To generate coverage you just need to run the following command:

```bash
$ make coverage
```
> Note: You must need to have `Golang` installed on your machine and build project before it.

## Author

[Henrique Holanda](https://henriqueholanda.dev) 

## The challenge

You can see the challenge requirements in this file: [CHALLENGE](CHALLENGE.md)