FROM quay.io/goswagger/swagger:v0.31.0 as swagger-build

# Switch to a clean build directory
WORKDIR /swagger-build

# Copy the swagger spec and existing generated code with any modifications and re-generate
COPY swagger.yml swagger.yml
COPY gen/server gen/server
COPY go.mod go.mod
COPY go.sum go.sum
RUN mkdir -p gen/server && swagger generate server --target gen/server --name KbdsRefRestapi --spec swagger.yml --principal interface{}

# Build stage should happen on the golang alpine image
FROM golang:1.21.11-alpine3.20 as build


# README: If any additional build dependencies are required for the project, add them to this RUN.
#         Do not add an additional RUN, as this will generate additional steps and caching layers
#         which will lengthen the build time.
RUN apk update \
    && apk add --no-cache git openssh ca-certificates \
    && update-ca-certificates

# Add SSH key and add known SSH hosts for fetching private Go packages
RUN mkdir -p /root/.ssh && \
    chmod 0700 /root/.ssh && \
    touch /root/.ssh/known_hosts  && \
    ssh-keyscan github.com > /root/.ssh/known_hosts  &&\
    git config --global url."git@github.com:".insteadOf "https://github.com/"

# Switch to the directory where our project will be stored
WORKDIR /src

# Tell Go which repo paths are private
ENV GOPRIVATE github.com/wangkebin

# Copy our project to the build container
COPY . .
RUN rm -rf gen
COPY --from=swagger-build /swagger-build/gen gen

# Build the project and make it executable
RUN mkdir /project
RUN go build -o /project/kbds-ref-restapi "./gen/server/cmd/..."
RUN chmod +x /project/kbds-ref-restapi

FROM alpine:3.20.0

# Binary will be run out of /project
WORKDIR /project

# Add public CA cert store for verifying outbound connections
# README: If any additional runtime dependencies are required for the project, add them to this RUN.
#         Do not add an additional run, as this will generate additional steps and caching layers
#         which will lengthen the build time.
# TODO:   Set this to the minimum set of certificates required to verify the minimal set of outbound services
RUN apk update \
    && apk add ca-certificates --no-cache \
    && update-ca-certificates

# Create an unprivileged user and switch to it so the process does not have access to root
RUN adduser --disabled-password --no-create-home --uid 1000 notroot notroot
USER notroot:notroot

# Copy the binary in
COPY --from=build /project/kbds-ref-restapi kbds-ref-restapi

# Start the binary
ENTRYPOINT ["/project/kbds-ref-restapi", "--write-timeout", "500s"]
