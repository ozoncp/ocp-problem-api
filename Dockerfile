FROM ubuntu AS builder

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

RUN echo developer | sudo -S DEBIAN_FRONTEND="noninteractive" apt install -y golang && \
    echo developer | sudo -S apt install -y ca-certificates && sudo update-ca-certificates && \
    echo developer | sudo -S apt install -y make git vim protobuf-compiler

ENV GOPATH /home/developer/go
ENV PATH $PATH:/home/developer/go/bin

COPY . /home/developer/go/src/github.com/ozoncp/ocp-problem-api
RUN echo developer | sudo -S chown -R developer /home/developer/

WORKDIR /home/developer/go/src/github.com/ozoncp/ocp-problem-api

RUN make deps && make build

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /home/developer/go/src/github.com/ozoncp/ocp-problem-api/bin/ocp-problem-api .
RUN chown root:root ocp-problem-api
EXPOSE 8082
EXPOSE 8083
CMD ["./ocp-problem-api", "-host", ""]