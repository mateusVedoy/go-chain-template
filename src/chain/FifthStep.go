package chain

import "fmt"

type FifthStep struct {
	next IStep
}

func (F *FifthStep) Exec(contract *IContract) {

	fmt.Printf("The result of a %s operation is %d\n", contract.Operation, contract.Result)

	//HERE IS OPTIONAL 'CAUSE WE KNOW THAT THERE'S NO STEP FOWARD
	if F.next != nil {
		F.next.Exec(contract)
	}

	fmt.Println("The chain has ended into fifth step")
}

func (F *FifthStep) SetNext(next IStep) {
	F.next = next
}
