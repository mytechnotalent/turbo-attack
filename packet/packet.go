// Terms Of Use
// ------------
// Do NOT use this on any computer you do not own or are not allowed to run this on.
// You may NEVER attempt to sell this, it is free and open source.
// The authors and publishers assume no responsibility.
// For educational purposes only.

// Package routine provides custom packet creation.
package packet

import (
	"github.com/mytechnotalent/turbo-attack/bit"
	"github.com/mytechnotalent/turbo-attack/convert"
	"github.com/mytechnotalent/turbo-attack/rand"
)

// TCP4 provides custom TCP4 packet creation.
// It will return a valid []byte or error.
func TCP4(size int, ipv4Byte, portByte []byte) (packet []byte, err error) {
	packet = make([]byte, size, size)
	// ethernet header - dst MAC addr [6 bytes]
	nUInt8 := rand.UInt8(255)
	nIntCB := bit.Clear(nUInt8, 0) // to ensure unicast, bit 0 is not set
	nIntCB = bit.Clear(nIntCB, 1)  // to ensure factory default, bit 1 is not set
	randDstMACAddr, _ := rand.Byte(5)
	nByte := convert.Int8ToByte(nIntCB)
	// ethernet header - dst MAC addr [6 bytes]
	packet[0] = *nByte
	packet[1] = randDstMACAddr[0]
	packet[2] = randDstMACAddr[1]
	packet[3] = randDstMACAddr[2]
	packet[4] = randDstMACAddr[3]
	packet[5] = randDstMACAddr[4]
	randSrcMACAddr, _ := rand.Byte(5)
	nByte = convert.Int8ToByte(nIntCB)
	// ethernet header - src MAC addr [6 bytes]
	packet[6] = *nByte
	packet[7] = randSrcMACAddr[0]
	packet[8] = randSrcMACAddr[1]
	packet[9] = randSrcMACAddr[2]
	packet[10] = randSrcMACAddr[3]
	packet[11] = randSrcMACAddr[4]
	// ethernet header - protocol type [2 bytes]
	packet[12] = 0x08
	packet[13] = 0x00
	// IP header - version & IHL [1 byte]
	packet[14] = 0x45
	// IP header - type of service [1 byte]
	packet[15] = 0x00
	// IP header - total length [2 bytes]
	packet[16] = 0x00
	packet[17] = 0x3c
	// IP header - identification [2 bytes]
	randIdentification, _ := rand.Byte(6)
	packet[18] = randIdentification[0]
	packet[19] = randIdentification[1]
	// IP header - flags, fragment, offset [2 bytes]
	packet[20] = 0x00
	packet[21] = 0x00
	// IP header - TTl [1 byte]
	packet[22] = 0x40
	// IP header - protocol [1 byte]
	packet[23] = 0x06
	// IP header - checksum [2 bytes]
	packet[24] = 0x00
	packet[25] = 0x00
	// IP header - src IP addr [4 bytes]
	randSrcIPAddr, _ := rand.Byte(4)
	packet[26] = randSrcIPAddr[0]
	packet[27] = randSrcIPAddr[1]
	packet[28] = randSrcIPAddr[2]
	packet[29] = randSrcIPAddr[3]
	// IP header - dst IP addr [4 bytes]
	packet[30] = ipv4Byte[12]
	packet[31] = ipv4Byte[13]
	packet[32] = ipv4Byte[14]
	packet[33] = ipv4Byte[15]
	// TCP header - src TCP port [2 bytes]
	randSrcTCPPort, _ := rand.Byte(2)
	packet[34] = randSrcTCPPort[0]
	packet[35] = randSrcTCPPort[1]
	// TCP header - dst TCP port [2 bytes]
	packet[36] = portByte[0]
	packet[37] = portByte[1]
	// TCP header - sequence number [4 bytes]
	randSeqNum, _ := rand.Byte(4)
	packet[38] = randSeqNum[0]
	packet[39] = randSeqNum[1]
	packet[40] = randSeqNum[2]
	packet[41] = randSeqNum[3]
	// TCP header - acknowledgement number [4 bytes]
	packet[42] = 0x00
	packet[43] = 0x00
	packet[44] = 0x00
	packet[45] = 0x00
	// TCP header - data offset reserved flags [2 bytes]
	packet[46] = 0xa0
	packet[47] = 0x02
	// TCP header - window size [2 bytes]
	packet[48] = 0x72
	packet[49] = 0x10
	// TCP header - checksum [2 bytes]
	packet[50] = 0x00
	packet[51] = 0x00
	// TCP header - urgent pointer [2 bytes]
	packet[52] = 0x00
	packet[53] = 0x00
	// TCP header option - mss [4 bytes]
	packet[54] = 0x02
	packet[55] = 0x04
	packet[56] = 0x05
	packet[57] = 0xb4
	// TCP header option - sack permitted [2 bytes]
	packet[58] = 0x04
	packet[59] = 0x02
	randTimeStamp, _ := rand.Byte(4)
	// TCP header option - time stamps [10 bytes]
	packet[60] = 0x08
	packet[61] = 0x0a
	packet[62] = randTimeStamp[0]
	packet[63] = randTimeStamp[1]
	packet[64] = randTimeStamp[2]
	packet[65] = randTimeStamp[3]
	packet[66] = 0x00
	packet[67] = 0x00
	packet[68] = 0x00
	packet[69] = 0x00
	// TCP header option - NOP
	packet[70] = 0x01
	// TCP header option - window scale
	packet[71] = 0x03
	packet[72] = 0x03
	packet[73] = 0x07
	return packet, err
}

