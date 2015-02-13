package usecases

import (
	"github.com/acazau/cloud-manager/domain"
)

type Instance struct {
	Id   domain.UUID `json:"id"`
	Name string      `json:"name"`
}

type InstanceInteractor struct {
	InstanceRepository domain.InstanceRepository
}

func (interactor *InstanceInteractor) ListInstances() ([]*Instance, error) {
	instances, err := interactor.InstanceRepository.ListInstances()
	if err != nil {
		return nil, err
	}

	domainInstances := make([]*Instance, len(instances), len(instances))
	for i := range instances {
		domainInstances[i] = mapFromDomainInstance(instances[i])
	}
	return domainInstances, nil
}

func mapFromDomainInstance(domainInstance *domain.Instance) (instance *Instance) {
	instance.Id = domainInstance.Id
	instance.Name = domainInstance.Name
	return
}
