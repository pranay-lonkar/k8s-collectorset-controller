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

type RestDevice struct {

	AlertStatus string `json:"alertStatus,omitempty"`

	NetflowCollectorGroupName string `json:"netflowCollectorGroupName,omitempty"`

	AzureState int32 `json:"azureState,omitempty"`

	RelatedDeviceId int32 `json:"relatedDeviceId,omitempty"`

	DisplayName string `json:"displayName"`

	Link string `json:"link,omitempty"`

	AwsState int32 `json:"awsState,omitempty"`

	Description string `json:"description,omitempty"`

	CanUseRemoteSession bool `json:"canUseRemoteSession,omitempty"`

	DisableAlerting bool `json:"disableAlerting,omitempty"`

	NetflowCollectorGroupId int32 `json:"netflowCollectorGroupId,omitempty"`

	CreatedOn int64 `json:"createdOn,omitempty"`

	AncestorHasDisabledLogicModule bool `json:"ancestorHasDisabledLogicModule,omitempty"`

	SystemProperties []NameAndValue `json:"systemProperties,omitempty"`

	ManualDiscoveryFlags ManualDiscoveryFlags `json:"manualDiscoveryFlags,omitempty"`

	HostStatus string `json:"hostStatus,omitempty"`

	AutoPropsUpdatedOn int64 `json:"autoPropsUpdatedOn,omitempty"`

	ScanConfigId int32 `json:"scanConfigId,omitempty"`

	Id int32 `json:"id,omitempty"`

	EnableNetflow bool `json:"enableNetflow,omitempty"`

	LastDataTime int64 `json:"lastDataTime,omitempty"`

	AlertStatusPriority int32 `json:"alertStatusPriority,omitempty"`

	AlertDisableStatus string `json:"alertDisableStatus,omitempty"`

	HostGroupIds string `json:"hostGroupIds"`

	UpTimeInSeconds int64 `json:"upTimeInSeconds,omitempty"`

	DeviceType int32 `json:"deviceType,omitempty"`

	CurrentCollectorId int32 `json:"currentCollectorId,omitempty"`

	NetflowCollectorDescription string `json:"netflowCollectorDescription,omitempty"`

	NetflowCollectorId int32 `json:"netflowCollectorId,omitempty"`

	UserPermission string `json:"userPermission,omitempty"`

	HasDisabledSubResource bool `json:"hasDisabledSubResource,omitempty"`

	AutoPropsAssignedOn int64 `json:"autoPropsAssignedOn,omitempty"`

	UpdatedOn int64 `json:"updatedOn,omitempty"`

	PreferredCollectorGroupName string `json:"preferredCollectorGroupName,omitempty"`

	SdtStatus string `json:"sdtStatus,omitempty"`

	PreferredCollectorGroupId int32 `json:"preferredCollectorGroupId,omitempty"`

	CustomProperties []NameAndValue `json:"customProperties,omitempty"`

	ToDeleteTimeInMs int64 `json:"toDeleteTimeInMs,omitempty"`

	CollectorDescription string `json:"collectorDescription,omitempty"`

	PreferredCollectorId int32 `json:"preferredCollectorId"`

	LastRawdataTime int64 `json:"lastRawdataTime,omitempty"`

	Name string `json:"name"`

	AlertingDisabledOn TreeNode `json:"alertingDisabledOn,omitempty"`

	DeletedTimeInMs int64 `json:"deletedTimeInMs,omitempty"`
}
