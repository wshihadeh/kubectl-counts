package resources

import (
	"bytes"
	"fmt"
	"text/tabwriter"

	"github.com/wshihadeh/kubectl-counts/pkg/constants"
	"github.com/wshihadeh/kubectl-counts/pkg/options"
	"github.com/wshihadeh/kubectl-counts/pkg/utils"
)

// Ingresses - a public function for searching ingresses with keyword
func Ingresses(opt *options.SearchOptions, by string) {
	switch by {
	case "ns":
		opt.AllNamespaces = true
		IngressesByNs(opt)
	case "class":
		IngressesByClass(opt)
	case "address":
		IngressesByAdress(opt)
	default:
		break
	}
}

func IngressesByNs(opt *options.SearchOptions) {
	var ingressInfo string

	ingressList := utils.IngressList(opt)
	utils.CheckFoundResources(len(ingressList.Items))

	buf := bytes.NewBuffer(nil)
	w := tabwriter.NewWriter(buf, 0, 0, 3, ' ', 0)

	fmt.Fprintln(w, constants.IngressesByNamespaceHeader)

	namespaces := make(map[string]int)

	for _, i := range ingressList.Items {
		var count = namespaces[i.Namespace]
		namespaces[i.Namespace] = count + 1
	}

	for ns, count := range namespaces {
		ingressInfo = fmt.Sprintf(constants.RowTemplate,
			ns,
			count,
		)
		fmt.Fprintln(w, ingressInfo)
	}

	w.Flush()
	fmt.Printf("%s", buf.String())
}

func IngressesByClass(opt *options.SearchOptions) {
	var ingressInfo string

	ingressList := utils.IngressList(opt)
	utils.CheckFoundResources(len(ingressList.Items))

	buf := bytes.NewBuffer(nil)
	w := tabwriter.NewWriter(buf, 0, 0, 3, ' ', 0)

	fmt.Fprintln(w, constants.IngressesByClasseHeader)

	classes := make(map[string]int)

	for _, i := range ingressList.Items {
		var count = classes[fmt.Sprintf("%s", i.Spec.IngressClassName)]
		classes[fmt.Sprintf("%s", i.Spec.IngressClassName)] = count + 1
	}

	for class, count := range classes {
		ingressInfo = fmt.Sprintf(constants.RowTemplate,
			class,
			count,
		)
		fmt.Fprintln(w, ingressInfo)
	}

	w.Flush()
	fmt.Printf("%s", buf.String())
}

func IngressesByAdress(opt *options.SearchOptions) {
	var ingressInfo string

	ingressList := utils.IngressList(opt)
	utils.CheckFoundResources(len(ingressList.Items))

	buf := bytes.NewBuffer(nil)
	w := tabwriter.NewWriter(buf, 0, 0, 3, ' ', 0)

	fmt.Fprintln(w, constants.IngressesByAdressHeader)

	namespaces := make(map[string]int)

	for _, i := range ingressList.Items {
		var count = namespaces[i.Status.LoadBalancer.Ingress[0].Hostname]
		namespaces[i.Status.LoadBalancer.Ingress[0].Hostname] = count + 1
	}

	for ns, count := range namespaces {
		ingressInfo = fmt.Sprintf(constants.RowTemplate,
			ns,
			count,
		)
		fmt.Fprintln(w, ingressInfo)
	}

	w.Flush()
	fmt.Printf("%s", buf.String())
}
