# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

LABEL name "DailyDog"

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/james2doyle/dailydog

ENV DOG_JSON "https://api.giphy.com/v1/gifs/random?api_key=dc6zaTOxFJmzC&tag=dog"
# bake in the SLACK_WEBHOOK?
# ENV SLACK_WEBHOOK ""

# Set the working directory to avoid relative paths after this
WORKDIR /go/src/github.com/james2doyle/dailydog

# Fetch the dependencies
RUN go get .

# build the binary to run later
RUN go build

# Run the dailydog command by default when the container starts.
ENTRYPOINT /go/bin/dailydog

# Document that the service listens on port 3000.
# nothing is exposed at the moment...
EXPOSE 3000