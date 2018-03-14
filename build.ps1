cd dist
go build -ldflags -H=windowsgui ..\cc-ui

cd bin
go build ..\..\cmd\cc
go build ..\..\cmd\ccutil

cd ..\..