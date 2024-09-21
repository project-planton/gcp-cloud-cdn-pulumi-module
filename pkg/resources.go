package pkg

import (
	"github.com/pkg/errors"
	"github.com/plantoncloud/planton/apis/zzgo/cloud/planton/apis/code2cloud/v1/gcp/gcpcloudcdn"
	"github.com/plantoncloud/pulumi-module-golang-commons/pkg/provider/gcp/pulumigoogleprovider"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ResourceStack struct {
	GcpLabels map[string]string
}

func Resources(ctx *pulumi.Context, stackInput *gcpcloudcdn.GcpCloudCdnStackInput) error {
	//create gcp provider using the credentials from the input
	_, err := pulumigoogleprovider.Get(ctx, stackInput.GcpCredential)
	if err != nil {
		return errors.Wrap(err, "failed to setup gcp provider")
	}
	return nil
}
