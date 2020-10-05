# Dependency Inversion Principle Example Application

![golangci-lint](https://github.com/idexter/dip-visa-app/workflows/golangci-lint/badge.svg)
![build](https://github.com/idexter/dip-visa-app/workflows/build/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/idexter/dip-visa-app)](https://goreportcard.com/report/github.com/idexter/dip-visa-app)
[![codecov](https://codecov.io/gh/idexter/dip-visa-app/branch/master/graph/badge.svg)](https://codecov.io/gh/idexter/dip-visa-app)

This application provides examples of DIP and how it helps to make application design better.

## Repository Structure

There are 3 branches with a bit different source code but which implements the same application.

Those branches are:
- [classic-approach](https://github.com/idexter/dip-visa-app/tree/classic-approach)
- [layers-separation](https://github.com/idexter/dip-visa-app/tree/layers-separation)
- [dip-principle (master)](https://github.com/idexter/dip-visa-app/tree/master)

Every version has business-logic within `CheckConfirmation(applicationID int)` function/method.
Don't focus on business-logic of this method. It's simple and made up from imagination. 
It should not prevent you from understanding code structure, but if you are interested, I described an idea
of application in the "Business Logic" section below.

### classic-approach

In "classic-approach" an application implemented as a bunch of functions which calls each other
to provide needed behaviour. There is no DI and no Layer Separation. Everything implemented
within one package and high-level details depends on low-level modules.

### layers-separation

In "layers-separation" branch low-level modules separated from high-level modules using packages.
At the same time high-level details also depends on low-level modules. There are 3 packages:
- `visa` - This package implements our business logic (it depends on 2 packages below).
- `report` - This package implements results storage.
- `storage` - This package implements file storage.

### dip-principle (master)

In "dip-principle (master)" we have final implementation of our application, but this time
we provide Dependency Injection, using Dependency Inversion Principle. So in this case
high-level modules don't depend on low-level modules.
This implementation also shows how we can connect our business logic with CLI and Web interface,
without changing logic itself. So it this case there are 2 binaries.

## Business Logic

We are tourist, who wants to get VISA to go abroad. 
Our application provides "Visa Application Service".
Using this service we can check will be the VISA approved.

- Visa Applications stored in `./data/applications.json`. It has some parameters.
- Old VISAs are stored in `./data/visas.json`. We need them for service business logic.

## Build

Run `make lint` to check if code-style locally.
Run `make build`. It will create 2 binaries inside project directory.
Run `make clean` to remove build artifacts.

- `dip-cli` - Command Line Interface Application
- `dip-srv` - HTTP Server Application

## Test

Run `make test` within project directory.

## Run

For CLI version:

```bash
$ ./dip-cli --id 1
$ ./dip-cli --id 0
```

For HTTP Server version:

```bash
$ ./dip-srv
$ curl -v -X GET "http://localhost:8080/?id=1"
$ curl -v -X GET "http://localhost:8080/?id=3"
```
