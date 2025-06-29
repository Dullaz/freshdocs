# Example Documentation

This is an example documentation file that demonstrates how to use FreshDocs annotations.

## API Reference

The following section documents the main API functions:

<!--- fresh:file core:cmd/update.go 345f884 -->
The `update` command allows you to update document hashes.

<!--- fresh:file core:cmd/validate.go 1bc4df0 -->
The `validate` command checks which documents are affected by code changes.

## Configuration

<!--- fresh:file core:config/config.go a5f6a81 -->
The configuration system supports multiple repositories and document groups.

## Usage Examples

Here are some examples of how to use the FreshDocs CLI:

```bash
# Initialize a new project
fresh init

# Check for stale documentation
fresh check

# Update all document hashes
fresh update
```

## Annotation Format

FreshDocs annotations follow this format:
- `<!--- fresh:file repo:path -->` (first time, no hash)
- `<!--- fresh:file repo:path abc1234 -->` (with Git commit hash) 