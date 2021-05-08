/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package rbac

import "strings"

//as a user of a backend service, he only understands resource of this service,
//to decouple authorization code from business code,
//a middleware should handle all the authorization logic, and this middleware only understand rest API,
//a resource mapping helps to maintain relations between api and resource.
var resourceMap = map[string]string{}

//PartialMap saves api partial matching
var PartialMap = map[string]string{}

func GetResource(api string) string {
	r, ok := resourceMap[api]
	if ok {
		return r
	}
	for partialAPI, resource := range PartialMap {
		if strings.Contains(api, partialAPI) {
			return resource
		}
	}
	return resourceMap["*"]
}

// MapResource saves the mapping from api to resource, it must be exactly match
func MapResource(api, resource string) {
	resourceMap[api] = resource
}

// PartialMapResource saves the mapping from api to resource, it is partial match
func PartialMapResource(api, resource string) {
	PartialMap[api] = resource
}
