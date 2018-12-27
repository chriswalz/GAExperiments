package genomes

import "math/rand"

// CrossUniformFloat64 crossover combines two individuals (the parents) into one
// (the offspring). Each parent's contribution to the Genome is determined by
// the value of a probability p. Each offspring receives a proportion of both of
// it's parents genomes. The new values are located in the hyper-rectangle
// defined between both parent's position in Cartesian space.
func CrossUniformFloat64(p1 []float64, p2 []float64, rng *rand.Rand) {
	for i := range p1 {
		var p = rng.Float64()
		p1[i] = p*p1[i] + (1-p)*p2[i]
		p2[i] = (1-p)*p1[i] + p*p2[i]
	}
}

func CrossUniformUint8(p1 []uint8, p2 []uint8, rng *rand.Rand) {
	for i := range p1 {
		var p = rng.Float64()
		p1[i] = uint8(p*float64(p1[i]) + (1-p)*float64(p2[i]))
		p2[i] = uint8((1-p)*float64(p1[i]) + p*float64(p2[i]))
	}
}