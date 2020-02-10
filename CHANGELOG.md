# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.6.7] - 2-10-2020
### Fixed
- docker command for tagging images with semver

## [0.6.5] - 2-10-2020
### Added
- condition for tagging containers with latest, commit hash and git tag

### Removed
- Cloud Build step for printing out directory listing

## [0.6.4] - 2-6-2020
### Added
- condition for tagging containers with latest, commit hash and git tag

### Changed
- run.sh to list directory files only

## [0.6.3] - 2-5-2020
### Added
- Case switch in bash script

### Changed
- renamed bash script to run.sh
- Cloud Build concurrent steps
- SHORT_SHA for tagging non master branches, TAG_NAME for master

### Removed
- Commented code in Dockerfile

## [0.6.2] - 2-4-2020
### Removed
- ENV vars from test.sh

## [0.6.1] - 2-4-2020
### Added
- Initial Cloud Build file
- Gosec scanning

### Removed
- Go mod

## [0.6.0] - 2-3-2020
### Added
- signal.Notify for closing the program

### Removed 
- iterations var for controlling the number of tests
- Jenkinsfile

## [0.5.0] - 1-28-2020
### Added
- Dockerfile
