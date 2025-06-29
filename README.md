# FreshDocs

[![Go Report Card](https://goreportcard.com/badge/github.com/Dullaz/freshdocs)](https://goreportcard.com/report/github.com/Dullaz/freshdocs)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](./LICENSE)

**Keep your documentation as fresh as your code!**

FreshDocs helps developers and teams maintain a real, actionable connection between documentation and source code. 
By embedding hidden comment annotations in doc files, you can link docs to specific code files or folders. 
Whenever the code changes, FreshDocs detects which documentation is now stale — so your docs never fall behind.

FreshDocs can also be used by developers to quickly find what documentation will need to be updated for the changes they are making!

---

## 🚀 Features

- **Code-to-Doc Linkage:** Use simple annotations to connect markdown files to code.
- **Stale Doc Detection:** Flags documentation that's out of date with committed code.
- **Uncommitted Change Detection:** Instantly see which docs are affected by your current (uncommitted) changes.
- **Flexible Linking:** Supports linking single files or entire directories.
- **Developer Friendly:** Integrates smoothly into existing workflows.

---

## 🛠️ Installation

```bash
brew install freshdocs
```

---

## 📖 How It Works

1. **Initialize freshdocs**  
   Navigate to where you want to store the config and run:
   ```bash
   fresh init
   ```
2. **Update the configuration**  
   The app will provide placeholder values that you can modify. More than one target repository and document folder can be configured.

   ```yaml
   version: 1
   repositories:
     core:
       path: ../my-service
     utils:
       path: ../shared-utils
   documentGroups:
     - path: ./docs/folder
       ext: .md
   ```
3. **Add annotations to your documents**  
   Place annotations at the start of a line in your markdown files:
   ```md
   <!--- fresh:file core:folder/to/file.go -->
   ```
   The first time you add an annotation, it will not have a hash. After running `fresh update`, a short git commit hash will be added:
   ```md
   <!--- fresh:file core:folder/to/file.go abc1234 -->
   ```
   Only annotations at the start of a line are recognized (not those used as examples in docs).

---

## 🧩 Annotation Format

- `<!--- fresh:file repo:path/to/file.go -->` (first time, no hash)
- `<!--- fresh:file repo:path/to/file.go abc1234 -->` (with git commit hash, after update)

---

## 🕹️ Commands

- `fresh check`  
  Lists all documentation files affected by **uncommitted changes** (staged or unstaged) in your code. Use this before you commit to see what docs will need updating for your current work.

- `fresh validate`  
  Lists all documentation files that are **stale** (out of date with the latest committed code). Use this after you commit to see what docs need updating.

- `fresh update`  
  Updates the hashes in your documentation to match the current git commit for each linked file.

- `fresh update <file.md>`  
  Updates hashes for a specific documentation file.

- `fresh find <path/to/code.go>`  
  Lists all documentation files that reference the given code file.

---

## 📝 Command Examples

```shell
fresh validate
# docs/example.md affected by cmd/update.go
# docs/example_2.md affected by cmd/check.go

fresh check
# docs/example_3.md affected by cmd/validate.go

fresh update
# (hashes are now updated)

fresh update docs/example.md
# (hashes for that file are updated)

fresh find cmd/update.go
# docs/example.md
```

---

## 📂 Repository Structure

- `cmd/` – CLI entry point
- `config/` – Configuration files
- `docs/` – Example documentation
- `processor/` – Core logic
- `main.go` – App entry

---

## 🤝 Contributing

PRs, bug reports, and feature ideas are welcome!  
Please open an issue to discuss major changes before submitting a PR.

---

## 📄 License

This project is licensed under the MIT License. See [LICENSE](./LICENSE) for details.

---

## 📬 Contact

Questions? Open an issue or contact [@Dullaz](https://github.com/Dullaz).

