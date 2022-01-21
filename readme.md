# What is this?

This is command line program to read text from standard in and write text to standard out converting github style emoji text to unicode characters.

This is go lang implementation of [mrowa44/emojify](https://github.com/mrowa44/emojify).

# installation

```bash
% go install github.com/byplayer/emojify/v2/cmd/emojify
```

# How to use it

## command line

```bash
% echo :+1: nice! | emojify
```

## git log

```bash
% git log | emojify | less -r
```

## git configuration

If you set below configuration in .gitconfig, then all git log uses emojify filter.

```
[pager]
  log = emojify | less -r
```
