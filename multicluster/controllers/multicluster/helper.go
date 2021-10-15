/*
Copyright 2021 Antrea Authors.
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

package multicluster

import "k8s.io/apimachinery/pkg/types"

const (
	antreaMcsLabel    = "antrea.io/multi-cluster"
	ServiceKind       = "Service"
	EndpointsKind     = "Endpoints"
	ServiceImportKind = "ServiceImport"
	SourceServiceType = "SourceServiceType"
)

const (
	Separator = '/'
)

// TODO: Use NamespacedName stringer method instead of this. eg nsName.String()
func NamespacedName(namespace, name string) string {
	return namespace + string(Separator) + name
}

// GetNamespacedName returns the objects name as NamespacedName struct.
func GetNamespacedName(namespace, name string) types.NamespacedName {
	return types.NamespacedName{
		Namespace: namespace,
		Name:      name,
	}
}
