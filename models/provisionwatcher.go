/*******************************************************************************
 * Copyright 2019 Dell Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License
 * is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing permissions and limitations under
 * the License.
 *******************************************************************************/

package models

import (
	"encoding/json"
)

type ProvisionWatcher struct {
	Timestamps
	Id             string            `json:"id"`
	Name           string            `json:"name"`           // unique name and identifier of the addressable
	Identifiers    map[string]string `json:"identifiers"`    // set of key value pairs that identify type of of address (MAC, HTTP,...) and address to watch for (00-05-1B-A1-99-99, 10.0.0.1,...)
	Profile        DeviceProfile     `json:"profile"`        // device profile that should be applied to the devices available at the identifier addresses
	Service        DeviceService     `json:"service"`        // device service that owns the watcher
	OperatingState OperatingState    `json:"operatingState"` // operational state - either enabled or disabled
}

// Custom marshaling to make empty strings null
func (pw ProvisionWatcher) MarshalJSON() ([]byte, error) {
	test := struct {
		Timestamps
		Id             string            `json:"id"`
		Name           *string           `json:"name"`           // unique name and identifier of the addressable
		Identifiers    map[string]string `json:"identifiers"`    // set of key value pairs that identify type of of address (MAC, HTTP,...) and address to watch for (00-05-1B-A1-99-99, 10.0.0.1,...)
		Profile        DeviceProfile     `json:"profile"`        // device profile that should be applied to the devices available at the identifier addresses
		Service        DeviceService     `json:"service"`        // device service that owns the watcher
		OperatingState OperatingState    `json:"operatingState"` // operational state - either enabled or disabled
	}{
		Id:             pw.Id,
		Timestamps:     pw.Timestamps,
		Profile:        pw.Profile,
		Service:        pw.Service,
		OperatingState: pw.OperatingState,
	}

	// Empty strings are null
	if pw.Name != "" {
		test.Name = &pw.Name
	}

	// Empty maps are null
	if len(pw.Identifiers) > 0 {
		test.Identifiers = pw.Identifiers
	}

	return json.Marshal(test)
}

/*
 * To String function for ProvisionWatcher
 */
func (pw ProvisionWatcher) String() string {
	out, err := json.Marshal(pw)
	if err != nil {
		return err.Error()
	}
	return string(out)
}
