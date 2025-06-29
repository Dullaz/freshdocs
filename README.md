# FreshDocs

[![Go Report Card](https://goreportcard.com/badge/github.com/Dullaz/freshdocs)](https://goreportcard.com/report/github.com/Dullaz/freshdocs)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](./LICENSE)

**Keep your documentation as fresh as your code!**

FreshDocs is an app that helps developers and users maintain a real, actionable connection between documentation and source code. 
By embedding hidden comment annotations in doc files, you can link docs to specific code files or folders. 
Whenever the code changes, FreshDocs detects which documentation is now stale — so your docs never fall behind.

FreshDocs can also be used by developers to quickly find what documentation will need to be updated for the changes they are making!

---

## 🚀 Features

- **Code-to-Doc Linkage:** Use simple annotations to connect markdown files to code.
- **Stale Doc Detection:** Automatically flags documentation that’s out of date when code changes.
- **Flexible Linking:** Supports linking single files or entire directories.
- **Developer Friendly:** Integrates smoothly into existing workflows.

---

## 🛠️ Installation


```bash
brew install freshdocs
```

---

## 📖 How It Works

1. **Initialise freshdocs**  
   Navigate to where you want to store the fresh config and run
   ```bash
   fresh init
   ```
2. **Update the configuration**  
   The app will provide placeholder values that you can modify
   
   More than one target repository can be specified, and more than one
   document folder can be configured too.

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
   
   ```md
   # the following is an annotation
   
   <!--- fresh:file core:folder/to/file.go -->
   ```
   
   For a full list of supported annotations, see here
---

## Command examples

```shell
fresh validate
docs/example.md:2 affected by cmd/update.go
docs/example_2.md:2 affected by cmd/check.go
```

```shell
fresh check
docs/example_3.md:2 affected by cmd/validate.go
```

```shell
fresh update
```
(hashes are now updated)

```shell
fresh update specific/file/path.md
```
(hashes for that path are updated)

```shell
fresh find path/to/some/code.go
some/docs/path.md
another/docs/path.md
```

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

