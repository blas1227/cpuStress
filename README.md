Project
=================
CPUStress is a tool that makes possible to test your cooling system by stressing 
each CPU core maintining them 100% busy , make sure to monitor system temperature in
every moment in order to avoid any hardware hazard.

How to run
=================
There is nothing additional packages neccesary in order to run this tool , just 
go to bin/ folder and download the binary compiled for your Operating System

To run it in Linux/macOS
`./cpuStress`

To run it in Windows just perform a double click in the executable file.

How to Build
=================
This tool was developed in Go Language and can be compiled for Macos, Windows, Linux and Android OS.

Before try to compile be sure to have go compiler installed in your system, if not please go 
to following link , download and install it . [GoLang](https://golang.org/)

Once Go compiler installed you can compile it by perfoming the following command :
- Go to cpuStress folder
to compile it for you current system execute: 
`go build sc/cpuStress.go`
to compile it for cross-platform execute:
`env GOOS=target GOARCH=arch go build sc/cpuStress.go`
where target and arch can be taked following table below : 
|	target	|	arch	|
| ----------| ----------|
|	windows |	amd64	|
|	windows |	386	    |
|   linux   |	amd64   |
|	linux   |   arm     |

you can find more platforms available to be compiled on the next link : [cross-platform](https://www.digitalocean.com/community/tutorials/how-to-build-go-executables-for-multiple-platforms-on-ubuntu-16-04)

Version History
=================
#### 1.0
- Initial Release