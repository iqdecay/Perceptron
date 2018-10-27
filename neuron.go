package main


type Perceptron struct {
  // A Perceptron is determined by its size, weights and activation function
  size int
  weights []float32{}
  activation func(float32) float32
  learningRate  float32
}
func (p Perceptron) updateWeights(input []float32{}, target float32) (error){
  // Update the weigth according to input, learningRate and activation
  newWeights := []float32
  if p.size != len(input) {
    return fmt.Errorf("The input hasn't the right dimension")
  }
  sum := 0
  for i, w_i := range(p.weights) {
    sum += w_i * input[i]
  }
  output := p.activation(sum)
  for i, w_i := range(p.weights) {
    newWeight := w_i + p.learning_rate * (target - output) * input[i]
    newWeights = append(newWeights, newWeight)
  }
  p.weights = newWeights
  fmt.Println("Weights updated !")
}

func activation(f float32) float32 {
  // Implements the Rectified linear unit activation function
  if f > 0 {
    return f
  } else {
    return 0
  }
}
