GOCMD=go
GOBUILD=$(GOCMD) build
NAME=demo
PATHTOMAIN=./cmd/$(NAME)

all: build

build:
	$(GOBUILD) -o $(NAME) $(PATHTOMAIN)/main.go
