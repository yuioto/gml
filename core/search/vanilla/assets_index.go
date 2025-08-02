package vanilla

import (
	"net/http"

	"github.com/yuioto/gml/internal/utils"
)

type AssetsIndex struct {
	Objects Objects `json:"objects"`
}

type Objects map[string]struct {
	Hash string `json:"hash"`
	Size int    `json:"size"`
}

func GetAssetIndex(version Version) (AssetsIndex, error) {
	url := version.AssetIndex.URL
	var assetsIndex AssetsIndex
	if err := utils.FetchJSON(http.DefaultClient, url, &assetsIndex); err != nil {
		return AssetsIndex{}, err
	}
	return assetsIndex, nil
}
