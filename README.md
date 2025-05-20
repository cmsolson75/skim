# skim

**skim** is a configurable CLI tool that aggregates source code files into a single output, optimized for LLM ingestion, code review, or static analysis. It supports extension filtering, directory skipping, and optional inclusion of `cloc` and `tree` metadata.

---

## 🚀 Features

- Flatten and concatenate code files into a context file
- Filter by allowed extensions
- Skip common junk/infra directories (e.g., `.git`, `node_modules`)
- Optional project structure (`tree`) and line count stats (`cloc`)
- Configurable via `config.yaml` or CLI overrides
- Packaged as a Go CLI binary

---

## 🔧 Installation

### 📦 Option 1: Go Install

```bash
go install github.com/cmsolson75/skim/cmd/skim@latest
````

Ensure `$GOBIN` is on your `PATH`:

```bash
export PATH="$HOME/go/bin:$PATH"
```

### 🛠 Option 2: Manual Build

```bash
git clone https://github.com/cmsolson75/skim.git
cd skim
go build -o skim ./cmd/skim
mv skim /usr/local/bin/skim  # optional
```

---

## 📁 Configuration

By default, skim looks for `config.yaml` in the current directory.

### Example:

```yaml
input_dir: ./project
output_dir: ./out
output_name: context.txt

allowed_extensions:
  - .go
  - .py
  - .yaml
  - .txt
  - .md

skip_dirs:
  - .git
  - __pycache__
  - .venv
  - node_modules
  - build
  - dist
```

---

## 🧪 Usage

### With `config.yaml`

```bash
skim
```

### With CLI overrides

```bash
skim \
  --input-dir ./src \
  --output-dir ./dump \
  --output-name llm.txt \
  --allowed-extensions .go,.py,.yaml \
  --skip-dirs .git,__pycache__ \
  --cloc --tree
```

---

## ⚙️ Flags

| Flag                   | Description                              |
| ---------------------- | ---------------------------------------- |
| `--input-dir`          | Directory to scan (default: `.`)         |
| `--output-dir`         | Output directory (default: `./out`)      |
| `--output-name`        | Output filename (default: `context.txt`) |
| `--allowed-extensions` | Comma-separated list of extensions       |
| `--skip-dirs`          | Comma-separated dirs to ignore           |
| `--cloc`               | Include line count analysis              |
| `--tree`               | Include directory structure output       |

---

## 📦 Output

The tool generates a flat text file with:

* Project root
* Optional `tree` and `cloc` output
* Concatenated file contents with clear headers

---

## 🛠 Requirements

* Go 1.20+
* Optional (if using flags):

  * [`tree`](https://linux.die.net/man/1/tree)
  * [`cloc`](https://github.com/AlDanial/cloc)

---

## 🤝 Contributing

1. Fork the repo
2. Make changes on a feature branch
3. Submit a PR

---

## 🪪 License

MIT License

---

## ✨ Example

```bash
skim --input-dir . --cloc --tree
```

Generates `./out/context.txt` with all `.go`, `.py`, `.yaml`, etc., files inlined, along with optional `cloc` and `tree` stats.
