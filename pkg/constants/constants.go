package constants

const (
	DeploymentByNamespaceHeader       = "NAMESPACE\tDEPLOYMENTS"

	PodsByNamespaceHeader       = "NAMESPACE\tPODS"
	PodsByNodeHeader            = "NODE\tPODS"
	PodsByStatusHeader          = "STATUS\tPODS"
	PodsByRestartsHeader        = "RESTARTS-COUNT\tPODS"

	ServicesByNamespaceHeader   = "NAMESPACE\tSERVICES"
	ServicesByTypeHeader        = "TYPE\tSERVICES"

	IngressesByNamespaceHeader   = "NAMESPACE\tINGRESSES"
	IngressesByClasseHeader      = "CLASS\tINGRESSES"
	IngressesByAdressHeader      = "ADRESS\tINGRESSES"

	JobsByNamespaceHeader        = "NAMESPACE\tJOBS"
	JobsByClasseHeader           = "CONTAINER\tJOBS"
	JobsByAdressHeader           = "IMAGE\tJOBS"

	RowTemplate  = "%s\t%d"
)
