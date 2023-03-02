# Gazebo - API - Contributing Guide

In this guide you'll get an overview of the contribution workflow for adding new features to the Gazebo API project.

[[_TOC_]]

## Getting started
The API repository contains a set of protobuf definitions used across multiple services and clients. These protobuf definitions can be found in the `proto` folder.
This folder includes:
- Domain definitions
- Simulation events
- gRPC services

Feel free to navigate our codebase to understand a little bit more about how things relate to each other, all our codebase includes documentation. Have some feedback? Feel free to open an issue!

We follow the [Feature Branch](https://www.atlassian.com/git/tutorials/comparing-workflows/feature-branch-workflow) workflow.

### Make changes
1. Create a working branch from `main`. New feature branches should be named following this syntax: `feature/<name>`.
2. Starting working on your changes.

### Merge Request
When you're finished with your changes, create a Merge Request, also known as MR. Adding a brief description of the changes and the reason why these changes were necessary will help reviewers understand your changes.

### Merging
The Gazebo Web team will review and merge your MR once it's been approved by a Reviewer. Your contribution will now be part of the upcoming release.

## Bugfixes
We have defined a workflow for bug-fixing that will guide you through the process of adding changes to not only the version that has the specific bug, but also for porting that fix into different releases.

Imagine you've found a bug in version `1.1.3`, here are the steps you need to follow in order to create a bugfix:

1. `git checkout tags/v1.1.3 -B fix/<name>`.
1. Work on your changes.
1. `git checkout tags/v1.1.3 -B release/v1.1.4`.
1. Create a MR from fix/<name> into `release/v1.1.4`.
1. Create a new git tag for `v1.1.4` following the `Releases guideline` defined in the [README](README.md).
1. Create a MR from `release/v1.1.4` into main. (And any other versions that this bugfix should be ported to).
