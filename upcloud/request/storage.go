package request

import (
	"encoding/xml"
	"fmt"
	"github.com/jalle19/upcloud-go-sdk/upcloud"
)

/**
GetStoragesRequest represents a request for retrieving all or some storages
*/
type GetStoragesRequest struct {
	// If specified, only storages with this access type will be retrieved
	Access string
	// If specified, only storages with this type will be retrieved
	Type string
	// If specified, only storages marked as favorite will be retrieved
	Favorite bool
}

/**
RequestURL() implements the Request interface
*/
func (r *GetStoragesRequest) RequestURL() string {
	if r.Access != "" {
		return fmt.Sprintf("/storage/%s", r.Access)
	}

	if r.Type != "" {
		return fmt.Sprintf("/storage/%s", r.Type)
	}

	if r.Favorite {
		return "/storage/favorite"
	}

	return "/storage"
}

/**
GetStorageDetailsRequest represents a request for retrieving details about a piece of storage
*/
type GetStorageDetailsRequest struct {
	UUID string
}

/**
RequestURL() implements the Request interface
*/
func (r *GetStorageDetailsRequest) RequestURL() string {
	return fmt.Sprintf("/storage/%s", r.UUID)
}

/**
CreateStorageRequest represents a request to create a storage device
*/
type CreateStorageRequest struct {
	XMLName xml.Name `xml:"storage"`

	Size       int                 `xml:"size"`
	Tier       string              `xml:"tier,omitempty"`
	Title      string              `xml:"title"`
	Zone       string              `xml:"zone"`
	BackupRule *upcloud.BackupRule `xml:"backup_rule,omitempty"`
}

/**
RequestURL() implements the Request interface
*/
func (r *CreateStorageRequest) RequestURL() string {
	return "/storage"
}

/**
ModifyStorageRequest represents a request to modify a storage device
*/
type ModifyStorageRequest struct {
	XMLName xml.Name `xml:"storage"`
	UUID    string   `xml:"-"`

	Title      string              `xml:"title,omitempty"`
	Size       int                 `xml:"size,omitempty"`
	BackupRule *upcloud.BackupRule `xml:"backup_rule,omitempty"`
}

/**
RequestURL() implements the Request interface
*/
func (r *ModifyStorageRequest) RequestURL() string {
	return fmt.Sprintf("/storage/%s", r.UUID)
}

/**
AttachStorageRequest represents a request to attach a storage device to a server
*/
type AttachStorageRequest struct {
	XMLName    xml.Name `xml:"storage_device"`
	ServerUUID string   `xml:"-"`

	Type        string `xml:"type,omitempty"`
	Address     string `xml:"address,omitempty"`
	StorageUUID string `xml:"storage,omitempty"`
}

/**
RequestURL() implements the Request interface
*/
func (r *AttachStorageRequest) RequestURL() string {
	return fmt.Sprintf("/server/%s/storage/attach", r.ServerUUID)
}

/**
DetachStorageRequest represents a request to detach a storage device from a server
*/
type DetachStorageRequest struct {
	XMLName    xml.Name `xml:"storage_device"`
	ServerUUID string   `xml:"-"`

	Address string `xml:"address"`
}

/**
RequestURL() implements the Request interface
*/
func (r *DetachStorageRequest) RequestURL() string {
	return fmt.Sprintf("/server/%s/storage/detach", r.ServerUUID)
}

/**
DeleteStorageRequest represents a request to delete a storage device
*/
type DeleteStorageRequest struct {
	UUID string
}

/**
RequestURL() implements the Request interface
*/
func (r *DeleteStorageRequest) RequestURL() string {
	return fmt.Sprintf("/storage/%s", r.UUID)
}
