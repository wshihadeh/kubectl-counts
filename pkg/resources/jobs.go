package resources

import (
	"bytes"
	"fmt"
	"text/tabwriter"

	"github.com/wshihadeh/kubectl-counts/pkg/constants"
	"github.com/wshihadeh/kubectl-counts/pkg/options"
	"github.com/wshihadeh/kubectl-counts/pkg/utils"
)

// Jobs - a public function for searching jobs with keyword
func Jobs(opt *options.SearchOptions, by string) {
	switch by {
	case "ns":
		opt.AllNamespaces = true
		JobsByNs(opt)
	case "container":
		JobsByConatiner(opt)
	case "image":
		JobsByImage(opt)
	default:
		break
	}
}

func JobsByNs(opt *options.SearchOptions) {
	var jobInfo string

	jobList := utils.JobList(opt)
	utils.CheckFoundResources(len(jobList.Items))

	buf := bytes.NewBuffer(nil)
	w := tabwriter.NewWriter(buf, 0, 0, 3, ' ', 0)

	fmt.Fprintln(w, constants.JobsByNamespaceHeader)

	namespaces := make(map[string]int)

	for _, i := range jobList.Items {
		var count = namespaces[i.Namespace]
		namespaces[i.Namespace] = count + 1
	}

	for ns, count := range namespaces {
		jobInfo = fmt.Sprintf(constants.RowTemplate,
			ns,
			count,
		)
		fmt.Fprintln(w, jobInfo)
	}

	w.Flush()
	fmt.Printf("%s", buf.String())
}

func JobsByConatiner(opt *options.SearchOptions) {
	var jobInfo string

	jobList := utils.JobList(opt)
	utils.CheckFoundResources(len(jobList.Items))

	buf := bytes.NewBuffer(nil)
	w := tabwriter.NewWriter(buf, 0, 0, 3, ' ', 0)

	fmt.Fprintln(w, constants.JobsByNamespaceHeader)

	namespaces := make(map[string]int)

	for _, i := range jobList.Items {
		var count = namespaces[i.Spec.Template.Spec.Containers[0].Name]
		namespaces[i.Spec.Template.Spec.Containers[0].Name] = count + 1
	}

	for ns, count := range namespaces {
		jobInfo = fmt.Sprintf(constants.RowTemplate,
			ns,
			count,
		)
		fmt.Fprintln(w, jobInfo)
	}

	w.Flush()
	fmt.Printf("%s", buf.String())
}

func JobsByImage(opt *options.SearchOptions) {
	var jobInfo string

	jobList := utils.JobList(opt)
	utils.CheckFoundResources(len(jobList.Items))

	buf := bytes.NewBuffer(nil)
	w := tabwriter.NewWriter(buf, 0, 0, 3, ' ', 0)

	fmt.Fprintln(w, constants.JobsByNamespaceHeader)

	namespaces := make(map[string]int)

	for _, i := range jobList.Items {
		var count = namespaces[i.Spec.Template.Spec.Containers[0].Image]
		namespaces[i.Spec.Template.Spec.Containers[0].Image] = count + 1
	}

	for ns, count := range namespaces {
		jobInfo = fmt.Sprintf(constants.RowTemplate,
			ns,
			count,
		)
		fmt.Fprintln(w, jobInfo)
	}

	w.Flush()
	fmt.Printf("%s", buf.String())
}
