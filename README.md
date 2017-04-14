# runeset

runset finds unique characters in a file (or somewhere else).

It reads stdin and prints all distinct Unicode characters found,
sorted, followed by a newline

```sh
$ go get github.com/thomasheller/runeset
$ echo aaccaabaadfa > /tmp/alpha.txt
$ cat /tmp/alpha.txt | runeset

abcdf
```
