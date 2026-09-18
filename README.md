# url-open

`url-open` is a small command-line utility that opens URLs in your system's default web browser.

URLs can be provided directly as command-line arguments, read from a file, or streamed through `stdin`. This makes it convenient to use with Unix pipelines and tools such as `xargs`.

## Installation

### From source

Requires Go 1.21 or later.

```bash
go install github.com/unsubble/url-open/cmd/url-open@latest
```

The binary will be installed to your Go binary directory.

Make sure your Go binary directory is in your `PATH`.

### Build from source

```bash
git clone https://github.com/unsubble/url-open.git
cd url-open

go build -o url-open ./cmd/url-open
```

## Usage

Open URLs from a file:

```bash
url-open urls.txt
```

Open URLs directly:

```bash
url-open https://example.com https://example.org
```

Read URLs from `stdin`:

```bash
cat urls.txt | url-open
```

Use with `xargs`:

```bash
cat urls.txt | xargs url-open
```

## Options

```text
-start
    Start index (inclusive). Default: 0.

-end
    End index (exclusive). Values <= 0 mean the end of the input.
    Default: -1.

-batch
    Number of URLs to open in each batch. Default: 1.
```

For example:

```bash
url-open -start 10 -end 20 urls.txt
```

opens URLs with indexes `10` through `19`.

To open URLs in batches:

```bash
url-open -batch 5 urls.txt
```

## Supported Platforms

* Linux
* FreeBSD
* macOS
* Windows

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
