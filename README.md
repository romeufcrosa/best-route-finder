
# Best Route Finder

## Description

TBD

## Specification

* For more information, please see the postman files at /docs/postman

## Instructions

### Pre-Requisites

***In order to run this locally, the Golang package dependency manager [dep](https://github.com/golang/dep) is required.***

### Running Locally

#### Install dependencies

```bash
make install
```

#### Start container dependencies (MySQL)

```bash
make local-env
```

#### Start the API

```bash
make local-run
```

>Please note that this cleans and then bootstraps the database with some data so that the user already has something to work with

#### Run Tests

```bash
make test
```

>Please note that running the tests will clean the local DB beforehand

### Running with Docker

#### Create the container

```bash
make docker-build
```

#### Start the container

```bash
make docker-local-run
```

It will then be exposed in **localhost:8080**. So any requests would work the same as they would in local run mode.

#### Run tests in container

```bash
make docker-tests
```

### Using

#### Postman

In folder `docs/postman` there is a JSON file of a Postman collection that contains all the possible requests the API currently supports.
Just import it to your Postman and it will create a *RouteFinder* collection.
