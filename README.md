# Valerie - validate HTML schemas with Valerie [![CI Status](https://circleci.com/gh/felicianotech/valerie.svg?style=shield)](https://app.circleci.com/pipelines/github/felicianotech/valerie) [![Software License](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/felicianotech/valerie/master/LICENSE)

Validate HTML with Valerie!

Valerie is an HTML schema validator that ensures schemas embedded in your HTML are correct.
Currently Valerie is very early software and only contains limited Open Graph support.
Further Open Graph support, and eventually support for Twitter Cards, JSON Feed, and more will come in the future.


## Table of Contents

- [Install Valerie](#install-valerie)
  - [Linux](#linux)
  - [macOS](#macos)
  - [Windows](#windows)
  - [Continuous Integration (CI) Systems](#continuous-integration-ci-systems)
- [Usage](#usage)


## Install Valerie

### Linux

There are a few ways you can install Valerie on a Linux amd64 or arm64 system.

#### Debian Package (.deb)
You can install Valerie into an Apt based computer by download the `.deb` file to the desired system.

For graphical systems, you can download it from the [GitHub Releases page][gh-releases].
Many distros allow you to double-click the file to install.
Via terminal, you can do the following:

```bash
wget https://github.com/felicianotech/valerie/releases/download/v0.1.0/valerie_0.1.0_amd64.deb
sudo dpkg -i valeri_0.1.0_amd64.deb
```

`0.1.0` and `amd64` may need to be replaced with your desired version and CPU architecture respectively.

#### Binary Install
You can download and run the raw Valerie binary from the [GitHub Releases page][gh-releases] if you don't want to use any package manager.
Simply download the tarball for your OS and architecture and extract the binary to somewhere in your `PATH`.
Here's one way to do this with `curl` and `tar`:

```bash
dlURL="https://github.com/felicianotech/valerie/releases/download/v0.1.0/valerie-v0.1.0-linux-amd64.tar.gz"
curl -sSL $dlURL | sudo tar -xz -C /usr/local/bin valerie
```

`0.1.0` and `amd64` may need to be replaced with your desired version and CPU architecture respectively.

### macOS

There are two ways you can install Valerie on a macOS (amd64) system.
Support for M1 macs (the arm64 chip) is coming later in 2021 if there's demand.

#### Brew (recommended)

Installing Valerie via brew is a simple one-liner:

```bash
brew install felicianotech/tap/valerie
```

#### Binary Install
You can download and run the raw Valerie binary from the [GitHub Releases page][gh-releases] if you don't want to use Brew.
Simply download the tarball for your OS and architecture and extract the binary to somewhere in your `PATH`.
Here's one way to do this with `curl` and `tar`:

```bash
dlURL="https://github.com/felicianotech/valerie/releases/download/v0.1.0/valerie-v0.1.0-macos-amd64.tar.gz"
curl -sSL $dlURL | sudo tar -xz -C /usr/local/bin valerie
```

`0.1.0` may need to be replaced with your desired version.

### Windows

Valerie supports Windows 10 by downloading and installing the binary.
Chocolately support is likely coming in the future.
If there's a Windows package manager you'd like support for (including Chocolately), please open a GItHub Issue and ask for it.

#### Binary Install (exe)
You can download and run the Valerie executable from the [GitHub Releases page][gh-releases].
Simply download the zip for architecture and extract the exe.

### Continuous Integration (CI) Systems

Valerie can be installed in a CI environment pretty much the same way you'd install it on your own computer.
There is 1st-party support for some CI platforms in order to make the process easier.

#### CircleCI
Valerie will be available as an orb in the near future.


#### GitHub Actions
Coming soon, probably.
Open an Issue to request it and demonstrate demand.


## Usage

The basic idea is to pass a URL you want to check like this:

```bash
valerie validate "https://www.feliciano.tech"
```

Run `valerie help` to see all commands available.


## License

This repository is licensed under the MIT license.
The license can be found [here](./LICENSE).



[gh-releases]: https://github.com/felicianotech/valerie/releases
