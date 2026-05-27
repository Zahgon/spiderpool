// Copyright 2025 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

// Note: The following source files are come from the latest version of https://github.com/k8snetworkplumbingwg/sriov-cni/blob/master/pkg/utils/packet.go.
// We can't directly go mod import package, it reports error:
// "require github.com/k8snetworkplumbingwg/sriov-cni: version “v2.8.0” invalid: should be v0 or v1, not v2.""

// So we copied the source files here and made some code changes.

package networking

import (
	"net"
	"time"

	"go.uber.org/zap"

	"github.com/vishvananda/netlink"

	"github.com/mdlayher/ndp"
)

var (
	arpPacketName    = "ARP"
	icmpV6PacketName = "ICMPv6"
)

// SetSocketTimeout sets the timeout for a socket in nanoseconds.
func SetSocketTimeout(sock int, timeout time.Duration) error { _ = "STUB: not implemented"; return nil }

// SendARPReuqest sends a gratuitous ARP packet with the provided source IP over the provided interface.
// UPDATE: the golang arp library requires an IPv4 address to exist for the NIC to send ARP request packets.
func SendARPReuqest(l netlink.Link, srcIP, dstIP net.IP) error {
	_ = "STUB: not implemented"
	/* As per RFC 5944 section 4.6, a gratuitous ARP packet can be sent by a node in order to spontaneously cause other nodes to update
	 * an entry in their ARP cache. In the case of SRIOV-CNI, an address can be reused for different pods. Each pod could likely have a
	 * different link-layer address in this scenario, which makes the ARP cache entries residing in the other nodes to be an invalid.
	 * The gratuitous ARP packet should update the link-layer address accordingly for the invalid ARP cache.
	 */return nil
}

// Construct the ARP packet following RFC 5944 section 4.6.

// Hardware Type: 1 is Ethernet

// Protocol Type: 0x0800 is IPv4

// Hardware address Length: 6 bytes for MAC address

// Protocol address length: 4 bytes for IPv4 address

// Operation: 1 is request, 2 is response

// Sender hardware address

// Sender protocol address

// Target hardware address is the Broadcast MAC.

// Target protocol address

// Ethertype of ARP (0x0806)
// Interface Index
// Hardware Type: 1 is Ethernet
// Packet Type.
// Hardware address Length: 6 bytes for MAC address
// Address is the broadcast MAC address.

// Create a socket such that the Ethernet header would constructed by the OS. The arpPacket only contains the ARP payload.

func SendUnsolicitedNeighborAdvertisement(dstIP net.IP, ifi *net.Interface, ndpClient *ndp.Conn) error {
	_ = "STUB: not implemented"
	return nil
}

// Always multicast the message to the target's solicited-node multicast
// group as if we have no knowledge of its MAC address.

// we send a gratuitous neighbor solicitation to checking if ip is conflict

// SendUnsolicitedNeighborAdvertisement sends an unsolicited neighbor advertisement packet with the provided source IP over the provided interface.
func SendUnsolicitedNeighborAdvertisement1(dstIP net.IP, l netlink.Link) error {
	_ = "STUB: not implemented"
	/* As per RFC 4861, a link-layer address change can multicast a few unsolicited neighbor advertisements to all nodes to quickly
	 * update the cached link-layer addresses that have become invalid. In the case of SRIOV-CNI, an address can be reused for
	 * different pods. Each pod could likely have a different link-layer address in this scenario, which makes the Neighbor Cache
	 * entries residing in the neighbors to be an invalid. The unsolicited neighbor advertisement should update the link-layer address
	 * accordingly for the IPv6 entry.
	 * However if any of these conditions are true:
	 *  - The IPv6 address was not reused for the new pod.
	 *  - No prior established communication with the neighbor.
	 * Then the neighbor receiving this unsolicited neighbor advertisement would be silently discard. This behavior is described
	 * in RFC 4861 section 7.2.5. This is acceptable behavior since the purpose of sending an unsolicited neighbor advertisement
	 * is not to create a new entry but rather update already existing invalid entries.
	 */return nil
}

