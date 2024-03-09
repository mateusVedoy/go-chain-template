package chain

type IStep interface {
	Exec(contract *IContract)
	SetNext(next IStep)
}
