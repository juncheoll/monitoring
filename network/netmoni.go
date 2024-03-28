package network

import (
	"fmt"

	"github.com/shirou/gopsutil/net"
)

var prevInfo map[string]map[string]uint64

func Init() error {
	prevInfo = make(map[string]map[string]uint64)

	ioCounters, err := net.IOCounters(true)
	if err != nil {
		return err
	}

	for _, io := range ioCounters {
		prevInfo[io.Name] = make(map[string]uint64)
		prevInfo[io.Name]["bSent"] = io.BytesSent
		prevInfo[io.Name]["bRecv"] = io.BytesRecv
		prevInfo[io.Name]["pSent"] = io.PacketsSent
		prevInfo[io.Name]["pRecv"] = io.PacketsRecv
	}

	return nil
}

func PrintNetworkInfo() error {
	fmt.Println("-----------Network Info-----------")
	ioCounters, err := net.IOCounters(true)
	if err != nil {
		return err
	}

	for _, io := range ioCounters {
		fmt.Printf("Interface: %s\n", io.Name)

		bSent := prevInfo[io.Name]["bSent"] - io.BytesSent
		bRecv := prevInfo[io.Name]["bRecv"] - io.BytesRecv
		pSent := prevInfo[io.Name]["pSent"] - io.PacketsSent
		pRecv := prevInfo[io.Name]["pRecv"] - io.PacketsRecv
		fmt.Printf("Bytes Sent: %d, Bytes Recv: %d\n", bSent, bRecv)
		fmt.Printf("Packets Sent: %d, Packets Recv: %d\n", pSent, pRecv)

		prevInfo[io.Name]["bSent"] = io.BytesSent
		prevInfo[io.Name]["bRecv"] = io.BytesRecv
		prevInfo[io.Name]["pSent"] = io.PacketsSent
		prevInfo[io.Name]["pRecv"] = io.PacketsRecv
	}

	return nil
}
