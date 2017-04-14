# runeset

runeset finds unique characters in a file (or somewhere else).

It reads stdin and prints all distinct Unicode characters found,
sorted, followed by a newline.

```sh
$ go get github.com/thomasheller/runeset
$ echo aaccaabaadfa > /tmp/alpha.txt
$ cat /tmp/alpha.txt | runeset

abcdf
```

Using the `-escape` flag, runeset prints single-quoted characters (Go
syntax).

```sh
$ cat /tmp/alpha.txt | runeset -escape
'\n''a''b''c''d''f'
```
