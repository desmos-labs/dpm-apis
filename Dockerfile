# This Docker allows running an instance of the DPM APIs.
#
# How to build the image:
# > docker build --tag desmoslabs/dpm-apis .
#
# How to run the image:
# > docker run desmoslabs/dpm-apis

FROM golang:1.20-alpine
ARG arch=x86_64

# Set up dependencies
ENV PACKAGES curl make git libc-dev bash gcc linux-headers eudev-dev python3 ca-certificates build-base
RUN set -eux; apk add --no-cache $PACKAGES;

# Set working directory for the build
WORKDIR /code

# Add sources files
COPY . /code/

# See https://github.com/CosmWasm/wasmvm/releases
ADD https://github.com/CosmWasm/wasmvm/releases/download/v1.2.3/libwasmvm_muslc.aarch64.a /lib/libwasmvm_muslc.aarch64.a
RUN sha256sum /lib/libwasmvm_muslc.aarch64.a | grep d6904bc0082d6510f1e032fc1fd55ffadc9378d963e199afe0f93dd2667c0160

ADD https://github.com/CosmWasm/wasmvm/releases/download/v1.2.3/libwasmvm_muslc.x86_64.a /lib/libwasmvm_muslc.x86_64.a
RUN sha256sum /lib/libwasmvm_muslc.x86_64.a | grep bb8ffda690b15765c396266721e45516cb3021146fd4de46f7daeda5b0d82c86

# Copy the library you want to the final location that will be found by the linker flag `-lwasmvm_muslc`
RUN cp /lib/libwasmvm_muslc.${arch}.a /usr/local/lib/libwasmvm_muslc.a

# Set the environment variables
ENV GIN_MODE=release

# Build the executable
RUN BUILD_TAGS=muslc GOOS=linux GOARCH=amd64 LINK_STATICALLY=true make build

# Move the executable inside the bin folder to make it runnable without specifying the full path
RUN cp /code/build/dpm-apis /usr/bin/apis

# Set the entrypoint, so that the user can set the config using the CMD
ENTRYPOINT ["apis"]