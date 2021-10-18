package utils

import (
	"context"

	log "github.com/sirupsen/logrus"
	appsv1 "k8s.io/api/apps/v1"

	"github.com/wshihadeh/kubectl-counts/pkg/client"
	"github.com/wshihadeh/kubectl-counts/pkg/options"
)

// DeploymentList - return a list of Deployment(s)
func DeploymentList(opt *options.SearchOptions) *appsv1.DeploymentList {
	clientset := client.InitClient()
	ns, o := setOptions(opt)
	list, err := clientset.AppsV1().Deployments(ns).List(context.TODO(), *o)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
		}).Debug("Unable to get Deployment List")
	}
	return list
}