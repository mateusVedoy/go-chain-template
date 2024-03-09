package chain

type StepBuilder struct {
	firstStep IStep
	nextStep  []IStep
}

func (S *StepBuilder) SetFirstStep(step IStep) {
	S.firstStep = step
}

func (S *StepBuilder) WithStep(step IStep) {
	S.nextStep = append(S.nextStep, step)
}

func (S *StepBuilder) setChain() {
	for x := 0; x < len(S.nextStep)-1; x++ {
		S.nextStep[x].SetNext(S.nextStep[x+1])
	}
	S.firstStep.SetNext(S.nextStep[0])
}

func (S *StepBuilder) Build() IStep {
	S.setChain()
	return S.firstStep
}
