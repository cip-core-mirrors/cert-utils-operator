package controller

import (
	"github.com/redhat-cop/cert-utils-operator/pkg/controller/secrettokeystore"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, secrettokeystore.Add)
}
