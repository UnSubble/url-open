package browser

type Browser struct {
	os BrowserOS
}

func (b *Browser) Open(url string) error {
	return b.os.Open(url)
}
