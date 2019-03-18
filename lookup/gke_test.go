// Copyright 2018 ReactiveOps
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package lookup

import (
	"testing"

	"github.com/stretchr/testify/assert"

	clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

func TestGetClusterInfo(t *testing.T) {
	config1 := clientcmdapi.Config{
		Contexts: map[string]*clientcmdapi.Context{
			"not-gke": {
				Cluster: "helloworld",
			},
			"actual-gke": {
				Cluster: "gke_reactiveopsio_us-central1-a_rbac-lookup-testing",
			},
		},
		CurrentContext: "bar",
	}

	assert.Equal(t, getClusterInfo(&config1, "").ClusterName, "")
	assert.Equal(t, getClusterInfo(&config1, "not-gke").ClusterName, "")
	assert.Equal(t, getClusterInfo(&config1, "actual-gke").ClusterName, "rbac-lookup-testing")
}
