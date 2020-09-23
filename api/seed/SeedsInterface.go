package seed

type Publisher interface {
	items()
	RunSeeds() bool
}
