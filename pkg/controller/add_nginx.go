package controller

import (
	"github.com/Layquark/chubaoConsole/pkg/controller/nginx"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, nginx.Add)
}
