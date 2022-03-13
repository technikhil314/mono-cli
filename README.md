# 1command

Project based aliases for streamline your build and development flow.

# Problem

Right now when you are about run/build a new project you have to do the following

1. Look for what command starts the dev env
2. Look for what command build the project
3. and so for install dependancies, clean project and many other tasks

# Solution

1. Add this code to your favourite shell init script
2. install 1command using `npm install 1command`
3. and then run following commands for each described task

   | command | Purpose                            | Supports                     |
   | ------- | ---------------------------------- | ---------------------------- |
   | install | install dependancies               | npm, yarn, pnpn, lerna       |
   | build   | build your project                 | webpack, vite and frameworks |
   | clean   | cleans your build directory        | webpack, vite and frameworks |
   | start   | start your development environment | webpack, vite and frameworks |
