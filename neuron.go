package main


type Perceptron struct {
  size int
  input_vector []float32{}
  weights_vector []float32{}
  activation func(float32) float32
}
func (p Perceptron) updateWeights() {
  new_weights = []float32
}

func (p Perceptron) activation() {

}
func vectorDifference(u  []float32{}, v []float32{}) ([]float32, error){
  if len(u) != len(v) {
    return nil, fmt.Errorf("The two vectors aren't the same size")
  }
  difference := []int
  for i, val_i := range u {
    difference = append(difference, val_i*v[i])
  }
  return difference
}
