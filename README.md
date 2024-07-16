# otel-sig-security-example-go
Example of a Go project following SIG Security's recommendations

# What's in here

* GitHub Actions that created "attestations" (including binary and image digests)
* CodeQL, for static analysis

Optionally, we have examples of the following, which are NOT official recommendations at this point:

* Pinning of versions for GitHub Actions, updated weekly by Renovate

# How to verify

## Using GitHub CLI

Using [GitHub CLI tool (`gh`)](https://github.com/cli/cli):
```command
gh attestation verify otel-sig-security-example-go_0.0.1_linux_amd64.tar.gz --owner jpkrohling
```

Note: `jpkrohling` refers to the organization, and would be `open-telemetry` on our actual repositories.

## Sigstore's Transparency logs

* Open the workflow that generated the release, like this: https://github.com/jpkrohling/otel-sig-security-example-go/actions/runs/9908788890
* Find the file you want to verify on that page, like: `otel-sig-security-example-go_0.0.1_linux_amd64.tar.gz`
* Open the attestation, and check the digest there, like: `sha256:a83a491c2dfdec1d5ce6c528bf53717602a50b7ee0c0c02fe4bf4611a33e9dee`
* Open Rekor's search and type in the hash there, resulting in this: https://search.sigstore.dev/?hash=sha256:a83a491c2dfdec1d5ce6c528bf53717602a50b7ee0c0c02fe4bf4611a33e9dee
* Check that it matches the expectations, like:
  * GitHub Workflow SHA: 40ba6ff5674c08f9b6474ca75620606be5620a11
  * GitHub Workflow Ref: refs/tags/v0.0.1
