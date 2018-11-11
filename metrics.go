package main

// NodeMetrics : A struct for holding the metrics of a single Node
type NodeMetrics struct {
	Timeslice   float64 `json:"timeslice"`
	NodeCPUTime float64 `json:"cpu_time"`
	NodeMemUsed float64 `json:"mem_used"`
}
