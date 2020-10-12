# operator

A command line tool for creating and deploying production machine learning functions.

These functions can be deployed on Google Cloud managed infrastructure as:

* [Cloud Functions](https://cloud.google.com/functions)
* [Cloud Run](https://cloud.google.com/run) containerized applications

Warning: this is a pre-release alpha version of this tool. Please send any bugs or feedback.

## Requirements

This project creates boiler plate code that uses:

* [Go](https://golang.org/doc/install): this CLI is written in Go (but it generates Python code!)
* [go-bindata](https://github.com/go-bindata/go-bindata)
* [pyenv](https://github.com/pyenv/pyenv) and [pyenv-virtualenv](https://github.com/pyenv/pyenv-virtualenv): the boiler plate generated by this tool enables you to easily create Python virtual environments, but assumes you have these tools installed already.
* [gcloud](https://cloud.google.com/sdk/gcloud) and an active GCP project in order to deploy your functions.

Specifically for Cloud Run:

* [Docker](https://docs.docker.com/get-docker/) to build and run cloud run containerized applications locally.

## Installing

Clone this repo and run:

```bash
❯ cd ~/src/github.com/operatorai/operator
❯ go get https://github.com/go-bindata/go-bindata
❯ make install
```

## Usage

Set up the CLI tool using `operator init`.

```bash
❯ operator init
Use the arrow keys to navigate: ↓ ↑ → ← 
? Deployment type: 
  ▸ Google Cloud Function
    Google Cloud Run
```

Create a new deployment with `operator create`:

```bash
❯ operator create hello-world 
```

... and set it up:

```bash
❯ cd hello-world 
❯ make install # To create a pyenv-virtualenv
```

Launch it locally:

```bash
❯ make localhost
```

... and, when you're ready, deploy it!

```bash
❯ operator deploy .
```

## Limitations

There are many! This is a version 0.

* All deployments go to `region=europe-west2`
* Only http-triggered cloud functions are supported
* This assumes that GCP APIs have been enabled and may not fail gracefully
* The templates/ directory has two templates which have high level of duplication
* There is no support for AWS yet

## Notes

This tool has been built using the [Cobra Generator](https://github.com/spf13/cobra/blob/master/cobra/README.md#cobra-generator).

To add a new command:

```bash
❯ cobra add <command-name>
```
