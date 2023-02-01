package util

import (
	"fmt"
	"runtime"
	"time"
)

type sysStatus struct {
	Uptime       string `json:"服务运行时间"`
	NumGoroutine int    `json:"当前 Goroutines 数量"`

	// General statistics.
	MemAllocated string `json:"当前内存使用量"`  // bytes allocated and still in use
	MemTotal     string `json:"所有被分配的内存"` // bytes allocated (even if freed)
	MemSys       string `json:"内存占用量"`    // bytes obtained from system (sum of XxxSys below)
	Lookups      uint64 `json:"指针查找次数"`   // number of pointer lookups
	MemMallocs   uint64 `json:"内存分配次数"`   // number of mallocs
	MemFrees     uint64 `json:"内存释放次数"`   // number of frees

	// Main allocation heap statistics.
	HeapAlloc    string `json:"当前 Heap 内存使用量"` // bytes allocated and still in use
	HeapSys      string `json:"Heap 内存占用量"`    // bytes obtained from system
	HeapIdle     string `json:"Heap 内存空闲量"`    // bytes in idle spans
	HeapInuse    string `json:"正在使用的 Heap 内存"` // bytes in non-idle span
	HeapReleased string `json:"被释放的 Heap 内存"`  // bytes released to the OS
	HeapObjects  uint64 `json:"Heap 对象数量"`     // total number of allocated objects

	// Low-level fixed-size structure allocator statistics.
	//	Inuse is bytes used now.
	//	Sys is bytes obtained from system.
	StackInuse  string `json:"启动 Stack 使用量"` // bootstrap stacks
	StackSys    string `json:"被分配的 Stack 内存"`
	MSpanInuse  string `json:"MSpan 结构内存使用量"` // mspan structures
	MSpanSys    string `json:"被分配的 MSpan 结构内存"`
	MCacheInuse string `json:"MCache 结构内存使用量"` // mcache structures
	MCacheSys   string `json:"被分配的 MCache 结构内存"`
	BuckHashSys string `json:"被分配的剖析哈希表内存"`   // profiling bucket hash table
	GCSys       string `json:"被分配的 GC 元数据内存"` // GC metadata
	OtherSys    string `json:"其它被分配的系统内存"`    // other system allocations

	// Garbage collector statistics.
	NextGC       string `json:"下次 GC 内存回收量"` // next run in HeapAlloc time (bytes)
	LastGC       string `json:"距离上次 GC 时间"`  // last run in absolute time (ns)
	PauseTotalNs string `json:"GC 暂停时间总量"`
	PauseNs      string `json:"上次 GC 暂停时间"` // circular buffer of recent GC pause times, most recent at [(NumGC+255)%256]
	NumGC        uint32 `json:"GC 执行次数"`
}

type SysInfo struct {
	SysStatus sysStatus
}

func UpdateSystemStatus(sysStatus *sysStatus) {
	sysStatus.Uptime = TimeSincePro(AppStartTime, time.Now())

	m := new(runtime.MemStats)
	runtime.ReadMemStats(m)
	sysStatus.NumGoroutine = runtime.NumGoroutine()

	sysStatus.MemAllocated = FileSize(int64(m.Alloc))
	sysStatus.MemTotal = FileSize(int64(m.TotalAlloc))
	sysStatus.MemSys = FileSize(int64(m.Sys))
	sysStatus.Lookups = m.Lookups
	sysStatus.MemMallocs = m.Mallocs
	sysStatus.MemFrees = m.Frees

	sysStatus.HeapAlloc = FileSize(int64(m.HeapAlloc))
	sysStatus.HeapSys = FileSize(int64(m.HeapSys))
	sysStatus.HeapIdle = FileSize(int64(m.HeapIdle))
	sysStatus.HeapInuse = FileSize(int64(m.HeapInuse))
	sysStatus.HeapReleased = FileSize(int64(m.HeapReleased))
	sysStatus.HeapObjects = m.HeapObjects

	sysStatus.StackInuse = FileSize(int64(m.StackInuse))
	sysStatus.StackSys = FileSize(int64(m.StackSys))
	sysStatus.MSpanInuse = FileSize(int64(m.MSpanInuse))
	sysStatus.MSpanSys = FileSize(int64(m.MSpanSys))
	sysStatus.MCacheInuse = FileSize(int64(m.MCacheInuse))
	sysStatus.MCacheSys = FileSize(int64(m.MCacheSys))
	sysStatus.BuckHashSys = FileSize(int64(m.BuckHashSys))
	sysStatus.GCSys = FileSize(int64(m.GCSys))
	sysStatus.OtherSys = FileSize(int64(m.OtherSys))

	sysStatus.NextGC = FileSize(int64(m.NextGC))
	sysStatus.LastGC = fmt.Sprintf("%.1fs", float64(time.Now().UnixNano()-int64(m.LastGC))/1000/1000/1000)
	sysStatus.PauseTotalNs = fmt.Sprintf("%.1fs", float64(m.PauseTotalNs)/1000/1000/1000)
	sysStatus.PauseNs = fmt.Sprintf("%.3fs", float64(m.PauseNs[(m.NumGC+255)%256])/1000/1000/1000)
	sysStatus.NumGC = m.NumGC
}
