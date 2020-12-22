/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package version

import (
	"fmt"
	"runtime"
)

var (
	// Number the number of version that are running at the moment.
	Number string

	// GitCommit compiled from which commit.
	GitCommit string

	// BuildDate date of the application was build.
	BuildDate = ""

	// GoVersion version of go.
	GoVersion = runtime.Version()

	// OsArch os architecture that used to run this application.
	OsArch = fmt.Sprintf("%s %s", runtime.GOOS, runtime.GOARCH)
)
