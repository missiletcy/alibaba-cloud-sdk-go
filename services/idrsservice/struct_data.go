package idrsservice

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

// Data is a nested struct in idrsservice response
type Data struct {
	CreatorName               string            `json:"CreatorName" xml:"CreatorName"`
	PhoneNumber               string            `json:"PhoneNumber" xml:"PhoneNumber"`
	CreateSuccess             bool              `json:"CreateSuccess" xml:"CreateSuccess"`
	RuleId                    string            `json:"RuleId" xml:"RuleId"`
	LiveRecordMaxClient       int               `json:"LiveRecordMaxClient" xml:"LiveRecordMaxClient"`
	Source                    string            `json:"Source" xml:"Source"`
	TotalElements             int64             `json:"TotalElements" xml:"TotalElements"`
	Content                   string            `json:"Content" xml:"Content"`
	Role                      string            `json:"Role" xml:"Role"`
	MqEndpoint                string            `json:"MqEndpoint" xml:"MqEndpoint"`
	Id                        string            `json:"Id" xml:"Id"`
	DepartmentName            string            `json:"DepartmentName" xml:"DepartmentName"`
	Email                     string            `json:"Email" xml:"Email"`
	MqInstanceId              string            `json:"MqInstanceId" xml:"MqInstanceId"`
	LiveRecordAll             bool              `json:"LiveRecordAll" xml:"LiveRecordAll"`
	CompletedTasks            int               `json:"CompletedTasks" xml:"CompletedTasks"`
	LiveRecordLayout          int               `json:"LiveRecordLayout" xml:"LiveRecordLayout"`
	DepartmentId              string            `json:"DepartmentId" xml:"DepartmentId"`
	MqSubscribe               bool              `json:"MqSubscribe" xml:"MqSubscribe"`
	Status                    string            `json:"Status" xml:"Status"`
	Name                      string            `json:"Name" xml:"Name"`
	LiveRecordTaskProfile     string            `json:"LiveRecordTaskProfile" xml:"LiveRecordTaskProfile"`
	RuleName                  string            `json:"RuleName" xml:"RuleName"`
	TotalTasks                int               `json:"TotalTasks" xml:"TotalTasks"`
	UpdatedAt                 string            `json:"UpdatedAt" xml:"UpdatedAt"`
	RecordingType             string            `json:"RecordingType" xml:"RecordingType"`
	ErrorMessage              string            `json:"ErrorMessage" xml:"ErrorMessage"`
	TotalPages                int               `json:"TotalPages" xml:"TotalPages"`
	LiveRecordVideoResolution int               `json:"LiveRecordVideoResolution" xml:"LiveRecordVideoResolution"`
	ClientQueueSize           int               `json:"ClientQueueSize" xml:"ClientQueueSize"`
	Disabled                  bool              `json:"Disabled" xml:"Disabled"`
	MqTopic                   string            `json:"MqTopic" xml:"MqTopic"`
	CreatedAt                 string            `json:"CreatedAt" xml:"CreatedAt"`
	Description               string            `json:"Description" xml:"Description"`
	LiveRecordEveryOne        bool              `json:"LiveRecordEveryOne" xml:"LiveRecordEveryOne"`
	Username                  string            `json:"Username" xml:"Username"`
	HasRole                   bool              `json:"HasRole" xml:"HasRole"`
	TaskItemQueueSize         int               `json:"TaskItemQueueSize" xml:"TaskItemQueueSize"`
	Channel                   string            `json:"Channel" xml:"Channel"`
	MqGroupId                 string            `json:"MqGroupId" xml:"MqGroupId"`
	Code                      int               `json:"Code" xml:"Code"`
	VideoUrl                  string            `json:"VideoUrl" xml:"VideoUrl"`
	MqEventList               []string          `json:"MqEventList" xml:"MqEventList"`
	TaskIds                   []string          `json:"TaskIds" xml:"TaskIds"`
	TokenData                 TokenData         `json:"TokenData" xml:"TokenData"`
	Departments               []DepartmentsItem `json:"Departments" xml:"Departments"`
	Items                     []ItemsItem       `json:"Items" xml:"Items"`
	Tasks                     []TasksItem       `json:"Tasks" xml:"Tasks"`
}