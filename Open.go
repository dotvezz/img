package img

import (
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

// Open opens a file at the provided path and returns an image.Image if it was able to open and decode
func Open(path string) (i image.Image, err error) {
	f, err := os.Open(path)

	i, _, err = image.Decode(f, )

	if err != nil {
		return
	}

	err = f.Close()
	return
}
