![Github CI/CD](https://img.shields.io/github/workflow/status/kingmidas74/golculator/Test:%20Base)
![GitHub last commit](https://img.shields.io/github/last-commit/kingmidas74/golculator)
![Go Report](https://goreportcard.com/badge/github.com/kingmidas74/golculator)
![Repository Top Language](https://img.shields.io/github/languages/top/kingmidas74/golculator)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/kingmidas74/golculator)
![Github Repository Size](https://img.shields.io/github/repo-size/kingmidas74/golculator)
![Github Open Issues](https://img.shields.io/github/issues/kingmidas74/golculator)
![Lines of code](https://img.shields.io/tokei/lines/github/kingmidas74/golculator)
![License](https://img.shields.io/badge/license-MIT-green)
![GitHub contributors](https://img.shields.io/github/contributors/kingmidas74/golculator)
![Simply the best ;)](https://img.shields.io/badge/simply-the%20best%20%3B%29-orange)

<img align="right" width="50%" src="./assets/logo.png">

# Golculator

## Task description

Write a simple calculator with support of expression parsing.

### Requirements

- Support any math correct expressions with a primitive base operations
- Support unary operations
- Support  operations with complex numbers
- The calculator should have web-GUI
- Test coverage

### Optional

- Support of Quaternions and Octonions algebra
- Support user-defined functions

## Solution notes

- :book: standard Go project layout (or not :neutral_face:)
- :cd: github CI/CD + docker compose + Makefile included
- :card_file_box: PostgreSQL database

## HOWTO

- run with `make` (rebuild images) and go to [localhost:3000](http://localhost:3000)
- start with `make start` (without rebuild) and go to [localhost:3000](http://localhost:3000)
- test with `make test`

## A picture is worth a thousand words

<img src="./assets/web-example.png">