package cmd

import (
	"fmt"
	"github.com/wshihadeh/kubectl-counts/pkg/resources"
	"os"

	"github.com/spf13/cobra"
	"strings"
)

var (
	// deploymentsCmd represents the pods command
	deploymentsCmd = &cobra.Command{
		Use:     "deployments",
		Aliases: []string{"deploy", "deployment"},
		Short:   "Get Deployments counts by namespace",
		Run: func(cmd *cobra.Command, args []string) {
			getResoourceCounts(args, "deployments")
		},
	}

	// podsCmd represents the pods command
	podsCmd = &cobra.Command{
		Use:     "pods",
		Aliases: []string{"po", "pod"},
		Short:   "Get Pods counts by namespace, node, restarts, and status",
		Run: func(cmd *cobra.Command, args []string) {
			getResoourceCounts(args, "pods")
		},
	}

	// ingressesCmd represents the ingresses command
	ingressesCmd = &cobra.Command{
		Use:     "ingresses",
		Aliases: []string{"ing", "ingress"},
		Short:   "Get Ingresses counts by namespace, class and address",
		Run: func(cmd *cobra.Command, args []string) {
			getResoourceCounts(args, "ingresses")
		},
	}

	// jobsCmd represents the jobs command
	jobsCmd = &cobra.Command{
		Use:     "jobs",
		Aliases: []string{"job"},
		Short:   "Get Jobs counts by namespace, container and image",
		Run: func(cmd *cobra.Command, args []string) {
			getResoourceCounts(args, "jobs")
		},
	}

	// servicesCmd represents the service command
	servicesCmd = &cobra.Command{
		Use:     "services",
		Aliases: []string{"svc", "service"},
		Short:   "Get Services counts by namespace and tyoe",
		Run: func(cmd *cobra.Command, args []string) {
			getResoourceCounts(args, "services")
		},
	}
)

func init() {
	rootCmd.AddCommand(deploymentsCmd)
	supportedOptions["deployments"] = append(supportedOptions["deployments"], "ns")
	rootCmd.AddCommand(podsCmd)
	supportedOptions["pods"] = append(supportedOptions["pods"], "ns")
	supportedOptions["pods"] = append(supportedOptions["pods"], "node")
	supportedOptions["pods"] = append(supportedOptions["pods"], "restarts")
	supportedOptions["pods"] = append(supportedOptions["pods"], "status")
	rootCmd.AddCommand(ingressesCmd)
	supportedOptions["ingresses"] = append(supportedOptions["ingresses"], "ns")
	supportedOptions["ingresses"] = append(supportedOptions["ingresses"], "class")
	supportedOptions["ingresses"] = append(supportedOptions["ingresses"], "address")
	rootCmd.AddCommand(jobsCmd)
	supportedOptions["jobs"] = append(supportedOptions["jobs"], "ns")
	supportedOptions["jobs"] = append(supportedOptions["jobs"], "container")
	supportedOptions["jobs"] = append(supportedOptions["jobs"], "image")
	rootCmd.AddCommand(servicesCmd)
	supportedOptions["services"] = append(supportedOptions["services"], "ns")
	supportedOptions["services"] = append(supportedOptions["services"], "type")

	deploymentsCmd.Flags().StringVarP(&by, "by", "b", "ns", fmt.Sprintf("By Counts options: %s", strings.Join(supportedOptions["deployments"], ", ")))
	podsCmd.Flags().StringVarP(&by, "by", "b", "ns", fmt.Sprintf("By Counts options: %s", strings.Join(supportedOptions["pods"], ", ")))
	ingressesCmd.Flags().StringVarP(&by, "by", "b", "ns", fmt.Sprintf("By Counts options: %s", strings.Join(supportedOptions["ingresses"], ", ")))
	jobsCmd.Flags().StringVarP(&by, "by", "b", "ns", fmt.Sprintf("By Counts options: %s", strings.Join(supportedOptions["jobs"], ", ")))
	servicesCmd.Flags().StringVarP(&by, "by", "b", "ns", fmt.Sprintf("By Counts options: %s", strings.Join(supportedOptions["services"], ", ")))

}

func getResoourceCounts(args []string, resourceType string) {

	if !supportedOptions[resourceType].Has(by) {
		fmt.Println(fmt.Sprintf("Unsported option value: The value %s is not supported by %s command. Supported values are %s", by, resourceType, strings.Join(supportedOptions[resourceType], ",")))
		os.Exit(1)
	}

	switch resourceType {
	case "deployments":
		resources.Deployments(searchOptions, by)
	case "pods":
		resources.Pods(searchOptions, by)
	case "ingresses":
		resources.Ingresses(searchOptions, by)
	case "jobs":
		resources.Jobs(searchOptions, by)
	case "services":
		resources.Services(searchOptions, by)
	default:
		break
	}
}
