package cloudtrailiface

import (
	"github.com/datacratic/aws-sdk-go/service/cloudtrail"
)

type CloudTrailAPI interface {
	CreateTrail(*cloudtrail.CreateTrailInput) (*cloudtrail.CreateTrailOutput, error)

	DeleteTrail(*cloudtrail.DeleteTrailInput) (*cloudtrail.DeleteTrailOutput, error)

	DescribeTrails(*cloudtrail.DescribeTrailsInput) (*cloudtrail.DescribeTrailsOutput, error)

	GetTrailStatus(*cloudtrail.GetTrailStatusInput) (*cloudtrail.GetTrailStatusOutput, error)

	LookupEvents(*cloudtrail.LookupEventsInput) (*cloudtrail.LookupEventsOutput, error)

	StartLogging(*cloudtrail.StartLoggingInput) (*cloudtrail.StartLoggingOutput, error)

	StopLogging(*cloudtrail.StopLoggingInput) (*cloudtrail.StopLoggingOutput, error)

	UpdateTrail(*cloudtrail.UpdateTrailInput) (*cloudtrail.UpdateTrailOutput, error)
}