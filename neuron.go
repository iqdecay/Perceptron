package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type Perceptron struct {
	// A Perceptron is determined by its size, weights and activation function
	size         int
	weights      []float64
	bias         float64
	activation   func(float64) float64
	learningRate float64
}

type dataPoint struct {
	// A data point is anything that can be used for binary classification
	dimension int
	data      []float64 // data should be of length dimension
}

type labeledDataPoint struct {
	dataPoint
	label float64
}

type DataSet struct {
	training []labeledDataPoint
	testing  []labeledDataPoint
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
	fmt.Println("Input : ", input)
	for i, w_i := range p.weights {
		sum += w_i * input[i]
	}
	output := p.activation(sum)
	fmt.Println("Prediction :", output)
	return output, nil
}

func (p *Perceptron) train(data []labeledDataPoint) {
	// Train the perceptron with labeled data
	for _, point := range data {
		inputVector := point.data
		target := point.label
		p.updateWeights(inputVector, target)
	}
}

func activation(f float64) float64 {
	// Implements the Rectified linear unit activation function
	if f > 0 {
		return 1
	} else {
		return 0
	}
}

func generateDataPoint() dataPoint {
	a, b, c := float64(rand.Intn(200))-100, float64(rand.Intn(200))-100, float64(rand.Intn(200))-100
	return dataPoint{dimension: 3, data: []float64{a, b, c}}
}

func (p dataPoint) computeLabel() float64 {
	x, y := p.data[0], p.data[1]
	// A point (x, y, z) is on the right side of the plane ((0,0,1),(1,1,0)) iff
	// -x+1 > y
	if -x + 1 > y {
		return 1
	} else {
		return 0
	}
}

func generateLabeledDataPoint() labeledDataPoint {
	var point dataPoint
	var labeledPoint labeledDataPoint
	point = generateDataPoint()
	label := point.computeLabel()
	labeledPoint.dataPoint = point
	labeledPoint.label = label
	return labeledPoint

}

func generateDataset(size, splitRate int) (d DataSet) {
	// splitRate is the proportion of the dataset that is used for testing
	seed := time.Now().UnixNano()
	rand.Seed(seed)
	fsize := float64(size)
	fsplitRate := float64(splitRate)

	trainingSize := int(math.RoundToEven(fsize * (1 - fsplitRate/100)))
	fmt.Println("trainingSize : ", trainingSize)
	trainingSet := make([]labeledDataPoint, trainingSize)

	testingSize := size - trainingSize
	fmt.Println("testingSize  : ", testingSize)
	testingSet := make([]labeledDataPoint, testingSize)

	// Generate trainingSet
	for i := 0; i < trainingSize; i++ {
		// Generate a labeled DataPoint and add it to the dataset
		trainingSet[i] = generateLabeledDataPoint()
	}

	// Generate testingSet
	for i := 0; i < testingSize; i++ {
		// Generate a labeled DataPoint and add it to the datase
		o := generateLabeledDataPoint()
		testingSet[i] = o

	}

	return DataSet{training: trainingSet, testing: testingSet}

}

func main() {
	seed := time.Now().UnixNano()
	rand.Seed(seed)
	a, b := rand.Float64(), rand.Float64()
	p := Perceptron{2, []float64{a, b}, 0, activation, 0.01}
	dataset := generateDataset(10, 20)
	//fmt.Println("Dataset :", dataset)
	// Train the perceptron
	p.train(dataset.training)

	// Test the classifier
	testCard := float64(len(dataset.testing))
	wrong := 0.0
	for _, point := range dataset.testing {
		fmt.Println(point)

		inputVector := point.data
		prediction, _ := p.predict(inputVector)
		if point.label != prediction {
			wrong += 1
		}
	}
	fmt.Println(1 - wrong/testCard)
}
