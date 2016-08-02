package main

import(
    "fmt"
    "flag"
    log "logger"
    "os"

    "github.com/verybluebot/tarinator-go"
)

type stringsSlice []string

func showHelp() {
    fmt.Println(`

Usage: CLI Template [OPTIONS]

Options:
    -c, --compress   Compress files or/and directories (Path must be relative or absolute).
    -f, --file       Tar File name. Using tar.gz suffix will create .tar.gz using only .tar will create .tar file
    -e, --extract    Extract .tar or tar.gz file to that path.

    -h, --help       prints this help info.


    `)
}


func setFlag(flag *flag.FlagSet) {
    flag.Usage = func() {
        showHelp()
    }
}

func (i *stringsSlice) String() string {
    return fmt.Sprintf("%s", *i)
}

func (i *stringsSlice) Set(value string) error {
    *i = append(*i, value)
    return nil
}

func main() {
    var files stringsSlice
    flag.Var(&files, "c", "")
    flag.Var(&files, "compress", "")

    var target string
    flag.StringVar(&target, "f", "", "")
    flag.StringVar(&target, "file", "", "")

    var extract string
    flag.StringVar(&extract, "e", "", "")
    flag.StringVar(&extract, "extract", "", "")

    var sHelp bool
    flag.BoolVar(&sHelp, "h", false, "")
    flag.BoolVar(&sHelp, "help", false, "")

    setFlag(flag.CommandLine)

    flag.Parse()

    if sHelp {
        showHelp()
        return
    }

    //header.GetHeader()

    log.InitLogger()

    if target == "" {
        log.Warningln("No target tar file was specified")
        os.Exit(1)
    }

    if len(files) > 0 {
        for _, f := range files {
            if _, err := os.Stat(f); os.IsNotExist(err) {
                log.Errorf("File or direcotroy does not exists: %s", f)
                os.Exit(1)
            }
        }

        log.Printf("Start taring files/directories: %v", files)
        err := tarinator.Tarinate(files, target)
        if err != nil {
            log.Errorf("Failed taring files: %s", err)
            os.Exit(1)
        }
    }

    if extract != "" {
        log.Printf("Start untaring files/directories: %v", files)
        err := tarinator.UnTarinate(extract, target)
        if err != nil {
            log.Errorf("Failed untaring file: %s", err)
            os.Exit(1)
        }
    }
}