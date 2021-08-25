FROM ubuntu AS builder

ARG DATABASE_URL

RUN apt update -y && \
    apt upgrade -y && \
    apt install -y locales && \
    apt install -y sudo && \
    echo "LC_ALL=en_US.UTF-8" >> /etc/environment && \
    echo "en_US.UTF-8 UTF-8" >> /etc/locale.gen && \
    echo "LANG=en_US.UTF-8" > /etc/locale.conf && \
    locale-gen en_US.UTF-8 && \
    useradd -m -G sudo developer && \
    echo 'developer:developer' | chpasswd
USER developer

RUN echo developer | sudo -S apt install -y ca-certificates && sudo update-ca-certificates && \
    echo developer | sudo -S apt install -y make git vim protobuf-compiler wget iputils-ping build-essential && \
    echo developer | sudo wget https://golang.org/dl/go1.17.linux-amd64.tar.gz && \
    echo developer | sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.17.linux-amd64.tar.gz && \
    echo developer | sudo rm go1.17.linux-amd64.tar.gz

ENV GOPATH /home/developer/go
ENV PATH $PATH:/home/developer/go/bin:/usr/local/go/bin
ENV DATABASE_URL $DATABASE_URL

COPY . /home/developer/go/src/github.com/ozoncp/ocp-problem-api
COPY ./run.sh .
RUN echo developer | sudo -S chown -R developer /home/developer/ ./run.sh && chmod +x run.sh

WORKDIR /home/developer/go/src/github.com/ozoncp/ocp-problem-api

RUN make deps && make build

CMD ["./run.sh"]