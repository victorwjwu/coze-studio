// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import "github.com/coze-dev/coze-studio/backend/domain/app/entity"

const TableNameAppReleaseRecord = "app_release_record"

// AppReleaseRecord Application Release Record
type AppReleaseRecord struct {
	ID            int64                          `gorm:"column:id;primaryKey;comment:Publish Record ID" json:"id"`                                              // Publish Record ID
	AppID         int64                          `gorm:"column:app_id;not null;comment:Application ID" json:"app_id"`                                           // Application ID
	SpaceID       int64                          `gorm:"column:space_id;not null;comment:Space ID" json:"space_id"`                                             // Space ID
	OwnerID       int64                          `gorm:"column:owner_id;not null;comment:Owner ID" json:"owner_id"`                                             // Owner ID
	IconURI       string                         `gorm:"column:icon_uri;not null;comment:Icon URI" json:"icon_uri"`                                             // Icon URI
	Name          string                         `gorm:"column:name;not null;comment:Application Name" json:"name"`                                             // Application Name
	Description   string                         `gorm:"column:description;comment:Application Description" json:"description"`                                 // Application Description
	ConnectorIds  []int64                        `gorm:"column:connector_ids;comment:Publish Connector IDs;serializer:json" json:"connector_ids"`               // Publish Connector IDs
	ExtraInfo     *entity.PublishRecordExtraInfo `gorm:"column:extra_info;comment:Publish Extra Info;serializer:json" json:"extra_info"`                        // Publish Extra Info
	Version       string                         `gorm:"column:version;not null;comment:Release Version" json:"version"`                                        // Release Version
	VersionDesc   string                         `gorm:"column:version_desc;comment:Version Description" json:"version_desc"`                                   // Version Description
	PublishStatus int32                          `gorm:"column:publish_status;not null;comment:Publish Status" json:"publish_status"`                           // Publish Status
	PublishAt     int64                          `gorm:"column:publish_at;not null;comment:Publish Time in Milliseconds" json:"publish_at"`                     // Publish Time in Milliseconds
	CreatedAt     int64                          `gorm:"column:created_at;not null;autoCreateTime:milli;comment:Create Time in Milliseconds" json:"created_at"` // Create Time in Milliseconds
	UpdatedAt     int64                          `gorm:"column:updated_at;not null;autoUpdateTime:milli;comment:Update Time in Milliseconds" json:"updated_at"` // Update Time in Milliseconds
}

// TableName AppReleaseRecord's table name
func (*AppReleaseRecord) TableName() string {
	return TableNameAppReleaseRecord
}
