// (c) Copyright 2022 Hewlett Packard Enterprise Development LP

package models

// CreateContentDataReq - struct representing CreateContentData request
type CreateContentDataReq struct {
	Title       string   `json:"title" validate:"required"`
	Description string   `json:"description" validate:"required"`
	Services    []string `json:"services,omitempty" validate:"dive"`
	Roles       []string `json:"roles,omitempty" validate:"dive"`
	ContentType string   `json:"contentType" validate:"required"`
	LocationURL string   `json:"locationURL" validate:"required"`
}

// CreateContentDataResp - struct representing CreateContentData response
type CreateContentDataResp struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Services    []string `json:"services"`
	Roles       []string `json:"roles"`
	ContentType string   `json:"contentType"`
	LocationURL string   `json:"locationURL"`
	CreatedAt   string   `json:"createdAt"`
}

// UpdateContentDataReq - struct representing UpdateContentData request
type UpdateContentDataReq struct {
	Title          string   `json:"title,omitempty"`
	Description    string   `json:"description,omitempty"`
	Services       []string `json:"services,omitempty"`
	Roles          []string `json:"roles,omitempty"`
	ContentType    string   `json:"contentType,omitempty"`
	LocationURL    string   `json:"locationURL,omitempty"`
	FieldsToUpdate []string `json:"fieldsToUpdate,omitempty"`
}

// UpdateContentDataResp - struct representing UpdateContentData response
type UpdateContentDataResp struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Services    []string `json:"services"`
	Roles       []string `json:"roles"`
	ContentType string   `json:"contentType"`
	LocationURL string   `json:"locationURL"`
	CreatedAt   string   `json:"createdAt"`
}

// GetContentDataReq - struct representing GetContentData request
type GetContentDataReq struct {
	Services     []string `json:"services"`
	Roles        []string `json:"roles"`
	ContentTypes []string `json:"contentTypes"`
	Search       string   `json:"search"`
	SortBy       string   `json:"sortby"`
}

// GetContentDataResp - struct representing GetContentData response
type GetContentDataResp struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Services    []string `json:"services"`
	Roles       []string `json:"roles"`
	ContentType string   `json:"contentType"`
	LocationURL string   `json:"locationURL"`
	CreatedAt   string   `json:"createdAt"`
	UpdatedAt   string   `json:"updatedAt"`
}
