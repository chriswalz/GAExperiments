package genomes

import "math/rand"

// MutNormalUint8 modifies a uint8 gene if a coin toss is under a defined
// mutation rate. The new gene value is a random value sampled from a normal
// distribution centered on the gene's current value and with a standard
// deviation proportional to the current value. It does so for each gene.
func MutNormalUint8(genome []uint8, rate float64, rng *rand.Rand) {
	for i := range genome {
		// Flip a coin and decide to mutate or not
		if rng.Float64() < rate {
			genome[i] += uint8(rng.NormFloat64() * float64(genome[i]))
		}
	}
}

func MutUint8(genome []uint8, rate float64, rng *rand.Rand) {
	for i := range genome {
		// Flip a coin and decide to mutate or not
		if rng.Float64() < rate {
			genome[i] = uint8(rng.Intn(255))
		}
	}
}