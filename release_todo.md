# Release Process TODO - Circular Dependency Issue

## Problem Description

There's a circular dependency in the release process:

1. **Examples need the latest goat package** - The examples should reference the new version being released
2. **Goat package isn't ready until release** - The version doesn't exist until CI publishes it
3. **Release should include updated examples** - The git tag should point to a commit with updated example dependencies
4. **But examples can't be updated until after release** - Because the version doesn't exist yet

## Current Issue

When running the release script, it updates the examples to use the new version (e.g., `v0.3.48`), but then:

- The examples reference a version that doesn't exist yet
- Any `go mod tidy` or Go build operations fail with:
  ```
  github.com/peterszarvas94/goat@v0.3.48: invalid version: unknown revision v0.3.48
  ```
- CI fails because it can't build/test the examples

## Potential Solutions to Investigate

### Option 1: Two-Phase Release
1. Create release with examples still on old version
2. After CI publishes, update examples in a separate commit
3. **Problem**: Release tag doesn't include updated examples

### Option 2: Use Replace Directives Permanently
1. Keep `replace github.com/peterszarvas94/goat => ../..` in example go.mod files
2. Examples always use local development version
3. **Problem**: Examples don't demonstrate real-world usage

### Option 3: Pre-release Workflow
1. Create a pre-release/RC tag first
2. Update examples to use the pre-release
3. Create final release
4. **Problem**: More complex workflow, still has timing issues

### Option 4: Separate Example Repository
1. Move examples to a separate repository
2. Update examples after each goat release
3. **Problem**: Examples separated from main codebase

### Option 5: Build-time Version Injection
1. Keep examples on a placeholder version (e.g., `v0.0.0-dev`)
2. Use build scripts to inject the correct version during CI
3. **Problem**: Complex build process

### Option 6: Skip Example Testing During Release
1. Temporarily disable example builds/tests during release process
2. Update examples after release, test separately
3. **Problem**: No validation that examples work with new version

## Current Workaround

The release script currently:
1. Uses `go mod edit -replace` to point to local code temporarily
2. Updates version with `go mod edit -require`
3. Removes replace directive
4. Commits and tags
5. **Skips `go mod tidy`** to avoid download errors

But CI still fails when trying to build/test the examples.

## Next Steps

1. **Investigate CI configuration** - Can we skip example testing during release builds?
2. **Consider workflow changes** - Maybe examples should be updated post-release?
3. **Explore Go module alternatives** - Are there other ways to handle this dependency cycle?
4. **Test pre-release approach** - Would using pre-release tags solve the timing issue?

## Files Involved

- `scripts/release.sh` - Release automation script
- `.github/workflows/release.yml` - CI release workflow  
- `examples/*/go.mod` - Example dependency declarations
- `go.mod` - Main project dependencies

### Option 7: No go.mod for Examples During Development

**Concept**: Remove go.mod files from examples during development, generate them only during release.

**Steps needed**:

1. **Remove example go.mod files from git tracking**:
   ```bash
   git rm examples/*/go.mod examples/*/go.sum
   echo "examples/*/go.mod" >> .gitignore
   echo "examples/*/go.sum" >> .gitignore
   ```

2. **Create go.mod template for each example**:
   ```bash
   # Create templates like examples/bare/go.mod.template
   module bare
   
   go 1.24.1
   
   require (
       github.com/a-h/templ v0.3.906
       github.com/peterszarvas94/goat {{VERSION}}
   )
   ```

3. **Update development workflow**:
   - Use `go work` for development (already have go.work)
   - Examples run using workspace, no individual go.mod needed
   - Add examples to go.work if not already there

4. **Modify release script**:
   ```bash
   # Generate go.mod files from templates during release
   for example_dir in examples/*/; do
       if [ -f "$example_dir/go.mod.template" ]; then
           sed "s/{{VERSION}}/$VERSION/g" "$example_dir/go.mod.template" > "$example_dir/go.mod"
           cd "$example_dir"
           go mod tidy  # This works because version exists after tag creation
           cd - > /dev/null
       fi
   done
   ```

5. **Update CI/build processes**:
   - Ensure go.work is used for development builds
   - Generate go.mod files before testing examples in CI
   - Or skip example testing during release, test them post-release

**Advantages**:
- No circular dependency during development
- Examples always use local development version
- Clean release process with proper versioning
- No complex replace directives

**Disadvantages**:
- Examples don't have go.mod during development (unusual)
- Need to maintain templates
- CI needs to handle both scenarios (dev vs release)

**Implementation checklist**:
- [ ] Verify go.work includes all examples
- [ ] Create go.mod.template files
- [ ] Update .gitignore
- [ ] Remove existing go.mod files from git
- [ ] Test development workflow without example go.mod files
- [ ] Update release script to generate go.mod from templates
- [ ] Update CI to handle missing go.mod files during development
- [ ] Document the new workflow for contributors

## Decision Needed

Choose which approach to implement based on:
- Complexity vs. maintainability
- Developer experience
- CI reliability
- Example accuracy/usefulness

### Option 8: Two-Step Release Process

**Concept**: Release the main package first, then update examples in a follow-up commit/release.

**Workflow**:

1. **Step 1 - Release main package**:
   ```bash
   # Release script only updates main package, leaves examples unchanged
   git tag v0.3.49
   git push origin v0.3.49
   # CI publishes the package
   ```

2. **Step 2 - Update examples after package is live**:
   ```bash
   # Wait for CI to publish (or check that version exists)
   # Then update examples to use the new version
   for example_dir in examples/*/; do
       cd "$example_dir"
       go mod edit -require="github.com/peterszarvas94/goat@v0.3.49"
       go mod tidy  # This works because v0.3.49 now exists
       cd - > /dev/null
   done
   git add examples/*/go.mod examples/*/go.sum
   git commit -m "chore: update examples to use goat v0.3.49"
   git push origin main
   ```

**Questions to consider**:
- **Do we need a new version for the example updates?** 
  - Probably not - it's just updating examples, not changing the main package
  - The examples update can be a regular commit on main
  - Or could be a patch version (v0.3.49 → v0.3.50) if you want releases to always include latest examples

**Advantages**:
- Simple and straightforward
- No circular dependency issues
- Examples always work and reference real published versions
- Can be automated with a two-step script

**Disadvantages**:
- Release tag doesn't include updated examples (they come in next commit)
- Two-step process is more complex
- Time gap between release and example updates

**Implementation options**:

**Option 8a - Single script with wait**:
```bash
# Release main package
git tag $VERSION && git push origin $VERSION

# Wait for package to be published
echo "Waiting for package to be published..."
sleep 60  # or poll until version exists

# Update examples
# ... update logic here
```

**Option 8b - Two separate commands**:
```bash
./scripts/release.sh v0.3.49        # Release main package
./scripts/update-examples.sh v0.3.49 # Update examples after CI finishes
```

**Option 8c - Examples get patch version**:
```bash
./scripts/release.sh v0.3.49         # Release main package  
./scripts/update-examples.sh v0.3.50 # Update examples + create new patch release
```

**Option 7 (no go.mod during dev) seems promising** because it:
- Eliminates the circular dependency entirely
- Uses Go workspaces as intended for monorepo development
- Generates proper go.mod files only when needed (release time)
- Keeps examples simple and always working during development

**Option 8 (two-step release) is also viable** because it:
- Keeps the current structure intact
- Solves the timing issue with a simple delay
- Could be automated into a single command that waits
- Examples always reference real published versions