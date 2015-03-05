package usecases

import (
	"errors"
	"github.com/acazau/cloud-manager/domain"
)

type IInstanceProvider interface {
	ListInstances() ([]Instance, error)
}

type Instance struct {
	Id   domain.UUID `json:"id"`
	Name string      `json:"name"`
}

type InstanceProviderManager struct {
	InjectedInstanceProvider IInstanceProvider
}

func (instanceProviderManager *InstanceProviderManager) ListInstances() ([]Instance, error) {
	if instanceProviderManager.InjectedInstanceProvider == nil {
		return nil, errors.New("Instance Provider Repository cannot be null")
	}
	instances, err := instanceProviderManager.InjectedInstanceProvider.ListInstances()
	if err != nil {
		return nil, err
	}

	domainInstances := make([]Instance, len(instances))
	for i := range instances {
		domainInstances[i] = mapFromDomainInstance(instances[i])
	}
	return domainInstances, nil
}

func mapFromDomainInstance(domainInstance Instance) (instance Instance) {
	instance.Id = domainInstance.Id
	instance.Name = domainInstance.Name
	return
}
