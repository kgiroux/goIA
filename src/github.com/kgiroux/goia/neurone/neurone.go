package neurone

import (
	"fmt"
	"math"

	"github.com/kgiroux/goia/config"
)

// Neurone that can resolve equation
type Neurone struct {
	weight              []float64
	bias                float64
	output              float64
	preactivationResult float64
	gradientBias        float64
	gradientWeight      []float64
	configuration       config.Config
}

// GetNeurone : Constructor a Neurone Object
func GetNeurone() Neurone {
	return Neurone{configuration: config.GetConfig()}
}

// GetWeight return the arrays
func (m *Neurone) GetWeight() []float64 {
	return m.weight
}

// SetWeight Define the value of the weight
func (m *Neurone) SetWeight(value []float64) {
	m.weight = value
}

// SetBias : defined the bias of the neurone
func (m *Neurone) SetBias(value float64) {
	m.bias = value
}

// GetBias return the bias of the neurone
func (m *Neurone) GetBias() float64 {
	return m.bias
}

// GetOutput return the bias of the neurone
func (m *Neurone) GetOutput() float64 {
	return m.output
}

// Preactivation for the compute of the PreactivationResult
func Preactivation(neurone Neurone, entries []float64) Neurone {
	var result float64
	for i := 0; i < neurone.configuration.GetEntriesNumber()-1; i++ {
		result += neurone.weight[i] * entries[i]
	}
	result += neurone.bias
	fmt.Printf("Value of the preactivation phase : %f\n", result)

	neurone.preactivationResult = result
	return neurone
}

// ActivationSignoide : Compute the result of the neurone
func ActivationSignoide(neurone Neurone) Neurone {
	var result float64

	fmt.Printf("Value of the signoide :   %f\n", (1 + math.Exp(-neurone.preactivationResult)))

	result = 1 / (1 + math.Exp(-neurone.preactivationResult))

	neurone.output = result
	return neurone
}

// DescentOfGradient : reverse the compute for the learning phase
func DescentOfGradient(neurone Neurone, target float64, entries []float64) Neurone {
	neurone.gradientWeight = make([]float64, neurone.configuration.GetEntriesNumber())
	neurone.gradientBias = (neurone.output - target) * (math.Exp(-neurone.preactivationResult) / math.Pow(2, (1+math.Exp(-neurone.preactivationResult))))
	fmt.Printf(" Gradient Bias : %f\n", neurone.gradientBias)
	for i := 0; i < (neurone.configuration.GetEntriesNumber() - 1); i++ {
		neurone.gradientWeight[i] = (neurone.output - target) * (math.Exp(-neurone.preactivationResult) / math.Pow(2, (1+math.Exp(-neurone.preactivationResult)))) * entries[i]
		fmt.Printf(" Neurone Gradien Weight %d : %f\n", i, neurone.gradientWeight[i])
	}

	//neurone.gradientWeight[0] = (neurone.output - target) * (math.Exp(-neurone.preactivationResult) / math.Pow(2, (1+math.Exp(-neurone.preactivationResult)))) * entries[0]
	//neurone.gradientWeight[1] = (neurone.output - target) * (math.Exp(-neurone.preactivationResult) / math.Pow(2, (1+math.Exp(-neurone.preactivationResult)))) * entries[1]

	for i := 0; i < neurone.configuration.GetEntriesNumber()-1; i++ {
		neurone.weight[i] = neurone.weight[i] - neurone.configuration.LearningRate()*neurone.gradientWeight[i]
		fmt.Printf("weight %d : %f\n", i, neurone.weight[i])
	}
	neurone.bias = neurone.bias - neurone.configuration.LearningRate()*neurone.gradientBias
	fmt.Printf("%f\n", neurone.bias)
	return neurone

}
