/*
	Copyright (c) 2016, Percona LLC and/or its affiliates. All rights reserved.

	This program is free software: you can redistribute it and/or modify
	it under the terms of the GNU Affero General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.

	This program is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU Affero General Public License for more details.

	You should have received a copy of the GNU Affero General Public License
	along with this program.  If not, see <http://www.gnu.org/licenses/>
*/

package pmm

import (
	"fmt"
	"time"
)

var VERSION = "1.0.6"

const (
	PMMBaseDir     = "/usr/local/percona/pmm-client"
	agentBaseDir   = "/usr/local/percona/qan-agent" // This is also hardcoded in mysql_queries.go
	qanAPIBasePath = "qan-api"
	emojiUnhappy   = "😡"
	emojiHappy     = "🙂"
	noMonitoring   = "No monitoring registered for this node identified as"
	apiTimeout     = 30 * time.Second
	NameRegex      = `^[-\w:\.]{2,60}$`
)

var (
	ConfigFile = fmt.Sprintf("%s/pmm.yml", PMMBaseDir)

	ErrDuplicate = fmt.Errorf("there is already one instance with this name under monitoring.")
	ErrNoService = fmt.Errorf("no service found.")
	ErrOneLinux  = fmt.Errorf("there could be only one instance of linux metrics being monitored for this system.")

	errNoInstance = fmt.Errorf("no instance found on QAN API.")
)

const nodeExporterArgs = "-collectors.enabled=diskstats,filefd,filesystem,loadavg,meminfo,netdev,netstat,stat,time,uname,vmstat"

var mysqldExporterArgs = []string{
	"-collect.auto_increment.columns=true",
	"-collect.binlog_size=true",
	"-collect.global_status=true",
	"-collect.global_variables=true",
	"-collect.info_schema.innodb_metrics=true",
	"-collect.info_schema.processlist=true",
	"-collect.info_schema.query_response_time=true",
	"-collect.info_schema.tables=true",
	"-collect.info_schema.tablestats=true",
	"-collect.info_schema.userstats=true",
	"-collect.perf_schema.eventswaits=true",
	"-collect.perf_schema.file_events=true",
	"-collect.perf_schema.indexiowaits=true",
	"-collect.perf_schema.tableiowaits=true",
	"-collect.perf_schema.tablelocks=true",
	"-collect.slave_status=true",
	//"-collect.engine_innodb_status=true",
	//"-collect.engine_tokudb_status=true",
	//"-collect.info_schema.clientstats=true",
	//"-collect.info_schema.innodb_tablespaces=true",
	//"-collect.perf_schema.eventsstatements=true",
}

// mysqld_exporter args to disable optionally.
var mysqldExporterDisableArgs = map[string][]string{
	"tablestats": {
		"-collect.auto_increment.columns=",
		"-collect.info_schema.tables=",
		"-collect.info_schema.tablestats=",
		"-collect.perf_schema.indexiowaits=",
		"-collect.perf_schema.tableiowaits=",
		"-collect.perf_schema.tablelocks=",
	},
	"userstats":   {"-collect.info_schema.userstats="},
	"binlogstats": {"-collect.binlog_size="},
	"processlist": {"-collect.info_schema.processlist="},
}
