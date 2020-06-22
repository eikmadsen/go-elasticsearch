package main

import (
	"github.com/elastic/go-elasticsearch/v5/internal/cmd/generate/commands"
	_ "github.com/elastic/go-elasticsearch/v5/internal/cmd/generate/commands/gensource"
	_ "github.com/elastic/go-elasticsearch/v5/internal/cmd/generate/commands/genstruct"
	_ "github.com/elastic/go-elasticsearch/v5/internal/cmd/generate/commands/gentests"
)

func main() {
	commands.Execute()
}
