# docker-terratest
[![CI](https://github.com/austincloudguru/docker-terratest/workflows/CI/badge.svg?event=push)](https://github.com/austincloudguru/docker-terratest/actions?query=workflow%3ACI)

This repo contains a docker image used to test Terraform resources with Terratest.  The container is run with the `user` terratest, uid = 1000 so that any terraform related files are not created as root.

GoLang Version: 1.16.5  
Terraform Version: Same as the docker tag

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
You will need to pass the terraform directory to the container as well as AWS credentials. By default, the container assumes that tests are in the `test` directory.  If your tests are in a different directory, you can add the `-w /go/src/app/$TEST_DIR` to your docker command.

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
    -v $(cd ~ && pwd)/.aws:/terratest/.aws \
    -e AWS_PROFILE=dev \
    austincloud/terratest go test -v
```


## Authors
Module is maintained by [Mark Honomichl](https://github.com/austincloudguru).

## License
MIT Licensed.  See LICENSE for full details
