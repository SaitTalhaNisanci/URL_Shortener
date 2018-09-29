package model

// URL is a model.
type URL struct {
	Short string
	Long  string
}

// NewURL returns a new url model object with the given short and long urls.
func NewURL(short string, long string) *URL {
	return &URL{Long: long, Short: short}
}
