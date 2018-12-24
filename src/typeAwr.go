package main

import "github.com/influxdata/influxdb/client/v2"

type MainTable struct {
	WorkInfo					WorkInfo
	SQLOrderByElapsedTime      	[]SQLOrderByElapsedTime
	SQLOrderedByCPUTime        	[]SQLOrderedByCPUTime
	SQLOrderedByUserIOWaitTime 	[]SQLOrderedByUserIOWaitTime
	SQLOrderedByGets			[]SQLOrderedByGets
	SQLOrderedByReads			[]SQLOrderedByReads
	SQLOrderedByExecutions		[]SQLOrderedByExecutions
	SQLOrderedByVersionCount	[]SQLOrderedByVersionCount
	TopSQLWithTopEvents        	[]TopSQLWithTopEvents
	TopSQLWithTopRowSources    	[]TopSQLWithTopRowSources
	OperatingSystemStatistics	[]OperatingSystemStatistics
	CompleteListOfSQLText      	[]CompleteListOfSQLText
	BufferPoolStatistics		BufferPoolStatistics
	WaitEventsStatistics		WaitEventsStatistics
	ReportSummary				ReportSummary
}
type WorkInfo struct {
	WIDatabaseInstanceInformation	[]WIDatabaseInstanceInformation
	WIHostInformation				[]WIHostInformation
	WISnapshotInformation			[]WISnapshotInformation
}
// This table displays database instance information
type WIDatabaseInstanceInformation struct{
	DBName		string
	DBId		float64
	Instance	string
	Instnum		float64
	StartupTime	string
	Release		string
	RAC			string
}
// This table displays host information
type WIHostInformation struct {
	HostName	string
	Platform	string
	CPUs		float64
	Cores		float64
	Sockets		float64
	MemoryGB	float64
}
// This table displays snapshot information
type WISnapshotInformation struct {
	Name			string
	SnapId			float64
	SnapTime		string
	Sessions		float64
	CursorsSession	float64
}
// TODO SQL ordered by Gets
type SQLOrderedByGets struct{
	BufferGets 			float64
	Executions			float64
	GetsPerExec 		float64
	Total				float64
	ElapsedTime			float64
	CPU					float64
	IO					float64
	SQLID				string
	SQLModule			string
	SQLText				string
}

