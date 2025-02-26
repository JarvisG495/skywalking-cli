// Licensed to Apache Software Foundation (ASF) under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Apache Software Foundation (ASF) licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package metrics

import (
	"github.com/urfave/cli/v2"

	"github.com/apache/skywalking-cli/internal/commands/metrics/aggregation"
	"github.com/apache/skywalking-cli/internal/commands/metrics/list"

	"github.com/apache/skywalking-cli/internal/commands/metrics/thermodynamic"

	"github.com/apache/skywalking-cli/internal/commands/metrics/linear"
	"github.com/apache/skywalking-cli/internal/commands/metrics/single"
)

var Command = &cli.Command{
	Name:  "metrics",
	Usage: "Query metrics defined in backend OAL",
	Subcommands: cli.Commands{
		single.Command,
		linear.Single,
		linear.Multiple,
		thermodynamic.Command,
		aggregation.TopN,
		list.Command,
	},
}
