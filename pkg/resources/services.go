package resources

import (
	"bytes"
	"fmt"
	"text/tabwriter"

	"github.com/wshihadeh/kubectl-counts/pkg/constants"
	"github.com/wshihadeh/kubectl-counts/pkg/options"
	"github.com/wshihadeh/kubectl-counts/pkg/utils"
)

// Services - a public function for searching services
func Services(opt *options.SearchOptions, by string) {
	switch by {
	case "ns":
		opt.AllNamespaces = true
		servicesByNs(opt)
	case "type":
		servicesByType(opt)
	default:
		break
	}
}

func servicesByNs(opt *options.SearchOptions) {
	var serviceInfo string

	serviceList := utils.ServiceList(opt)
	utils.CheckFoundResources(len(serviceList.Items))

	buf := bytes.NewBuffer(nil)
	w := tabwriter.NewWriter(buf, 0, 0, 3, ' ', 0)

	fmt.Fprintln(w, constants.ServicesByNamespaceHeader)

	namespaces := make(map[string]int)

	for _, s := range serviceList.Items {
		var count = namespaces[s.Namespace]
		namespaces[s.Namespace] = count + 1
	}

	for ns, count := range namespaces {
		serviceInfo = fmt.Sprintf(constants.RowTemplate,
			ns,
			count,
		)
		fmt.Fprintln(w, serviceInfo)
	}

	w.Flush()
	fmt.Printf("%s", buf.String())
}

func servicesByType(opt *options.SearchOptions) {
	var serviceInfo string

	serviceList := utils.ServiceList(opt)
	utils.CheckFoundResources(len(serviceList.Items))

	buf := bytes.NewBuffer(nil)
	w := tabwriter.NewWriter(buf, 0, 0, 3, ' ', 0)

	fmt.Fprintln(w, constants.ServicesByTypeHeader)

	types := make(map[string]int)

	for _, s := range serviceList.Items {
		var count = types[string(s.Spec.Type)]
		types[string(s.Spec.Type)] = count + 1
	}

	for t, count := range types {
		serviceInfo = fmt.Sprintf(constants.RowTemplate,
			t,
			count,
		)
		fmt.Fprintln(w, serviceInfo)
	}

	w.Flush()
	fmt.Printf("%s", buf.String())
}
