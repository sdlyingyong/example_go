下载器



[用法] 
ty.exe ../tmp.mp4




[测速]


测速链接(100MB)


ty https://speedtest2.niutk.com:8080/download?size=100000000&r=0.5807917751032634


[编译]

编译为ty.exe

linux

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o ty main.go
