package utils

import (

	log "github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"os"
	"fmt"

	"github.com/wshihadeh/kubectl-counts/pkg/client"
	"github.com/wshihadeh/kubectl-counts/pkg/options"
)

// setOptions - set common options for clientset
func setOptions(opt *options.SearchOptions) (string, *metav1.ListOptions) {
	// set default namespace as "default"
	namespace := "default"

	// override `namespace` if `--all-namespaces` exist
	if opt.AllNamespaces {
		namespace = ""
	} else {
		if len(opt.Namespace) > 0 {
			namespace = opt.Namespace
		} else {
			ns, _, err := client.ClientConfig().Namespace()
			if err != nil {
				log.WithFields(log.Fields{
					"err": err.Error(),
				}).Debug("Failed to resolve namespace")
			} else {
				namespace = ns
			}
		}
	}

	// retrieve listOptions from meta
	listOptions := &metav1.ListOptions{}

	return namespace, listOptions
}

func CheckFoundResources(count int) {
	if count <= 0 {
		fmt.Printf("%s", "No resources found\n")
		os.Exit(1)
	}
}