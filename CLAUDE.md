# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

G3 is a CLI tool that uses GitHub Gist as a storage backend. It provides S3-like commands (`ls`, `cp`, `get`, `rm`) to store and retrieve files using GitHub's infrastructure. Files are encrypted with AES-256-CBC before being uploaded to Gist.

## Commands

```bash
# Development
npm run dev                    # Run CLI in development mode (ts-node)
npm run dev -- ls              # Run specific command
npm run dev -- cp ./file.txt   # Test copy command

# Build
npm run build                  # Clean and compile TypeScript to dist/
npm run clean                  # Remove dist folder

# Testing
npm test                       # Run all tests
node --import tsx --test test/config/index.spec.ts  # Run single test file

# Code Quality
npm run lint                   # Check for lint errors
npm run lint:fix               # Check and auto-fix lint errors
npm run format                 # Format with Biome
npm run check                  # Run both lint and format (with auto-fix)
```

## Architecture

### Factory Pattern for Dependencies

All core operations use factory functions that receive `G3Dependecies` (containing `G3Config`). This enables dependency injection and testability:

- `createGistFactory`, `deleteGistFactory` - Gist API operations
- `createG3FileFactory`, `parseG3FileFactory` - G3File creation/parsing
- `getG3EntriesFactory` - File entry retrieval

### Data Flow

1. **Upload (`cp`)**: File → chunk into pieces → encrypt each chunk → upload as separate Gists → store metadata in `~/.local/share/g3/files/<filename>.g3.json`
2. **Download (`get`)**: Read metadata → clone Gist repos to temp folder → decrypt chunks → concatenate into original file
3. **Metadata files**: JSON files with `.g3.json` extension track Gist IDs and URLs for each stored file

### Key Modules

- `src/config/` - Configuration parsing with Zod validation, defaults to `~/.config/g3/config.json`
- `src/gist/` - GitHub Gist API client using undici
- `src/crypto/` - AES-256-CBC encryption/decryption
- `src/pool/` - Piscina worker pool for parallel chunk processing and Gist cloning
- `src/g3file/` - G3File class representing stored files and their Gist metadata
- `src/fsdata/` - Filesystem operations for metadata JSON files
- `src/cmd/` - CLI command handlers

### Worker Pool

Uses Piscina for parallel processing:

- `processGistChunk` - Encrypts and uploads file chunks
- `cloneRepo` - Clones Gist repositories
- `decryptFile` - Decrypts downloaded files

### Configuration Schema

Required: `GITHUB_TOKEN`, `AES_KEY` (32-byte hex)
Optional: `DATA_FOLDER` (default: `~/.local/share/g3/files/`), `CHUNK_SIZE` (default: 10MB, min: 5MB)

### Test Workflow

Focus the testing approach to unit testing and focusing on unit tests for functions mainly.
