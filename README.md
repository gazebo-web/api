# Gazebo - API

This repository contains the API definitions of different services used in Gazebo projects.

## Quick start

### From releases

#### Go
```bash
go get github.com/gazebo-web/api
```

You can also download a specific version:
```bash
go get github.com/gazebo-web/api@v1.0.0
```

#### C++ and TypeScript
Download from our [Releases](https://github.com/gazebo-web/api/releases) page.

### Code generation (source)

#### Buf
This project uses [Buf](https://buf.build/) to build protobuf files. In order to install Buf, refer to their 
[installation docs](https://docs.buf.build/installation).

#### Building

A makefile is provided to build files. Running `make` generates code for all the languages we currently support.

The following commands are available to build files for specific languages.

##### Go
```bash
buf generate --template buf.gen.go.yaml
```

##### C++
```bash
buf generate --template buf.gen.cpp.yaml
```

##### TypeScript
```bash
buf generate --template buf.gen.ts.yaml
```

## Contributing
See [the contributing guide](CONTRIBUTING.md) for detailed instructions on how to get started with our development workflow.

## Releases
The Gazebo API project uses Semantic Versioning for its releases: `<Major>.<minor>.<patch>`.

For Gazebo API v1.2.3:
- `1` represents the major version.
- `2` represents the minor version.
- `3` represents the patch version.

For release candidate versions such as: `v1.3.3-rc0`
- `0` represents the release candidate version, it increases every time a new draft version is created for `v1.3.3`

NOTE: We prepend a `v` to the version.

### Policy

Our current release policy is:
- Bug fixes increase patch versions.
- New features increase minor versions.
- Breaking changes increase major versions.

### Guideline

1. Merge all the new features into the `main` branch.
2. [Create a new tag](https://github.com/gazebo-web/api/tags) increasing the minor or patch version depending on our policy. The `Release notes` section must be filled in with the features included in this new release. The tag created for this release should be: `vM.m.p`
3. Go to the latest Pipeline executed in the `main` branch, there will be a set of `build` jobs for the different languages that the Gazebo API support.
    1. Click on each of the jobs listed there and press the `Keep` button in order to avoid GitLab from auto-cleaning the artifacts generated by each job.
    2. After pressing the `Keep` button on each job, copy the URL from the `Download` button.
3. If the `Release notes` section was filled in correctly, you'll find a new `Release` in the [Releases](https://github.com/gazebo-web/api/releases) page. If not, go back to step 2 and fill out the respective `Release notes` section by editing the created tag.
4. Click on `Edit` button next to the Release title.
5. In the edit page, go to the `Release assets` section, add a new entry for each link copied in step 3.
    1. URL: The URL copied in step 3.
    2. Link title: Name of the language, example: C++.
    3. Type: Other.
6. Press `Save changes`, the release is now official.
7. Let the different teams consuming the Gazebo API know that a new release has been created.