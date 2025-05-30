package compare

import (
	"fmt"

	"kompare/DAO"
	"kompare/cli"
	"kompare/query"
	"kompare/tools"

	routeclient "github.com/openshift/client-go/route/clientset/versioned"
)

func CompareRoutes(clientsetToSource, clientsetToTarget *routeclient.Clientset, namespaceName string, TheArgs cli.ArgumentsReceivedValidated) ([]DAO.DiffWithName, error){
	var TheDiff []DAO.DiffWithName
	sourceRoutes, err := query.ListRoutes(clientsetToSource, namespaceName)
	if err != nil {
		fmt.Printf("Error getting routes list: %v\n", err)
		return TheDiff, err
	}
	targetRoutes, err := query.ListRoutes(clientsetToTarget, namespaceName)
	if err != nil {
		fmt.Printf("Error getting deployments list: %v\n", err)
		return TheDiff, err
	}
	var diffCriteria []string
	if TheArgs.FiltersForObject == "" {
		diffCriteria = []string{"Spec.Host", "Name"}
	} else {
		diffCriteria = tools.ParseCommaSeparateList(TheArgs.FiltersForObject)
	}
	return CompareVerboseVSNonVerbose(sourceRoutes, targetRoutes, diffCriteria, TheArgs)
}
