package tccpn

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/giantswarm/microerror"
	"github.com/giantswarm/operatorkit/v5/pkg/controller/context/finalizerskeptcontext"

	"github.com/giantswarm/aws-operator/service/controller/controllercontext"
	"github.com/giantswarm/aws-operator/service/controller/key"
)

func (r *Resource) EnsureDeleted(ctx context.Context, obj interface{}) error {
	cr, err := key.ToControlPlane(obj)
	if err != nil {
		return microerror.Mask(err)
	}
	cc, err := controllercontext.FromContext(ctx)
	if err != nil {
		return microerror.Mask(err)
	}

	{
		r.logger.Debugf(ctx, "disabling the termination protection of the tenant cluster's control plane nodes cloud formation stack")

		i := &cloudformation.UpdateTerminationProtectionInput{
			EnableTerminationProtection: aws.Bool(false),
			StackName:                   aws.String(key.StackNameTCCPN(&cr)),
		}

		_, err = cc.Client.TenantCluster.AWS.CloudFormation.UpdateTerminationProtection(i)
		if IsDeleteInProgress(err) {
			r.logger.Debugf(ctx, "the tenant cluster's control plane nodes cloud formation stack is being deleted")
			r.event.Emit(ctx, &cr, "CFDelete", fmt.Sprintf("the tenant cluster's control plane nodes cloud formation stack has stack status %#q", cloudformation.StackStatusDeleteInProgress))

			r.logger.Debugf(ctx, "keeping finalizers")
			finalizerskeptcontext.SetKept(ctx)

			r.logger.Debugf(ctx, "canceling resource")

			return nil

		} else if IsDeleteFailed(err) {
			r.logger.Debugf(ctx, "the tenant cluster's control plane nodes cloud formation stack failed to delete")
			r.event.Emit(ctx, &cr, "CFDeleteFailed", fmt.Sprintf("the tenant cluster's control plane nodes cloud formation stack has stack status %#q", cloudformation.StackStatusDeleteFailed))

			r.logger.Debugf(ctx, "keeping finalizers")
			finalizerskeptcontext.SetKept(ctx)

			r.logger.Debugf(ctx, "canceling resource")

			return nil

		} else if IsNotExists(err) {
			r.logger.Debugf(ctx, "the tenant cluster's control plane nodes cloud formation stack does not exist")
			r.event.Emit(ctx, &cr, "CFDeleted", fmt.Sprintf("the tenant cluster's control plane nodes cloud formation stack has stack status %#q", cloudformation.StackStatusDeleteComplete))
			r.logger.Debugf(ctx, "canceling resource")

			return nil

		} else if err != nil {
			return microerror.Mask(err)
		}

		r.logger.Debugf(ctx, "disabled the termination protection of the tenant cluster's control plane nodes cloud formation stack")
	}

	{
		r.logger.Debugf(ctx, "requesting the deletion of the tenant cluster's control plane nodes cloud formation stack")

		i := &cloudformation.DeleteStackInput{
			StackName: aws.String(key.StackNameTCCPN(&cr)),
		}

		_, err = cc.Client.TenantCluster.AWS.CloudFormation.DeleteStack(i)
		if IsUpdateInProgress(err) {
			r.logger.Debugf(ctx, "the tenant cluster's control plane nodes cloud formation stack is being updated")

			r.logger.Debugf(ctx, "keeping finalizers")
			finalizerskeptcontext.SetKept(ctx)

			r.logger.Debugf(ctx, "canceling resource")

			return nil

		} else if err != nil {
			return microerror.Mask(err)
		}

		r.logger.Debugf(ctx, "requested the deletion of the tenant cluster's control plane nodes cloud formation stack")
		r.event.Emit(ctx, &cr, "CFDeleteRequested", "requested the deletion of the tenant cluster's control plane nodes cloud formation stack")

		r.logger.Debugf(ctx, "keeping finalizers")
		finalizerskeptcontext.SetKept(ctx)
	}

	return nil
}
