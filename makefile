
all: wx

wx: *.go
	go build github.com/morya/wx
	cp ${GOPATH}/src/github.com/morya/wx/wx ~/bin/
