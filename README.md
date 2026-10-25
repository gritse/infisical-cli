Fork of [Infisical/cli](https://github.com/Infisical/cli) with a few fixes for machine-identity use
of PAM, kept as a branch to be offered upstream.

- `main` mirrors upstream.
- `pam-machine-identity` is upstream's latest release tag plus the fix commits, nothing else.
- `fork` (this branch, the default) only holds the release plumbing: `.github/workflows/release.yml`
  rebases `pam-machine-identity` onto each new upstream release, tests it, publishes
  `ghcr.io/gritse/infisical-cli:<upstream version>-gritse.N` (linux amd64 + arm64, also used as
  the PAM gateway image) and attaches a macOS arm64 binary to a GitHub release.

Run the workflow by hand with *force* to cut a new `-gritse.N` for the same upstream version after
adding fixes. A failed rebase opens an issue.
