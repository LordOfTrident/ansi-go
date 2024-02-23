<a name="readme-top"></a>
<div align="center">
	<a href="./LICENSE">
		<img alt="License" src="https://img.shields.io/badge/license-MIT-e8415e?style=for-the-badge">
	</a>
	<a href="https://github.com/LordOfTrident/ansi-go/graphs/contributors">
		<img alt="Contributors" src="https://img.shields.io/github/contributors/LordOfTrident/ansi-go?style=for-the-badge&color=f36a3b">
	</a>
	<a href="https://github.com/LordOfTrident/ansi-go/stargazers">
		<img alt="Stars" src="https://img.shields.io/github/stars/LordOfTrident/ansi-go?style=for-the-badge&color=efb300">
	</a>
	<a href="https://github.com/LordOfTrident/ansi-go/issues">
		<img alt="Issues" src="https://img.shields.io/github/issues/LordOfTrident/ansi-go?style=for-the-badge&color=0fae5e">
	</a>
	<a href="https://github.com/LordOfTrident/ansi-go/pulls">
		<img alt="Pull requests" src="https://img.shields.io/github/issues-pr/LordOfTrident/ansi-go?style=for-the-badge&color=4f79e4">
	</a>
	<br><br><br>
	<img src="./res/logo.png" width="350px">
	<h1 align="center">ansi-go</h1>
	<p align="center">✨ ANSI escape sequences wrapper library for Go</p>
	<p align="center">
		<a href="#documentation">Documentation</a>
		·
		<a href="https://github.com/LordOfTrident/ansi-go/issues">Report Bug</a>
		·
		<a href="https://github.com/LordOfTrident/ansi-go/issues">Request Feature</a>
	</p>
	<br>
</div>

<details>
	<summary>Table of contents</summary>
	<ul>
		<li><a href="#introduction">Introduction</a></li>
		<li><a href="#usage">Usage</a></li>
		<li><a href="#example">Example</a></li>
		<li><a href="#documentation">Documentation</a></li>
		<li><a href="#bugs">Bugs</a></li>
	</ul>
</details>

## Introduction
Since i like to always reinvent the wheel, heres the [1000th](https://pkg.go.dev/search?q=ansi&m=)
wrapper library around [ANSI escape sequences](https://en.wikipedia.org/wiki/ANSI_escape_code) for Go.

## Usage
Add this package to your project
```
$ go get github.com/LordOfTrident/ansi-go
```

And import it
```go
package main

import "github.com/LordOfTrident/ansi-go"

func main() {
	style := ansi.Bold + ansi.LightMagenta

	ansi.Println(ansi.ClearScreen, ansi.Goto(1, 1), style, "Hello, world!")
}
```

## Example
```
$ cd example
$ go build
$ ./example
```

## Documentation
This library is very small and the code itself is pretty much documentation. Just look around
the source files.

## Bugs
If you find any bugs, please, [create an issue and report them](https://github.com/LordOfTrident/ansi-go/issues).

<br>
<h1></h1>
<br>

<div align="center">
	<a href="https://go.dev/">
		<img alt="Go" src="https://img.shields.io/badge/Go-007d9c?style=for-the-badge&logo=go&logoColor=white">
	</a>
	<a href="https://www.cse.psu.edu/~kxc104/class/cse472/09f/hw/hw7/vt100ansi.htm">
		<img alt="ANSI" src="https://img.shields.io/badge/ANSI%2FVT100-d9653f?style=for-the-badge&logo=ansi&logoColor=white">
	</a>
	<p align="center">Made with ❤️ love</p>
</div>

<p align="right">(<a href="#readme-top">Back to top</a>)</p>
