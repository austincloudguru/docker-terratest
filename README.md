# docker-terratest
[![CI](https://github.com/austincloudguru/docker-terratest/workflows/CI/badge.svg?event=push)](https://github.com/austincloudguru/docker-terratest/actions?query=workflow%3ACI)

This repo containes a docker image used to test Terraform resources with Terratest.  

GoLang Version: 1.16.5
Terraform Version: 1.0.0

## Building the Container
Checkout the repository.
```bash
git clone git@github.com:austincloudguru/docker-terratest
```

Build the Docker Image:
```bash
 docker build -t austincloud/terratest .
```

## Usage:
You will need to pass the terraform directory to the container as well as AWS credentials.

__Running it with AWS environmental variables__
```bash
docker run --rm -it \
    -v $PWD:/go/src/app \
    --env-file <(env | grep "^AWS") \
    austincloud/terratest go test -v
```

__Running it using AWS credential files__ 
```bash
docker run --rm -it \
    -v $PWD:/go/src/app \
    -v $(cd ~ && pwd)/.aws/credentials:/root/.aws/credentials \
    -v $(cd ~ && pwd)/.aws/config:/root/.aws/config \
    -e AWS_SHARED_CREDENTIALS_FILE=/root/.aws/credentials \
    -e AWS_CONFIG_FILE=/root/.aws/credentials \
    -e AWS_PROFILE=cdp-dev \
    austincloud/terratest go test -v
```

## Authors
Module is maintained by [Mark Honomichl](https://github.com/austincloudguru).

## License
MIT Licensed.  See LICENSE for full details
