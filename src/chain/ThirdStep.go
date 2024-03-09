package chain

type ThirdStep struct {
	next IStep
}

func (T *ThirdStep) Exec(contract *IContract) {

	contract.Operation = "SUM"

	if T.next != nil {
		T.next.Exec(contract)
	}
}

func (T *ThirdStep) SetNext(next IStep) {
	T.next = next
}
