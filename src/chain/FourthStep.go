package chain

type FourthStep struct {
	next IStep
}

func (F *FourthStep) Exec(contract *IContract) {

	if contract.Operation == "SUM" {
		contract.Result = contract.FirstValue + contract.Secondvalue
	}

	if F.next != nil {
		F.next.Exec(contract)
	}
}

func (F *FourthStep) SetNext(next IStep) {
	F.next = next
}
