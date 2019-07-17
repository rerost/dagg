package cmd

import (
	"html/template"
	"io/ioutil"
	"os"

	"github.com/izumin5210/clig/pkg/clib"
	"github.com/pkg/errors"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	yaml "gopkg.in/yaml.v2"
	"k8s.io/utils/exec"

	"github.com/rerost/dagg/pkg/dagg"
	"github.com/rerost/dagg/pkg/dagg/definition"
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
		Use:   ctx.Build.AppName,
		Short: "Dag generate tool",
		PersistentPreRunE: func(c *cobra.Command, args []string) error {
			if err := ctx.Init(); err != nil {
				return errors.WithStack(err)
			}
			return nil
		},
	}

	clib.AddLoggingFlags(cmd)

	cmd.AddCommand(
		clib.NewVersionCommand(ctx.IO, ctx.Build),
		&cobra.Command{
			Use:   "gen",
			Short: "Generate dag by template",
			Args:  cobra.ExactArgs(2),
			RunE: func(_ *cobra.Command, args []string) error {
				dagFile, err := ioutil.ReadFile(args[0])
				if err != nil {
					return errors.WithStack(err)
				}
				tpl := template.Must(template.ParseFiles(args[1]))

				dag := definition.Dag{}
				err = yaml.Unmarshal(dagFile, &dag)
				if err != nil {
					return errors.WithStack(err)
				}
				err = tpl.Execute(os.Stdout, dag)
				if err != nil {
					return errors.WithStack(err)
				}
				return nil
			},
		},
	)

	return cmd
}
