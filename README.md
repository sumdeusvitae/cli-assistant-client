[![Go Reference](https://pkg.go.dev/badge/github.com/SumDeusVitae/cli-assistant-client#section-directories.svg)](https://pkg.go.dev/github.com/SumDeusVitae/cli-assistant-client#section-directories)
[![Go Report Card](https://goreportcard.com/badge/github.com/SumDeusVitae/cli-assistant-client)](https://goreportcard.com/report/github.com/SumDeusVitae/cli-assistant-client)
![CI Status](https://github.com/SumDeusVitae/cli-assistant-client/actions/workflows/CI.yml/badge.svg)
<p>Server for our assistant: <a href="https://github.com/sumdeusvitae/cli-assistant#readme" target="_blank">Server - CLI Assistant</a>
</p>
<div align="center">
  <img src="https://github.com/SumDeusVitae/cli-assistant-client/blob/main/assistantHeader.png" />
</div>

## Table of Contents
- [Introduction](#introduction)
- [Example](#example)
- [Installation](#installation)
- [Usage](#usage)
- [Commands](#commands)
- [Uninstall](#uninstall)
- [Contributing](#contributing-)


## Introduction
CLI Assistant is a command-line tool that lets you ask questions and get answers directly in the terminal. It's perfect for situations where you don’t have access to a web browser, but still need quick answers or help solving problems. Whether you need a quick reminder, a bit of troubleshooting, or just some general knowledge, CLI Assistant has you covered—all from the comfort of your command line.

## Example
<p>
    <img src="https://github.com/SumDeusVitae/cli-assistant-client/blob/main/cli-assistant.gif" width="100%" alt="CLI ASSISTANT Example">
</p>


## Installation
This assumes you already have a working Go environment, if not please see
[this page](https://golang.org/doc/install) first.

To build, and install the `cli-assistant-client` command into your Go toolchain's `bin` directory. Go ahead and run:
```bash
go install github.com/SumDeusVitae/cli-assistant-client@latest
```
After this, rename the binary:
```bash
mv $(go env GOPATH)/bin/cli-assistant-client $(go env GOPATH)/bin/qs
```
Now, you run your app using qs instead of cli-assistant-client.
Run this command to check:
```bash
qs version
```
If it didn't work try:
```bash
cli-assistant-client version
```
If it works
Perhaps rm didn't work due to permission issues.
Try using cp to copy the file first, and then remove the original:
```bash
cp $(go env GOPATH)/bin/cli-assistant-client $(go env GOPATH)/bin/qs
rm $(go env GOPATH)/bin/cli-assistant-client
```



## Usage

To run CLI ASSISTANT, use following commands.

### For New Users:
Please register first. **Note:** Email is optional. It may be helpful for password recovery if you forget your password.

```bash
qs register
```

### For Existing Users:
If you already have an account, simply log in:

```bash
qs login
```

### Asking Questions:
After registering or logging in, you can ask a question using the following command:

```bash
qs q <your question here>
```

This will return the response to your query. This is the core functionality of the tool.

To see a list of available commands, use:

```bash
qs help
```

Alternatively, you can find the available commands in the "Commands" section below.


## Commands 
- `help`:  Displays a help message
- `register`: Registers new user
- `login`:  Login as existent user
- `q <question>`:  Asks AI question 
- `whoami`:  Checks if you logged in
- `helth`:  Checking server status
- `version`:  Checks current version
- `env`:  Shows saved environmental variables
- `update`:  Updates CLI to the latest version


## Uninstall
To uninstall CLI Assistant, follow the steps :

  Locate the binary: Find the binary in your GOBIN or GOPATH/bin directory.
  Remove the binary: Delete it from the terminal:

```bash
rm $(go env GOPATH)/bin/qs
```
  Remove the source code (optional):
```bash
rm -rf $(go env GOPATH)/src/github.com/SumDeusVitae/cli-assistant-client
```

## Contributing 🤝 

### Clone the repo

```bash
git clone https://github.com/sumdeusvitae/cli-assistant-client@latest
cd cli-assistant-client
```

### Build the project

```bash
go build -o qs
```

### Run the project

```bash
./qs register
./qs q <your question>
```

### Run the tests

```bash
go test ./...
```

### Submit a pull request

If you'd like to contribute, please fork the repository and open a pull request to the `main` branch.
