# P2W - Page To What

[中文](README_zh.md)

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

## Install

If you're not using Docker, you'll need to have *Chrome* installed before you can use it.

### Executable File

First you need to download the P2W executable file according to your operating system.

If you are using Linux, place the executable in the `/usr/local/bin` directory.

If you are using Windows, place the executable in any directory in the `PATH` environment variable.

### Docker

> Docker image not yet released, 
> Please use the `Dockerfile` under the project to package it yourself.

```shell
docker run -d -p 12800:12800 --name p2w skye-z/p2w
```

## Use

### Use from the command line

P2W provides `pdf` and `image` commands.

* Public Flags
    * --url/-u: URL to be converted.
    * --path/-p: output path after conversion (default `. /`)
    * --code/-c: the code that identifies the conversion task
    * ---send/-s: the address to send the converted file to.
* `image` Dedicated Flags
    * --element/-e: intercept element
    * --quality/-q: image quality (default 90)

> Note that `path` and `send` can only be used interchangeably, with the `send` tag taking precedence.

```shell
# Export PDF to current path
p2w pdf -u="https://github.com" -p="./"
# Send PDF to specified address
p2w pdf -u="https://github.com" -c="github" -s="http://localhost:8080/test"

# Output full page image to current path
p2w image -u="https://github.com" -p="./" -q="90"
# Send the full page image to the specified address
p2w image -u="https://github.com" -q="90" -c="github" -s="http://localhost:8080/test"

# Outputs an image of the specified element to the current path
p2w image -u="https://github.com" -p="./" -q="90" -e=".application-main"
# Sends an image of a specified element to a specified address
p2w image -u="https://github.com" -q="90" -e=".application-main" -c="github" -s="http://localhost:8080/test"
```

### Used through the HTTP interface

P2W provides two `GET` interfaces, `/api/pdf` and `/api/img`.

* public parameters
    * url: the url to be converted
    * code: the code of the conversion task
    * send: the address to send the converted file to
* `image` special parameters
    * element: the element to be captured
    * quality: quality of the image (default 90)

Please note that you need to start the HTTP server with the command line command `server` to access it.

``` shell
p2w server -p="12800"
```

## Packaging and compilation

### Compile the executable

```shell
go mod download
go mod tidy
# Packaging for the current platform
go build -o p2w -ldflags '-s -w'

# MacOS
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o p2w -ldflags '-s -w'
# Linux
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o p2w -ldflags '-s -w'
# Windows
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o p2w -ldflags '-s -w'
```

### Packaging Docker images

```shell
# Compile the Linux version first
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o p2w -ldflags '-s -w'
# Then build the image in the directory
docker build -t skye-z/p2w:1.0.0 .
```