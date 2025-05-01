package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/webercoder/go-dalle/client"
)

func usage(msg string) {
	usage := fmt.Sprintf(`Usage: %s "prompt"`, os.Args[0])
	log.Fatalf("%s\n%s\n", msg, usage)
}

func browser(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start", url}
	case "darwin":
		cmd = "open"
		args = []string{url}
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
		args = []string{url}
	}

	command := exec.Command(cmd, args...)
	return command.Start()
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		usage("no args provided")
	}

	prompt := args[0]

	c := client.NewDallEClient(OpenAIAPIKEY)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	resp, err := c.Request(ctx, client.DallERequest{
		Model:   client.ModelDallE3,
		Prompt:  prompt,
		Size:    client.Size1792x1024,
		Quality: client.QualityStandard,
		Count:   1,
	})

	if err != nil {
		log.Fatalf("error sending request: %v", err)
	}

	for i, data := range resp.Data {
		log.Printf("Result %d: %s (revised prompt: %s)", i+1, data.URL, data.RevisedPrompt)
		if err := browser(data.URL); err != nil {
			log.Printf("could not open browser to %v: %v", data.URL, err)
		}
	}
}
