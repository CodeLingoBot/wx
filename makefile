target:=wx

all: $(target)

$(target): *.go
	go build .

install:$(target)
	@rm $(HOME)/bin/$(target)
	@cp $(GOPATH)/src/github.com/morya/$(target)/$(target) ~/bin/
	@cd $(HOME); supervisorctl restart $(target):
