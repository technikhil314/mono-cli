# mono-cli

A single command to do all your tasks across any project. This cli aims to supports any runtime, language, framework, package manager, test framework etc.

# Problem

Right now to run new project (where you are new to the project) you have to do the following

1. read the docs just to find what should you run to get your development environment up
2. But over the years dues tools like npm,yarn,maven,gradle the commands have some de-facto standard viz
   1. with `npm` you will run `npm start` or `npm run dev`
   2. with `maven` you will run `maven start`
   3. and so on
3. yet every time you have to read throgh docs to find what command to run which is not good DX

# Solution

1. This cli tries to abstract away the actual commands as these commands are very tightly coupled with the framework or tech stack you are choosing.
2. The idea is to give you abstraction on top of some industry standard commands to get rid of your hustle.

# How does it work

1. It tries to detect multiple deciding factors like runtime, package manager, framework, bundler, linter, test runner, test framework
2. Then it runs the default commands for each of mentioned factor like `npm install` or `webpack serve` etc

# Advantages of using mono-cli

1. imagine you have to jump between multiple projects like nextjs, create-react-app, remix.run etc.
2. When you jump from nextjs to create-react-app the command change from `npm run dev` to `npm start`
3. And you have to either recall that or check in the docs
4. With mono-cli you have to remember just the verbs like install, build, clean, start etc and prefix them with `mono` and you have one command doing same job across projects.
5. mono-cli internally detects things like what runtime, package manager, dev start commands to run and runs those commands on your behalf.
6. See the [usage](#usage) section below to see the commands and their defaults

# installation

### Using brew

```bash
brew install technikhil314/mono-cli/mono-cli
```

### Manual

Download latest and specific artifact for your OS and arch from the [releases](https://github.com/technikhil314/mono-cli/releases) section.

and run following commands using any POSIX shell viz sh, bash, ksh, fish, zsh etc

#### For UNIX like OS

```bash
tar -xvzf <path to the zip you downloaded>
cp ./mono /usr/local/bin
```

# Default subcommand nomenclature

| command         | Purpose                                    | Current support    |
| --------------- | ------------------------------------------ | ------------------ |
| bootstrap       | Install all dependancies mentioned in file | :white_check_mark: |
| install         | Install individual depndancy               | TO DO              |
| build           | build for production deployment            | TO DO              |
| devbuild        | build with robust debugging enables        | TO DO              |
| lint            | run linter                                 | TO DO              |
| test            | run all tests                              | TO DO              |
| testfile        | run tests in written in a file             | TO DO              |
| testpattern     | run all tests that match a pattern         | TO DO              |
| cover           | run all tests and generate coverage report | TO DO              |
| clean           | delete all dependancies                    | TO DO              |
| cleancache      | clear package managers cache               | TO DO              |
| cleanbuildcache | clear build cache                          | TO DO              |

# Usage

1. Add this code to your favourite shell init script
2. This cli follows a Unix standard command pattern so any arguments after `--` will be passed further down to actual commands like `mvn`, `npm` etc
3. and then run following commands for each described task

# Current support

| command         | npm                | yarn               | pnpm               | lerna              |
| --------------- | ------------------ | ------------------ | ------------------ | ------------------ |
| bootstrap       | :white_check_mark: | :white_check_mark: | :white_check_mark: | :white_check_mark: |
| install         | :white_check_mark: | :white_check_mark: | :white_check_mark: | :x:                |
| build           | TODO               | TODO               | TODO               | TODO               |
| devbuild        | TODO               | TODO               | TODO               | TODO               |
| lint            | TODO               | TODO               | TODO               | TODO               |
| test            | TODO               | TODO               | TODO               | TODO               |
| testfile        | TODO               | TODO               | TODO               | TODO               |
| testpattern     | TODO               | TODO               | TODO               | TODO               |
| cover           | TODO               | TODO               | TODO               | TODO               |
| clean           | TODO               | TODO               | TODO               | TODO               |
| cleancache      | TODO               | TODO               | TODO               | TODO               |
| cleanbuildcache | TODO               | TODO               | TODO               | TODO               |

# ?????? Notice

This is still a WIP project. Feel free to raise issues, contribute or suggest. TIA.

# License

MIT
