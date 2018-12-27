package genomes

import (
	"github.com/MaxHalford/eaopt"
	"log"
	"math/rand"

	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"os"
)


var canonImg image.Image
var obama image.Image
func init()  {
	image.RegisterFormat("jpeg", "jpeg", jpeg.Decode, jpeg.DecodeConfig)
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)
	file, err := os.Open("assets/left.png")
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

type ImageRandomGenome struct {
	Img *image.RGBA
}

// tri color
func (genome ImageRandomGenome) Evaluate() (float64, error) {
	sum := 0.0
	for i := 0; i < genome.Img.Bounds().Dx(); i++ {
		for j := 0; j < genome.Img.Bounds().Dy(); j++ {
			r, _, _, _ := genome.Img.At(i,j).RGBA()
			mult := i * j
			if mult % 3 == 0 {
				sum += float64(r)
			} else if mult % 3 == 1 {
				sum += float64(r)
			} else {
				sum += float64(r)
			}
		}
	}
	return -sum, nil
}


// Mutate X% of genome
func (genome ImageRandomGenome) Mutate(rng *rand.Rand) {
	img := genome.Img
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

// No crossover
func (genome ImageRandomGenome) Crossover(Y eaopt.Genome, rng *rand.Rand) {
}

// Clone a Vector to produce a new one that points to a different slice.
func (genome ImageRandomGenome) Clone() eaopt.Genome {
	newImg := image.NewRGBA(image.Rect(0, 0, genome.Img.Bounds().Dx(), genome.Img.Bounds().Dy()))
	copy(newImg.Pix, genome.Img.Pix)
	var newGenome = ImageRandomGenome{
		Img: newImg,
	}
	return newGenome
}

// starts with a random image
func RandomImage(rng *rand.Rand) eaopt.Genome {
	img := image.NewRGBA(image.Rect(0, 0, canonImg.Bounds().Dx(), canonImg.Bounds().Dy()))
	dx := img.Bounds().Dx()
	dy := img.Bounds().Dy()

	for i := 0; i < dx; i++ {
		for j := 0; j < dy; j++ {
			img.Set(i,j, color.RGBA{
				R: uint8(rng.Intn(255)) ,
				G: uint8(rng.Intn(255)),
				B: uint8(rng.Intn(255)),
				A: 255,
			})
		}
	}
	ih := ImageRandomGenome{
		Img: img,
	}
	return ih
}




func (genome ImageRandomGenome) SaveImage(path, label string) {
	// add generation label
	clone := genome.Clone().(ImageRandomGenome)
	addLabel(clone.Img, 10, clone.Img.Bounds().Dy()-10, label)

	// save image
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	err = jpeg.Encode(f, clone.Img, nil)
	if err != nil {
		log.Print(err)
	}
}