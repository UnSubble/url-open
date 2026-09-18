package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/unsubble/url-open/internal/browser"
	"github.com/unsubble/url-open/internal/input"
)

var version = "dev"

func main() {
	startFlag := flag.Int("start", 0, "start index (inclusive)")
	endFlag := flag.Int("end", -1, "end index (exclusive, [index <= 0] = last)")
	batchFlag := flag.Int("batch", 1, "number of URLs to open in each batch")

	versionFlag := flag.Bool("V", false, "print version")
	verbose := flag.Bool("v", false, "enable verbose output")

	flag.Parse()

	if *versionFlag {
		fmt.Printf("url-open %s\n", version)
		return
	}

	if *startFlag < 0 {
		fmt.Fprintln(os.Stderr, "start index cannot be negative")
		os.Exit(1)
	}

	if *batchFlag < 1 {
		fmt.Fprintln(os.Stderr, "batch must be greater than 0")
		os.Exit(1)
	}

	source, err := createSource(flag.Args())
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	b := browser.New()

	index := 0
	batch := make([]string, 0, *batchFlag)

	err = source.Read(func(url string) error {
		if index < *startFlag {
			index++
			return nil
		}

		if *endFlag > 0 && index >= *endFlag {
			return nil
		}

		index++

		batch = append(batch, url)

		if len(batch) < *batchFlag {
			return nil
		}

		if err := openBatch(b, batch, *verbose); err != nil {
			return err
		}

		batch = batch[:0]

		return nil
	})

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if len(batch) > 0 {
		if err := openBatch(b, batch, *verbose); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}

func openBatch(b *browser.Browser, urls []string, verbose bool) error {
	for _, url := range urls {
		if verbose {
			fmt.Printf("opening %s\n...", url)
		}

		if err := b.Open(url); err != nil {
			fmt.Fprintf(
				os.Stderr,
				"failed to open %q: %v\n",
				url,
				err,
			)
		}
	}

	return nil
}

func createSource(args []string) (input.Source, error) {
	switch len(args) {
	case 0:
		return input.NewStdinSource(), nil

	case 1:
		_, err := os.Stat(args[0])

		if os.IsNotExist(err) {
			return input.NewArgsSource(args), nil
		}

		if err != nil {
			return nil, fmt.Errorf("failed to access %q: %w", args[0], err)
		}

		return input.NewFileSource(args[0])

	default:
		return input.NewArgsSource(args), nil
	}
}
