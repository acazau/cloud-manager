package infrastructure

import (
	"github.com/acazau/cloud-manager/domain"
	api "github.com/acazau/cloud-manager/usecases/api/v0"
)

type DigitalOceanRepository struct {
	Logger domain.Logger
}

func (repo *DigitalOceanRepository) ListInstances() ([]api.Instance, error) {
	var instances []api.Instance

	repo.Logger.Info("Listed all instances from digital ocean")
	return instances, nil
}
