package vanilla

type AssetsIndex struct {
	Objects Objects `json:"objects"`
}

type Objects map[string]struct {
	Hash string `json:"hash"`
	Size int    `json:"size"`
}
