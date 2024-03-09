package chain

type IBuilder interface {
	SetFirstStep(step IStep)
	WithStep(IStep)
	Build() IStep
}
