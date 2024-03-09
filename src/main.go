package main

import "github.com/golesson/go-chain-pattern/src/chain"

func main() {

	//chain contract
	contract := &chain.IContract{}

	//STEPS
	first := &chain.FirstStep{}
	second := &chain.SecondStep{}
	third := &chain.ThirdStep{}
	fourth := &chain.FourthStep{}
	fifth := &chain.FifthStep{}

	//BUILDER
	builder := &chain.StepBuilder{}

	builder.SetFirstStep(first)
	builder.WithStep(second)
	builder.WithStep(third)
	builder.WithStep(fourth)
	builder.WithStep(fifth)

	begin := builder.Build()

	begin.Exec(contract)
}
