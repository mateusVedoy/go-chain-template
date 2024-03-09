package chain

type FirstStep struct {
	next IStep
}

func (F *FirstStep) Exec(contract *IContract) {

	contract.FirstValue = 1

	if F.next != nil {
		F.next.Exec(contract)
	}
}

func (F *FirstStep) SetNext(next IStep) {
	F.next = next
}
