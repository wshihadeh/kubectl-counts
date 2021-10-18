package utils

import (
	"context"

	log "github.com/sirupsen/logrus"
	networkingv1beta1 "k8s.io/api/networking/v1beta1"

	"github.com/wshihadeh/kubectl-counts/pkg/client"
	"github.com/wshihadeh/kubectl-counts/pkg/options"
)

// IngressList - return a list of Ingresses
func IngressList(opt *options.SearchOptions) *networkingv1beta1.IngressList {
	clientset := client.InitClient()
	ns, o := setOptions(opt)
	list, err := clientset.NetworkingV1beta1().Ingresses(ns).List(context.TODO(), *o)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Debug("Unable to get Ingress List")
	}
	return list
}
