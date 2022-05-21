package commands

import (
	"github.com/rldw/s3recovery/flags"
	"github.com/rldw/s3recovery/services"
)

func InfoCmd(cmdFlags *flags.InfoCmdFlags) error {
	cloudwatchUtils := new(services.CloudwatchUtils)
	_, err := cloudwatchUtils.GetBucketMetrics(cmdFlags.BucketName)

	if err != nil {
		return err
	}

	return nil
}
