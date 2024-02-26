// Copyright 2022 Google LLC
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

package contexts

func init() {
	resourceContextMap["mysqlinstance"] = ResourceContext{
		// SQL instances names are reserved for 1 week after use: https://cloud.google.com/sql/docs/mysql/delete-instance
		SkipDriftDetection: true,
		ResourceKind:       "SQLInstance",
	}

	resourceContextMap["sqlserverinstance"] = ResourceContext{
		// SQL instances names are reserved for 1 week after use: https://cloud.google.com/sql/docs/mysql/delete-instance
		SkipDriftDetection: true,
		ResourceKind:       "SQLInstance",
	}

	resourceContextMap["sqlserverclone"] = ResourceContext{
		// SQL instances names are reserved for 1 week after use: https://cloud.google.com/sql/docs/mysql/delete-instance
		SkipDriftDetection: true,
		ResourceKind:       "SQLInstance",
		SkipUpdate:         true,
	}

	resourceContextMap["sqldatabase"] = ResourceContext{
		ResourceKind: "SQLDatabase",
	}

	resourceContextMap["sqlsslcert"] = ResourceContext{
		ResourceKind: "SQLSSLCert",
		SkipUpdate:   true,
	}

	resourceContextMap["sqluser"] = ResourceContext{
		ResourceKind: "SQLUser",
		// Skip update test, as the only field that is updatable is "password", which we
		// can't verify on the underlying API due to it not being passed back in a GET.
		SkipUpdate: true,
	}
}
