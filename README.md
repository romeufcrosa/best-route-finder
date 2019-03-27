
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

#### Run Tests

```bash
make test
```

>Please note that running the tests will clean the local DB beforehand

#### Start the API

```bash
make local-run
```

>Please note that this cleans and then bootstraps the database with some data so that the user already has something to work with

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

>Please note that like with the local test run, this is a destructive operation on the database
>So if you intend to boostrap it again with the default data please run the next command

#### Bootstrapping MySQL container with data

```bash
make docker-boostrap
```

### Using

#### Postman

In folder `docs/postman` there is a JSON file of a Postman collection that contains all the possible requests the API currently supports.
Just import it to your Postman and it will create a *RouteFinder* collection.

#### cURL

##### Add Node

```bash
curl -X POST \
  http://localhost:8080/api/v1/nodes \
  -H 'Content-Type: application/json' \
  -H 'cache-control: no-cache' \
  -d '{"name": "Node_A"}'
```

##### Add Edge

```bash
curl -X POST \
  http://localhost:8080/api/v1/edges \
  -H 'Content-Type: application/json' \
  -H 'cache-control: no-cache' \
  -d '{"from_id": 1,"to_id": 2,"cost": 12,"duration": 1}'
```

##### Get Best Route from Origin to Destination

```bash
curl -X GET \
  http://localhost:8080/api/v1/routes/from/1/to/2 \
  -H 'Content-Type: application/json' \
  -H 'cache-control: no-cache'
```

### TODO List

Due to time constraints this first iteration is lacking in features and requiring some technical improvements.

- [ ] Use the actual queue implementation which has concurrency safety instead of the current solution with a makeshift queue using maps.
- [ ] Add more tests, current ones are mostly aiming for happy path.
- [ ] DRY it up, there's probably some methods that are either very similar or identical.
- [ ] SOLIDify it more, some methods are doing a bit more than I want them to... some extraction and refactoring is required.
- [ ] Find a better bootstrap of the database for tests without compromising the data for local run and vice-versa.
- [ ] Add constraints on the database structure so that edges/nodes can't be duplicated.
- [ ] Complete the CRUD.
