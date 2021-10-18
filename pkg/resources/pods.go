package resources

import (
	"bytes"
	"fmt"
	"sort"
	"strconv"
	"text/tabwriter"

	"github.com/wshihadeh/kubectl-counts/pkg/constants"
	"github.com/wshihadeh/kubectl-counts/pkg/options"
	"github.com/wshihadeh/kubectl-counts/pkg/utils"
)

// Pods - a public function for searching pods with keyword
func Pods(opt *options.SearchOptions, by string) {
	switch by {
	case "ns":
		opt.AllNamespaces = true
		podsByNs(opt)
	case "node":
		podsByNode(opt)
	case "restarts":
		podsByRestarts(opt)
	case "status":
		podsByStatus(opt)
	default:
		break
	}
}

func podsByRestarts(opt *options.SearchOptions) {
	var podInfo string

	podList := utils.PodList(opt)
	utils.CheckFoundResources(len(podList.Items))

	buf := bytes.NewBuffer(nil)
	w := tabwriter.NewWriter(buf, 0, 0, 3, ' ', 0)

	fmt.Fprintln(w, constants.PodsByRestartsHeader)
	restarts := make(map[string]int)

	for _, p := range podList.Items {
		var restartCount int32 = 0

		for _, c := range p.Status.ContainerStatuses {
			restartCount = c.RestartCount + restartCount
		}
		var r = fmt.Sprintf("%d", restartCount)
		var count = restarts[r]
		restarts[r] = count + 1
	}

	keys := make([]int, 0, len(restarts))
	for k := range restarts {
		i, _ := strconv.Atoi(k)
		keys = append(keys, i)
	}

	sort.Ints(keys)
	//sort.Sort(sort.Reverse(sort.IntSlice(keys)))

	for _, k := range keys {
		podInfo = fmt.Sprintf(constants.RowTemplate,
			fmt.Sprintf("%d", k),
			restarts[fmt.Sprintf("%d", k)],
		)
		fmt.Fprintln(w, podInfo)
	}

	w.Flush()
	fmt.Printf("%s", buf.String())
}

func podsByStatus(opt *options.SearchOptions) {
	var podInfo string

	podList := utils.PodList(opt)
	utils.CheckFoundResources(len(podList.Items))

	buf := bytes.NewBuffer(nil)
	w := tabwriter.NewWriter(buf, 0, 0, 3, ' ', 0)

	fmt.Fprintln(w, constants.PodsByStatusHeader)
	statuses := make(map[string]int)

	for _, p := range podList.Items {
		var count = statuses[string(p.Status.Phase)]
		statuses[string(p.Status.Phase)] = count + 1
	}

	for status, count := range statuses {
		podInfo = fmt.Sprintf(constants.RowTemplate,
			status,
			count,
		)
		fmt.Fprintln(w, podInfo)
	}

	w.Flush()
	fmt.Printf("%s", buf.String())
}

func podsByNs(opt *options.SearchOptions) {
	var podInfo string

	podList := utils.PodList(opt)
	utils.CheckFoundResources(len(podList.Items))

	buf := bytes.NewBuffer(nil)
	w := tabwriter.NewWriter(buf, 0, 0, 3, ' ', 0)

	fmt.Fprintln(w, constants.PodsByNamespaceHeader)
	namespaces := make(map[string]int)

	for _, p := range podList.Items {
		var count = namespaces[p.Namespace]
		namespaces[p.Namespace] = count + 1
	}

	for ns, count := range namespaces {
		podInfo = fmt.Sprintf(constants.RowTemplate,
			ns,
			count,
		)
		fmt.Fprintln(w, podInfo)
	}

	w.Flush()
	fmt.Printf("%s", buf.String())
}

func podsByNode(opt *options.SearchOptions) {
	var podInfo string

	podList := utils.PodList(opt)
	utils.CheckFoundResources(len(podList.Items))

	buf := bytes.NewBuffer(nil)
	w := tabwriter.NewWriter(buf, 0, 0, 3, ' ', 0)

	fmt.Fprintln(w, constants.PodsByNodeHeader)
	nodes := make(map[string]int)

	for _, p := range podList.Items {
		var count = nodes[p.Spec.NodeName]
		nodes[p.Spec.NodeName] = count + 1
	}

	for node, count := range nodes {
		podInfo = fmt.Sprintf(constants.RowTemplate,
			node,
			count,
		)
		fmt.Fprintln(w, podInfo)
	}

	w.Flush()
	fmt.Printf("%s", buf.String())
}
