// cmd/artificialmintstudio/main.go
package main

import (
"flag"
"log"
"os"

"artificialmintstudio/internal/artificialmintstudio"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := artificialmintstudio.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
