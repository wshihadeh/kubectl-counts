package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/wshihadeh/kubectl-counts/pkg/options"
)

type optionList []string

var (
	by                      string
	supportedOptions        = make(map[string]optionList)
	rootCmdDescriptionShort = "Get k8s resource counts"
	rootCmdDescriptionLong  = `Get k8s resource counts

More info: https://github.com/wshihadeh/kubectl-counts
`

	rootCmdExamples = `
Get pod counts
$ kubectl counts pod -by=ns
$ kubectl counts pod -by=node
$ kubectl counts pod -by=restarts
$ kubectl counts pod -by=status

Get service counts
$ kubectl counts services -by=ns
$ kubectl counts services -by=type

Get jobs counts
$ kubectl counts jobs -by=ns
$ kubectl counts jobs -by=container
$ kubectl counts jobs -by=image

Get deployment counts
$ kubectl counts deployments -by=ns

Get ingress counts
$ kubectl counts ingresses -by=ns
$ kubectl counts ingresses -by=class
$ kubectl counts ingresses -by=address
`
)

// generic search options handler
var searchOptions = options.NewSearchOptions()

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "kubectl-counts",
	Aliases: []string{"c", "count"},
	Short:   rootCmdDescriptionShort,
	Long:    rootCmdDescriptionLong,
	Example: rootCmdExamples,
}


func (list optionList) Has(option string) bool {
	for _, o := range list {
		if o == option {
			return true
		}
	}
	return false
}

func init() {
	// Global Flags
	rootCmd.PersistentFlags().StringVarP(
		&searchOptions.Namespace, "namespace", "n", "",
		"Namespace for search. (default: \"default\")")
	rootCmd.PersistentFlags().BoolVarP(
		&searchOptions.AllNamespaces, "all-namespaces", "A", false,
		"If present, list the requested object(s) across all namespaces.")
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
