package example

import "github.com/UtkarshGupta-CS/go-terminal-prompt"

var langs = []string{
	"c",
	"c++",
	"lua",
	"go",
	"js",
	"ruby",
	"python",
}

func main() {
	i := terminalprompt.Choose("What's your favorite language?", langs)
	println("picked: " + langs[i])
}
