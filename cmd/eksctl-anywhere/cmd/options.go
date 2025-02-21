package cmd

import (
	"fmt"

	"github.com/aws/eks-anywhere/pkg/cluster"
	"github.com/aws/eks-anywhere/pkg/version"
)

type clusterOptions struct {
	fileName             string
	bundlesOverride      string
	managementKubeconfig string
}

func newClusterSpec(options clusterOptions) (*cluster.Spec, error) {
	var specOpts []cluster.SpecOpt
	if options.bundlesOverride != "" {
		specOpts = append(specOpts, cluster.WithOverrideBundlesManifest(options.bundlesOverride))
	}
	if options.managementKubeconfig != "" {
		managementCluster, err := cluster.LoadManagement(options.managementKubeconfig)
		if err != nil {
			return nil, fmt.Errorf("unable to get management cluster from kubeconfig: %v", err)
		}
		specOpts = append(specOpts, cluster.WithManagementCluster(managementCluster))
	}

	clusterSpec, err := cluster.NewSpec(options.fileName, version.Get(), specOpts...)
	if err != nil {
		return nil, fmt.Errorf("unable to get cluster config from file: %v", err)
	}

	return clusterSpec, nil
}
