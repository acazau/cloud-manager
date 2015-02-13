package infrastructure

import (
	"github.com/acazau/cloud-manager/domain"
	api "github.com/acazau/cloud-manager/usecases/api/v0"
	"github.com/goamz/goamz/aws"
)

type AuthFunc func() (auth aws.Auth, err error)

type AWSRepository struct {
	Logger domain.Logger
}

func (repo *AWSRepository) ListInstances() ([]api.Instance, error) {
	var instances []api.Instance

	repo.Logger.Info("Listed all instances from aws")
	return instances, nil
}
