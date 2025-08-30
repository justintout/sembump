# sembump

[![CI](https://github.com/justintout/sembump/actions/workflows/ci.yml/badge.svg)](https://github.com/justintout/sembump/actions/workflows/ci.yml)

Easily bump MAJOR/MINOR/PATCH numbers given a specific [SemVer](https://semver.org) version.

Forked from [jessfraz's sumbump](https://github.com/jessfraz/junk/tree/master/sembump), with the extended capability of bumping "release candidate" pre-release version numbers (ex: v4.0.1-rc.1, v1.2.3-4).

Incrementing major/minor/patch will handle any SemVer-valid scheme. Bumping the pre-release tag will only handle exactly one pre-release version either of the form `-tag.number` or `-number`. 

Bumping a non-prerelease tagged version with the `pre` kind will append `-rc.1` to the version.

Note: The default branch is now `main`.


## Installation

#### Binaries (latest release)

- **darwin** [amd64](https://github.com/justintout/sembump/releases/latest/download/sembump-darwin-amd64) / [arm64](https://github.com/justintout/sembump/releases/latest/download/sembump-darwin-arm64)
- **freebsd** [386](https://github.com/justintout/sembump/releases/latest/download/sembump-freebsd-386) / [amd64](https://github.com/justintout/sembump/releases/latest/download/sembump-freebsd-amd64)
- **linux** [386](https://github.com/justintout/sembump/releases/latest/download/sembump-linux-386) / [amd64](https://github.com/justintout/sembump/releases/latest/download/sembump-linux-amd64) / [arm](https://github.com/justintout/sembump/releases/latest/download/sembump-linux-arm) / [arm64](https://github.com/justintout/sembump/releases/latest/download/sembump-linux-arm64)
- **solaris** [amd64](https://github.com/justintout/sembump/releases/latest/download/sembump-solaris-amd64)
- **windows** [386](https://github.com/justintout/sembump/releases/latest/download/sembump-windows-386) / [amd64](https://github.com/justintout/sembump/releases/latest/download/sembump-windows-amd64)

## Release Automation

Merges to `main` publish a GitHub Release using the version in the `VERSION` file and build cross‑platform binaries (darwin amd64/arm64, linux 386/amd64/arm/arm64, freebsd 386/amd64, solaris amd64, windows 386/amd64).

The download links above always reference the latest release via `releases/latest/download/*`, so the README does not need to be updated each time.

To cut a new release:

1. Edit `VERSION` (ensure it starts with `v`, e.g. `v0.2.1`).
2. Commit and merge to `main` (or open a PR then merge).
3. The workflow will tag (if missing), build binaries, and publish a release with assets.

Add `[skip release]` to a commit message on `main` if you need to push changes without creating a release.
