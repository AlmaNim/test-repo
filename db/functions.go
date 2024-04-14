package db

import (
	"errors"
	"go_banners/api"

	"github.com/lib/pq"
	"gorm.io/datatypes"
)

func CreateBanner(banner *api.PostBannerRequest) (*uint, error) { //создает новый баннер на основе предоставленного запроса, возвращает id баннера или ошибку

	result := db.Create(&Banner{
		FeatureId: banner.FeatureId,
		Tags:      banner.TagIds,
		Content:   datatypes.JSON(banner.Content),
	})
	if result.Error != nil {
		return nil, errors.New("can't save") //в случае возникновения ошибки при сохранении возвращается ошибка
	}
	return &banner.ID, nil //возврат id созданного баннера
}

func GetBanner(tagId int, featureId int, includeInactive bool) (*Banner, error) { //возвращает баннер по предоставленному id тега, id фичи и флагу, указывающему, должны ли включаться неактивные баннеры. Возвращает найденный баннер или ошибку.

	banner := Banner{
		FeatureId: featureId,
		Tags:      pq.Int32Array{int32(tagId)},
	}
	if !includeInactive {
		banner.IsActive = true //Если не нужно включать неактивные баннеры, установить IsActive в true, будет использовано, когда будет прикручен токен авторизации
	}
	result := db.First(&banner)
	if result.Error != nil {
		return nil, errors.New("not found")
	}
	return &banner, nil
}

func GetBanners(featureId *int, tagId *int, limit *int, offset *int) (*[]Banner, error) { //возвращает массив баннеров в соответствии с заданными параметрами: featureId, tagId и ограничениями на количество и смещение. Возвращает массив баннеров или ошибку.

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
	var banners *[]Banner //переменная для хранения результатов запроса
	result := query.Find(banners)
	if result.Error != nil {
		return nil, errors.New("internal error")
	}
	return banners, nil
}

func UpdateBanner(bannerId int, banner Banner) error { //обновляет существующий баннер с заданным id на основе предоставленного баннера. Возвращает ошибку в случае неудачи.
	var dbBanner *Banner
	result := db.First(dbBanner, bannerId)
	if result.Error != nil {
		return errors.New("internal error") //неизвестная ошибка при поиске
	}
	if dbBanner == nil {
		return errors.New("not found") //баннер не найден
	}
	//обновление полей баннера:
	dbBanner.IsActive = banner.IsActive
	dbBanner.Content = banner.Content
	dbBanner.FeatureId = banner.FeatureId
	dbBanner.Tags = banner.Tags
	db.Save(dbBanner)

	return nil
}

func DeleteBanner(bannerId int) error { //удаляет баннер с заданным id из бд. Возвращает ошибку, если баннер не найден или при возникновении внутренней ошибки.
	result := db.Delete(&Banner{}, bannerId)
	if result.Error != nil {
		return errors.New("internal error")
	}
	if result.RowsAffected == 0 {
		return errors.New("not found")
	}
	return nil
}
