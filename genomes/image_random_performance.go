package genomes

import (
"github.com/MaxHalford/eaopt"
"golang.org/x/image/font"
"golang.org/x/image/font/basicfont"
"golang.org/x/image/math/fixed"
"log"
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

type ImageRandomGenomePerformance []uint8

// tri color
func (genome ImageRandomGenomePerformance) Evaluate() (float64, error) {
	var sum float64 = 0
	for i := 0; i < len(genome); i++ {
		if i % 4 == 0 {
			sum += float64(genome[i])
		} else if i % 4 == 3 {
			sum += float64(genome[i])
		} else {
			sum -= float64(genome[i])
		}
	}
	return -sum, nil
}


// Mutate X% of genome
func (genome ImageRandomGenomePerformance) Mutate(rng *rand.Rand) {
	MutUint8(genome, 0.01, rng)
	for i := 0; i < len(genome); i++ {
		if i % 4 == 0 {

		} else if i % 4 == 3 {
			genome[i] = 255
		} else {
			//genome[i] = 0
		}
	}
}

// No crossover
func (genome ImageRandomGenomePerformance) Crossover(Y eaopt.Genome, rng *rand.Rand) {
	CrossUniformUint8(genome, Y.(ImageRandomGenomePerformance), rng)
}

// Clone a Vector to produce a new one that points to a different slice.
func (genome ImageRandomGenomePerformance) Clone() eaopt.Genome {
	newGenome := make(ImageRandomGenomePerformance, len(genome))
	copy(newGenome, genome)
	return newGenome
}

// starts with a random image
func RandomImagePerformance(rng *rand.Rand) eaopt.Genome {
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

	gp := make(ImageRandomGenomePerformance, len(img.Pix))
	copy(gp, img.Pix)
	return gp
}


func addLabel(img *image.RGBA, x, y int, label string) {
	black := color.RGBA{0, 0, 0, 255}
	white := color.RGBA{255, 255, 255, 255}
	point := fixed.Point26_6{fixed.Int26_6(x * 64), fixed.Int26_6(y * 64)}
	point2 := fixed.Point26_6{fixed.Int26_6((x + 1)* 64), fixed.Int26_6((y +1)* 64)}

	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(black),
		Face: basicfont.Face7x13,
		Dot:  point,
	}
	d2 := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(white),
		Face: basicfont.Face7x13,
		Dot:  point2,
	}
	d.DrawString(label)
	d2.DrawString(label)

}

func (genome ImageRandomGenomePerformance) SaveImage(path, label string) {
	newImg := image.NewRGBA(image.Rect(0, 0, canonImg.Bounds().Dx(), canonImg.Bounds().Dy()))
	copy(newImg.Pix, genome)

	addLabel(newImg, 10, newImg.Bounds().Dy()-10, label)

	// save image
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	err = jpeg.Encode(f, newImg, nil)
	if err != nil {
		log.Print(err)
	}
}