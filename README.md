# perm

Print human-friendly descriptions of a file's permissions

## Usage:

```
$ ls perm/perm.go
-rw-r--r-- 1 tiago tiago 906 ago 24 22:29 perm/perm.go
$ go run main.go --file perm/perm.go
owner can read and write to it
group members can read from it
others can read from it
```