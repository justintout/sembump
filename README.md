# sembump

[![Travis CI](https://travis-ci.org/justintout/sembump.svg?branch=master)](https://travis-ci.org/justintout/sembump)

Easily bump MAJOR/MINOR/PATCH numbers given a specific [SemVer](https://semver.org) version. 

Basically forked from [jessfraz's sumbump](https://github.com/jessfraz/junk/tree/master/sembump), with the extended capability of bumping "release candidate" pre-release version numbers (ex: v4.0.1-rc.1, v1.2.3-4). 

Incrementing major/minor/patch will handle any SemVer-valid scheme. Bumping the pre-release tag will only handle exactly one pre-release version either of the form `-tag.number` or `-number`. 


## Installation

#### Binaries

- **darwin** [386](https://github.com/justintout/sembump/releases/download/v0.0.0/sembump-darwin-386) / [amd64](https://github.com/justintout/sembump/releases/download/v0.0.0/sembump-darwin-amd64)
- **freebsd** [386](https://github.com/justintout/sembump/releases/download/v0.0.0/sembump-freebsd-386) / [amd64](https://github.com/justintout/sembump/releases/download/v0.0.0/sembump-freebsd-amd64)
- **linux** [386](https://github.com/justintout/sembump/releases/download/v0.0.0/sembump-linux-386) / [amd64](https://github.com/justintout/sembump/releases/download/v0.0.0/sembump-linux-amd64) / [arm](https://github.com/justintout/sembump/releases/download/v0.0.0/sembump-linux-arm) / [arm64](https://github.com/justintout/sembump/releases/download/v0.0.0/sembump-linux-arm64)
- **solaris** [amd64](https://github.com/justintout/sembump/releases/download/v0.0.0/sembump-solaris-amd64)
- **windows** [386](https://github.com/justintout/sembump/releases/download/v0.0.0/sembump-windows-386) / [amd64](https://github.com/jessfraz/sembump/releases/download/v0.0.0/sembump-windows-amd64)
