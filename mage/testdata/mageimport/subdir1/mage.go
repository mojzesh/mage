package mage

import (
	"fmt"

	"github.com/mojzesh/mage/mg"
)

// BuildSubdir Builds stuff.
func BuildSubdir() {
	fmt.Println("buildsubdir")
}

// NS is a namespace.
type NS mg.Namespace

// Deploy deploys stuff.
func (NS) Deploy() {
	fmt.Println("deploy")
}
