FROM golang:1.16
LABEL maintainer="AustinCloudGuru"

ENV TERRAFORM_VERSION 1.0.0

RUN apt-get update && apt-get install -y gnupg software-properties-common curl && \
    curl -fsSL https://apt.releases.hashicorp.com/gpg | apt-key add - && \
    apt-add-repository "deb [arch=amd64] https://apt.releases.hashicorp.com $(lsb_release -cs) main" && \
    apt-get update && \
    apt-get install terraform=${TERRAFORM_VERSION}

WORKDIR $GOPATH/src/app/test/
