# pword

[![Go Reference](https://pkg.go.dev/badge/github.com/norwd/pword.svg)](https://pkg.go.dev/github.com/norwd/pword)
[![License](https://img.shields.io/github/license/norwd/pword)](https://github.com/norwd/pword/blob/main/LICENSE)
[![Go version](https://img.shields.io/github/go-mod/go-version/norwd/pword)](https://github.com/norwd/pword/blob/main/go.mod)
[![Latest release](https://img.shields.io/github/v/release/norwd/pword?include_prereleases)](https://github.com/norwd/pword/releases/latest)
[![Go Report Card](https://goreportcard.com/badge/github.com/norwd/pword)](https://goreportcard.com/report/github.com/norwd/pword)
[![CodeFactor](https://www.codefactor.io/repository/github/norwd/pword/badge)](https://www.codefactor.io/repository/github/norwd/pword)
[![Coverage Status](https://coveralls.io/repos/github/norwd/pword/badge.svg?branch=main)](https://coveralls.io/github/norwd/pword?branch=main)

A small, command line password generator.

Inspired by the [xkcd](https://xkcd.com/936) comic on password strength, `pword`
is little password generator in the spirit of fun and interesting password
strategies. It supports a variety of different password generation backends,
including `pword xkcd`. For a full list see `pword --help`.

## Installation

Installing `pword` is as simple as using the Go module installer on the command
line. If you already have Go installed, just run:

```sh
go install github.com/norwd/pword@latest
```

If you don't want to install the entire Go tool-chain, you can just download the
binary directly from [https://gobinaries.com](https://gobinaries.com) for a
smaller installation footprint using the following command:

```sh
curl -sf https://gobinaries.com/norwd/pword@latest | sh
```

### Installing From Source

If you want to build `pword` locally or make modifications, you can clone this
repo or [download](https://github.com/norwd/pword/archive/refs/heads/main.zip)
the source onto your local machine.

```sh
git clone git@github.com:norwd/pword.git # Clone pword (or download source zip)
cd pword                                 # Change directories to the source root
go build                                 # Build the source
./pword                                  # Run your local binary!
cp pword ~/bin                           # (Optional) Install the binary
```

## Security

*Be sure to also read my [security policy]*

This a relatively small project, intended to explore password generation
techniques, especially in relation to memorability while maintaining high
entropy. However, security is still a high priority. I am not qualified  to say
that passwords produced by `pword` are secure to one degree or another, but I do
believe that they are secure enough for low risk applications; throw-away reddit
accounts or a one time signup to a random web site.

[security policy]: https://github.com/norwd/pword/security/policy
