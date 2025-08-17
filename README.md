
# FreshDocs

A tool for developers to keep documentation up to date with code.



[![GPLv3 License](https://img.shields.io/badge/License-GPL%20v3-yellow.svg)](https://opensource.org/licenses/)



## Installation


Linux

```bash
  brew tap dullaz/freshdocs
  brew install freshdocs --cask
```
    
Apple (until I figure out how signing works)

```bash
brew tap dullaz/freshdocs
brew install freshdocs --cask --no-quarantine
```
## Usage

Create a config file that points to where your docs and code lives. freshdocs needs to be run in the same folder as this file.

Example config: 
```yaml
version: 1
repositories:
    core:
        path: .
documentGroups:
    - path: ./docs
      ext: .md
```

Add annotations to your documents
```markdown
# Some document you use
<!--- fresh:file core:path/to/code.go -->
```

Run `freshdocs update` to add a hash to the annotations

```markdown
# Some document you use
<!--- fresh:file core:path/to/code.go f275337 -->
```

Run `freshdocs validate` periodically or as a precommit check

```bash
~/workspace/freshdocs main* ❯ freshdocs validate
docs/howto.md:1 annotation is stale
```

Update your docs and then run `freshdocs update` to freshen the commit hashes in your annotations.


## Contributing

Contributions are always welcome!

I wrote this project in golang to learn golang, so expect things to be a bit messy.

Feel free to raise a PR
## License

[GPL v3.0](https://choosealicense.com/licenses/gpl-3.0)

