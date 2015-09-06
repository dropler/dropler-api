# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

# Copy the local package files to the container's workspace.
ADD . /go/src/dropler

# Deps
RUN go get dropler

# Build the outyet command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN go install dropler

# Port
ENV PORT=8080

# Document that the service listens on port 8080.
EXPOSE 8080

CMD []
ENTRYPOINT /go/bin/dropler

