# this image uses the same binary that is built in the CI, they have the same checksum
# to build the binary required by this Dockerfile, use: 
# 
# $ CGO_ENABLED=0 go build -o otel-sig-security-example-go .
# $ docker build -t jpkroehling/otel-sig-security-example-go:latest .
#
# or, if you are building with goreleaser:
# 
# $ goreleaser --snapshot --clean
# $ docker build -f Dockerfile -t jpkroehling/otel-sig-security-example-go:latest dist/otel-sig-security-example-go_linux_amd64_v1

FROM alpine:3.17
ENTRYPOINT ["/otel-sig-security-example-go"]

# Note that we use the exact same binary that will be built in the CI
# The binary from the image and from CI have the same checksum
COPY otel-sig-security-example-go /