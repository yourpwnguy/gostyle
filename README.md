<div align="center">

![GoStyle LOGO](https://i.imgur.com/pnwMlty.png)

</div>

</div>
<h4 align="center">Simple, fast, and efficient tool for decoding JWT tokens, supporting base64 and JSON Web Token (JWT) formats..</h4>
<p align="center">
<img src="https://img.shields.io/github/go-mod/go-version/yourpwnguy/gostyle">
<!-- <a href="https://github.com/yourpwnguy/gostyle/releases"><img src="https://img.shields.io/github/downloads/yourpwnguy/gostyle/total"> -->
<a href="https://github.com/yourpwnguy/gostyle/graphs/contributors"><img src="https://img.shields.io/github/contributors-anon/yourpwnguy/gostyle">
<!-- <a href="https://github.com/yourpwnguy/gostyle/releases/"><img src="https://img.shields.io/github/release/yourpwnguy/gostyle"> -->
<a href="https://github.com/yourpwnguy/gostyle/issues"><img src="https://img.shields.io/github/issues-raw/yourpwnguy/gostyle">
<a href="https://github.com/yourpwnguy/gostyle/stars"><img src="https://img.shields.io/github/stars/yourpwnguy/gostyle">
<!-- <a href="https://github.com/yourpwnguy/gostyle/discussions"><img src="https://img.shields.io/github/discussions/yourpwnguy/gostyle"> -->
</p>

---
      
## Installation

To install GoStyle, use `go get`:

```sh
go get "github.com/yourpwnguy/gostyle"
```

## Supported Styles

Following styles are supported

![Styles](https://i.imgur.com/kvtZzh6.png)

Following colors are supported

![Colors](https://i.imgur.com/l97CagL.png)

Following labels are supported

![Labels](https://i.imgur.com/U3j4dJv.png)

### Usage
Firstly, Import the library and Instantiate it as follows
```golang
package main

import (
  "github.com/yourpwnguy/gostyle"
)

func main() {
  g := gostyle.New()
}
```

Printing colored text as follows
```golang
fmt.Println(g.Red("This is Red Text!"))
```

You can also apply styles:
```golang
fmt.Println(g.Italic("This string is in italic"))
```

Combining styles and colors:
```golang
fmt.Println(g.Bold(g.Red("This string is bold and red")))
```

### Labels

Labels are useful for printing minimal output schemas for different types of messages:
```golang
fmt.Println(g.LogGood("Good"))
fmt.Println(g.LogBad("Bad"))
fmt.Println(g.LogInfo("Informational"))
fmt.Println(g.LogRun("Processing"))
fmt.Println(g.LogQue("Question"))
```

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
LogInfo, LogQue, LogRun, LogBad, LogGood
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
  "github.com/yourpwnguy/gostyle"
)

func main() {
  g := gostyle.New()
  fmt.Println(g.Bold(g.Red("This string is red")))
}
```

- Provides a cleaner and more user-friendly interface for applying various text styles.
- Users can easily chain multiple styles together, making it simpler to create richly formatted text output.
- Offers a wider range of supported styles, including italic, underline, strike-through, and more, providing greater flexibility for customizing text appearance.


## Why GoStyle?

I mean Why Not? It is **simple**, **fast** and **awesome** at the same time !!!

1. **Simplicity First**: GoStyle keeps thing simple. With a clean and intuitive API, it's easy to style text without wrestling with complex configurations.

2. **Flexibility Matters**: Whether you want bold, italic, or striking text, GoStyle has you covered. Mix and match styles effortlessly to create the perfect look.

3. **Easy to Learn**: New to Go? No problem. GoStyle's straightforward interface makes it a breeze for beginners to start styling text right away.

4. **Rich Styling Options**: From colors to styles to effects, GoStyle offers a wide range of customization options to make your text stand out.

## Contribution

We believe that GoStyle has immense potential, and your contributions can help make it even better. Whether you're fixing bug, adding new features, or improving documentation, your efforts are highly valued.

If you think GoStyle could benefit from additional colors, styles, or effects, don't hesitate to start a pull request. Together, we can continue to enhance GoStyle and make it the go-to choice for text styling in Go applications.

For more information about the library, please refer to [Documentation](https://pkg.go.dev/github.com/yourpwnguy/gostyle).
