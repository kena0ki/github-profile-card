FROM golang:1.13

WORKDIR /workspace
# COPY . .
RUN apt-get update \
    # Install
    && apt-get install apt-file -y && apt-file update \
    && apt-get install vim git -y \
    # Clean up
    && apt-get autoremove -y \
    && apt-get clean -y \
    # Build
    && go build

CMD cd api && go run main.go
