```go
package main

import (
    "fmt"
    "github.com/jpadilla/ivona"
)

func main() {
    client := ivona.New("ACCESS_KEY", "SECRET_KEY")
    options := ivona.NewSpeechOptions("Hello World")
    r, err := client.CreateSpeech(options)
    fmt.Println(r, err)
}
```
