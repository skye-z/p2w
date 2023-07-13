# P2W - Page To What

P2W is a web page converted to PDF or image tool, provides command and http api

[![](https://img.shields.io/badge/Go-1.20+-%2300ADD8?style=flat&logo=go)](go.work)
[![](https://img.shields.io/badge/P2W-1.0.0-green)](control)
[![CodeQL](https://github.com/skye-z/p2w/workflows/CodeQL/badge.svg)](https://github.com/skye-z/p2w/security/code-scanning)

[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=skye-z_p2w&metric=security_rating)](https://sonarcloud.io/summary/new_code?id=skye-z_p2w)
[![Maintainability Rating](https://sonarcloud.io/api/project_badges/measure?project=skye-z_p2w&metric=sqale_rating)](https://sonarcloud.io/summary/new_code?id=skye-z_p2w)
[![Reliability Rating](https://sonarcloud.io/api/project_badges/measure?project=skye-z_p2w&metric=reliability_rating)](https://sonarcloud.io/summary/new_code?id=skye-z_p2w)
[![Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=skye-z_p2w&metric=vulnerabilities)](https://sonarcloud.io/summary/new_code?id=skye-z_p2w)
[![Bugs](https://sonarcloud.io/api/project_badges/measure?project=skye-z_p2w&metric=bugs)](https://sonarcloud.io/summary/new_code?id=skye-z_p2w)

## TODO

* [x] Command
    * [x] PDF
    * [x] Image
    * [x] Server
* [ ] API
    * [ ] PDF
    * [ ] Image

## Compile and package
```shell
go mod download
go mod tidy

go build -o p2w -ldflags '-s -w'

# MacOS
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o p2w -ldflags '-s -w'
# Linux
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o p2w -ldflags '-s -w'
# Windows
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o p2w -ldflags '-s -w'
```

## Firewall pass

If you find that you cannot access other devices after startup, please check whether the firewall is enabled. If so, please pass the port

### Firewall

```shell
firewall-cmd --add-port=12800/tcp --permanent
firewall-cmd --reload
```