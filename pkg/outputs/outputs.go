package outputs

import (
	"github.com/plantoncloud/planton-cloud-apis/zzgo/cloud/planton/apis/code2cloud/v1/gcp/gcpcloudcdn"
	"github.com/pulumi/pulumi/sdk/v3/go/auto"
)

func PulumiOutputsToStackOutputsConverter(stackOutput auto.OutputMap,
	input *gcpcloudcdn.GcpCloudCdnStackInput) *gcpcloudcdn.GcpCloudCdnStackOutputs {
	return &gcpcloudcdn.GcpCloudCdnStackOutputs{}
}
