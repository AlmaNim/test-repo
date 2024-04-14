package api

import (
	"encoding/json"
	"go_banners/db"
)

type GetBannerRequest struct { //определяет параметры запроса для получения баннера
	FeatureId *int //уникалный id фичи
	TagId     *int //уникальный id тега
	Limit     *int //максимальное кол-во баннеров при отображении
	Offset    *int //число, определяющее, с какого "номера" начать показ баннеров
}

type PostBannerRequest struct { //определят параметры запроса для создания баннера
	TagIds    []int           `json:"tag_ids"`    //список идентификаторов тегов, ассоциированных с баннером
	FeatureId int             `json:"feature_id"` //идентификатор фичи, ассоциированной с баннером
	Content   json.RawMessage `json:"content"`    //содержимое баннера в формате JSON, json.RawMessage используется, чтобы задержать декодирование, пока не будет точно известна структура
	IsActive  bool            `json:"is_active"`  //указывает, активен ли баннер
}

type PatchBanner struct { //структура запроса для обновления баннера
	Id     int       //уникальный id обновляемого баннера
	Banner db.Banner //объект баннера с обновлёнными данными
}

type DeleteBanner struct { //структура запроса для удаления баннера
	Id int //уникальный id удаляемого баннера
}