// TODO SQL ordered by Reads
type SQLOrderedByReads struct{

}
// TODO SQL ordered by Executions
type SQLOrderedByExecutions struct{

}
// TODO SQL ordered by Version Count
type SQLOrderedByVersionCount struct{

}
// TODO Buffer Pool Statistics
type BufferPoolStatistics struct{

}
// Report Summary
type ReportSummary struct{
	TopADDMFindingsByAverageActiveSessions	[]TopADDMFindingsByAverageActiveSessions
	LoadProfile								[]LoadProfile
	InstanceEfficiencyPercentages			[]InstanceEfficiencyPercentages
	Top10ForegroundEventsByTotalWaitTime	[]TopForegroundEventsByTotalWaitTime
	WaitClassesByTotalWaitTime				[]WaitClassesByTotalWaitTime
	HostCPU 								[]HostCPU
	InstanceCPU 							[]InstanceCPU
	IOProfile								[]IOProfile
	MemoryStatistics 						[]MemoryStatistics
	CacheSizes 								[]CacheSizes
	SharedPoolStatistics 					[]SharedPoolStatistics
}
// Wait Events Statistics
type WaitEventsStatistics struct{
	ForegroundWaitClass		[]ForegroundWaitClass
	ForegroundWaitEvents	[]ForegroundWaitEvents
	BackgroundWaitEvents	[]BackgroundWaitEvents
}
// Foreground Wait Class
type ForegroundWaitClass struct {
	WaitClass		string
	Waits 			float64
	PerTime			float64
	TotalWaitTime	float64
	AvgWait			float64
	PerDBTime		float64
}
// Foreground Wait Events
type ForegroundWaitEvents struct {
	Event 			string
	Waits 			float64
	PerTime			float64
	TotalWaitTime 	float64
	AvgWait			float64
	WaitsTxn		float64
	PerDBTime		float64
}
// Background Wait Events
type BackgroundWaitEvents struct {
	Event 			string
	Waits 			float64
	PerTime			float64
	TotalWaitTime 	float64
	AvgWait			float64
	WaitsTxn		float64
	PerbgTime		float64
}
// Top ADDM Findings by Average Active Sessions
type TopADDMFindingsByAverageActiveSessions struct{
	FindingName					string
	AvgActiveSessionsTask		float64
	PerActiveSessionsFinding	float64
	TaskName					string
	BeginSnapTime				string
	EndSnapTime					string
}
// Load Profile
type LoadProfile struct{
	Name			string
	PerSecond		float64
	PerTransaction	float64
	PerExec			float64
	PerCall			float64
}
// Instance Efficiency Percentages (Target 100%)
type InstanceEfficiencyPercentages struct{ // TODO перейти на map[string]float64
	Name	string
	Value	float64
}
//Top 10 Foreground Events by Total Wait Time
type TopForegroundEventsByTotalWaitTime struct{
	Event			string
	Waits 			float64
	TotalWaitTime 	float64
	WaitAvg			float64
	PerDBTime		float64
	WaitClass		string
}
//Wait Classes by Total Wait Time
type WaitClassesByTotalWaitTime struct{
	WaitClass			string
	Waits				float64
	TotalWaitTime		float64
	AvgWait				float64
	PerDBTime			float64
	AvgActiveSessions	float64
}
// Host CPU
type HostCPU struct{
	CPUs	float64
	Cores	float64
	Sockets	float64
	LABegin	float64
	LAEnd	float64
	PerUser	float64
	PerSystem	float64
	PerWIO	float64
	PerIDLE	float64
}
// Instance CPU
type InstanceCPU struct{
	PerTotalCPU				float64
	PerBysuCPU				float64
	PerDBTimeWaiting		float64
}
// IO Profile
type IOProfile struct {
	Name			string
	RWPerSecond		float64
	ReadPerSecond	float64
	WritePerSecond	float64
}
// Memory Statistics
type MemoryStatistics struct {
	Name 	string
	Begin	float64
	End		float64
}
// Cache Sizes
type CacheSizes struct {
	Name 	string
	Begin	float64
	End	float64
}
// Shared Pool Statistics
type SharedPoolStatistics struct {
	Name	string
	Begin	float64
	End		float64
}
// SQL ordered by Elapsed Time
type SQLOrderByElapsedTime struct{
	ElapsedTime			float64
	Executions			float64
	ElapsedTimePerExec	float64
	Total				float64
	CPU					float64
	IO					float64
	SQLID				string
	SQLModule			string
	SQLText				string
}
// SQL ordered by CPU Time
type SQLOrderedByCPUTime struct{
	CPUTime				float64
	Executions			float64
	CPUPerExec			float64
	Total				float64
	ElapsedTime			float64
	CPU					float64
	IO					float64
	SQLID				string
	SQLModule			string
	SQLText				string
}
// SQL ordered by User I/O Wait Time
type SQLOrderedByUserIOWaitTime struct{
	UserIOTime			float64
	Executions			float64
	UIOPerExec			float64
	Total				float64
	ElapsedTime			float64
	CPU					float64
	IO					float64
	SQLID				string
	SQLModule			string
	SQLText				string
}
// Top SQL with Top Events
type TopSQLWithTopEvents struct{
	SQLID        string
	PlanHash     float64
	Executions   float64
	Activity     float64
	Event        string
	EventPer     float64
	RowSource    string
	RowSourcePer float64
	SQLText      string
}
// Top SQL with Top Row Sources
type TopSQLWithTopRowSources struct{
	SQLID				string
	PlanHash			float64
	Executions			float64
	Activity			float64
	RowSource			string
	RowSourcePer		float64
	TopEvent			string
	EventPer			float64
	SQLText				string
}
//Operating System Statistics
type OperatingSystemStatistics struct{
	Statistic		string
	Value			float64
	EndValue		float64
}
// Complete List of SQL Text
type CompleteListOfSQLText struct{// TODO перейти на map[string]string
	SQLID		string
	SQLText		string
}
// Type for work with HTML
type PageData struct {
	PageTitle           						string
	AttributeUploadFile 						bool
	NonParseCPU 								string	// Instance Efficiency Percentages
	ParseCPUElapsd 								string	// Parse CPU to Parse Elapsd %
	SoftParse 									string	// Soft Parse % %
	SharedPoolStatistics 						string	// Memory Usage %
	SQLWithExecution 							string	// % SQL with executions>1
	WorkInfo									WorkInfo
	WaitEventsStatistics						WaitEventsStatistics
	BufferPoolStatistics						BufferPoolStatistics
	ListSQLText         						[]ListSQLText

	SQLOrderByElapsedTime      	[]SQLOrderByElapsedTime
	SQLOrderedByCPUTime        	[]SQLOrderedByCPUTime
	SQLOrderedByUserIOWaitTime 	[]SQLOrderedByUserIOWaitTime
	SQLOrderedByGets			[]SQLOrderedByGets
	SQLOrderedByReads			[]SQLOrderedByReads
	SQLOrderedByExecutions		[]SQLOrderedByExecutions
	SQLOrderedByVersionCount	[]SQLOrderedByVersionCount
	TopSQLWithTopEvents        	[]TopSQLWithTopEvents
	TopSQLWithTopRowSources    	[]TopSQLWithTopRowSources
	TopForegroundEventsByTotalWaitTime			[]TopForegroundEventsByTotalWaitTime
}
// List real need SQL
type ListSQLText struct {
	SQLId 		string
	SQLDescribe string
	SQLText		string
	TextUI		string
}

// work with DB
type Config struct {
	NameDB      string
	UrlDB       string
	Username    string
	Password    string
	Measurement string
	Debug       string
	Client      client.Client
	Results 	[]client.Result
}