// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package multuscniconfig

import (
	"context"
	"time"

	netv1 "github.com/k8snetworkplumbingwg/network-attachment-definition-client/pkg/apis/k8s.cni.cncf.io/v1"
	"go.uber.org/zap"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/spidernet-io/spiderpool/pkg/election"
	spiderpoolv2beta1 "github.com/spidernet-io/spiderpool/pkg/k8s/apis/spiderpool.spidernet.io/v2beta1"
	crdclientset "github.com/spidernet-io/spiderpool/pkg/k8s/client/clientset/versioned"
	informers "github.com/spidernet-io/spiderpool/pkg/k8s/client/informers/externalversions/spiderpool.spidernet.io/v2beta1"
	listers "github.com/spidernet-io/spiderpool/pkg/k8s/client/listers/spiderpool.spidernet.io/v2beta1"
)

var informerLogger *zap.Logger

type MultusConfigController struct {
	MultusConfigControllerConfig
	client                client.Client
	multusConfigLister    listers.SpiderMultusConfigLister
	multusConfigSynced    cache.InformerSynced
	multusConfigWorkqueue workqueue.RateLimitingInterface
}

type MultusConfigControllerConfig struct {
	ControllerWorkers             int
	WorkQueueMaxRetries           int
	WorkQueueRequeueDelayDuration time.Duration
	LeaderRetryElectGap           time.Duration
	ResyncPeriod                  time.Duration
}

func NewMultusConfigController(multusConfigControllerConfig MultusConfigControllerConfig, client client.Client) *MultusConfigController {
	_ = "STUB: not implemented"
	return nil
}

func (mcc *MultusConfigController) SetupInformer(ctx context.Context, client crdclientset.Interface, leader election.SpiderLeaseElector) error {
	_ = "STUB: not implemented"
	return nil
}

func (mcc *MultusConfigController) addEventHandlers(multusConfigInformer informers.SpiderMultusConfigInformer) error {
	_ = "STUB: not implemented"
	return nil
}

func (mcc *MultusConfigController) enqueueMultusConfig(obj interface{}) {
	_ = "STUB: not implemented"
	return
}

func (mcc *MultusConfigController) Run(stopCh <-chan struct{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (mcc *MultusConfigController) runWorker() { _ = "STUB: not implemented"; return }

func (mcc *MultusConfigController) processNextWorkItem() bool {
	_ = "STUB: not implemented"
	return false
}

// discard some wrong input items

func (mcc *MultusConfigController) syncHandler(ctx context.Context, multusConfig *spiderpoolv2beta1.SpiderMultusConfig) error {
	_ = "STUB: not implemented"
	return nil
}

// use the annotation specified name as the CNI configuration name if set

// we need to wait and let the kubernetes delete this Net-Attach-Def first.

// the annotations updated

// the MultusConfig CNI configuration changed

// the net-attach-def ownerRef was removed

func generateNetAttachDef(netAttachName string, multusConf *spiderpoolv2beta1.SpiderMultusConfig) (*netv1.NetworkAttachmentDefinition, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// with Kubernetes OpenAPI validation, multusConfSpec.EnableCoordinator must not be nil

// head insertion later

// we'll use the default CNI version 0.3.1 if the annotation doesn't have it.
// the annotation custom CNI version is already validated by webhook.

// with Kubernetes OpenAPI validation, multusConfSpec.CniType must not be nil and default to "custom"

// head insertion

// we need to set Subvlan as first at the CNI plugin chain

// head insertion

// we need to set Subvlan as first at the CNI plugin chain

// SRIOV special annotation

// head insertion

// head insertion

// SRIOV special annotation

// head insertion

// head insertion

// It's impossible get into the default branch

func generateMacvlanCNIConf(disableIPAM bool, multusConfSpec spiderpoolv2beta1.MultusCNIConfigSpec) interface{} {
	_ = "STUB: not implemented"
	return nil

	// choose interface basement name
}

// set vlanID for interface basement name

// set default IPPools for spiderpool cni configuration

// if multusConfSpec.MacvlanConfig.SpiderpoolConfigPools.MatchMasterSubnet != nil {
// 	netConf.IPAM.MatchMasterSubnet = *multusConfSpec.MacvlanConfig.SpiderpoolConfigPools.MatchMasterSubnet
// }

func generateIPvlanCNIConf(disableIPAM bool, multusConfSpec spiderpoolv2beta1.MultusCNIConfigSpec) interface{} {
	_ = "STUB: not implemented"
	return nil

	// choose interface basement name
}

// set default IPPools for spiderpool cni configuration

// if multusConfSpec.IPVlanConfig.SpiderpoolConfigPools.MatchMasterSubnet != nil {
// 	netConf.IPAM.MatchMasterSubnet = *multusConfSpec.IPVlanConfig.SpiderpoolConfigPools.MatchMasterSubnet
// }

func generateVlanCNIConf(disableIPAM bool, multusConfSpec spiderpoolv2beta1.MultusCNIConfigSpec) interface{} {
	_ = "STUB: not implemented"
	return nil
}

func generateSriovCNIConf(disableIPAM bool, multusConfSpec spiderpoolv2beta1.MultusCNIConfigSpec) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// set default IPPools for spiderpool cni configuration

func generateIBSriovCNIConf(disableIPAM bool, multusConfSpec spiderpoolv2beta1.MultusCNIConfigSpec) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// set default IPPools for spiderpool cni configuration

func generateIpoibCNIConf(disableIPAM bool, multusConfSpec spiderpoolv2beta1.MultusCNIConfigSpec) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// set default IPPools for spiderpool cni configuration

func generateOvsCNIConf(disableIPAM bool, multusConfSpec *spiderpoolv2beta1.MultusCNIConfigSpec) interface{} {
	_ = "STUB: not implemented"
	return nil
}

func generateIfacer(master []string, vlanID int32, bond *spiderpoolv2beta1.BondConfig) interface{} {
	_ = "STUB: not implemented"
	return nil
}

func generateCoordinatorCNIConf(coordinatorSpec *spiderpoolv2beta1.CoordinatorSpec) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// coordinatorSpec could be nil, and we just need the coorinator CNI specified and use the default configuration

func marshalCniConfig2String(netAttachName, cniVersion string, plugins interface{}) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
