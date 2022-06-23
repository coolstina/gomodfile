// Copyright 2022 helloshaohua <wu.shaohua@foxmail.com>;
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gomodfile

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/coolstina/expression"
	"github.com/coolstina/fsfire"
)

const GoModuleFilename = "go.mod"

func IsGoProject() (bool, error) {
	dir := fsfire.RuntimeMasterDir()

	gomod := filepath.Join(dir, GoModuleFilename)
	notExists := fsfire.IsNotExists(gomod)
	if notExists {
		return false, os.ErrNotExist
	}

	var buffer = &bytes.Buffer{}

	data, err := ioutil.ReadFile(gomod)
	if err != nil {
		return false, err
	}

	buffer.Write(data)
	slice, err := fsfire.GetFileContentStringSliceWithBuffer(buffer)
	if err != nil {
		return false, err
	}

	first := strings.TrimSpace(slice[0])
	matched, err := expression.RegularMatchString(`module .*`, first)
	if err != nil {
		return false, err
	}

	return matched, nil
}
