<div align="center">

# GoStyle

</div>


<h4 align="center">GoStyle is a Go-To library for styling text output in terminals. It provides functions for coloring text, applying styles, adding text effects, and printing formatted log messages..</h4>
<p align="center">
<img src="https://img.shields.io/github/go-mod/go-version/iaakanshff/gostyle">
<!-- <a href="https://github.com/iaakanshff/gostyle/releases"><img src="https://img.shields.io/github/downloads/iaakanshff/gostyle/total"> -->
<a href="https://github.com/iaakanshff/gostyle/graphs/contributors"><img src="https://img.shields.io/github/contributors-anon/iaakanshff/gostyle">
<!-- <a href="https://github.com/iaakanshff/gostyle/releases/"><img src="https://img.shields.io/github/release/iaakanshff/gostyle"> -->
<a href="https://github.com/iaakanshff/gostyle/issues"><img src="https://img.shields.io/github/issues-raw/iaakanshff/gostyle">
<a href="https://github.com/iaakanshff/gostyle/stars"><img src="https://img.shields.io/github/stars/iaakanshff/gostyle">
<!-- <a href="https://github.com/iaakanshff/gostyle/discussions"><img src="https://img.shields.io/github/discussions/iaakanshff/gostyle"> -->
</p>

---
      
## Installation

To install GoStyle, use `go get`:

```sh
go get "github.com/iaakanshff/gostyle"
```

## Supported Styles


### Usage
Firstly, Import the library and Instantiate it as follows
```golang
package main

import (
  "github.com/iaakanshff/gostyle"
)

func main() {
  gs := GoStyle.New()
}
```

Printing colored text as follows
```golang
fmt.Println(gs.Red("This is Red Text!"))
```

You can also apply styles:
```golang
fmt.Println(gs.Italic("This string is in italic"))
```

Combining styles and colors:
```golang
fmt.Println(gs.Bold(gs.Red("This string is bold and red")))
```

### Labels

Labels are useful for printing minimal output schemas for different types of messages:
```golang
fmt.Println(gs.Info("This is an informational message"))
fmt.Println(gs.Que("This is a question"))
fmt.Println(gs.Bad("This is a bad message"))
fmt.Println(gs.Good("This is a good message"))
fmt.Println(gs.Run("This is a running message"))
```

### Output
![Labels](https://imgur.com/jXcHlhs)

## What's Included?

#### List of all colors


Normal Colors:
```
Black, Red, Green, Yellow, Blue, Magenta, Cyan, Gray, White
```

Bright Colors:
```
BrBlack, BrRed, BrGreen, BrYellow, BrBlue, BrMagenta, BrCyan, BrGray
```

#### List of all styles
```
Bold, BGround, Underline, Strike, Italic
```

#### List of all labels

```
Info, Que, Run, Bad, Good
```

#### List of all Effects

```
Blink, Reverse, Hidden
```

**Note:** Some terminals don't support special effects and Windows versions below windows 10 do not support ANSI escape sequences so the colors will not be printed in command prompt.

## Comparison with Other Go Libraries
When comparing GoStyle with the fatih/color library, GoStyle offers a more straightforward and intuitive API for text styling. While fatih/color provides a robust set of functionalities for colorizing text, it requires a more complex approach for applying styles like bold.

### fatih/color

```go
package main

import (
  "github.com/fatih/color"
)

func main() {
  red := color.New(color.FgRed).Add(color.Bold)
  red.Println("This string is bold and red")
}
```

### GoStyle
```golang
package main

import (
  "github.com/iaakanshff/gostyle"
)

func main() {
  gs := GoStyle.New()
  fmt.Println(gs.Bold(gs.Red("This string is red")))
}
```

- In contrast, GoStyle provides a cleaner and more user-friendly interface for applying various text styles.
- With GoStyle, users can easily chain multiple styles together, making it simpler to create richly formatted text output.
- Additionally, GoStyle offers a wider range of supported styles, including italic, underline, strike-through, and more, providing greater flexibility for customizing text appearance.


## Why GoStyle?

1. **Simplicity First**: GoStyle keeps things simple. With a clean and intuitive API, it's easy to style text without wrestling with complex configurations.

2. **Flexibility Matters**: Whether you want bold, italic, or striking text, GoStyle has you covered. Mix and match styles effortlessly to create the perfect look.

3. **Easy to Learn**: New to Go? No problem. GoStyle's straightforward interface makes it a breeze for beginners to start styling text right away.

4. **Rich Styling Options**: From colors to styles to effects, GoStyle offers a wide range of customization options to make your text stand out.

## Contribution

We believe that GoStyle has immense potential, and your contributions can help make it even better. Whether you're fixing bugs, adding new features, or improving documentation, your efforts are highly valued.

If you think GoStyle could benefit from additional colors, styles, or effects, don't hesitate to start a pull request. Together, we can continue to enhance GoStyle and make it the go-to choice for text styling in Go applications.
