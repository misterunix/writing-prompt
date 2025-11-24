# Contributing to writing-prompt

Thanks for your interest in contributing! This document explains how to report issues, propose changes, and submit pull requests for this repository.

**Be Respectful:** Please follow a respectful and constructive tone in issues and PRs.

Getting the code
- Fork the repository on GitHub and clone your fork locally.

  ```bash
  git clone https://github.com/<your-username>/writing-prompt.git
  cd writing-prompt
  git remote add upstream https://github.com/misterunix/writing-prompt.git
  ```

Create a branch
- Make feature/topic branches; don't work on `main` directly.

  ```bash
  git checkout -b feat/your-feature
  ```

Build and run
- The project is a small Go CLI. To run from source:

  ```bash
  go run main.go -n 25
  ```

- To build release binaries (Linux amd64/arm64 and Windows) the repo includes `build.sh`:

  ```bash
  ./build.sh
  ```

Code style
- This is a Go project. Use the standard Go tooling for formatting and simple checks:

  ```bash
  go fmt ./...
  go vet ./...
  ```

- If you add Go code, keep it simple and idiomatic. Avoid unnecessary complexity.

Tests
- There are currently no automated tests in the repository. If you add functionality, include tests and run them with:

  ```bash
  go test ./...
  ```

Committing
- Write clear, concise commit messages describing the change.
- Squash or rebase related work before opening a PR so the history is easy to review.

Pull requests
- Open a PR from your branch to `misterunix:main` with a descriptive title and explanation.
- Include the motivation, what you changed, and any manual testing steps.
- If the change is non-trivial, add examples or a small README update.

Reporting issues
- Please open issues for bugs, feature requests, or other improvements. Provide:
  - A short summary
  - Steps to reproduce (if applicable)
  - Expected vs actual behavior
  - Environment details (OS, Go version)

License
- By contributing, you agree that your contributions will be licensed under the project's BSD 3-Clause License. See `LICENSE` for details.

Questions
- If you're unsure where to start, open an issue describing what you'd like to work on or ask for guidance.

Thank you for contributing!

## All this is much easier if you use **vscode**
