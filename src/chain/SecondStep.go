package chain

type SecondStep struct {
	next IStep
}

func (S *SecondStep) Exec(contract *IContract) {

	contract.Secondvalue = 2

	if S.next != nil {
		S.next.Exec(contract)
	}
}

func (S *SecondStep) SetNext(next IStep) {
	S.next = next
}
