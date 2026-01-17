# Agent Instructions

This project uses **bd** (beads) for issue tracking. Run `bd onboard` to get started.

## Environment Setup

This project uses **Flox** for reproducible development environments. All tools are managed via flox.

```bash
flox activate        # Activate the environment (installs go, templ, goreleaser, prek, just)
```

The justfile uses `flox activate` automatically, so `just <recipe>` works without manual activation.

## Quick Reference

```bash
bd ready              # Find available work
bd show <id>          # View issue details
bd update <id> --status in_progress  # Claim work
bd close <id>         # Complete work
bd sync               # Sync with git
```

## Development Commands

```bash
just                  # List all recipes
just test             # Run tests (generates templ first)
just build            # Build the project
just fmt              # Format code (templ + go)
just pre-commit       # Run pre-commit on all files
just watch            # Watch for templ changes
```

## Pre-commit Hooks

Pre-commit is installed via `prek` (flox package). Always run before committing:

```bash
just pre-commit       # Or: prek run --all-files
```

To install git hooks:
```bash
just pre-commit-install
```

## Releasing

This project uses **goreleaser** for multi-platform releases (darwin/linux, amd64/arm64).

### Release Commands

```bash
just release-check       # Validate .goreleaser.yaml
just release-snapshot    # Build local snapshot (no publish)
just release 0.1.0       # Tag and push v0.1.0 (triggers GitHub Actions)
just release-local       # Build and publish locally (uses 1Password for token)
```

### Release Workflow

1. **Via GitHub Actions (recommended)**:
   ```bash
   just release 0.1.0    # Tags v0.1.0 and pushes
   ```
   GitHub Actions will build and publish automatically.

2. **Local release** (requires 1Password CLI):
   ```bash
   just release-local
   ```
   Uses token from: `op://homelab/pico GitHub Personal Access Token/token`

### Versioning

Use semantic versioning (semver):
- `0.1.0` - Initial release
- `0.1.1` - Patch (bug fixes)
- `0.2.0` - Minor (new features, backwards compatible)
- `1.0.0` - Major (breaking changes)

## CI/CD

- **CI workflow** (`ci.yaml`): Runs on feature branches and PRs
  - Runs `prek run --all-files`
  - Runs `just test`

- **Release workflow** (`release.yaml`): Runs on version tags (`v*`)
  - Builds binaries for darwin/linux (amd64/arm64)
  - Creates GitHub release with assets

## Landing the Plane (Session Completion)

**When ending a work session**, you MUST complete ALL steps below. Work is NOT complete until `git push` succeeds.

**MANDATORY WORKFLOW:**

1. **Run quality gates** (if code changed):
   ```bash
   just pre-commit
   just test
   ```
2. **File issues for remaining work** - Create issues for anything that needs follow-up
3. **Update issue status** - Close finished work, update in-progress items
4. **PUSH TO REMOTE** - This is MANDATORY:
   ```bash
   git pull --rebase
   bd sync
   git push
   git status  # MUST show "up to date with origin"
   ```
5. **Verify** - All changes committed AND pushed
6. **Hand off** - Provide context for next session

**CRITICAL RULES:**
- Work is NOT complete until `git push` succeeds
- NEVER stop before pushing - that leaves work stranded locally
- NEVER say "ready to push when you are" - YOU must push
- If push fails, resolve and retry until it succeeds
