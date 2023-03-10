# rollee-test-1

[![Build](https://github.com/shoriwe/rollee-test-1/actions/workflows/build.yaml/badge.svg)](https://github.com/shoriwe/rollee-test-1/actions/workflows/build.yaml)
[![codecov](https://codecov.io/gh/shoriwe/rollee-test-1/branch/main/graph/badge.svg?token=X4TANIHG6Q)](https://codecov.io/gh/shoriwe/rollee-test-1)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/shoriwe/rollee-test-1)
[![Go Report Card](https://goreportcard.com/badge/github.com/shoriwe/rollee-test-1)](https://goreportcard.com/report/github.com/shoriwe/rollee-test-1)

Test 1 for Rollee.

## Project structure

| File/Directory                                               | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ |
| [docs/Readme.md](docs/Readme.md)                             | This file contains the challenge to be implemented using `Go` |
| [.github/workflows/build.yaml](.github/workflows/build.yaml) | CI pipeline for unit building, testing, versioning and releasing the repository latest version. |
| [.github/workflows/coverage.yaml](.github/workflows/coverage.yaml) | CI pipeline for making the coverage for the entire repository. |
| [CHANGELOG.md](CHANGELOG.md)                                 | Autogenerated changelog, by using `standard-version` and `fix`, `feat` and `docs` commits. |
| [CONTRIBUTING.md](CONTRIBUTING.md)                           | This file describe all the specification to properly contribute for this project, including the commits convention to use. |
| [docs/CICD.md](docs/CICD.md)                                 | This document describes the CICD pipelines present in the repository. |

## Run tests

```go
go test -count=1 -v ./...
```

## Coverage

|                           Sunburst                           |                             Grid                             |
| :----------------------------------------------------------: | :----------------------------------------------------------: |
| [![](https://codecov.io/gh/shoriwe/rollee-test-1/branch/main/graphs/sunburst.svg?token=X4TANIHG6Q)](https://app.codecov.io/gh/shoriwe/rollee-test-1) | [![codecov-chart](https://codecov.io/gh/shoriwe/rollee-test-1/branch/main/graphs/tree.svg?token=X4TANIHG6Q)](https://app.codecov.io/gh/shoriwe/rollee-test-1) |
