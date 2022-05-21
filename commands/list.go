package commands

import (
	"github.com/rldw/s3recovery/flags"
	"github.com/rldw/s3recovery/services"
	"github.com/sirupsen/logrus"
)

func ListCmd(cmdFlags *flags.ListCmdFlags) error {
	s3Utils := new(services.S3Utils)
	buckets, err := s3Utils.ListBuckets()
	if err != nil {
		return err
	}

	for _, bucket := range buckets {
		logrus.Infof("%v", *bucket.Name)
	}

	return nil
}
