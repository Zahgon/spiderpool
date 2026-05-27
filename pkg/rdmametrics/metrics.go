// Copyright 2024 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package rdmametrics

import (
	"context"
	"os"

	"github.com/vishvananda/netlink"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"go.uber.org/zap"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/utils/exec"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/spidernet-io/spiderpool/pkg/lock"
	"github.com/spidernet-io/spiderpool/pkg/podownercache"
	"github.com/spidernet-io/spiderpool/pkg/rdmametrics/oteltype"
)

var netnsPathList = []string{"/var/run/netns", "/var/run/docker/netns"}

var (
	readDir                = os.ReadDir
	rdmaSystemGetNetnsMode = netlink.RdmaSystemGetNetnsMode

	rdmaMetricsPrefix = "rdma_"

	knownMetricsKeyDescription = map[string]string{
		"rx_write_requests":               "The number of received WRITE requests for the associated QPs.",
		"rx_read_requests":                "The number of received read requests",
		"rx_atomic_requests":              "The number of received atomic requests",
		"rx_dct_connect":                  "The number of received DCT connect requests",
		"out_of_buffer":                   "The number of out of buffer errors",
		"out_of_sequence":                 "The number of out-of-order arrivals",
		"duplicate_request":               "The number of duplicate requests",
		"rnr_nak_retry_err":               "The number of received RNR NAK packets did not exceed the QP retry limit",
		"packet_seq_err":                  "The number of packet sequence errors",
		"implied_nak_seq_err":             "The number of implied NAK sequence errors",
		"local_ack_timeout_err":           "The number of times QP's ack timer expired for RC, XRC, DCT QPs at the sender side",
		"resp_local_length_error":         "The number of times responder detected local length errors",
		"resp_cqe_error":                  "The number of response CQE errors",
		"req_cqe_error":                   "The number of times requester detected CQEs completed with errors",
		"req_remote_invalid_request":      "The number of times requester detected remote invalid request errors",
		"req_remote_access_errors":        "The number of request remote access errors",
		"resp_remote_access_errors":       "The number of response remote access errors",
		"resp_cqe_flush_error":            "The number of response CQE flush errors",
		"req_cqe_flush_error":             "The number of request CQE flush errors",
		"roce_adp_retrans":                "The number of RoCE adaptive retransmissions",
		"roce_adp_retrans_to":             "The number of RoCE adaptive retransmission timeouts",
		"roce_slow_restart":               "The number of RoCE slow restart",
		"roce_slow_restart_cnps":          "The number of times RoCE slow restart generated CNP packets",
		"roce_slow_restart_trans":         "The number of times RoCE slow restart changed state to slow restart",
		"rp_cnp_ignored":                  "The number of CNP packets received and ignored by the Reaction Point HCA",
		"rp_cnp_handled":                  "The number of CNP packets handled by the Reaction Point HCA to throttle the transmission rate",
		"np_ecn_marked_roce_packets":      "The number of RoCEv2 packets received by the notification point which were marked for experiencing the congestion (ECN bits where '11' on the ingress RoCE traffic)",
		"np_cnp_sent":                     "The number of CNP packets sent by the Notification Point when it noticed congestion experienced in the RoCEv2 IP header (ECN bits)",
		"rx_icrc_encapsulated":            "The number of RoCE packets with ICRC errors",
		"rx_vport_rdma_unicast_packets":   "The number of unicast RDMA packets received on the virtual port.",
		"tx_vport_rdma_unicast_packets":   "The number of unicast RDMA packets transmitted from the virtual port.",
		"rx_vport_rdma_multicast_packets": "The number of multicast RDMA packets received on the virtual port.",
		"tx_vport_rdma_multicast_packets": "The number of multicast RDMA packets transmitted from the virtual port.",
		"rx_vport_rdma_unicast_bytes":     "The number of bytes received in unicast RDMA packets on the virtual port.",
		"tx_vport_rdma_unicast_bytes":     "The number of bytes transmitted in unicast RDMA packets from the virtual port.",
		"rx_vport_rdma_multicast_bytes":   "The number of bytes received in multicast RDMA packets on the virtual port.",
		"tx_vport_rdma_multicast_bytes":   "The number of bytes transmitted in multicast RDMA packets from the virtual port.",
		"vport_speed_mbps":                "The speed of the virtual port expressed in megabits per second (Mbps).",
		"rx_discards":                     "The number of packets discarded by the device.",
		"tx_discards":                     "The number of packets discarded by the device.",
		"rx_pause":                        "The number of packets dropped by the device.",
		"tx_pause":                        "The number of packets dropped by the device.",
		"device_tos":                      "RDMA device traffic class (TOS) value.",
	}
)

