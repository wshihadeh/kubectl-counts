package resources

import (
	"bytes"
	"fmt"
	"text/tabwriter"

	"github.com/wshihadeh/kubectl-counts/pkg/constants"
	"github.com/wshihadeh/kubectl-counts/pkg/options"
	"github.com/wshihadeh/kubectl-counts/pkg/utils"
)

// Deployments - a public function for searching deployments with keyword
func Deployments(opt *options.SearchOptions, by string) {
	switch by {
	case "ns":
		opt.AllNamespaces = true
		deploymentsByNs(opt)
	default:
		break
	}
}

func deploymentsByNs(opt *options.SearchOptions) {
	var deploymentInfo string

	deploymentList := utils.DeploymentList(opt)
	utils.CheckFoundResources(len(deploymentList.Items))

	buf := bytes.NewBuffer(nil)
	w := tabwriter.NewWriter(buf, 0, 0, 3, ' ', 0)

	fmt.Fprintln(w, constants.DeploymentByNamespaceHeader)
	namespaces := make(map[string]int)

	for _, d := range deploymentList.Items {
		var count = namespaces[d.Namespace]
		namespaces[d.Namespace] = count + 1
	}

	for namespace, count := range namespaces {
		deploymentInfo = fmt.Sprintf(constants.RowTemplate,
			namespace,
			count,
		)
		fmt.Fprintln(w, deploymentInfo)
	}

	w.Flush()
	fmt.Printf("%s", buf.String())
}
