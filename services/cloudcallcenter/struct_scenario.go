package cloudcallcenter

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// Scenario is a nested struct in cloudcallcenter response
type Scenario struct {
	Name                string         `json:"Name" xml:"Name"`
	ScenarioId          string         `json:"ScenarioId" xml:"ScenarioId"`
	ScenarioDescription string         `json:"ScenarioDescription" xml:"ScenarioDescription"`
	BeebotVersion       string         `json:"BeebotVersion" xml:"BeebotVersion"`
	IsTemplate          bool           `json:"IsTemplate" xml:"IsTemplate"`
	Id                  string         `json:"Id" xml:"Id"`
	Description         string         `json:"Description" xml:"Description"`
	ScenarioName        string         `json:"ScenarioName" xml:"ScenarioName"`
	Type                string         `json:"Type" xml:"Type"`
	Strategy            Strategy       `json:"Strategy" xml:"Strategy"`
	Surveys             []Survey       `json:"Surveys" xml:"Surveys"`
	Variables           []KeyValuePair `json:"Variables" xml:"Variables"`
}
