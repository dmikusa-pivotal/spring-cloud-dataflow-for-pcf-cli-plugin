/*
 * Copyright (C) 2017-Present Pivotal Software, Inc. All rights reserved.
 *
 * This program and the accompanying materials are made available under
 * the terms of the under the Apache License, Version 2.0 (the "License”);
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package dataflow_test

import (
	. "github.com/pivotal-cf/spring-cloud-dataflow-for-pcf-cli-plugin/dataflow"

	"os/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("DataflowShellCommand", func() {
	const (
		fileName = "filename"
		url      = "https://some.host/path"
	)

	var (
		skipSslValidation bool
		cmd               *exec.Cmd
	)

	JustBeforeEach(func() {
		cmd = DataflowShellCommand(fileName, url, skipSslValidation)
	})

	Context("when SSL validation is to be performed", func() {
		BeforeEach(func() {
			skipSslValidation = false
		})

		It("should produce the correct command", func() {
			Expect(cmd.Args).To(Equal([]string{"java", "-jar", fileName, "--dataflow.uri=" + url, "--dataflow.credentials-provider-command=cf oauth-token", "--dataflow.mode=skipper"}))
		})
	})

	Context("when SSL validation is to be skipped", func() {
		BeforeEach(func() {
			skipSslValidation = true
		})

		It("should produce the correct command", func() {
			Expect(cmd.Args).To(Equal([]string{"java", "-jar", fileName, "--dataflow.uri=" + url, "--dataflow.credentials-provider-command=cf oauth-token", "--dataflow.mode=skipper", "--dataflow.skip-ssl-validation=true"}))
		})
	})

})