type GetObservable func(string) (metric.Int64ObservableCounter, bool)

type EthtoolImpl struct {
	Stats   func(netIfName string) ([]oteltype.Metrics, error)
	BusInfo func(string) (string, error)
}

type NetlinkImpl struct {
	RdmaLinkList func() ([]*netlink.RdmaLink, error)
	LinkList     func() ([]netlink.Link, error)
}

type RDMADevice struct {
	NetDevName   string
	NodeGUID     string
	SysImageGUID string
	IsRoot       bool
}

type NetnsItem struct {
	ID string
	Fd string
}

func Register(ctx context.Context, meter metric.Meter, cache podownercache.CacheInterface) error {
	_ = "STUB: not implemented"
	return nil
}

type exporter struct {
	nodeName              attribute.KeyValue
	meter                 metric.Meter
	lock                  lock.Mutex
	log                   *zap.Logger
	ch                    chan struct{}
	netns                 func(netns NetnsItem, toRun func() error) error
	netlinkImpl           NetlinkImpl
	ethtool               EthtoolImpl
	exec                  exec.Interface
	registration          metric.Registration
	waitToRegisterMetrics map[string]struct{}
	observableMap         map[string]metric.Int64ObservableCounter
	cache                 podownercache.CacheInterface
}

func (e *exporter) registerMetrics(meter metric.Meter) error { _ = "STUB: not implemented"; return nil }

// register known metrics

// register discovered metrics

func (e *exporter) reRegisterMetrics() error { _ = "STUB: not implemented"; return nil }

func (e *exporter) daemon(ctx context.Context) { _ = "STUB: not implemented"; return }

func (e *exporter) Callback(ctx context.Context, observer metric.Observer) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *exporter) updateUnregisteredMetrics(unRegistrationMetric []string) {
	_ = "STUB: not implemented"
	return
}

func (e *exporter) processNetNS(netns NetnsItem,
	vfToPfNameMap map[string]string,
	observer metric.Observer, getObservable GetObservable,
) error {
	_ = "STUB: not implemented"
	return nil
}

// host netns, don't need get default ip to mapping pod metadata to metrics

func listNodeNetNS() ([]NetnsItem, error) { _ = "STUB: not implemented"; return nil, nil }

// skip default netns, default netns is a host netns

func getIPToPodMap(ctx context.Context, cli client.Client) (map[string]types.NamespacedName, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// getNodeGUIDNetDeviceNameMap get map of node guid to rdma name
func getNodeGUIDNetDeviceNameMap(nl NetlinkImpl) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getVfToPfNameMap() (map[string]string, error) { _ = "STUB: not implemented"; return nil, nil }

func getVfToPfNameMapAt(devicesRoot string) (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getIfnameNetDevMap(nl NetlinkImpl) (map[string]RDMADevice, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// for example:
// ib device hardware addr 00:00:01:af:fe:80:00:00:00:00:00:00:03:a7:83:7a:20:bf:ed:2f
// node guid = addr[36:] =                                     03:a7:83:7a:20:bf:ed:2f

// 3a:b0:33:ff:fe:1a:0d:70
// 3a:b0:33:      1a:0d:70

// getDefaultIP returns the default IP address of the host/pod
func getDefaultIP(e exec.Interface) (string, error) {
	_ = "STUB: not implemented"
	// Check for IPv4 default route
	return "", nil
}

// Check for IPv6 default route

func extractSrcIPStringIndex(raw string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

// camelToSnake converts a camelCase string to snake_case
func camelToSnake(camel string) string { _ = "STUB: not implemented"; return "" }

func reverseMACAddress(mac string) string { _ = "STUB: not implemented"; return "" }
