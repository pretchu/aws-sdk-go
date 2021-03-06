package cognitosynciface

import (
	"github.com/datacratic/aws-sdk-go/service/cognitosync"
)

type CognitoSyncAPI interface {
	BulkPublish(*cognitosync.BulkPublishInput) (*cognitosync.BulkPublishOutput, error)

	DeleteDataset(*cognitosync.DeleteDatasetInput) (*cognitosync.DeleteDatasetOutput, error)

	DescribeDataset(*cognitosync.DescribeDatasetInput) (*cognitosync.DescribeDatasetOutput, error)

	DescribeIdentityPoolUsage(*cognitosync.DescribeIdentityPoolUsageInput) (*cognitosync.DescribeIdentityPoolUsageOutput, error)

	DescribeIdentityUsage(*cognitosync.DescribeIdentityUsageInput) (*cognitosync.DescribeIdentityUsageOutput, error)

	GetBulkPublishDetails(*cognitosync.GetBulkPublishDetailsInput) (*cognitosync.GetBulkPublishDetailsOutput, error)

	GetIdentityPoolConfiguration(*cognitosync.GetIdentityPoolConfigurationInput) (*cognitosync.GetIdentityPoolConfigurationOutput, error)

	ListDatasets(*cognitosync.ListDatasetsInput) (*cognitosync.ListDatasetsOutput, error)

	ListIdentityPoolUsage(*cognitosync.ListIdentityPoolUsageInput) (*cognitosync.ListIdentityPoolUsageOutput, error)

	ListRecords(*cognitosync.ListRecordsInput) (*cognitosync.ListRecordsOutput, error)

	RegisterDevice(*cognitosync.RegisterDeviceInput) (*cognitosync.RegisterDeviceOutput, error)

	SetIdentityPoolConfiguration(*cognitosync.SetIdentityPoolConfigurationInput) (*cognitosync.SetIdentityPoolConfigurationOutput, error)

	SubscribeToDataset(*cognitosync.SubscribeToDatasetInput) (*cognitosync.SubscribeToDatasetOutput, error)

	UnsubscribeFromDataset(*cognitosync.UnsubscribeFromDatasetInput) (*cognitosync.UnsubscribeFromDatasetOutput, error)

	UpdateRecords(*cognitosync.UpdateRecordsInput) (*cognitosync.UpdateRecordsOutput, error)
}