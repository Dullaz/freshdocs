# Example Documentation

This is an example documentation file that demonstrates how to use FreshDocs annotations.

## API Reference

<!--- fresh:file core:cmd/*.go f275337 -->
The following section documents the main API functions:

<!--- fresh:file core:cmd/update.go c8ba97d -->
The `update` command allows you to update document hashes.

<!--- fresh:file core:cmd/validate.go f275337 -->
The `validate` command checks which documents are affected by code changes.

## Configuration

<!--- fresh:file core:config/config.go c8ba97d -->
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