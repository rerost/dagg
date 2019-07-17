package cmd

import (
	"github.com/izumin5210/clig/pkg/clib"
	"github.com/pkg/errors"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"k8s.io/utils/exec"

	"github.com/rerost/dagg/pkg/dagg"
)

func NewDefaultDaggCommand(wd clib.Path, build clib.Build) *cobra.Command {
	return NewDaggCommand(&dagg.Ctx{
		WorkingDir: wd,
		IO:         clib.Stdio(),
		FS:         afero.NewOsFs(),
		Viper:      viper.New(),
		Exec:       exec.New(),
		Build:      build,
	})
}

func NewDaggCommand(ctx *dagg.Ctx) *cobra.Command {
	cmd := &cobra.Command{
		Use: ctx.Build.AppName,
		PersistentPreRunE: func(c *cobra.Command, args []string) error {
			return errors.WithStack(ctx.Init())
		},
	}

	clib.AddLoggingFlags(cmd)

	cmd.AddCommand(
		clib.NewVersionCommand(ctx.IO, ctx.Build),
	)

	return cmd
}
