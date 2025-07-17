# Simplebank

### Export PATH for go binaries

```bash
export PATH=$PATH:$(go env GOPATH)/bin
```

### Create migration files

```bash
migrate create -ext sql -dir db/migration -seq init_schema
```
