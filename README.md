aeqd
====

[![Build Status](https://travis-ci.org/mc-aeq/aeqd.png?branch=master)](https://travis-ci.org/mc-aeq/aeqd)
[![ISC License](http://img.shields.io/badge/license-ISC-blue.svg)](http://copyfree.org)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/mc-aeq/aeqd)

aeqd is fork of [decred/dcrd]. aeqd is a full node Decred implementation written in Go (golang).

This acts as a chain daemon for the [Æquator](https://aequator.iog) cryptocurrency.
aeqd maintains the entire past transactional ledger of Æquator and allows relaying of transactions to other Æquator nodes across the world. 

Note: To send or receive funds and join Proof-of-Stake mining, you will also need [aeqw](https://github.com/mc-aeq/aeqw) - the Æquator wallet.

This project is currently under active development and is in a Beta state.  It is extremely stable since it's based on Decred which has been in production use since February 2016.

## Requirements

[Go](http://golang.org) 1.9 or newer.

## Getting Started

- aeqd (and utilities) will now be installed in either ```$GOROOT/bin``` or  ```$GOPATH/bin``` depending on your configuration.  If you did not already add the bin directory to your system path during Go installation, we recommend you do so now.

## Updating

#### Windows

Install a newer MSI

#### Linux/BSD/MacOSX/POSIX - Build from Source

- **Dep**

  Dep is used to manage project dependencies and provide reproducible builds.
  To install:

  `go get -u github.com/golang/dep/cmd/dep`

Unfortunately, the use of `dep` prevents a handy tool such as `go get` from automatically downloading, building, and installing the source in a single command.  Instead, the latest project and dependency sources must be first obtained manually with `git` and `dep`, and then `go` is used to build and install the project.

**Getting the source**:

For a first time installation, the project and dependency sources can be obtained manually with `git` and `dep` (create directories as needed):

```
git clone https://github.com/mc-aeq/aeqd $GOPATH/src/github.com/mc-aeq/aeqd
cd $GOPATH/src/github.com/mc-aeq/aeqd
dep ensure
go install . ./cmd/...
```

To update an existing source tree, pull the latest changes and install the matching dependencies:

```
cd $GOPATH/src/github.com/mc-aeq/aeqd
git pull
dep ensure
go install . ./cmd/...
```

For more information about Æquator and how to set up your software please visit our website or ask our support staff.

## Docker

All tests and linters may be run in a docker container using the script `run_tests.sh`.  This script defaults to using the current supported version of
go.  You can run it with the major version of Go you would like to use as the only arguement to test a previous on a previous version of Go (generally Decred
supports the current version of Go and the previous one).

```
./run_tests.sh 1.9
```

To run the tests locally without docker:

```
./run_tests.sh local
```

## Contact

If you want to get in touch with us, please visit our website https://aequator.io - all possible communication channels are listed there.

## Issue Tracker

We're using a private JIRA bugtracking system. Please foward us any issues and we'll assign them internally. If you want to join our bug tracking team, contact us and you'll get direct access to our JIRA instance.

## Documentation

Currently documentation is available in docs/ folder, but please be aware that those documentations may change at any given time.

## License

dcrd is licensed under the [copyfree](http://copyfree.org) ISC License.
aeqd is a fork from dcrd and also under the ISC License.

[//]: # (These are reference links used in the body of this note and get stripped out when the markdown processor does its job. There is no need to format nicely because it shouldn't be seen. Thanks SO - http://stackoverflow.com/questions/4823468/store-comments-in-markdown-syntax)

   [decred/dcrd]: <https://github.com/decred/dcrd> 