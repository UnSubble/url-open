//go:build freebsd

package browser

import "os/exec"

type FreeBSDOS struct{}

func (FreeBSDOS) Open(url string) error {
	return exec.Command("xdg-open", url).Start()
}

func New() *Browser {
	return &Browser{
		os: FreeBSDOS{},
	}
}
