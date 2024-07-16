# otel-sig-security-example-go
Example of a Go project following SIG Security's recommendations

# What's in here

* GitHub Actions that created "attestations" (including binary and image digests)
* CodeQL, for static analysis

Optionally, we have examples of the following, which are NOT official recommendations at this point:

* Pinning of versions for GitHub Actions, updated weekly by Renovate

# How to verify

## Using GitHub CLI

Verifying the binary using [GitHub CLI tool (`gh`)](https://github.com/cli/cli):
```command
gh attestation verify otel-sig-security-example-go_0.0.3_linux_amd64.tar.gz --owner jpkrohling
```

Verifying the container image using `gh`:
```
gh attestation verify oci://ghcr.io/jpkrohling/otel-sig-security-example-go:0.0.3 --owner jpkrohling
```

Note: `jpkrohling` refers to the organization, and would be `open-telemetry` on our actual repositories.

## Sigstore's Transparency logs

Verifying the binary:
* Get the `sha256sum` for the file: `sha256sum otel-sig-security-example-go_0.0.3_linux_amd64.tar.gz`
* Open the transparency logs and get the attestation for the resulting SHA256 sum by using a query like this: `sha256:0d1d5c2255d7420e2561543cf74acdd98e6800f034a64bd771f6f83d2588ca26`
* It should yield one result, with details about the process that generated that binary: https://search.sigstore.dev/?hash=sha256:0d1d5c2255d7420e2561543cf74acdd98e6800f034a64bd771f6f83d2588ca26

Verifying the container image:
* Pull the image: `docker pull ghcr.io/jpkrohling/otel-sig-security-example-go:0.0.3`
* Obtain the checksum for the image: `docker inspect --format='{{index .RepoDigests 0}}' ghcr.io/jpkrohling/otel-sig-security-example-go:0.0.3`
* Open the transparency logs and get the attestation for the resulting SHA256 sum by using a query like this: `sha256:b2a3ff71b9a4fff62d96b51e3ef7e009abdf513c597680dd6e80b4e2b4940ed4`
* It should yield one result, with details about the process that generated that binary: https://search.sigstore.dev/?hash=sha256:b2a3ff71b9a4fff62d96b51e3ef7e009abdf513c597680dd6e80b4e2b4940ed4
