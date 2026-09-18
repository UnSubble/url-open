package browser

type BrowserOS interface {
	Open(url string) error
}
