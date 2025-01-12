package awsclient

import (
	"context"

	"github.com/giantswarm/microerror"
	"github.com/giantswarm/operatorkit/v5/pkg/controller/context/reconciliationcanceledcontext"
)

func (r *Resource) EnsureCreated(ctx context.Context, obj interface{}) error {
	cr, err := r.toClusterFunc(ctx, obj)
	if IsNotFound(err) {
		r.logger.Debugf(ctx, "cluster cr not available yet")
		r.logger.Debugf(ctx, "canceling reconciliation")
		reconciliationcanceledcontext.SetCanceled(ctx)

		return nil
	} else if err != nil {
		return microerror.Mask(err)
	}

	err = r.addAWSClientsToContext(ctx, cr)
	if err != nil {
		return microerror.Mask(err)
	}

	return nil
}
