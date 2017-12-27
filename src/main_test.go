package main
import (
	"testing"
)

// TODO coverage = 100%
func TestParser(t *testing.T) {

	type testPair struct {
		input string
		text string
		error error
	}

	var tests = []testPair{
		{"test/test.html", `<meta http-equiv="content-type" content="text/html; charset=UTF-8"><html lang="en"><head><title>AWR Report for DB: ZP, Inst: k10szp_1, Snaps: 111755-111758</title>    <style type="text/css">        body.awr {font:bold 10pt Arial,Helvetica,Geneva,sans-serif;color:black; background:White;}        pre.awr  {font:8pt Courier;color:black; background:White;}            .hidden   {position:absolute;left:-10000px;top:auto;width:1px;height:1px;overflow:hidden;}        .pad   {margin-left:17px;}        .doublepad {margin-left:34px;}    </style></head><body class="awr"><h1 class="awr">    WORKLOAD REPOSITORY report for</h1><p /><table border="0" width="600" class="tdiff" summary="This table displays database instance information">    <tr><th class="awrbg" scope="col">DB Name</th><th class="awrbg" scope="col">DB Id</th><th class="awrbg" scope="col">Instance</th><th class="awrbg" scope="col">Inst num</th><th class="awrbg" scope="col">Startup Time</th><th class="awrbg" scope="col">Release</th><th class="awrbg" scope="col">RAC</th></tr>    <tr><td scope="row" class='awrnc'>ZP</td><td align="right" class='awrnc'>2966226569</td><td class='awrnc'>sdf_test1</td><td align="right" class='awrnc'>1</td><td class='awrnc'>08-Ноя-17 16:11</td><td class='awrnc'>12.1.0.2.0</td><td class='awrnc'>NO</td></tr></table><p /><p /><table border="0" width="600" class="tdiff" summary="This table displays host information">    <tr><th class="awrbg" scope="col">Host Name</th><th class="awrbg" scope="col">Platform</th><th class="awrbg" scope="col">CPUs</th><th class="awrbg" scope="col">Cores</th><th class="awrbg" scope="col">Sockets</th><th class="awrbg" scope="col">Memory (GB)</th></tr>    <tr><td scope="row" class='awrnc'>sdfeede</td><td class='awrnc'>Solaris[tm] OE (32-bit)</td><td align="right" class='awrnc'> 256</td><td align="right" class='awrnc'>  32</td><td align="right" class='awrnc'>   1</td><td align="right" class='awrnc'>  478.50</td></tr></table><p /><table border="0" width="600" class="tdiff" summary="This table displays snapshot information">    <tr><th class="awrnobg" scope="col"></th><th class="awrbg" scope="col">Snap Id</th><th class="awrbg" scope="col">Snap Time</th><th class="awrbg" scope="col">Sessions</th><th class="awrbg" scope="col">Cursors/Session</th></tr>    <tr><td scope="row" class='awrnc'>Begin Snap:</td><td align="right" class='awrnc'>111755</td><td align="center" class='awrnc'>10-Дек-17 13:30:55</td><td align="right" class='awrnc'>951</td><td align="right" class='awrnc'>     12.0</td></tr>    <tr><td scope="row" class='awrc'>End Snap:</td><td align="right" class='awrc'>111758</td><td align="center" class='awrc'>10-Дек-17 14:00:00</td><td align="right" class='awrc'>1130</td><td align="right" class='awrc'>     13.0</td></tr>    <tr><td scope="row" class='awrnc'>Elapsed:</td><td class='awrnc'>&#160;</td><td align="center" class='awrnc'>              29.09 (mins)</td><td class='awrnc'>&#160;</td><td class='awrnc'>&#160;</td></tr>    <tr><td scope="row" class='awrc'>DB Time:</td><td class='awrc'>&#160;</td><td align="center" class='awrc'>             204.95 (mins)</td><td class='awrc'>&#160;</td><td class='awrc'>&#160;</td></tr></table><p /><h3 class="awr"><a class="awr" name="99999"></a>Report Summary</h3><p />Top ADDM Findings by Average Active Sessions<p /><ul></ul><table border="0" width="600" class="tdiff" summary="This table displays top ADDM findings by average active sessions"><tr><th class="awrbg" scope="col">Finding Name</th><th class="awrbg" scope="col">Avg active sessions of the task</th><th class="awrbg" scope="col">Percent active sessions of finding</th><th class="awrbg" scope="col">Task Name</th><th class="awrbg" scope="col">Begin Snap Time</th><th class="awrbg" scope="col">End Snap Time</th></tr>    <tr><td class='awrc'>Фиксации и откаты</td><td align="right" class='awrc'>8.08</td><td align="right" class='awrc'>50.71</td><td scope="row" class='awrc'>ADDM:2926569_1_111758</td><td class='awrc'>10-Дек-17 13:50</td><td class='awrc'>10-Дек-17 14:00</td></tr>    <tr><td class='awrnc'>Фиксации и откаты</td><td align="right" class='awrnc'>7.07</td><td align="right" class='awrnc'>57.14</td><td scope="row" class='awrnc'>ADDM:2926569_1_111758</td><td class='awrnc'>10-Дек-17 13:40</td><td class='awrnc'>10-Дек-17 13:50</td></tr>    <tr><td class='awrc'>Фиксации и откаты</td><td align="right" class='awrc'>6.09</td><td align="right" class='awrc'>48.90</td><td scope="row" class='awrc'>ADDM:2926569_1_111758</td><td class='awrc'>10-Дек-17 13:30</td><td class='awrc'>10-Дек-17 13:40</td></tr>    <tr><td class='awrnc'>Наиболее часто используемые операторы SQL</td><td align="right" class='awrnc'>8.08</td><td align="right" class='awrnc'>31.41</td><td scope="row" class='awrnc'>ADDM:2966226569_1_111758</td><td class='awrnc'>10-Дек-17 13:50</td><td class='awrnc'>10-Дек-17 14:00</td></tr>    <tr><td class='awrc'>Наиболее часто используемые операторы SQL</td><td align="right" class='awrc'>7.07</td><td align="right" class='awrc'>25.59</td><td scope="row" class='awrc'>ADDM:2966226569_1_111757</td><td class='awrc'>10-Дек-17 13:40</td><td class='awrc'>10-Дек-17 13:50</td></tr></table><p /><p />Load Profile<p /><table border="0" width="600" class="tdiff" summary="This table displays load profile">    <tr><th class="awrnobg" scope="col"></th><th class="awrbg" scope="col">Per Second</th><th class="awrbg" scope="col">Per Transaction</th><th class="awrbg" scope="col">Per Exec</th><th class="awrbg" scope="col">Per Call</th></tr>    <tr><td scope="row" class='awrc'>DB Time(s):</td><td align="right" class='awrc'>               7.1</td><td align="right" class='awrc'>               0.0</td><td align="right" class='awrc'>      0.00</td><td align="right" class='awrc'>      0.00</td></tr>    <tr><td scope="row" class='awrnc'>DB CPU(s):</td><td align="right" class='awrnc'>               2.8</td><td align="right" class='awrnc'>               0.0</td><td align="right" class='awrnc'>      0.00</td><td align="right" class='awrnc'>      0.00</td></tr>    <tr><td scope="row" class='awrc'>Background CPU(s):</td><td align="right" class='awrc'>               0.4</td><td align="right" class='awrc'>               0.0</td><td align="right" class='awrc'>      0.00</td><td align="right" class='awrc'>      0.00</td></tr>    <tr><td scope="row" class='awrnc'>Redo size (bytes):</td><td align="right" class='awrnc'>       5,048,931.5</td><td align="right" class='awrnc'>           8,538.4</td><td class='awrnc'>&#160;</td><td class='awrnc'>&#160;</td></tr>    <tr><td scope="row" class='awrc'>L</tr></table></p><h3 class="awr">Background Wait Events</h3><ul>    <li class="awr"> ordered by wait time desc, waits desc (idle events last) </li>    <li class="awr"> Only events with Total Wait Time (s) &gt;= .001 are shown </li>    <li class="awr"> %Timeouts: value of 0 indicates value was &lt; .5%.  Value of null is truly 0</li></ul><table border="0" class="tdiff" summary="This table displays background wait events statistics"><tr><th class="awrbg" scope="col">Event</th><th class="awrbg" scope="col">Waits</th><th class="awrbg" scope="col">%Time -outs</th><th class="awrbg" scope="col">Total Wait Time (s)</th><th class="awrbg" scope="col">Avg wait (ms)</th><th class="awrbg" scope="col">Waits /txn</th><th class="awrbg" scope="col">% bg time</th></tr>    <tr><td scope="row" class='awrc'>log file parallel write</td><td align="right" class='awrc'>1,085,250</td><td align="right" class='awrc'>0</td><td align="right" class='awrc'>482</td><td align="right" class='awrc'>0.44</td><td align="right" class='awrc'>1.05</td><td align="right" class='awrc'>39.78</td></tr>    <tr><td scope="row" class='awrnc'>target log write size</td><td align="right" class='awrnc'>498,563</td><td align="right" class='awrnc'>0</td><td align="right" class='awrnc'>164</td><td align="right" class='awrnc'>0.33</td><td align="right" class='awrnc'>0.48</td><td align="right" class='awrnc'>13.53</td></tr>    <tr><td scope="row" class='awrc'>LGWR worker group ordering</td><td align="right" class='awrc'>202</td><td align="right" class='awrc'>0</td><td align="right" class='awrc'>88</td><td align="right" class='awrc'>437.54</td><td align="right" class='awrc'>0.00</td><td align="right" class='awrc'>7.29</td></tr>    <tr><td scope="row" class='awrnc'>LGWR any worker group</td><td align="right" class='awrnc'>199</td><td align="right" class='awrnc'>0</td><td align="right" class='awrnc'>86</td><td align="right" class='awrnc'>432.82</td><td align="right" class='awrnc'>0.00</td><td align="right" class='awrnc'>7.11</td></tr>    <tr><td scope="row" class='awrc'>db file parallel write</td><td align="right" class='awrc'>617,149</td><td align="right" class='awrc'>0</td><td align="right" class='awrc'>82</td><td align="right" class='awrc'>0.13</td><td align="right" class='awrc'>0.60</td><td align="right" class='awrc'>6.80</td></tr>    <tr><td scope="row" class='awrnc'>oracle thread bootstrap</td><td align="right" class='awrnc'>103</td><td align="right" class='awrnc'>0</td><td align="right" class='awrnc'>4</td><td align="right" class='awrnc'>40.41</td><td align="right" class='awrnc'>0.00</td><td align="right" class='awrnc'>0.34</td></tr>    <tr><td scope="row" class='awrc'>control file sequential read</td><td align="right" class='awrc'>5,960</td><td align="right" class='awrc'>0</td><td align="right" class='awrc'>2</td><td align="right" class='awrc'>0.28</td><td align="right" class='awrc'>0.01</td><td align="right" class='awrc'>0.14</td></tr>    <tr><td scope="row" class='awrnc'>cell single block physical read</td><td align="right" class='awrnc'>1,188</td><td align="right" class='awrnc'>0</td><td align="right" class='awrnc'>1</td><td align="right" class='awrnc'>1.05</td><td align="right" class='awrnc'>0.00</td><td align="right" class='awrnc'>0.10</td></tr>    <tr><td scope="row" class='awrc'>os thread creation</td><td align="right" class='awrc'>103</td><td align="right" class='awrc'>0</td><td align="right" class='awrc'>1</td><td align="right" class='awrc'>9.57</td><td align="right" class='awrc'>0.00</td><td align="right" class='awrc'>0.08</td></tr>    <tr><td scope="row" class='awrnc'>control file parallel write</td><td align="right" class='awrnc'>1,043</td><td align="right" class='awrnc'>0</td><td align="right" class='awrnc'>1</td><td align="right" class='awrnc'>0.57</td><td align="right" class='awrnc'>0.00</td><td align="right" class='awrnc'>0.05</td></tr>    <tr><td scope="row" class='awrc'>Disk file Mirror Read</td><td align="right" class='awrc'>1,283</td><td align="right" class='awrc'>0</td><td align="right" class='awrc'>1</td><td align="right" class='awrc'>0.44</td><td align="right" class='awrc'>0.00</td><td align="right" class='awrc'>0.05</td></tr>    <tr><td scope="row" class='awrnc'>reliable message</td><td align="right" class='awrnc'>2,330</td><td align="right" class='awrnc'>0</td><td align="right" class='awrnc'>0</td><td align="right" class='awrnc'>0.13</td><td align="right" class='awrnc'>0.00</td><td align="right" class='awrnc'>0.02</td></tr>    <tr><td scope="row" class='awrc'>ASM file metadata operation</td><td align="right" class='awrc'>541</td><td align="right" class='awrc'>0</td><td align="right" class='awrc'>0</td><td align="right" class='awrc'>0.53</td><td align="right" class='awrc'>0.00</td><td align="right" class='awrc'>0.02</td></tr>    <tr><td scope="row" class='awrnc'>CGS wait for IPC msg</td><td align="right" class='awrnc'>16,988</td><td align="right" class='awrnc'>100</td><td align="right" class='awrnc'>0</td><td align="right" class='awrnc'>0.01</td><td align="right" class='awrnc'>0.02</td><td align="right" class='awrnc'>0.02</td></tr>    <tr><td scope="row" class='awrc'>db file async I/O submit</td><td align="right" class='awrc'>528,782</td><td align="right" class='awrc'>0</td><td align="right" class='awrc'>0</td><td align="right" class='awrc'>0.00</td><td align="right" class='awrc'>0.51</td><td align="right" class='awrc'>0.02</td></tr>    <tr><td scope="row" class='awrnc'>log file sync</td><td align="right" class='awrnc'>8</td><td align="right" class='awrnc'>0</td><td align="right" class='awrnc'>0</td><td align="right" class='awrnc'>18.77</td><td align="right" class='awrnc'>0.00</td><td align="right" class='awrnc'>0.01</td></tr>    <tr><td scope="row" class='awrc'>latch free</td><td align="right" class='awrc'>207</td><td align="right" class='awrc'>0</td><td align="right" class='awrc'>0</td><td align="right" class='awrc'>0.62</td><td align="right" class='awrc'>0.00</td><td align="right" class='awrc'>0.01</td></tr>    <tr><td scope="row" class='awrnc'>log file sequential read</td><td align="right" class='awrnc'>24</td><td align="right" class='awrnc'>0</td><td align="right" class='awrnc'>0</td><td align="right" class='awrnc'>5.17</td><td align="right" class='awrnc'>0.00</td><td align="right" class='awrnc'>0.01</td></tr>    <tr><td scope="row" class='awrc'>cell statistics gather</td><td align="right" class='awrc'>288</td><td align="right" class='awrc'>0</td><td align="right" class='awrc'>0</td><td align="right" class='awrc'>0.41</td><td align="right" class='awrc'>0.00</td><td align="right" class='awrc'>0.01</td></tr>    <tr><td scope="row" class='awrnc'>LGWR wait for redo copy</td><td align="right" class='awrnc'>1,941</td><td align="right" class='awrnc'>0</td><td align="right" class='awrnc'>0</td><td align="right" class='awrnc'>0.04</td><td align="right" class='awrnc'>0.00</td><td align="right" class='awrnc'>0.01</td></tr>    <tr><td scope="row" class='awrc'>undo segment extension</td><td align="right" class='awrc'>7</td><td align="right" class='awrc'>86</td><td align="right" class='awrc'>0</td><td align="right" class='awrc'>5.74</td><td align="right" class='awrc'>0.00</td><td align="right" class='awrc'>0.00</td></tr>    <tr><td scope="row" class='awrnc'>CSS operation: data update</td><td align="right" class='awrnc'>54</td><td align="right" class='awrnc'>0</td><td align="right" class='awrnc'>0</td><td align="right" class='awrnc'>0.70</td><td align="right" class='awrnc'>0.00</td><td align="right" class='awrnc'>0.00</td></tr>    <tr><td scope="row" class='awrc'>CSS operation: data query</td><td align="right" class='awrc'>54</td><td align="right" class='awrc'>0</td><td align="right" class='awrc'>0</td><td align="right" class='awrc'>0.46</td><td align="right" class='awrc'>0.00</td><td align="right" class='awrc'>0.00</td></tr>    <tr><td scope="row" class='awrnc'>CSS operation: action</td><td align="right" class='awrnc'>70</td><td align="right" class='awrnc'>0</td><td align="right" class='awrnc'>0</td><td align="right" class='awrnc'>0.34</td><td align="right" class='awrnc'>0.00</td><td align="right" class='awrnc'>0.00</td></tr>    <tr><td scope="row" class='awrc'>CSS initialization</td><td align="right" class='awrc'>6</td><td align="right" class='awrc'>0</td><td align="right" class='awrc'>0</td><td align="right" class='awrc'>3.05</td><td align="right" class='awrc'>0.00</td><td align="right" class='awrc'>0.00</td></tr>    <tr><td scope="row" class='awrnc'>log file single write</td><td align="right" class='awrnc'>24</td><td align="right" class='awrnc'>0</td><td align="right" class='awrnc'>0</td><td align="right" class='awrnc'>0.28</td><td align="right" class='awrnc'>0.00</td><td align="right" class='awrnc'>0.00</td></tr>    <tr><td scope="row" class='awrc'>buffer busy waits</td><td align="right" class='awrc'>251</td><td align="right" class='awrc'>0</td><td align="right" class='awrc'>0</td><td align="right" class='awrc'>0.03</td><td align="right" class='awrc'>0.00</td><td align="right" class='awrc'>0.00</td></tr>    <tr><td scope="row" class='awrnc'>KSV master wait</td><td align="right" class='awrnc'>107</td><td align="right" class='awrnc'>0</td><td align="right" class='awrnc'>0</td><td align="right" class='awrnc'>0.04</td><td align="right" class='awrnc'>0.00</td><td align="right" class='awrnc'>0.00</td></tr>    <tr><td scope="row" class='awrc'>direct path write</td><td align="right" class='awrc'>11</td><td align="right" class='awrc'>0</td><td align="right" class='awrc'>0</td><td align="right" class='awrc'>0.39</td><td align="right" class='awrc'>0.00</td><td align="right" class='awrc'>0.00</td></tr>    <tr><td scope="row" class='awrnc'>db file single write</td><td align="right" class='awrnc'>12</td><td align="right" class='awrnc'>0</td><td align="right" class='awrnc'>0</td><td align="right" class='awrnc'>0.35</td><td align="right" class='awrnc'>0.00</td><td align="right" class='awrnc'>0.00</td></tr>    <tr><td scope="row" class='awrc'>cell multiblock physical read</td><td align="right" class='awrc'>1</td><td align="right" class='awrc'>0</td><td align="right" class='awrc'>0</td><td align="right" class='awrc'>4.23</td><td align="right" class='awrc'>0.00</td><td align="right" class='awrc'>0.00</td></tr>    <tr><td scope="row" class='awrnc'>direct path read</td><td align="right" class='awrnc'>1</td><td align="right" class='awrnc'>0</td><td align="right" class='awrnc'>0</td><td align="right" class='awrnc'>3.65</td><td align="right" class='awrnc'>0.00</td><td align="right" class='awrnc'>0.00</td></tr>    <tr><td scope="row" class='awrc'>CSS operation: query</td><td align="right" class='awrc'>18</td><td align="right" class='awrc'>0</td><td align="right" class='awrc'>0</td><td align="right" class='awrc'>0.16</td><td align="right" class='awrc'>0.00</td><td align="right" class='awrc'>0.00</td></tr>    <tr><td scope="row" class='awrnc'>Disk file operations I/O</td><td align="right" class='awrnc'>525</td><td align="right" class='awrnc'>0</td><td align="right" class='awrnc'>0</td><td align="right" class='awrnc'>0.00</td><td align="right" class='awrnc'>0.00</td><td align="right" class='awrnc'>0.00</td></tr>    <tr><td scope="row" class='awrc'>rdbms ipc message</td><td align="right" class='awrc'>1,463,811</td><td align="right" class='awrc'>2</td><td align="right" class='awrc'>37,964</td><td align="right" class='awrc'>25.94</td><td align="right" class='awrc'>1.42</td><td align="right" class='awrc'>&#160;</td></tr>    <tr><td scope="row" class='awrnc'>Space Manager: slave idle wait</td><td align="right" class='awrnc'>2,650</td><td align="right" class='awrnc'>91</td><td align="right" class='awrnc'>12,569</td><td align="right" class='awrnc'>4742.92</td><td align="right" class='awrnc'>0.00</td><td align="right" class='awrnc'>&#160;</td></tr>    <tr><td scope="row" class='awrc'>EMON slave idle wait</td><td align="right" class='awrc'>1,745</td><td align="right" class='awrc'>100</td><td align="right" class='awrc'>8,725</td><td align="right" class='awrc'>5000.04</td><td align="right" class='awrc'>0.00</td><td align="right" class='awrc'>&#160;</td></tr>    <tr><td scope="row" class='awrnc'>GCR sleep</td><td align="right" class='awrnc'>12,273</td><td align="right" class='awrnc'>7</td><td align="right" class='awrnc'>3,455</td><td align="right" class='awrnc'>281.48</td><td align="right" class='awrnc'>0.01</td><td align="right" class='awrnc'>&#160;</td></tr>    <tr><td scope="row" class='awrc'>DIAG idle wait</td><td align="right" class='awrc'>18,337</td><td align="right" class='awrc'>100</td><td align="right" class='awrc'>3,422</td><td align="right" class='awrc'>186.60</td><td align="right" class='awrc'>0.02</td><td align="right" class='awrc'>&#160;</td></tr>    <tr><td scope="row" class='awrnc'>LGWR worker group idle</td><td align="right" class='awrnc'>1,085,206</td><td align="right" class='awrnc'>0</td><td align="right" class='awrnc'>2,881</td><td align="right" class='awrnc'>2.65</td><td align="right" class='awrnc'>1.05</td><td align="right" class='awrnc'>&#160;</td></tr>    <tr><td scope="row" class='awrc'>AQPC idle</td><td align="right" class='awrc'>59</td><td align="right" class='awrc'>100</td><td align="right" class='awrc'>1,770</td><td align="right" class='awrc'>30000.36</td><td align="right" class='awrc'>0.00</td><td align="right" class='awrc'>&#160;</td></tr>    <tr><td scope="row" class='awrnc'>heartbeat redo informer</td><td align="right" class='awrnc'>1,746</td><td align="right" class='awrnc'>0</td><td align="right" class='awrnc'>1,746</td><td align="right" class='awrnc'>1000.03</td><td align="right" class='awrnc'>0.00</td><td align="right" class='awrnc'>&#160;</td></tr>    <tr><td scope="row" class='awrc'>pmon timer</td><td align="right" class='awrc'>582</td><td align="right" class='awrc'>100</td><td align="right" class='awrc'>1,746</td><td align="right" class='awrc'>3000.04</td><td align="right" class='awrc'>0.00</td><td align="right" class='awrc'>&#160;</td></tr>    <tr><td scope="row" class='awrnc'>ges remote message</td><td align="right" class='awrnc'>1,344</td><td align="right" class='awrnc'>100</td><td align="right" class='awrnc'>1,746</td><td align="right" class='awrnc'>1298.94</td><td align="right" class='awrnc'>0.00</td><td align="right" class='awrnc'>&#160;</td></tr>    <tr><td scope="row" class='awrc'>class slave wait</td><td align="right" class='awrc'>429</td><td align="right" class='awrc'>0</td><td align="right" class='awrc'>1,745</td><td align="right" class='awrc'>4068.00</td><td align="right" class='awrc'>0.00</td><td align="right" class='awrc'>&#160;</td></tr>    <tr><td scope="row" class='awrnc'>wait for unread message on broadcast channel</td><td align="right" class='awrnc'>581</td><td align="right" class='awrnc'>100</td><td align="right" class='awrnc'>1,745</td><td align="right" class='awrnc'>3003.58</td><td align="right" class='awrnc'>0.00</td><td align="right" class='awrnc'>&#160;</td></tr>    <tr><td scope="row" class='awrc'>PING</td><td align="right" class='awrc'>403</td><td align="right" class='awrc'>33</td><td align="right" class='awrc'>1,745</td><td align="right" class='awrc'>4329.67</td><td align="right" class='awrc'>0.00</td><td align="right" class='awrc'>&#160;</td></tr>    <tr><td scope="row" class='awrnc'>ASM background timer</td><td align="right" class='awrnc'>350</td><td align="right" class='awrnc'>0</td><td align="right" class='awrnc'>1,744</td><td align="right" class='awrnc'>4983.02</td><td align="right" class='awrnc'>0.00</td><td align="right" class='awrnc'>&#160;</td></tr>    <tr><td scope="row" class='awrc'>lreg timer</td><td align="right" class='awrc'>581</td><td align="right" class='awrc'>100</td><td align="right" class='awrc'>1,743</td><td align="right" class='awrc'>3000.48</td><td align="right" class='awrc'>0.00</td><td align="right" class='awrc'>&#160;</td></tr>    <tr><td scope="row" class='awrnc'>smon timer</td><td align="right" class='awrnc'>656</td><td align="right" class='awrnc'>0</td><td align="right" class='awrnc'>1,740</td><td align="right" class='awrnc'>2652.90</td><td align="right" class='awrnc'>0.00</td><td align="right" class='awrnc'>&#160;</td></tr>    <tr><td scope="row" class='awrc'>Streams AQ: emn coordinator idle wait</td><td align="right" class='awrc'>174</td><td align="right" class='awrc'>100</td><td align="right" class='awrc'>1,740</td><td align="right" class='awrc'>10000.11</td><td align="right" class='awrc'>0.00</td><td align="right" class='awrc'>&#160;</td></tr>    <tr><td scope="row" class='awrnc'>Streams AQ: qmn slave idle wait</td><td align="right" class='awrnc'>62</td><td align="right" class='awrnc'>0</td><td align="right" class='awrnc'>1,736</td><td align="right" class='awrnc'>28000.43</td><td align="right" class='awrnc'>0.00</td><td align="right" class='awrnc'>&#160;</td></tr>    <tr><td scope="row" class='awrc'>Streams AQ: qmn coordinator idle wait</td><td align="right" class='awrc'>124</td><td align="right" class='awrc'>50</td><td align="right" class='awrc'>1,736</td><td align="right" class='awrc'>14000.21</td><td align="right" class='awrc'>0.00</td><td align="right" class='awrc'>&#160;</td></tr>    <tr><td scope="row" class='awrnc'>SQL*Net message from client</td><td align="right" class='awrnc'>3,400</td><td align="right" class='awrnc'>0</td><td align="right" class='awrnc'>4</td><td align="right" class='awrnc'>1.18</td><td align="right" class='awrnc'>0.00</td><td align="right" class='awrnc'>&#160;</td></tr></table><p /><hr align="left" width="20%" /><p /><a class="awr" href="#21">Back to Wait Events Statistics</a><br /><a class="awr" href="#top">Back to Top</a><p /><a class="exa" href="#CELL_TOPDB">Back to Exadata Top Database Consumers</a><br/><a class="exa" href="#CELL_STATISTICS">Back to Exadata Statistics</a></div><br /><a class="awr" href="#top">Back to Top</a><p /><p />End of Report</body></html>`, nil},
	//	{"fdsfsdf", "", errors.New(` open fdsfsdf: The system cannot find the file specified. `)},
	}


	for _, pair := range tests {
		v, err := readFile(pair.input)
		if (v != pair.text) || (err != pair.error)  {
			t.Error(
				"For", pair.input,
				"\n expected:", pair.text,
				"\n got text:", v,
				"\n got err:", err,
			)
		}
	}

}
func TestCreateMaps(t *testing.T) {

	mapsGood := make(map[string]string)
	mapsGood["Report Summary"] = `<p />Top ADDM Findings by Average Active Sessions<p /><ul></ul><table border="0" width="600" class="tdiff" summary="This table displays top ADDM findings by average active sessions"><tr><th class="awrbg" scope="col">Finding Name</th><th class="awrbg" scope="col">Avg active sessions of the task</th><th class="awrbg" scope="col">Percent active sessions of finding</th><th class="awrbg" scope="col">Task Name</th><th class="awrbg" scope="col">Begin Snap Time</th><th class="awrbg" scope="col">End Snap Time</th></tr>    <tr><td class='awrc'>Фиксации и откаты</td><td align="right" class='awrc'>8.08</td><td align="right" class='awrc'>50.71</td><td scope="row" class='awrc'>ADDM:2926569_1_111758</td><td class='awrc'>10-Дек-17 13:50</td><td class='awrc'>10-Дек-17 14:00</td></tr>    <tr><td class='awrnc'>Фиксации и откаты</td><td align="right" class='awrnc'>7.07</td><td align="right" class='awrnc'>57.14</td><td scope="row" class='awrnc'>ADDM:2926569_1_111758</td><td class='awrnc'>10-Дек-17 13:40</td><td class='awrnc'>10-Дек-17 13:50</td></tr>    <tr><td class='awrc'>Фиксации и откаты</td><td align="right" class='awrc'>6.09</td><td align="right" class='awrc'>48.90</td><td scope="row" class='awrc'>ADDM:2926569_1_111758</td><td class='awrc'>10-Дек-17 13:30</td><td class='awrc'>10-Дек-17 13:40</td></tr>    <tr><td class='awrnc'>Наиболее часто используемые операторы SQL</td><td align="right" class='awrnc'>8.08</td><td align="right" class='awrnc'>31.41</td><td scope="row" class='awrnc'>ADDM:2966226569_1_111758</td><td class='awrnc'>10-Дек-17 13:50</td><td class='awrnc'>10-Дек-17 14:00</td></tr>    <tr><td class='awrc'>Наиболее часто используемые операторы SQL</td><td align="right" class='awrc'>7.07</td><td align="right" class='awrc'>25.59</td><td scope="row" class='awrc'>ADDM:2966226569_1_111757</td><td class='awrc'>10-Дек-17 13:40</td><td class='awrc'>10-Дек-17 13:50</td></tr></table><p /><p />Load Profile<p /><table border="0" width="600" class="tdiff" summary="This table displays load profile">    <tr><th class="awrnobg" scope="col"></th><th class="awrbg" scope="col">Per Second</th><th class="awrbg" scope="col">Per Transaction</th><th class="awrbg" scope="col">Per Exec</th><th class="awrbg" scope="col">Per Call</th></tr>    <tr><td scope="row" class='awrc'>DB Time(s):</td><td align="right" class='awrc'>               7.1</td><td align="right" class='awrc'>               0.0</td><td align="right" class='awrc'>      0.00</td><td align="right" class='awrc'>      0.00</td></tr>    <tr><td scope="row" class='awrnc'>DB CPU(s):</td><td align="right" class='awrnc'>               2.8</td><td align="right" class='awrnc'>               0.0</td><td align="right" class='awrnc'>      0.00</td><td align="right" class='awrnc'>      0.00</td></tr>    <tr><td scope="row" class='awrc'>Background CPU(s):</td><td align="right" class='awrc'>               0.4</td><td align="right" class='awrc'>               0.0</td><td align="right" class='awrc'>      0.00</td><td align="right" class='awrc'>      0.00</td></tr>    <tr><td scope="row" class='awrnc'>Redo size (bytes):</td><td align="right" class='awrnc'>       5,048,931.5</td><td align="right" class='awrnc'>           8,538.4</td><td class='awrnc'>&#160;</td><td class='awrnc'>&#160;</td></tr>    <tr><td scope="row" class='awrc'>L</tr></table></p>`
	v, err := readFile("test/test.html")
	if err != nil{
		t.Error(err)
	}

	maps := make(map[string]string)
	err = createMaps(v, maps)
	if err != nil{
		t.Error(err)
	}

	if mapsGood["Report Summary"] != maps["Report Summary"] && len(maps) != 0  {
		t.Error(
			"For Report Summary", mapsGood["Report Summary"] ,
				"Expected", maps["Report Summary"] ,
		)
	}

}
func TestFixDot(t *testing.T) {
	type testPair struct {
		input string
		output float64
	}

	var tests = []testPair{
		{"100.00", 100.00},
		{"1,041,385", 1041385},
		{"1,041.7385", 1041.7385},
		{"78", 78},
		{"asdasd", 0},
	}

	for _, pair := range tests {
		v := fixDot(pair.input)
		if (v != pair.output)  {
			t.Error(
				"For", pair.input,
				"\n expected:", pair.output,
				"\n got text:", v,
			)
		}
	}


}