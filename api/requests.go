package api

import (
	"encoding/json"
	"go_banners/db"
)

type GetBannerRequest struct {
	FeatureId *int
	TagId     *int
	Limit     *int
	Offset    *int
}

type PostBannerRequest struct {
	TagIds    []int           `json:"tag_ids"`
	FeatureId int             `json:"feature_id"`
	Content   json.RawMessage `json:"content"`
	IsActive  bool            `json:"is_active"`
}

type PatchBanner struct {
	Id     int
	Banner db.Banner
}

type DeleteBanner struct {
	Id int
}
