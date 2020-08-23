# glow
This is the backend behind [glow]() web application (available soon).
For its frontend see [picard](https://github.com/smeruelo/picard).

## What is glow?
__glow__ is a simple time tracking tool.
It lets you define and tag your different projects and it tracks the time you spend working on each of them.

Current features / WIP:
* Add / edit / delete projects
* Track time dedications
* See how much time you've dedicated to each project today / this week / in total

Features to be added in the sort run:
* simple goals (daily, weekly)
* user accounts and authentication
* reports
* enter time dedications manually

Features to be added in the long run:
* complex goals
* graphical reports
* archive non-active projects
* calendar
* alerts

## Build
```bash
docker build -t glow_server .
```

## Disclaimer
The main purpose of this project is to learn.
The reasoning behind some design decisions might be just to learn about some specific approach,
even if it's not the best one.
