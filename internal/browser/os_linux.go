//go:build linux

package browser

import "os/exec"

type LinuxOS struct{}

func (LinuxOS) Open(url string) error {
	return exec.Command("xdg-open", url).Start()
}

func New() *Browser {
	return &Browser{
		os: LinuxOS{},
	}
}
