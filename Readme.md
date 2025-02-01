# mut

[![GitHub commits](https://img.shields.io/github/commit-activity/m/dmitrorezn/go-mut)](https://github.com/dmitrorezn/go-mut/commits/main)
[![GitHub last commit](https://img.shields.io/github/last-commit/dmitrorezn/go-mut)](https://github.com/dmitrorezn/go-mut/commits/main)

The `mut` package provides a generic, thread-safe mutable container for any type. It encapsulates a pointer to a value of type `T` along with a mutex to safely allow concurrent mutations.

## Features

- **Generic Container**: Works with any type.
- **Safe Mutations**: Uses a mutex to control access to the contained value.
- **Flexible Locking**: Supports both blocking (`Mut`) and non-blocking (`TryMut`) locking.

## Installation

To use the package, simply copy the `mut` package into your project or install it using your favorite dependency management tool.

If you are using Go modules:

```bash
go get github.com/dmitrorezn/go-mut
