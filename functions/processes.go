package functions

import "github.com/shirou/gopsutil/v3/process"

type ProcessInfo struct {
	PID         int32
	Name        string
	CPU         float64
	Memory      float64
	MemoryUsage float32
	MemorySize  uint64
}

func GetProcessInfo() []ProcessInfo {
	processes, _ := process.Processes()

	var processInfo []ProcessInfo

	for _, p := range processes {
		name, _ := p.Name()
		pid := p.Pid
		cpuUsage, _ := p.CPUPercent()
		mem, _ := p.MemoryInfo()
		memUsage, _ := p.MemoryPercent()

		if mem != nil {
			processInfo = append(processInfo, ProcessInfo{
				PID:         pid,
				Name:        name,
				CPU:         cpuUsage,
				Memory:      float64(mem.RSS),
				MemoryUsage: memUsage,
				MemorySize:  mem.RSS,
			})
		}
	}
	return processInfo
}

func SortByMemory(processes []ProcessInfo) []ProcessInfo {
	for i := 0; i < len(processes); i++ {
		for j := i + 1; j < len(processes); j++ {
			if float64(processes[i].Memory) > float64(processes[j].Memory) {
				processes[i], processes[j] = processes[j], processes[i]
			}
		}
	}

	if len(processes) > 10 {
		processes = processes[:10]
	}

	return processes
}

func SortByCPU(processes []ProcessInfo) []ProcessInfo {
	for i := 0; i < len(processes); i++ {
		for j := i + 1; j < len(processes); j++ {
			if processes[i].CPU < processes[j].CPU {
				processes[i], processes[j] = processes[j], processes[i]
			}
		}
	}

	if len(processes) > 10 {
		processes = processes[:10]
	}

	return processes
}
