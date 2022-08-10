## Go Test Command
```bash
go test ./... -v -coverprofile=coverage.out
```
## Sonar Scanner Command
```bash
docker network create sonar-network
docker run --network="sonar-network" -p 9000:9000 --name sonarqube -d sonarqube:lts-community
docker run -e SONAR_HOST_URL=http://sonarqube:9000 --network="sonar-network" --user="$(id -u):$(id -g)" -it -v "$PWD:/usr/src" sonarsource/sonar-scanner-cli -D sonar.login=aa484fce6b9a0dfeecf7b787237f1b7ae90b503f -D sonar.projectKey=go-test -D sonar.sources=. -D sonar.exclusions=**/*_test.g o -D sonar.tests=. -D sonar.test.inclusions=**/*_test.go -D sonar.go.coverage.reportPaths=coverage.out
```