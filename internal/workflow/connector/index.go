package connector

import (
	"dip/internal/workflow/filter"
	"time"
)

type DipResult struct {
	StartTime     time.Time
	EndTime       time.Time
	Duration      int
	Input         string
	Output        string
	ErrorMsg      string
	Status        string
	Configuration string
}

type DipContext struct {
}

type Trigger struct {
	EndPoint string
	Type     string
}

type Executor interface {
	Execute(ctx *DipContext) DipResult
}

type ConnectExecutor struct {
	Type         string
	Name         string
	ServiceName  string
	PreFilters   []filter.FilterExecute
	PostFilters  []filter.FilterExecute
	Executor     Executor
	ExecuteOrder int
}

type WorkFlow struct {
	ServiceName string
	Name        string
	Trigger     Trigger
	Connectors  []ConnectExecutor
}

func (wf WorkFlow) Execute(ctx *DipContext) DipResult {
	//startTime := time.Now()
	//for _, connector := range wf.Connectors {
	//	connectorExecuteStartTime := time.Now()
	//	err := doFilter(ctx, connector.PreFilters)
	//	result := connector.Executor.Execute(ctx)
	//	err := doFilter(ctx, connector.PreFilters)
	//	connectorExecuteEndTime := time.Now()
	//}
	//endTime := time.Now()
	return DipResult{}
}

func doFilter(ctx *DipContext, filters []filter.FilterExecute) error {
	for _, filter := range filters {
		return filter.Executor.DoFilter("")
	}
	return nil
}
