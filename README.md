# ITerminal
A very basic and easy to use Go library to print colored outputs.

### Installation
Install and update this go package with `go get -u github.com/egeoz/iterminal`

### Usage
Example of cleaning the terminal screen and changing the font color to red with white as the background:

```go
package main

import (
	"fmt"
	"github.com/egeoz/iterminal"
)

func main() {
    iterminal.ClearScreen()
    
    // MakeColor() method
    fmt.Println(iterminal.MakeColor(iterminal.ForegroundColor.Red, iterminal.BackgroundColor.White)+"Hello World!")
   
    // Don't change the background color
    fmt.Println(iterminal.MakeColor(iterminal.ForegroundColor.Red, "")+"Hello World!")

    // SetColor() method
    iterminal.SetColor(iterminal.ForegroundColor.Red, iterminal.BackgroundColor.White)
    fmt.Println("Hello World!")

}
```


Example of font attributes:
```go
    // Bold
    fmt.Println(iterminal.MakeFont(iterminal.Font.Bold)+"Hello World!")

    // Italic
    fmt.Println(iterminal.MakeFont(iterminal.Font.Italic)+"Hello World!")
    
    // Underline
    fmt.Println(iterminal.MakeFont(iterminal.Font.Underline)+"Hello World!")
    
    // Blinking
    fmt.Println(iterminal.MakeFont(iterminal.Font.Blinking)+"Hello World!")
    
    // Default
    fmt.Println(iterminal.MakeFont(iterminal.Font.Default)+"Hello World!")
    
    // via SetFont() method
    iterminal.SetFont(iterminal.MakeFont(iterminal.Font.Bold))
    fmt.Println("Hello World!")
```


Example of moving the terminal cursor:
```go
    // Moves the cursor to the very beginning of the terminal window
    iterminal.MoveCursor(0,0)

```
