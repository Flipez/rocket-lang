# How to release a new RocketLang version

In order to be ready for the release, have all changes in the main branch.

## 1. Create a tag
```
git tag -a v0.23.0 -m v0.23.0
```

This will run the tests and on success build and publish binaries to a draft release.

## 2. Update the changelog
```
docker run -it --rm -e CHANGELOG_GITHUB_TOKEN -v "$(pwd)":/usr/local/src/your-app githubchangeloggenerator/github-changelog-generator -u flipez -p rocket-lang
```

Locally on main, run the changelog generator and commit + push the new Changelog to main. This can be used for a more comprehensive text in the final release.

## 3. Update the documentation
Make sure the documentation is up to date
```
go run docs/generate.go
```

Create a new versioned doc, matching the created tag
```
yarn docusaurus docs:version v0.23.0
```

Once created, check `docs/versions.json` if any version can be removed.
If you decide to remove a version from the list, deleted the corresponding documentation and sidebar
```
rm -rf versioned_docs/version-v0.22.0/
rm versioned_sidebars/version-v0.22.0-sidebars.json
```

In `docs/docusaurus.config.js`, update the `lastVersion` value with the tag just created.

Commit and push the changes to master.

## 4. Publish the release
On the GitHub release page, update the text of the draft release with the last section of the CHANGELOG.md and publish the release.