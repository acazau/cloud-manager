package domain

type InstanceRepository interface {
	ListInstances() ([]*Instance, error)
	Start(instance *Instance) error
	Stop(id UUID) (*Instance, error)
}

type Instance struct {
	Id   UUID
	Name string
	test InstanceRepository
}
