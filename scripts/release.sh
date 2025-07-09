#!/bin/bash
set -e

VERSION=$1
if [ -z "$VERSION" ]; then
    echo "Usage: $0 <version>"
    echo "Example: $0 v1.2.3"
    exit 1
fi

# Validate version format
if [[ ! $VERSION =~ ^v[0-9]+\.[0-9]+\.[0-9]+$ ]]; then
    echo "Error: Version must be in format v1.2.3"
    exit 1
fi

echo "Preparing release: $VERSION"

# Check if tag already exists locally
if git tag -l | grep -q "^$VERSION$"; then
    echo "Error: Tag $VERSION already exists locally"
    exit 1
fi

# Check if tag exists on remote
if git ls-remote --tags origin | grep -q "refs/tags/$VERSION$"; then
    echo "Error: Tag $VERSION already exists on remote"
    exit 1
fi

# Check for uncommitted changes
if ! git diff-index --quiet HEAD --; then
    echo "Error: You have uncommitted changes"
    git status --porcelain
    exit 1
fi

# Check if we're on main branch (optional safety check)
CURRENT_BRANCH=$(git branch --show-current)
if [ "$CURRENT_BRANCH" != "main" ] && [ "$CURRENT_BRANCH" != "master" ]; then
    echo "Warning: You're not on main/master branch (current: $CURRENT_BRANCH)"
    read -p "Continue anyway? (y/N): " -n 1 -r
    echo
    if [[ ! $REPLY =~ ^[Yy]$ ]]; then
        exit 1
    fi
fi

# First update examples to use the new version (before creating tag)
echo "Updating goat dependency to $VERSION in examples..."
for example_dir in examples/*/; do
    if [ -f "$example_dir/go.mod" ]; then
        echo "  Updating $(basename "$example_dir")"
        cd "$example_dir"
        # Use local replace temporarily since the version doesn't exist yet
        go mod edit -replace="github.com/peterszarvas94/goat=../.."
        go mod edit -require="github.com/peterszarvas94/goat@$VERSION"
        # Remove the replace directive
        go mod edit -dropreplace="github.com/peterszarvas94/goat"
        cd - > /dev/null
    fi
done

# Remove go.sum files to avoid checksum mismatches
git rm examples/*/go.sum 2>/dev/null || true

# Add go.sum to gitignore if not already there
if ! grep -q "examples/\*/go.sum" .gitignore; then
    echo "examples/*/go.sum" >> .gitignore
fi

# Stage and commit the updated go.mod files
git add examples/*/go.mod .gitignore
if ! git diff --cached --quiet; then
    git commit -m "chore: update goat dependency to $VERSION in examples"
    echo "Updated example dependencies"
else
    echo "No dependency updates needed"
fi

# Now create and push tag (so it includes the updated dependencies)
echo "Creating tag $VERSION..."
git tag -a "$VERSION" -m "Release $VERSION"

echo "Pushing tag and main branch to remote..."
git push origin main
git push origin "$VERSION"

echo "✅ Release $VERSION created successfully!"
echo "GitHub Actions will now build and publish the release."
echo "Check: https://github.com/peterszarvas94/goat/actions"