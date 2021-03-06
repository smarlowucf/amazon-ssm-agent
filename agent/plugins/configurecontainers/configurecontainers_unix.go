// Copyright 2016 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may not
// use this file except in compliance with the License. A copy of the
// License is located at
//
// http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
// either express or implied. See the License for the specific language governing
// permissions and limitations under the License.
//
// +build darwin freebsd linux netbsd openbsd

package configurecontainers

import (
	"errors"

	"github.com/aws/amazon-ssm-agent/agent/contracts"
	"github.com/aws/amazon-ssm-agent/agent/log"
)

func runInstallCommands(log log.T, pluginInput ConfigureContainerPluginInput, orchestrationDirectory string) (out contracts.PluginOutput) {
	//TODO Unix install

	out.MarkAsFailed(log, errors.New("Configure Docker is not currently supported on unix"))

	return out
}

func runUninstallCommands(log log.T, pluginInput ConfigureContainerPluginInput, orchestrationDirectory string) (out contracts.PluginOutput) {
	//TODO Unix uninstall

	out.MarkAsFailed(log, errors.New("Configure Docker is not currently supported on unix"))

	return out
}
