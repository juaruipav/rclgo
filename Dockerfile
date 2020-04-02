# Starting from golang image
FROM golang:latest as go-docker

# Our final image should have ros:eloquent already installed
FROM ros:eloquent

# Get the binaries from previous docker stage:
COPY --from=go-docker /usr/local/go /usr/local/go

# Create the main user:
ARG user=rclgo
RUN useradd -ms /bin/bash ${user}
USER ${user}
WORKDIR /home/${user}
RUN  mkdir -p /home/${user}/go/bin && \
     mkdir -p /home/${user}/go/src

# Environment variables
ENV PATH $PATH:/usr/local/go/bin
ENV GOPATH /home/${user}/go
ENV GOROOT /usr/local/go

# Checkout rclgo code:
RUN git clone https://github.com/richardrigby/rclgo.git /home/${user}/go/src/rclgo