
# Best Route Finder

## Description

This API finds the optimal route (cheapest and fastest) between an origin and a destination point.
These points are identified by ID and Name, although right now the API only uses ID as the parameter for origin and destination in the GetRoute endpoint.

The system is represented by Nodes and Edges, in a directed Graph structure.
Each edge has two weights, identified by time (that it takes to travel between them) and cost (i.e. shipping cost).

## Specification

The API Specification is detailed in OpenAPI 3.0.0, the file can be found in `/docs/spec/`.

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

### TODO List
Due to time constraints this first iteration is lacking in features and requiring some technical improvements.

- [ ] Add more tests, current ones are mostly aiming for happy path.
- [ ] DRY it up, there's probably some methods that are either very similar or identical. 
- [ ] SOLIDify it more, some methods are doing a bit more than I want them to... some extraction and refactoring is required.
- [ ] Find a better bootstrap of the database for tests without compromising the data for local run and vice-versa.
- [ ] Add constraints on the database structure so that edges/nodes can't be duplicated.
- [ ] Complete the CRUD.
