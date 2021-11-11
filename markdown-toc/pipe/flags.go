package pipe

import (
	"github.com/urfave/cli"
)

var Flags = []cli.Flag{
	cli.StringSliceFlag{
		Name:   "markdown-toc.pattern",
		Usage:  "Pattern for markdown. Supports multiple patterns with comma-separated values.",
		Value:  &cli.StringSlice{"README.md"},
		EnvVar: "MARKDOWN_TOC_PATTERNS,PLUGIN_MARKDOWN_TOC_PATTERNS",
	},
	cli.StringFlag{
		Name:        "markdown-toc.arguments",
		Usage:       "Pass in the arguments for markdown-toc.",
		Value:       "--bullets='-'",
		EnvVar:      "MARKDOWN_TOC_ARGUMENTS,PLUGIN_MARKDOWN_TOC_ARGUMENTS",
		Destination: &Pipe.Markdown.Arguments,
	},
}
