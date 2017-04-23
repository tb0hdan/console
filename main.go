package main

import (
        "flag"

        "github.com/fatih/color"
       )

var flagNoColor = flag.Bool("no-color", false, "Disable color output")


// called by console.Console
func Handle_data(data string) (result string) {
    if data != "" {
        result = "You said " + data
    }
    return
}

func main() {
    flag.Parse()
    if *flagNoColor {
        color.NoColor = true // disables colorized output
    }

    green := color.New(color.FgGreen).SprintFunc()
    Console(green(">")+ " ")
}
