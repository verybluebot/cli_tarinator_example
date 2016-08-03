# Golang CLI Tarinator Example
## Genaral
Cli Tarinator is a simple example for usage of the tarinator-go package to
compress and decompress files and directories in Golang.

The package and its documentation can be found here:

https://github.com/verybluebot/tarinator-go

Here is a tutorial for createing this cli example:

https://www.youtube.com/watch?v=Gg0qfDglwhs

## Installing and Running

```
// get project files
git clone https://github.com/verybluebot/cli_tarinator_example.git
```

Then `cd` into project root directory and use the `make` utility to build the binary:
```
make all
```

Now run it (use -h to get help)
```
bin/main -h

Usage: CLI Tarinator [OPTIONS]

Options:
    -c, --compress   Compress files or/and directories (Path must be relative or absolute).
    -f, --file       Tar File name. Using tar.gz suffix will create .tar.gz using only .tar will create .tar file
    -e, --extract    Extract .tar or tar.gz file to that path.

    -h, --help       prints this help info.
```

## Licence
[MIT](https://github.com/verybluebot/cli_tarinator_example/blob/master/LICENCE.md)