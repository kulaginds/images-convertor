package main

import (
	"fmt"
	"image/jpeg"
	"image/png"
	"os"

	"github.com/pkg/errors"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("jpeg2bmp input.jpeg output.bmp")
		return
	}

	inFn := os.Args[1]
	outFn := inFn + ".bmp"

	if len(os.Args) >= 3 {
		outFn = os.Args[2]
	}

	inF, err := os.Open(inFn)
	if err != nil {
		fmt.Println(errors.Wrapf(err, "cannot open input file: %s", inFn))
		os.Exit(1)
	}

	img, err := png.Decode(inF)
	if err != nil {
		fmt.Println(errors.Wrapf(err, "cannot decode input file: %s", inFn))
		os.Exit(1)
	}

	outF, err := os.Create(outFn)
	if err != nil {
		fmt.Println(errors.Wrapf(err, "cannot create output file: %s", outFn))
		os.Exit(1)
	}

	if err = jpeg.Encode(outF, img, nil); err != nil {
		fmt.Println(errors.Wrap(err, "cannot encode bmp"))
		os.Exit(1)
	}

	fmt.Println("done")
}
