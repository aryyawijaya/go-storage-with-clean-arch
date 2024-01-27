## Description

Go project that provide storage (files) use cases with trying to structure code that accomplish Clean Architecture way.

## Architecture diagram

![architecture diagram](https://github.com/aryyawijaya/go-storage-with-clean-arch/blob/dev/architecture_diagram.png)

The diagram above depicts the structure of the code in this app/system that following Dependency Rule: *Source code dependencies must point only inward, toward higher-level policies*

## Running the app

```bash
# prod
$ make up-prod
# dev
$ make up-dev
```

## Log the app

```bash
$ make logs-api
```

## Test

```bash
$ make test
```