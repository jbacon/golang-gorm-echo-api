# A GoLang Web Http Api Server
A GoLang (`1.15.3`) HTTP API Server with SQLite backend, built using Gorm ORM and Echo Web Framework. 
For reference/example purposes.

## TL;DR
```bash
go run main.go
```
* Test command: `curl http://localhost:1323/`
* Database is created/seeded with dummy data for some tables.
### Post Report
```bash
curl \
--header "Content-Type: application/json" \
--request POST \
--data '{
  "createdDate": "2019-12-01T12:00:00.0-07:00",
  "tests": [
    {
      "short_name": "world_readable",
      "vulnerable": true,
      "data": [
        {
          "path": "/data/data/file1"
        },
        {
          "path": "/data/data/file2"
        },
        {
          "path": "/data/data/file3"
        },
        {
          "path": "/data/data/file4"
        }
      ]
    }
  ]
}' http://localhost:1323/apps/app1/versions/1.0/reports
```
#### Get Reports
```bash
curl \
--header "Content-Type: application/json" \
--request GET \
http://localhost:1323/apps/app1/versions/1.0/reports
```

## DOCKER STUFF (just for reference)

### 0. Docker Work
```bash
docker run -i -t -p 1323 \
-v ${PWD}:/app/ \
--workdir /app/ \
golang:1.15.3 go run main.go
```

### 1. Docker Build
```bash
docker build -t josh-bacon-golang-api:latest .
```

### 2. Docker Run
```bash
docker run -p 1323:1323 josh-bacon-golang-api:latest
```

### 3. Docker Publish
```bash
docker tag josh-bacon-golang-api:latest gcr.io/rich-involution-105622/josh-bacon-golang-api:latest
docker push gcr.io/rich-involution-105622/josh-bacon-golang-api:latest
```

## DEPLOYMENT (just for reference)

### 1. Google Cloud Run 
```bash
gcloud run deploy josh-bacon-golang-api \
--image gcr.io/rich-involution-105622/josh-bacon-golang-api:latest \
--platform=managed \
--region=us-west1
```

## OTHER USEFUL GOLANG COMMANDS  (just for reference)

### 1. Linting - Prints out style mistakes
```
go get -u golang.org/x/lint/golint
~/go/bin/golint -set_exit_status ./...
```

### 2. Vet - Examines Go source code and reports suspicious constructs (`go doc cmd/vet`)
```
go vet ./... 2> go-vet-report.out
```

### 3. Vendor  - Adds/Updates all dependencies to the `vendor/` directory
```
go mod vendor
```

### 4. Test - Run all tests (for all included go modules)
```bash
go test -mod vendor ./... -json > go-test-report.json
go test -mod vendor ./... -coverprofile=go-test-coverage.out
```
- Test logs are printed to `STDOUT` and `go-test-report.json`
- Test coverage report is printed to `go-test-coverage.json`

### 5. Build Executable
```bash
go build -o executable -mod vendor ./...
```