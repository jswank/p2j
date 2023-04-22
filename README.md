# p2j
Convert a Maven POM file to JSON.

Notable Features:
- input can be STDIN or specified by filename
- all keys are converted to lowercase
- POM elements which may have multiple children (like licenses or developers)
  will be converted to JSON arrays, even is there is only a single element.

## Install
```console
$ go install install github.com/jswank/p2j@latest
```

## Usage
```console
$ p2j foo.pom

$ cat foo.pom | p2j
```

## License
MIT
