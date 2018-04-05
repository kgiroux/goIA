package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/kgiroux/goia/config"
	"github.com/kgiroux/goia/hello"
	"github.com/kgiroux/goia/neurone"
)

func main() {
	fmt.Println(hello.BuildHello())
	configuration := config.GetConfig()
	var lenghtEntry = configuration.GetEntriesNumber()
	var entries = make([]float64, lenghtEntry)
	var dataSetEntries = make([]float64, lenghtEntry)
	var weight = make([]float64, lenghtEntry)
	neuroneTest := neurone.GetNeurone()
	rand.Seed(time.Now().UTC().UnixNano())

	// Entries Initialisation
	for i := 0; i < configuration.GetEntriesNumber()-1; i++ {
		entries[i] = rand.Float64()
	}

	for i := 0; i < configuration.GetEntriesNumber()-1; i++ {
		weight[i] = rand.Float64()
	}
	neuroneTest.SetWeight(weight)

	// Definition of the bias
	neuroneTest.SetBias(rand.Float64())
	neuroneTest = neurone.Preactivation(neuroneTest, entries)
	neuroneTest = neurone.ActivationSignoide(neuroneTest)

	fmt.Println("Initialisation weight : ")
	for i := 0; i < configuration.GetEntriesNumber(); i++ {
		fmt.Printf("%f\n", neuroneTest.GetWeight()[i])
	}
	fmt.Println("Fin Initialisation weight : ")
	fmt.Println("Bias : ")
	fmt.Printf("%f\n", neuroneTest.GetBias())

	dataSets := neurone.GetDataSetArray(configuration.GetDataSetNumber())
	fmt.Printf("Length of dataset %d \n", len(dataSets))
	for i := 0; i < (configuration.GetDataSetNumber() / 2); i++ {
		for j := 0; j < configuration.GetEntriesNumber()-1; j++ {
			dataSetEntries[j] = rand.Float64()
			fmt.Printf("%f\n", dataSetEntries[j])
		}
		dataSets[i].SetEntries(dataSetEntries)
		dataSets[i].SetTarget(1)
	}
	fmt.Println("Initialisation entries : ")
	for i := (configuration.GetDataSetNumber() / 2); i < configuration.GetDataSetNumber(); i++ {
		for j := 0; j < configuration.GetEntriesNumber()-1; j++ {
			dataSetEntries[j] = -rand.Float64()
			fmt.Printf("%f\n", dataSetEntries[j])
		}
		dataSets[i].SetEntries(dataSetEntries)
		dataSets[i].SetTarget(0)
	}
	fmt.Println("Fin Initialisation entries")
	for i := 0; i < configuration.EpochNumber(); i++ {
		fmt.Println("Interation : ", i)
		for j := 0; j < len(dataSets); j++ {
			neuroneTest = neurone.Preactivation(neuroneTest, dataSets[j].GetEntries())
			neuroneTest = neurone.ActivationSignoide(neuroneTest)
			if configuration.Log() {
				fmt.Println("------------------------------------------")
				fmt.Printf("Y %f\n", neuroneTest.GetOutput())
				fmt.Printf("T %f\n", dataSets[j].GetTarget())
				for k := 0; k < configuration.GetEntriesNumber(); k++ {
					fmt.Printf("X%d : %f\n", k, dataSets[j].GetEntries()[k])
				}
				fmt.Println("------------------------------------------")
			}
			neuroneTest = neurone.DescentOfGradient(neuroneTest, dataSets[j].GetTarget(), dataSets[j].GetEntries())
		}
	}

	var entriesVerification []float64

	entriesVerification = make([]float64, 2)

	entriesVerification[0] = 15
	entriesVerification[1] = 15

	neuroneTest = neurone.Preactivation(neuroneTest, entriesVerification)
	neuroneTest = neurone.ActivationSignoide(neuroneTest)
	fmt.Printf("Cas 15 Y %f\n", neuroneTest.GetOutput())

	entriesVerification[0] = -0.5
	entriesVerification[1] = -0.5

	neuroneTest = neurone.Preactivation(neuroneTest, entriesVerification)
	neuroneTest = neurone.ActivationSignoide(neuroneTest)
	fmt.Printf("Cas -15 Y %f\n", neuroneTest.GetOutput())
}
