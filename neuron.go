package main

import (
	"fmt"
)

type Perceptron struct {
	// A Perceptron is determined by its size, weights and activation function
	size         int
	weights      []float64
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
	p.weights = newWeights
	return nil
}

func (p Perceptron) predict(input []float64) (float64, error) {
	if p.size != len(input) {
		return 0, fmt.Errorf("The input hasn't the right dimension")
	}
	sum := 0.0
	for i, w_i := range p.weights {
		sum += w_i * input[i]
	}
	output := p.activation(sum)
	return output, nil
}

func activation(f float64) float64 {
	// Implements the Rectified linear unit activation function
	if f > 0 {
		return f
	} else {
		return 0
	}
}

type Point struct {
	X float64
	Y float64
}

type colouredPoint struct {
	Point
	colour float64
}

func main() {
	var points []colouredPoint
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			var colour float64
			if i < 0 {
				colour = 1
				points = append(points, colouredPoint{Point{float64(i), float64(j)}, colour})
			} else if  i > 0 {
				colour = 2
				points = append(points, colouredPoint{Point{float64(i), float64(j)}, colour})
			}
		}
	}
	p := Perceptron{2,[]float64{1.5, 0.5}, activation, 0.01 }
	training := points[:80]
	testing := points[80:]
	for _, point := range training {
		input_vector := []float64{point.X, point.Y}
		target := point.colour
		(&p).updateWeights(input_vector, target)
	}
	fmt.Println(p.weights)

	for _, point := range testing {
		input_vector := []float64{point.X, point.Y}
		prediction, _ := p.predict(input_vector)
		fmt.Println(prediction, point.colour)
	}


}
