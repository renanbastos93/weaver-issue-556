# weaver-issue-556
Simulate issue opened

## Pull request to fix this issue
https://github.com/ServiceWeaver/weaver/pull/772

## Issue
https://github.com/ServiceWeaver/weaver/issues/556

panic stack trace
```sh
weaverdemo git:(main) ✗ go run .                                                    
╭───────────────────────────────────────────────────╮
│ app        : weaverdemo                           │
│ deployment : 5535fd76-3c62-43ad-af45-44c2b9e44f2a │
╰───────────────────────────────────────────────────╯
runtime: goroutine stack exceeds 1000000000-byte limit
runtime: sp=0x14020612350 stack=[0x14020612000, 0x14040612000]
fatal error: stack overflow
...
exitt status 2
```