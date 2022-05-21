package main

import (
	"github.com/alecthomas/kingpin"
	"github.com/rldw/s3recovery/commands"
	"github.com/rldw/s3recovery/flags"
	"github.com/sirupsen/logrus"
	"log"
	"os"
)

var (
	// The Version of the app
	Version = "1.0.0"
)

func main() {
	app := kingpin.New("s3restore", "A tool to restore S3 Bucket objects")
	app.Version(Version)

	verbose := app.Flag("verbose", "Enable verbose logging").Bool()

	// `list` command returns a list of all buckets
	cmdList := app.Command("list", "Returns a list of all S3 Buckets.")
	listFlags := new(flags.ListCmdFlags)

	// `info` command lists latest possible point to restore and general metrics about the bucket
	cmdInfo := app.Command("info", "Returns the earliest possible point to restore and general metrics about the bucket.")
	infoFlags := new(flags.InfoCmdFlags)
	cmdInfo.Flag("bucket-name", "Name of the S3 Bucket.").Required().StringVar(&infoFlags.BucketName)

	// `restore` command sets bucket versioning of all objects in a bucket
	cmdRestore := app.Command("restore", "Restores all objects in the bucket to a specified point in time.")
	restoreFlags := new(flags.RestoreCmdFlags)
	cmdRestore.Flag("bucket-name", "Name of the S3 Bucket.").Required().StringVar(&restoreFlags.BucketName)
	cmdRestore.Flag("point-in-time", "The point in time to restore to in Format (TODO)").Required().StringVar(&restoreFlags.RestorePoint)

	command := kingpin.MustParse(app.Parse(os.Args[1:]))

	errtpl := "%v\n"
	if *verbose {
		logrus.SetLevel(logrus.DebugLevel)
		errtpl = "%+v\n"
	}

	var err error
	switch command {
	case cmdList.FullCommand():
		err = commands.ListCmd(listFlags)
	case cmdInfo.FullCommand():
		err = commands.InfoCmd(infoFlags)
	case cmdRestore.FullCommand():
		err = commands.RestoreCmd(restoreFlags)
	}

	if err != nil {
		log.Printf(errtpl, err)
		os.Exit(1)
	}
}
