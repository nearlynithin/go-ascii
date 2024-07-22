<div align="center">


# ASCII art
*Converting images into ascii art!*


### Built with GO


![GO](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)


`ASCII art` means printing an image using ASCII characters with a monospace font.
This project explains the methods used to convert an image into ASCII art mainly using [Relative luminance](https://en.wikipedia.org/wiki/Relative_luminance) method.


</div>


---


![screenshot](/assets/screenshot2.png)
_Example from_ `assets/moe.png`


---


# Build


1. Clone the repository


```
git clone https://github.com/nearlynithin/go-ascii.git
```


2. cd into `go-ascii `


```
cd go-ascii/
```


3. install dependencies


```
go mod download
```


4. build command


```
go build -o go-ascii
```


5. Add an image filename after the executable file


```
./go-ascii <image.png>
```


6. You can list all the images in the `assets` directory


```
ls assets/
```
---

# Flags
- `-color` to print with color
```
./go-ascii -color <image.png>
```

---
# File Structure


```
.
├── assets      //All your images go here
│   ├── eyes.jpg
│   ├── golang.png
│   ├── moe.png
│   ├── monogatari.jpeg
│   └── senjougahara.png
├── go.mod
├── go.sum
├── image_loader.go
├── main.go
├── printer.go
├── README.md
└── transformer.go
```
---
# Usage


- [x] Make sure the terminal is zoomed out to get a better output
- [x] Supported file formats are `jpg`,`jpeg`,`png`.
- [x] To get rid of the blank space on the right side of the terminal, remove the lines `24` to `35` from `transformer.go`
- [x] _Smaller the terminal, Better the image :)_


---


Feel free to improve things and make a pull request,
*Thanks in Advance!*


> [!TIP]
> Pending stuff :P
> - Better algorithm to calculate aspect ratio according to the terminal width
> - Instructions for building on windows



> [!NOTE]
> This was built on `Ubuntu 24.04 LTS`

