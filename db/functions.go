package db

import (
	"errors"
	"go_banners/api"

	"github.com/lib/pq"
	"gorm.io/datatypes"
)

func CreateBanner(banner *api.PostBannerRequest) (*uint, error) {

	result := db.Create(&Banner{
		FeatureId: banner.FeatureId,
		Tags:      banner.TagIds,
		Content:   datatypes.JSON(banner.Content),
	})
	if result.Error != nil {
		return nil, errors.New("can't save")
	}
	return &banner.ID, nil
}

func GetBanner(tagId int, featureId int, includeInactive bool) (*Banner, error) {
	banner := Banner{
		FeatureId: featureId,
		Tags:      pq.Int32Array{int32(tagId)},
	}
	if !includeInactive {
		banner.IsActive = true
	}
	result := db.First(&banner)
	if result.Error != nil {
		return nil, errors.New("not found")
	}
	return &banner, nil
}

func GetBanners(featureId *int, tagId *int, limit *int, offset *int) (*[]Banner, error) {
	query := db.Order("ID DESC")
	if limit != nil {
		query.Limit(*limit)
	}
	if limit != nil {
		query.Offset(*offset)
	}
	if featureId != nil {
		query.Where("feature_id = ?", featureId)
	}
	if tagId != nil {
		query.Where("tag_id = ?", tagId)
	}
	var banners *[]Banner
	result := query.Find(banners)
	if result.Error != nil {
		return nil, errors.New("internal error")
	}
	return banners, nil
}

func UpdateBanner(bannerId int, banner Banner) error {
	var dbBanner *Banner
	result := db.First(dbBanner, bannerId)
	if result.Error != nil {
		return errors.New("internal error")
	}
	if dbBanner == nil {
		return errors.New("not found")
	}
	dbBanner.IsActive = banner.IsActive
	dbBanner.Content = banner.Content
	dbBanner.FeatureId = banner.FeatureId
	dbBanner.Tags = banner.Tags
	db.Save(dbBanner)

	return nil
}

func DeleteBanner(bannerId int) error {
	result := db.Delete(&Banner{}, bannerId)
	if result.Error != nil {
		return errors.New("internal error")
	}
	if result.RowsAffected == 0 {
		return errors.New("not found")
	}
	return nil
}
