package main

import (
    "io"
    "io/ioutil"
    "os"
    "strings"
    "log"
)

/**
 *  Reads all txt files in the dictionary directory  
 *  and encodes them as strings literals in textfiles.go
 */
func main() {
    dirname := "./dictionary/"
    fs, _ := ioutil.ReadDir(dirname)
    out, _ := os.Create("textfiles.go")
    out.Write([]byte("package main \n\nconst (\n"))
    for _, f := range fs {
        if strings.HasSuffix(f.Name(), ".txt") {
            out.Write([]byte(strings.TrimSuffix(f.Name(), ".txt") + " = `"))
            if f, err := os.Open(dirname + f.Name()); err != nil {
               log.Fatal(err)

            } else {
                io.Copy(out, f)
                out.Write([]byte("`\n"))
            }
        }
    }
    out.Write([]byte(")\n"))
}