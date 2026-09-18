//go:build windows

package browser

import "os/exec"

type WindowsOS struct{}

func (WindowsOS) Open(url string) error {
	return exec.Command(
		"rundll32",
		"url.dll,FileProtocolHandler",
		url,
	).Start()
}

func New() *Browser {
	return &Browser{
		os: WindowsOS{},
	}
}
