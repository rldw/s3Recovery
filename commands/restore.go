package commands

import (
	"github.com/rldw/s3recovery/flags"
	"github.com/sirupsen/logrus"
)

func RestoreCmd(cmdFlags *flags.RestoreCmdFlags) error {
	var err error

	logrus.WithFields(logrus.Fields{
		"bucketName":   cmdFlags.BucketName,
		"restorePoint": cmdFlags.RestorePoint,
	}).Debug("inside RestoreCmd")

	return err
}
