# Zone Identifier Cleanup

`Zone identifier cleanup` is a simple CLI tool written in Go to recursively remove files containing `"Zone.Identifier"` from a specified directory.

> ⚠️ **Note:** This tool was created to fix a very specific issue: WSL sometimes creates `Zone.Identifier` files when copying files from a Windows directory. It is a personal, one-off solution and may **not adhere to best coding standards**. It will **probably not be maintained**. Use at your own risk.

---

## Installation

### macOS / Linux (one-liner)

```bash
git clone https://github.com/assynu/zone-identifier-cleanup zone-identifier-cleanup-temp && \
cd zone-identifier-cleanup-temp && \
go build -o zone-identifier-cleanup && \
chmod +x zone-identifier-cleanup && \
sudo mv zone-identifier-cleanup /usr/local/bin/ && \
cd .. && rm -rf zone-identifier-cleanup-temp
```

This will:

1. Clone the repository into a temporary folder.
2. Build the `zone-identifier-cleanup` binary.
3. Make sure it has execute permissions.
4. Move it to `/usr/local/bin/` so it's available globally.
5. Clean up the temporary folder.

### Windows

1. Clone the repository:

```powershell
git clone <your-repo-url>
cd zone-identifier-cleanup
go build -o zone-identifier-cleanup.exe
```

2. Move `zone-identifier-cleanup.exe` to a folder in your PATH, e.g., `C:\Users\<YourUser>\bin\`.
3. Make sure the folder is included in your PATH environment variable.

> ⚠️ **Note:** This tool wasn't tested on windows, just on wsl linux instance running on windows. It may not work with windows powershell or cmd.

---

## Usage

```bash
zone-identifier-cleanup <directory>
```

Examples:

```bash
# Remove Zone.Identifier files in the current directory
zone-identifier-cleanup .

# Remove Zone.Identifier files in a specific folder
zone-identifier-cleanup /path/to/folder
```

* `.` refers to the **current working directory**.
* Absolute paths are also supported.
