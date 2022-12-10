# redl

A cli tool for downloading courses inside Mixin ecosystem.

```
NAME:
   redl - A simple powerful cli tool for downloading courses inside Mixin ecosystem.

USAGE:
   redl [global options] command [command options] [arguments...]

VERSION:
   0.1

COMMANDS:
   single, s  Download a single course
   range, r   Download all courses
   all, a     Download all courses
   help, h    Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --token value, -t value  collected after OAuth on the website (with Bearer prefix) (default: "Bearer ")
   --base value, -b value   base URL for downloading (default: "xuexi-courses-api.songy.info")
   --dir value, -d value    the output data directory
   --help, -h               show help (default: false)
   --version, -v            print the version (default: false)

```
## Example

```
$ ./redl s -t "Bearer xxx" -i 6000 -d ~/Downloads
```

## dev state

- [x] single
- [ ] range
- [ ] all
- [ ] resume
