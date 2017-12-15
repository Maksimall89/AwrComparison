package main
import (
	"testing"
)

// TODO coverage = 100%

func TestParser(t *testing.T) {
	t.Parallel()

	type testPair struct {
		input string
		text string
		error error
	}

	var tests = []testPair{
		{"test/test.html", `<meta http-equiv="content-type" content="text/html; charset=UTF-8"><html lang="en"><head><title>AWR Report for DB: ZP, Inst: k10szp_1, Snaps: 111755-111758</title>
    <style type="text/css">
        body.awr {font:bold 10pt Arial,Helvetica,Geneva,sans-serif;color:black; background:White;}
        pre.awr  {font:8pt Courier;color:black; background:White;}
            .hidden   {position:absolute;left:-10000px;top:auto;width:1px;height:1px;overflow:hidden;}
        .pad   {margin-left:17px;}
        .doublepad {margin-left:34px;}
    </style></head><body class="awr">
<h1 class="awr">
    WORKLOAD REPOSITORY report for

</h1>
<p />
<table border="0" width="600" class="tdiff" summary="This table displays database instance information">
    <tr><th class="awrbg" scope="col">DB Name</th><th class="awrbg" scope="col">DB Id</th><th class="awrbg" scope="col">Instance</th><th class="awrbg" scope="col">Inst num</th><th class="awrbg" scope="col">Startup Time</th><th class="awrbg" scope="col">Release</th><th class="awrbg" scope="col">RAC</th></tr>
    <tr><td scope="row" class='awrnc'>ZP</td><td align="right" class='awrnc'>2966226569</td><td class='awrnc'>sdf_test1</td><td align="right" class='awrnc'>1</td><td class='awrnc'>08-Ноя-17 16:11</td><td class='awrnc'>12.1.0.2.0</td><td class='awrnc'>NO</td></tr>
</table>
<p />
<p />
<table border="0" width="600" class="tdiff" summary="This table displays host information">
    <tr><th class="awrbg" scope="col">Host Name</th><th class="awrbg" scope="col">Platform</th><th class="awrbg" scope="col">CPUs</th><th class="awrbg" scope="col">Cores</th><th class="awrbg" scope="col">Sockets</th><th class="awrbg" scope="col">Memory (GB)</th></tr>
    <tr><td scope="row" class='awrnc'>sdfeede</td><td class='awrnc'>Solaris[tm] OE (32-bit)</td><td align="right" class='awrnc'> 256</td><td align="right" class='awrnc'>  32</td><td align="right" class='awrnc'>   1</td><td align="right" class='awrnc'>  478.50</td></tr>
</table>
<p />
<table border="0" width="600" class="tdiff" summary="This table displays snapshot information">
    <tr><th class="awrnobg" scope="col"></th><th class="awrbg" scope="col">Snap Id</th><th class="awrbg" scope="col">Snap Time</th><th class="awrbg" scope="col">Sessions</th><th class="awrbg" scope="col">Cursors/Session</th></tr>
    <tr><td scope="row" class='awrnc'>Begin Snap:</td><td align="right" class='awrnc'>111755</td><td align="center" class='awrnc'>10-Дек-17 13:30:55</td><td align="right" class='awrnc'>951</td><td align="right" class='awrnc'>     12.0</td></tr>
    <tr><td scope="row" class='awrc'>End Snap:</td><td align="right" class='awrc'>111758</td><td align="center" class='awrc'>10-Дек-17 14:00:00</td><td align="right" class='awrc'>1130</td><td align="right" class='awrc'>     13.0</td></tr>
    <tr><td scope="row" class='awrnc'>Elapsed:</td><td class='awrnc'>&#160;</td><td align="center" class='awrnc'>              29.09 (mins)</td><td class='awrnc'>&#160;</td><td class='awrnc'>&#160;</td></tr>
    <tr><td scope="row" class='awrc'>DB Time:</td><td class='awrc'>&#160;</td><td align="center" class='awrc'>             204.95 (mins)</td><td class='awrc'>&#160;</td><td class='awrc'>&#160;</td></tr>
</table>
<p />
<h3 class="awr"><a class="awr" name="99999"></a>Report Summary</h3>
<p />Top ADDM Findings by Average Active Sessions<p />
<ul>
</ul>
<table border="0" width="600" class="tdiff" summary="This table displays top ADDM findings by average active sessions"><tr><th class="awrbg" scope="col">Finding Name</th><th class="awrbg" scope="col">Avg active sessions of the task</th><th class="awrbg" scope="col">Percent active sessions of finding</th><th class="awrbg" scope="col">Task Name</th><th class="awrbg" scope="col">Begin Snap Time</th><th class="awrbg" scope="col">End Snap Time</th></tr>
    <tr><td class='awrc'>Фиксации и откаты</td><td align="right" class='awrc'>8.08</td><td align="right" class='awrc'>50.71</td><td scope="row" class='awrc'>ADDM:2926569_1_111758</td><td class='awrc'>10-Дек-17 13:50</td><td class='awrc'>10-Дек-17 14:00</td></tr>
    <tr><td class='awrnc'>Фиксации и откаты</td><td align="right" class='awrnc'>7.07</td><td align="right" class='awrnc'>57.14</td><td scope="row" class='awrnc'>ADDM:2926569_1_111758</td><td class='awrnc'>10-Дек-17 13:40</td><td class='awrnc'>10-Дек-17 13:50</td></tr>
    <tr><td class='awrc'>Фиксации и откаты</td><td align="right" class='awrc'>6.09</td><td align="right" class='awrc'>48.90</td><td scope="row" class='awrc'>ADDM:2926569_1_111758</td><td class='awrc'>10-Дек-17 13:30</td><td class='awrc'>10-Дек-17 13:40</td></tr>
    <tr><td class='awrnc'>Наиболее часто используемые операторы SQL</td><td align="right" class='awrnc'>8.08</td><td align="right" class='awrnc'>31.41</td><td scope="row" class='awrnc'>ADDM:2966226569_1_111758</td><td class='awrnc'>10-Дек-17 13:50</td><td class='awrnc'>10-Дек-17 14:00</td></tr>
    <tr><td class='awrc'>Наиболее часто используемые операторы SQL</td><td align="right" class='awrc'>7.07</td><td align="right" class='awrc'>25.59</td><td scope="row" class='awrc'>ADDM:2966226569_1_111757</td><td class='awrc'>10-Дек-17 13:40</td><td class='awrc'>10-Дек-17 13:50</td></tr>
</table><p />
<p />Load Profile<p />
<table border="0" width="600" class="tdiff" summary="This table displays load profile">
    <tr><th class="awrnobg" scope="col"></th><th class="awrbg" scope="col">Per Second</th><th class="awrbg" scope="col">Per Transaction</th><th class="awrbg" scope="col">Per Exec</th><th class="awrbg" scope="col">Per Call</th></tr>
    <tr><td scope="row" class='awrc'>DB Time(s):</td><td align="right" class='awrc'>               7.1</td><td align="right" class='awrc'>               0.0</td><td align="right" class='awrc'>      0.00</td><td align="right" class='awrc'>      0.00</td></tr>
    <tr><td scope="row" class='awrnc'>DB CPU(s):</td><td align="right" class='awrnc'>               2.8</td><td align="right" class='awrnc'>               0.0</td><td align="right" class='awrnc'>      0.00</td><td align="right" class='awrnc'>      0.00</td></tr>
    <tr><td scope="row" class='awrc'>Background CPU(s):</td><td align="right" class='awrc'>               0.4</td><td align="right" class='awrc'>               0.0</td><td align="right" class='awrc'>      0.00</td><td align="right" class='awrc'>      0.00</td></tr>
    <tr><td scope="row" class='awrnc'>Redo size (bytes):</td><td align="right" class='awrnc'>       5,048,931.5</td><td align="right" class='awrnc'>           8,538.4</td><td class='awrnc'>&#160;</td><td class='awrnc'>&#160;</td></tr>
    <tr><td scope="row" class='awrc'>L</tr>
</table>
</p>
<a class="exa" href="#CELL_TOPDB">Back to Exadata Top Database Consumers</a>
<br/>
<a class="exa" href="#CELL_STATISTICS">Back to Exadata Statistics</a>
</div>
<br /><a class="awr" href="#top">Back to Top</a><p />
<p />
End of Report
</body></html>
`, nil},
	//	{"fdsfsdf", "", errors.New(` open fdsfsdf: The system cannot find the file specified. `)},
	}

	//readFile()

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
