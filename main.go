package main

import (
	"fmt"
	"io"
	"os"

	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/google/go-containerregistry/pkg/name"
	"github.com/google/go-containerregistry/pkg/v1/remote"
)

func main() {
	srcImage := os.Args[1]
	directory := os.Args[2]

	ref, _ := name.ParseReference(srcImage)
	img, _ := remote.Image(ref, remote.WithAuthFromKeychain(authn.DefaultKeychain))

	l, _ := img.Layers()

	i := 0
	for _, layer := range l {
		r, _ := layer.Uncompressed()
		b, _ := io.ReadAll(r)
		os.WriteFile(fmt.Sprintf(directory+"/layer%d.tar", i), b, 0777)
		i++

	}

}
