package main

import (
	"log"

	"github.com/LaJase/servcli/cmd"
	"github.com/spf13/cobra/doc"
)

func main() {
	cmd := cmd.GetServCliCmd()
	cmd.DisableAutoGenTag = true
	cmd.InitDefaultCompletionCmd()
	err := doc.GenMarkdownTree(cmd, "./")
	if err != nil {
		log.Fatal(err)
	}
}
