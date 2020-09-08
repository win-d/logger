# Simple logger

Cross-platform logger for personal use. 
The module creates log files in the selected directory and writes your messages.
The names of the subdirectories and log files depend on the current date.

## Example usage

```go
package main

import (
    "github.com/win-d/logger" // import this module
    "log"
)

func main() {
    err := logger.SetDir("/var/log/project") // set main log directory
    if err != nil {
        log.Fatalln(err.Error())
    }

    logger.Write("First line")
    logger.Write("Second line")
}
```

Here we specify the directory that will be the main for subdirectories and log files.
A subdirectory with the current month number will be created in the main directory.
A log file with the name of the current date will be created in the subdirectory.

---

Example of the resulting file path:
```
/var/log/project/05/15055020.log
---
/var/log/project - main directory
05 - subdirectory with the current month number
15052020.log - log file with the name of the current date
```

In the log file will we see something like:
```
2020/05/15 12:34:56 First line
2020/05/15 12:34:56 Second line
```