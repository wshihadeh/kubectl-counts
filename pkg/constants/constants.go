package constants

const (
	// DeploymentByNamespaceHeader : Header for the deployment command
	DeploymentByNamespaceHeader = "NAMESPACE\tDEPLOYMENTS"
	// PodsByNamespaceHeader : Header for the pods command
	PodsByNamespaceHeader = "NAMESPACE\tPODS"
	// PodsByNodeHeader : Header for the pods command
	PodsByNodeHeader = "NODE\tPODS"
	// PodsByStatusHeader : Header for the pods command
	PodsByStatusHeader = "STATUS\tPODS"
	// PodsByRestartsHeader : Header for the pods command
	PodsByRestartsHeader = "RESTARTS-COUNT\tPODS"

	// ServicesByNamespaceHeader : Header for the services command
	ServicesByNamespaceHeader = "NAMESPACE\tSERVICES"
	// ServicesByTypeHeader : Header for the services command
	ServicesByTypeHeader = "TYPE\tSERVICES"

	// IngressesByNamespaceHeader : Header for the ingresses command
	IngressesByNamespaceHeader = "NAMESPACE\tINGRESSES"
	// IngressesByClasseHeader : Header for the ingresses command
	IngressesByClasseHeader = "CLASS\tINGRESSES"
	// IngressesByAdressHeader : Header for the ingresses command
	IngressesByAdressHeader = "ADRESS\tINGRESSES"

	// JobsByNamespaceHeader : Header for the jobs command
	JobsByNamespaceHeader = "NAMESPACE\tJOBS"
	// JobsByClasseHeader : Header for the jobs command
	JobsByClasseHeader = "CONTAINER\tJOBS"
	// JobsByAdressHeader : Header for the jobs command
	JobsByAdressHeader = "IMAGE\tJOBS"

	// RowTemplate : Output template
	RowTemplate = "%s\t%d"
)
