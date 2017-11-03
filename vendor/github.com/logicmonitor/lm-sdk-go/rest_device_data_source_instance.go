/* 
 * 
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * OpenAPI spec version: 1.0.0
 * 
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package logicmonitor

type RestDeviceDataSourceInstance struct {

	AlertStatus string `json:"alertStatus,omitempty"`

	DisplayName string `json:"displayName"`

	GroupId int32 `json:"groupId,omitempty"`

	Description string `json:"description,omitempty"`

	DisableAlerting bool `json:"disableAlerting,omitempty"`

	DeviceId int32 `json:"deviceId,omitempty"`

	DataSourceId int32 `json:"dataSourceId,omitempty"`

	SdtAt string `json:"sdtAt,omitempty"`

	LockDescription bool `json:"lockDescription,omitempty"`

	LastUpdatedTime int64 `json:"lastUpdatedTime,omitempty"`

	Id int32 `json:"id,omitempty"`

	AlertStatusPriority int32 `json:"alertStatusPriority,omitempty"`

	AlertDisableStatus string `json:"alertDisableStatus,omitempty"`

	WildValue string `json:"wildValue"`

	StopMonitoring bool `json:"stopMonitoring,omitempty"`

	DeviceDataSourceId int32 `json:"deviceDataSourceId,omitempty"`

	WildValue2 string `json:"wildValue2,omitempty"`

	SdtStatus string `json:"sdtStatus,omitempty"`

	DeviceDisplayName string `json:"deviceDisplayName,omitempty"`

	GroupsDisabledThisSource []TreeNode `json:"groupsDisabledThisSource,omitempty"`

	LastCollectedTime int64 `json:"lastCollectedTime,omitempty"`

	GroupName string `json:"groupName,omitempty"`

	CustomProperties []NameAndValue `json:"customProperties,omitempty"`

	Name string `json:"name,omitempty"`

	AlertingDisabledOn TreeNode `json:"alertingDisabledOn,omitempty"`

	DataSourceType string `json:"dataSourceType,omitempty"`
}
