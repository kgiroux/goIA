package main

import (
	"fmt"
	"math"
)

type Neurone struct {
	weight              [2]float64
	bias                float64
	output              float64
	preactivationResult float64
	gradientBias        float64
	gradientWeight      [2]float64
}

type DataSet struct {
	entries [2]float64
	target  float64
}

func main() {
	fmt.Println("Hello world!")
	fmt.Println("Programme 1 : Neurone")

	var entries [2]float64
	entries[0] = 2
	entries[1] = 2

	var neurone Neurone
	neurone.weight[0] = 0.1
	neurone.weight[1] = 0.2
	neurone.bias = -0.6
	neurone.preactivationResult = preactivation(neurone, entries)
	neurone.output = activationSignoide(neurone)
	//fmt.Printf("Output %f\n", neurone.output)

	fmt.Println("Initialisation coefficient : ")
	fmt.Printf("%f\n", neurone.weight[0])
	fmt.Printf("%f\n", neurone.weight[1])
	fmt.Printf("%f\n", neurone.bias)

	var dataSets [5]DataSet
	dataSets[0].entries[0] = 2
	dataSets[0].entries[1] = 2
	dataSets[0].target = 1
	dataSets[1].entries[0] = 1
	dataSets[1].entries[1] = 1
	dataSets[1].target = 1
	dataSets[2].entries[0] = -2
	dataSets[2].entries[1] = -2
	dataSets[2].target = 0
	dataSets[3].entries[0] = -1
	dataSets[3].entries[1] = -1
	dataSets[3].target = 0
	dataSets[4].entries[0] = -0.5
	dataSets[4].entries[1] = -0.5
	dataSets[4].target = 0
	for i := 0; i < 75000; i++ {
		//fmt.Println("Interation : ", i)
		for j := 0; j < len(dataSets); j++ {
			neurone.preactivationResult = preactivation(neurone, dataSets[j].entries)
			neurone.output = activationSignoide(neurone)
			fmt.Println("------------------------------------------")
			fmt.Printf("Y %f\n", neurone.output)
			fmt.Printf("T %f\n", dataSets[j].target)
			fmt.Printf("X1 :  %f\n", dataSets[j].entries[0])
			fmt.Printf("X2 :  %f\n", dataSets[j].entries[1])
			fmt.Println("------------------------------------------")
			neurone = descentOfGradient(neurone, dataSets[j].target, dataSets[j].entries)
		}
	}

	fmt.Println("Coefficient end: ")
	fmt.Printf("%f\n", neurone.weight[0])
	fmt.Printf("%f\n", neurone.weight[1])
	fmt.Printf("%f\n", neurone.bias)

	var entriesVerification [2]float64

	entriesVerification[0] = 15
	entriesVerification[1] = 15

	neurone.preactivationResult = preactivation(neurone, entriesVerification)
	neurone.output = activationSignoide(neurone)
	fmt.Printf("Cas 15 Y %f\n", neurone.output)

	entriesVerification[0] = -0.5
	entriesVerification[1] = -0.5

	neurone.preactivationResult = preactivation(neurone, entriesVerification)
	neurone.output = activationSignoide(neurone)
	fmt.Printf("Cas -15 Y %f\n", neurone.output)
}

func preactivation(neurone Neurone, entries [2]float64) float64 {

	var result float64

	for i := 0; i < len(neurone.weight); i++ {
		result += neurone.weight[i] * entries[i]
	}
	result += neurone.bias
	//fmt.Printf("Result preactiviation %f\n", result)
	return result
}

func activationSignoide(neurone Neurone) float64 {
	var result float64

	//fmt.Printf("Activation signoide %f\n", (1 + math.Exp(-neurone.preactivationResult)))

	result = 1 / (1 + math.Exp(-neurone.preactivationResult))

	return result
}

func descentOfGradient(neurone Neurone, target float64, entries [2]float64) Neurone {
	neurone.gradientBias = (neurone.output - target) * (math.Exp(-neurone.preactivationResult) / math.Pow(2, (1+math.Exp(-neurone.preactivationResult))))
	//fmt.Printf(" Gradient Bias : %f\n", neurone.gradientBias)
	//fmt.Printf(" Debut du calcul Bias : %f\n", (neurone.output - target))
	//fmt.Printf(" Deuxiement part of the calcul %f\n", math.Exp(-neurone.preactivationResult))
	//fmt.Printf(" Last part du calcul %f\n", math.Pow(2, (1+math.Exp(-neurone.preactivationResult))))

	neurone.gradientWeight[0] = (neurone.output - target) * (math.Exp(-neurone.preactivationResult) / math.Pow(2, (1+math.Exp(-neurone.preactivationResult)))) * entries[0]
	//fmt.Printf("%e\n", neurone.gradientWeight[0])

	neurone.gradientWeight[1] = (neurone.output - target) * (math.Exp(-neurone.preactivationResult) / math.Pow(2, (1+math.Exp(-neurone.preactivationResult)))) * entries[1]
	//fmt.Printf("%e\n", neurone.gradientWeight[1])

	neurone.weight[0] = neurone.weight[0] - 0.08*neurone.gradientWeight[0]
	neurone.weight[1] = neurone.weight[1] - 0.08*neurone.gradientWeight[1]
	neurone.bias = neurone.bias - 0.08*neurone.gradientBias

	//fmt.Println("Nouveau coefficient : ")
	//fmt.Printf("%f\n", neurone.weight[0])
	//fmt.Printf("%f\n", neurone.weight[1])
	//fmt.Printf("%f\n", neurone.bias)

	return neurone

}
