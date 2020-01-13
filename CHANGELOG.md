# Change Log

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/)
and this project adheres to [Semantic Versioning](http://semver.org/).

## Unreleased

### Known Issues

### Added

### Changed

## 0.4.0

### Changed

* Remove dummy input from `list` action now that the [bug fix](https://github.com/choria-io/mcorpc-agent-provider/issues/126) has been released. Make a note of the choria version this requires.

## 0.3.0

### Changed

* Remove the test action `echo`
* Add summary to `list` output

## 0.2.2

### Changed

* `list`: Preserve spaces when showing the command run in a session
* `list`: don't send response if there are no user sessions

## 0.2.1

### Changed

* Make `kill` action errors more informative
* Ensure host is reported when running `list`.

## 0.2.0

### Added

* `kill` action

### Changed

* Rework `user list` so command is fully captured.

## 0.1.0

### Known Issues

* user list: command is not fully captured (truncated at first space)

### Added

* Basic working version
* Basic `user list` action
