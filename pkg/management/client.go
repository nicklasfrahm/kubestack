package management

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/nicklasfrahm/kubestack/api/v1alpha1"
	"github.com/nicklasfrahm/kubestack/pkg/management/common"
	"github.com/nicklasfrahm/kubestack/pkg/management/ssh"
)

const (
	// OSNexus is the probed name of Cisco NX-OS.
	OSNexus = "Nexus"
	// OSUbuntu is the probed name of Ubuntu.
	OSUbuntu = "Ubuntu"
)

var clientFactories = map[v1alpha1.Protocol]common.ClientFactory{
	v1alpha1.ProtocolSSH: ssh.NewClient,
	// TODO: Add support for SNMP.
}

// NewClient returns a new client for the given connection. Client will also
// implicitly connect to the appliance without an explicit call to Connect().
func NewClient(connRef types.NamespacedName, options ...common.Option) (common.Client, error) {
	opts, err := common.GetDefaultOptions().Apply(options...)
	if err != nil {
		return nil, err
	}

	if opts.KubernetesClient == nil {
		return nil, fmt.Errorf("missing required option: WithKubernetesClient()")
	}

	// TODO: Get context from options.
	ctx := context.TODO()

	conn := new(v1alpha1.Connection)
	err = opts.KubernetesClient.Get(ctx, connRef, conn)
	if err != nil {
		if client.IgnoreNotFound(err) == nil {
			return nil, fmt.Errorf("failed to read Connection: %s/%s", connRef.Namespace, connRef.Name)
		}
		return nil, err
	}

	newClient := clientFactories[conn.Spec.Protocol]
	if newClient == nil {
		return nil, fmt.Errorf("unknown protocol: %s", conn.Spec.Protocol)
	}

	mgmt, err := newClient(conn, common.WithKubernetesClient(opts.KubernetesClient))
	if err != nil {
		return nil, err
	}

	if err := mgmt.Connect(); err != nil {
		return nil, err
	}

	return mgmt, nil
}
