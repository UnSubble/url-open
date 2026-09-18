//go:build darwin

package browser

import "os/exec"

type DarwinOS struct{}

func (DarwinOS) Open(url string) error {
	return exec.Command("open", url).Start()
}

func New() *Browser {
	return &Browser{
		os: DarwinOS{},
	}
}