// TCP6 provides custom TCP6 packet creation.
// It will return a []byte slice or error.
func TCP6(size int, ipv6Byte []byte, portByte []byte) (packet []byte, err error) {
	packet = make([]byte, size, size)
	// ethernet header - dst MAC addr [6 bytes]
	nUInt8 := rand.UInt8(255)
	nIntCB := bit.Clear(nUInt8, 0) // to ensure unicast, bit 0 is not set
	nIntCB = bit.Clear(nIntCB, 1)  // to ensure factory default, bit 1 is not set
	randDstMACAddr, _ := rand.Byte(5)
	nByte := convert.Int8ToByte(nIntCB)
	// ethernet header - dst MAC addr [6 bytes]
	packet[0] = *nByte
	packet[1] = randDstMACAddr[0]
	packet[2] = randDstMACAddr[1]
	packet[3] = randDstMACAddr[2]
	packet[4] = randDstMACAddr[3]
	packet[5] = randDstMACAddr[4]
	randSrcMACAddr, _ := rand.Byte(5)
	nByte = convert.Int8ToByte(nIntCB)
	// ethernet header - src MAC addr [6 bytes]
	packet[6] = *nByte
	packet[7] = randSrcMACAddr[0]
	packet[8] = randSrcMACAddr[1]
	packet[9] = randSrcMACAddr[2]
	packet[10] = randSrcMACAddr[3]
	packet[11] = randSrcMACAddr[4]
	// ethernet header - protocol type [2 bytes]
	packet[12] = 0x86
	packet[13] = 0xdd
	// IP header - version [4-bits] + traffic class [8-bits] + flow label [20-bits]
	packet[14] = 0x60
	packet[15] = 0x00
	packet[16] = 0x00
	packet[17] = 0x00
	// IP header - payload length [2 bytes]
	packet[18] = 0x00
	packet[19] = 0x14
	// IP header - next header [1 byte]
	packet[20] = 0x06
	// IP header - hop limit [1 byte]
	packet[21] = 0xff
	// IP header - src IP addr [16 bytes]
	randSrcIPAddr, _ := rand.Byte(16)
	packet[22] = randSrcIPAddr[0]
	packet[23] = randSrcIPAddr[1]
	packet[24] = randSrcIPAddr[2]
	packet[25] = randSrcIPAddr[3]
	packet[26] = randSrcIPAddr[4]
	packet[27] = randSrcIPAddr[5]
	packet[28] = randSrcIPAddr[6]
	packet[29] = randSrcIPAddr[7]
	packet[30] = randSrcIPAddr[8]
	packet[31] = randSrcIPAddr[9]
	packet[32] = randSrcIPAddr[10]
	packet[33] = randSrcIPAddr[11]
	packet[34] = randSrcIPAddr[12]
	packet[35] = randSrcIPAddr[13]
	packet[36] = randSrcIPAddr[14]
	packet[37] = randSrcIPAddr[15]
	// IP header - dst IP addr [16 bytes]
	packet[38] = ipv6Byte[0]
	packet[39] = ipv6Byte[1]
	packet[40] = ipv6Byte[2]
	packet[41] = ipv6Byte[3]
	packet[42] = ipv6Byte[4]
	packet[43] = ipv6Byte[5]
	packet[44] = ipv6Byte[6]
	packet[45] = ipv6Byte[7]
	packet[46] = ipv6Byte[8]
	packet[47] = ipv6Byte[9]
	packet[48] = ipv6Byte[10]
	packet[49] = ipv6Byte[11]
	packet[50] = ipv6Byte[12]
	packet[51] = ipv6Byte[13]
	packet[52] = ipv6Byte[14]
	packet[53] = ipv6Byte[15]
	// TCP header - src TCP port [2 bytes]
	randSrcTCPPort, _ := rand.Byte(2)
	packet[54] = randSrcTCPPort[0]
	packet[55] = randSrcTCPPort[1]
	// TCP header - dst TCP port [2 bytes]
	packet[56] = portByte[0]
	packet[57] = portByte[1]
	// TCP header - sequence number [4 bytes]
	randSeqNum, _ := rand.Byte(4)
	packet[58] = randSeqNum[0]
	packet[59] = randSeqNum[1]
	packet[60] = randSeqNum[2]
	packet[61] = randSeqNum[3]
	// TCP header - acknowledgement number [4 bytes]
	packet[62] = 0x00
	packet[63] = 0x00
	packet[64] = 0x00
	packet[65] = 0x00
	// TCP header - data offset reserved flags [2 bytes]
	packet[66] = 0x50
	packet[67] = 0x02
	// TCP header - window size [2 bytes]
	packet[68] = 0xff
	packet[69] = 0xff
	// TCP header - checksum [2 bytes]
	packet[70] = 0x1b
	packet[71] = 0x82
	// TCP header - urgent pointer [2 bytes]
	packet[72] = 0x00
	packet[73] = 0x00
	return packet, err
}
