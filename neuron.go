package main

import (
	"fmt"
	"math/rand"
)

type Perceptron struct {
	// A Perceptron is determined by its size, weights and activation function
	size         int
	weights      []float64
	bias         float64
	activation   func(float64) float64
	learningRate float64
}

func (p *Perceptron) updateWeights(input []float64, target float64) error {
	// Update the weigth according to input, learningRate and activation
	var newWeights []float64
	output, err := p.predict(input)
	if err != nil {
		return err
	}
	for i, w_i := range p.weights {
		newWeight := w_i + p.learningRate*(target-output)*input[i]
		newWeights = append(newWeights, newWeight)
	}
	p.bias = p.bias + p.learningRate*(target-output)
	p.weights = newWeights
	return nil
}

func (p Perceptron) predict(input []float64) (float64, error) {
	if p.size != len(input) {
		return 0, fmt.Errorf("The input hasn't the right dimension")
	}
	sum := 0.0
	sum += p.bias
	for i, w_i := range p.weights {
		sum += w_i * input[i]
	}
	output := p.activation(sum)
	return output, nil
}

func activation(f float64) float64 {
	// Implements the Rectified linear unit activation function
	if f > 0 {
		return 1
	} else {
		return 0
	}
}

func generateRGB() (float64, float64, float64){
	r := float64(rand.Intn(256))
	g := float64(rand.Intn(256))
	b := float64(rand.Intn(256))
	if r == g || r == b || b == g {
		return generateRGB()
	}
	return r, g, b
}

type RGB struct {
	r, g, b float64
}

type Pixel struct {
	RGB
	colour int
}

func (p *Pixel) determineColour() {
	if p.r > p.b && p.r > p.g {
		p.colour = 0
	}
	if p.g > p.b && p.g > p.r {
		p.colour = 1
	}
	if p.b > p.r && p.b > p.g {
		p.colour = 2
	}
}


func main() {
	var pixels []Pixel
	n := 100
	// Generate the training set
	for i := 0; i< n; i++ {
		r, g, b := generateRGB()
		pixel := Pixel{RGB{r, g, b}, [3]int{0, 0, 0}}
		pixel.determineColour()
		pixels = append(pixels, pixel)
	}
	for i := range []int{0,1,2} {
		p := Perceptron{3, []float64{0.0, 0.0}, activation, 0.01}
		for _, pixel := range pixels {
			input_vector := []float64{pixel.r, pixel.g, pixel.b}
			switch
		}

	}

}









}
