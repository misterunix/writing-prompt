# Security Policy

## Supported Versions

This project is a simple command-line tool. Security updates are applied to the latest version available on the `main` branch.

| Version | Supported          |
| ------- | ------------------ |
| latest (main) | :white_check_mark: |
| older releases | :x: |

## Reporting a Vulnerability

If you discover a security vulnerability in this project, please report it responsibly:

1. **Do not open a public issue** for the vulnerability.
2. Email the maintainer at **misterunix@gmaiol.com** with:
   - A description of the vulnerability
   - Steps to reproduce or proof-of-concept
   - Potential impact
   - Suggested fix (if available)
3. Allow a reasonable time for the maintainer to respond and address the issue before public disclosure.

## Security Best Practices

When using `writing-prompt`:

- **Build from source**: If security is a concern, build binaries yourself from the latest source rather than using pre-built binaries.
  
  ```bash
  git clone https://github.com/misterunix/writing-prompt.git
  cd writing-prompt
  go build -o bin/writing-prompt main.go
  ```

- **Dependencies**: This project uses Go's standard library and embeds data files. Review `go.mod` and embedded files if you have specific security requirements.

- **Data files**: The tool reads from embedded text files in `data/`. If you modify or add data files, ensure they don't contain malicious content.

- **Output files**: The tool writes CSV files to `~/writing-prompts/`. Be aware of file permissions and avoid running the tool with elevated privileges unless necessary.

## Disclosure Policy

- The maintainer will acknowledge receipt of a vulnerability report within **5 business days**.
- A fix or mitigation will be prioritized based on severity.
- Once a fix is available, a security advisory or release note will be published.

## Contact

For security-related questions or concerns, contact the maintainer at **misterunix@gmail.com**.

Thank you for helping keep this project secure!
