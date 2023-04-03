FROM golang:1.18
LABEL maintainer="AustinCloudGuru"

# Change was made to debian packages to include the revision.
ARG tf_version=1.4.4-1
ARG uid=1000
ARG gid=1000
ARG user=terratest
ARG group=terratest
ARG terratest_home=/terratest

RUN apt-get update && apt-get install -y gnupg software-properties-common curl \
    && wget -O- https://apt.releases.hashicorp.com/gpg | gpg --dearmor | tee /usr/share/keyrings/hashicorp-archive-keyring.gpg \
    && echo "deb [signed-by=/usr/share/keyrings/hashicorp-archive-keyring.gpg] https://apt.releases.hashicorp.com $(lsb_release -cs) main" | tee /etc/apt/sources.list.d/hashicorp.list \
    && apt-get update \
    && apt-get install terraform=${tf_version} \
    && mkdir -p ${terratest_home}/.aws \
    && chown -R ${uid}:${gid} $terratest_home \
    && groupadd -g ${gid} ${group} \
    && useradd -d "$terratest_home" -u ${uid} -g ${gid} -m -s /bin/bash ${user}

USER ${user}

WORKDIR $GOPATH/src/app/test/
