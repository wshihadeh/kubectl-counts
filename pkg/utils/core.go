package utils

import (
	"context"

	log "github.com/sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"

	"github.com/wshihadeh/kubectl-counts/pkg/client"
	"github.com/wshihadeh/kubectl-counts/pkg/options"
)

// PodList - return a list of Pod(s)
func PodList(opt *options.SearchOptions) *corev1.PodList {
	clientset := client.InitClient()
	ns, o := setOptions(opt)
	list, err := clientset.CoreV1().Pods(ns).List(context.TODO(), *o)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Debug("Unable to get Pod List")
	}
	return list
}

// ServiceList - return a list of Services
func ServiceList(opt *options.SearchOptions) *corev1.ServiceList {
	clientset := client.InitClient()
	ns, o := setOptions(opt)
	list, err := clientset.CoreV1().Services(ns).List(context.TODO(), *o)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Debug("Unable to get Service List")
	}
	return list
}
