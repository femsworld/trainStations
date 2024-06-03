package functions

import (
    "fmt"
    "os"
)

func HandleError(err error) {
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        os.Exit(1)
    }
}
