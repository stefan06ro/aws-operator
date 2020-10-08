package cpvpc

import (
	"context"
	"sync"

	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"

	"github.com/giantswarm/aws-operator/service/controller/controllercontext"
)

const (
	Name = "cpvpc"
)

type Config struct {
	Logger micrologger.Logger

	InstallationName string
}

type Resource struct {
	logger micrologger.Logger

	cachedCidr string
	cachedID   string
	mutex      sync.Mutex

	installationName string
}

func New(config Config) (*Resource, error) {
	if config.Logger == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}

	if config.InstallationName == "" {
		return nil, microerror.Maskf(invalidConfigError, "%T.InstallationName must not be empty", config)
	}

	r := &Resource{
		logger: config.Logger,

		cachedCidr: "",
		cachedID:   "",
		mutex:      sync.Mutex{},

		installationName: config.InstallationName,
	}

	return r, nil
}

func (r *Resource) Name() string {
	return Name
}

func (r *Resource) addVPCInfoToContext(ctx context.Context) error {
	cc, err := controllercontext.FromContext(ctx)
	if err != nil {
		return microerror.Mask(err)
	}

	cc.Status.ControlPlane.VPC.CIDR = "0.0.0.0/0"

	return nil
}
