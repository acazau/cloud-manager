package domain

type IInstance interface {
	Start(instance *Instance) error
	Stop(id UUID) (*Instance, error)
}

type Instance struct {
	Id       UUID
	Name     string
	Instance IInstance
}

type InstanceProvider struct {
	Id   UUID
	Name string
}
