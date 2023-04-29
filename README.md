# Worder

Worder is a CLI tool written in Golang for generating files and counting words. It includes two commands: `count` and `generate`.

## Concurrency

Worder uses a worker pool and concurrency to speed up the word counting and file generation processes.
By default, the worker pool has a size of 10, **TODO:**

## Installation

To install Worder, you need to have Go installed on your machine. After that, you can install Worder by running:

```bash
go get github.com/ezzaldeeen/worder
go install worder
```

## Usage

To use Worder, navigate to the directory where you installed it and run the worder command followed by the desired command:
