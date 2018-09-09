# CLI

## Start development environment (docker + docker-compose)

* rok4go-bin: container watching gofile and building binary using dev mode (Assets are not bundled in a gofile)
* rok4go-ui: container watching and building ui using webpack
* mongo: mongo... ;)
* mongo-express: ui for mongo

```
docker-compose -f dev-env/docker-compose.yml up -d
```

URLs:

* UI: http://127.0.0.1:8080
* WMTS: http://127.0.0.1:8080/wmts
* API: http://127.0.0.1:8080/api/v1/config/global
* Mongo Express: http://127.0.0.1:8081


## Build

```
make
```

# Useful links

## Go

### Tests

Assert: https://godoc.org/github.com/stretchr/testify/assert

## WMTS

- OGC: http://www.opengeospatial.org/standards/wmts
- XSD schema: http://schemas.opengis.net/wmts/1.0/

## Mongo

### Sharding

https://docs.mongodb.com/manual/tutorial/deploy-shard-cluster/

### CRUD

- Doc: https://docs.mongodb.com/manual/crud/
- API: https://godoc.org/github.com/globalsign/mgo#Collection

### Grid FS

- GridFS Doc: https://docs.mongodb.com/manual/core/gridfs/
- Go lib for GridFS: https://godoc.org/github.com/globalsign/mgo#GridFS