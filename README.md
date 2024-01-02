[![License: MIT](https://img.shields.io/badge/license-MIT-purple.svg)](https://opensource.org/licenses/MIT)
[![Github action](../../actions/workflows/go.yml/badge.svg?branch=main)](../../actions/workflows/go.yml)

# snapshot-diff

## :book: Introduction

snapshot-diff is tool for compare two snapshots and find out what file had been changed between these snapshots

## :pushpin: Features

- [x] list volumes & snapshots
- [ ] compare two snapshots (timestamp/size)
- [x] hash snapshots files
- [ ] compare two snapshots (hash)
- [x] cache snapshots hash
- [ ] export files hash
- [ ] background compute new snapshots hash

# How to Use

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

## Requirement

snapshot-diff is currently planned to work on QNAP nas with docker through the [container-station](https://www.qnap.com/fr-fr/software/container-station) implementation

## :floppy_disk: Install

- install [container-station](https://www.qnap.com/fr-fr/software/container-station) using QNAP App Center
- Creating an Application using the content of `docker-compose.yml`

## :sos: Troubleshooting

### :space_invader: Known Issues

# :package: dependency

- [Go](https://go.dev/)

# :dvd: Versioning

We use [SemVer](https://semver.org/) for versioning. For the versions available, see the tags on this repository.

# :hearts: Show your support

Give a :star: star if this project helped you !

# 💁🏽 Contribute

Your contribution is always welcome!
Please read [contribution guideline](CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests to us.

# :closed_book: License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
