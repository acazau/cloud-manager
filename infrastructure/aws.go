package infrastructure

import (
	"github.com/acazau/cloud-manager/domain"
	api "github.com/acazau/cloud-manager/usecases/api/v0"
)

type AWSRepository struct {
	LogManager domain.LogManager
}

func (repo *AWSRepository) ListInstances() ([]api.Instance, error) {
	var instances []api.Instance

	err := repo.LogManager.Log(domain.Info, "Listed all instances from aws")
	return instances, err
}
