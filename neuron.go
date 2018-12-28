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

type Point struct {
	X, Y float64
}

type colouredPoint struct {
	Point
	colour float64
}

func main() {
	p := Perceptron{2, []float64{0., 0.0}, 0, activation, 0.01}

	// Generate the dataset
	var points []colouredPoint
	n := 100
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			var colour float64
			if i < j {
				colour = 1
				points = append(points, colouredPoint{Point{float64(i), float64(j)}, colour})
			} else if i > j {
				colour = 0
				points = append(points, colouredPoint{Point{float64(i), float64(j)}, colour})
			}
		}
	}
	// Fisher-Yates shuffle
	for i := range points {
		j := rand.Intn(i + 1)
		points[i], points[j] = points[j], points[i]
	}

	// Partition the dataset in training and testing set
	n_points := len(points)
	partition := (n_points * 8) / 10
	training := points[:partition]
	testing := points[partition:]

	// Train the classifier
	for _, point := range training {
		inputVector := []float64{point.X, point.Y}
		target := point.colour
		(&p).updateWeights(inputVector, target)
	}

	// Test the classifier
	testCard := float64(len(testing))
	wrong := 0.0
	for _, point := range testing {
		inputVector := []float64{point.X, point.Y}
		prediction, _ := p.predict(inputVector)
		if point.colour != prediction {
			wrong += 1
		}
	}
	fmt.Println(1 - wrong/testCard)
}
