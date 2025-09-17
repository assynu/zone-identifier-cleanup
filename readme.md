# zic

`zic` is a simple CLI tool written in Go to recursively remove files containing `"Zone.Identifier"` from a specified directory.

> ⚠️ **Note:** This tool was created to fix a very specific issue: WSL sometimes creates `Zone.Identifier` files when copying files from a Windows directory. It is a personal, one-off solution and may **not adhere to best coding standards**. It will **probably not be maintained**. Use at your own risk.

---

## Installation

### macOS / Linux (one-liner)

```bash
git clone https://github.com/assynu/zone-identifier-cleanup zic-temp && \
cd zic-temp && \
go build -o zic && \
chmod +x zic && \
sudo mv zic /usr/local/bin/ && \
cd .. && rm -rf zic-temp
```

This will:

1. Clone the repository into a temporary folder.
2. Build the `zic` binary.
3. Make sure it has execute permissions.
4. Move it to `/usr/local/bin/` so it's available globally.
5. Clean up the temporary folder.

### Windows

1. Clone the repository:

```powershell
git clone <your-repo-url>
cd zic
go build -o zic.exe
```

2. Move `zic.exe` to a folder in your PATH, e.g., `C:\Users\<YourUser>\bin\`.
3. Make sure the folder is included in your PATH environment variable.

---

## Usage

```bash
zic <directory>
```

Examples:

```bash
# Remove Zone.Identifier files in the current directory
zic .

# Remove Zone.Identifier files in a specific folder
zic /path/to/folder
```

* `.` refers to the **current working directory**.
* Absolute paths are also supported.
