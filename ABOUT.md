# About writing-prompt

writing-prompt is a small command-line tool that generates short writing prompts by combining elements from simple word lists (characters, actions, descriptions, settings, and plot twists).

Purpose
- Provide quick, varied writing prompts for warm-ups, exercises, or inspiration.

Quick start
- Run the bundled binary (examples):
  - `./bin/writing-prompt-linux-amd64 -n 25`
  - `./bin/writing-prompt-linux-arm64 -n 25`
- Build and run from source (requires Go):
  - `go run main.go -n 25`
  - Or run `./build.sh` then the binary in `bin/`

Repository layout
- `main.go` — program entrypoint.
- `data/` — text files used as word lists: `characters.txt`, `actions.txt`, `descriptions.txt`, `names.txt`, `plottwists.txt`, `rawtitles.txt`, `settings.txt`.
- `bin/` — pre-built binaries for supported platforms.
- `build.sh` — build helper script.
- `LICENSE` — BSD 3-Clause license.

License
- This project is distributed under the BSD 3-Clause License. See `LICENSE` for details.

Author / Maintainer
- Owner: `misterunix` (repository owner)
- Copyright shown in `LICENSE`.

