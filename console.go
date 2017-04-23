package main

import (
       "bufio"
       "fmt"
       "os"
       "time"
       )


func get_prompt(prompt string) (result string) {
    t := time.Now()
    // Mon Jan 2 15:04:05 MST 2006, really?
    result = "[" + t.Format("15:04:05") + "]" + prompt
    return
}

func Console(prompt string) {
    new_prompt := get_prompt(prompt)
    fmt.Print(new_prompt)
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        response := Handle_data(scanner.Text())
        if response != "" {
            fmt.Println(response)
        }
        new_prompt = get_prompt(prompt)
        fmt.Print(new_prompt)
    }
}
