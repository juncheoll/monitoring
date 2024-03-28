package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"monitoring/network"

	"github.com/shirou/gopsutil/process"
)

func Monitoring() {
	pid := int32(os.Getpid())
	log.Printf("현재 프로세스 ID: %d\n", pid)

	proc, err := process.NewProcess(pid)
	if err != nil {
		log.Println("프로세스를 찾을 수 없습니다:", err)
		return
	}

	err = network.Init()
	if err != nil {
		log.Println("네트워크 정보 읽기 에러:", err)
	}

	for {
		time.Sleep(time.Second * 3)

		if err = printMonitoringInfo(proc); err != nil {
			log.Println("모니터링 에러:", err)
			return
		}
	}
}

func printMonitoringInfo(proc *process.Process) error {
	fmt.Println("---------Monitoring Info---------")

	cpuPercent, err := proc.CPUPercent()
	if err != nil {
		return err
	}
	fmt.Printf("CPU 사용률: %.3f%%\n", cpuPercent)

	memInfo, err := proc.MemoryInfo()
	if err != nil {
		return err
	}
	fmt.Printf("메모리 사용량: %v bytes\n", memInfo.RSS)

	err = network.PrintNetworkInfo()
	if err != nil {
		return err
	}

	return nil
}
