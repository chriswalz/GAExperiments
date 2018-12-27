package genomes

/*
import (
	"image"
	"image/jpeg"
	"image/png"
)

import (
	"github.com/MaxHalford/eaopt"
	"math/rand"

	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"os"
)



func init()  {
	image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
	file, err := os.Open("assets/harambe.jpeg")
	if err != nil {
		panic(err)
	}
	canonImg, _, err = image.Decode(file)
	if err != nil {
		panic(err)
	}
	file, err = os.Open("assets/thedonald.jpeg")
	if err != nil {
		panic(err)
	}
	obama, _, err = image.Decode(file)
	if err != nil {
		panic(err)
	}

}

// A Vector contains float64s.
type ImageStartGenome struct {
	Img *image.RGBA
}

func (ih ImageStartGenome) Evaluate() (float64, error) {
	sum := 0.0
	for i := 0; i < ih.Img.Bounds().Dx(); i++ {
		for j := 0; j < ih.Img.Bounds().Dy(); j++ {
			r, g, b, _ := ih.Img.At(i,j).RGBA()
			mult := i * j
			if mult % 3 == 0 {
				sum += float64(r) - float64(g) - float64(b)
			} else if mult % 3 == 1 {
				sum += -float64(r) + float64(g) - float64(b)
			} else {
				sum += -float64(r) - float64(g) + float64(b)
			}
		}
	}
	return -sum, nil
}


func (ih ImageStartGenome) Mutate(rng *rand.Rand) {
	img := ih.Img
	rate := 0.1
	dx := img.Bounds().Dx()
	dy := img.Bounds().Dy()
	count := int(rate * float64(dx) * float64(dy)) + 1
	for c := 0; c < count; c++ {
		i := rng.Intn(dx)
		j := rng.Intn(dy)
		mut := color.RGBA{
			R: uint8(rng.Intn(255)) ,
			G: uint8(rng.Intn(255)),
			B: uint8(rng.Intn(255)),
			A: 255,
		}
		img.Set(i, j, mut)
	}
}

// Crossover a Genome with another Genome .
func (ih ImageStartGenome) Crossover(Y eaopt.Genome, rng *rand.Rand) {
}
//const size = 400
// Clone a Vector to produce a new one that points to a different slice.
func (ih ImageStartGenome) Clone() eaopt.Genome {
	img := image.NewRGBA(image.Rect(0, 0, canonImg.Bounds().Dx(), canonImg.Bounds().Dy()))
	copy(img.Pix, ih.Img.Pix)
	var imh2 = ImageStartGenome{
		Img: img,
	}
	return imh2
}

// starts with a specific image
func StartImage(rng *rand.Rand) eaopt.Genome {
	img := image.NewRGBA(image.Rect(0, 0, canonImg.Bounds().Dx(), canonImg.Bounds().Dy()))
	dx := img.Bounds().Dx()
	dy := img.Bounds().Dy()

	for i := 0; i < dx; i++ {
		for j := 0; j < dy; j++ {
			r, g, b, _ := obama.At(i,j).RGBA()
			img.Set(i,j, color.RGBA{
				R: uint8(r / 0x101),
				G: uint8(g / 0x101),
				B: uint8(b / 0x101),
				A: 255,
			})
		}
	}
	ih := ImageStartGenome{
		Img: img,
	}
	return ih
}
*/