// Construct the ICMPv6 Neighbor Advertisement packet following RFC 4861.
// payload := new(bytes.Buffer)
// ICMPv6 Flags: As per RFC 4861, the solicited flag must not be set and the override flag should be set (to
// override existing cache entry) for unsolicited advertisements.
// if writeErr := binary.Write(payload, binary.BigEndian, uint32(0x20000000)); writeErr != nil {
// 	return formatPacketFieldWriteError("Flags", icmpV6PacketName, writeErr)
// }
// if _, writeErr := payload.Write(dstIP.To16()); writeErr != nil { // ICMPv6 Target IPv6 Address.
// 	return formatPacketFieldWriteError("Target IPv6 Address", icmpV6PacketName, writeErr)
// }
// if writeErr := binary.Write(payload, binary.BigEndian, uint8(2)); writeErr != nil { // ICMPv6 Option Type: 2 is target link-layer address.
// 	return formatPacketFieldWriteError("Option Type", icmpV6PacketName, writeErr)
// }
// if writeErr := binary.Write(payload, binary.BigEndian, uint8(1)); writeErr != nil { // ICMPv6 Option Length. Units of 8 bytes.
// 	return formatPacketFieldWriteError("Option Length", icmpV6PacketName, writeErr)
// }
// if _, writeErr := payload.Write(l.Attrs().HardwareAddr); writeErr != nil { // ICMPv6 Option Link-layer Address.
// 	return formatPacketFieldWriteError("Option Link-layer Address", icmpV6PacketName, writeErr)
// }
// Construct ICMPv6 Neighbor Solicitation message

// Target IPv6 address
//"failed to write target IPv6 address: %v", writeErr)

// Source link-layer address option type

// Option length

// Source link-layer address

// ICMPv6 type is neighbor advertisement.
// ICMPv6 Code: As per RFC 4861 section 7.1.2, the code is always 0.
// Checksum is calculated later.

// Get the byte array of the ICMPv6 Message.

// Create a socket such that the Ethernet header and IPv6 header would constructed by the OS.

// As per RFC 4861 section 7.1.2, the IPv6 hop limit is always 255.

// Set the destination IPv6 address to the IPv6 link-local all nodes multicast address (ff02::1).

// Blocking wait for interface ifName to have carrier (!NO_CARRIER flag).
func WaitForCarrier(l netlink.Link, waitTime time.Duration) bool {
	_ = "STUB: not implemented"
	return false
}

/* Grow wait time exponentionally (factor 1.5). */

/* Wait for carrier, i.e. IFF_UP|IFF_RUNNING. Note that there is also
 * IFF_LOWER_UP, but we follow iproute2 ([1]).
 *
 * [1] https://git.kernel.org/pub/scm/network/iproute2/iproute2.git/tree/ip/ipaddress.c?id=f9601b10c21145f76c3d46c163bac39515ed2061#n86
 */

// htons converts an uint16 from host to network byte order.
func htons(i uint16) uint16 { _ = "STUB: not implemented"; return 0 }

// formatPacketFieldWriteError builds an error string for the cases when writing to a field of a packet fails.
func formatPacketFieldWriteError(field string, packetType string, writeErr error) error {
	_ = "STUB: not implemented"
	return nil
}

// NewSock returns a new raw socket to listen for ARP packets on the specified network interface.
func NewARPSockRAW(l netlink.Link) (fd int, err error) {
	_ = "STUB: not implemented"
	// Create a raw socket to listen for ARP packets.
	return 0, nil
}

// defer syscall.Close(sock)

// Bind the socket to the network interface.

func NewNDPSockRaw(iface string) (int, error) {
	_ = "STUB: not implemented"
	// Create a raw socket for ICMPv6
	return 0, nil
}

// Bind the socket to the network interface

func ParseIPv6NeighborAdvertisementMsg(n int, buf []byte) (srcIP net.IP, mac net.HardwareAddr, err error) {
	_ = "STUB: not implemented"
	return *new(net.IP), *new(net.HardwareAddr), nil
}

// this isn't a ICMPv6 NA message

// Extract the source IP address (offset 8-24 in the IPv6 header)

// Options start after the 24-byte ICMPv6 header

// Check if the target link-layer address option is present
// Iterate over options to find the target link-layer address option

// Length is in units of 8 octets

// Type 2 is the target link-layer address

func AnnounceIPs(logger *zap.Logger, iface string, ips []net.IP) error {
	_ = "STUB: not implemented"
	return nil
}

// send an gratuitous arp to announce the new mac address
