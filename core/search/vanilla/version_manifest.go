package vanilla

import (
	"net/http"
	"time"

	"github.com/yuioto/gml/internal/utils"
)

const VersionManifestV2URL = "https://launchermeta.mojang.com/mc/game/version_manifest_v2.json"

type Latest struct {
	Release  string `json:"release"`
	Snapshot string `json:"snapshot"`
}

/* version_manifest_v2 VersionManifest.Versions.VersionAll */
type Versions struct {
	ID              string      `json:"id"`
	Type            ReleaseType `json:"type"`
	URL             string      `json:"url"`
	Time            time.Time   `json:"time"`
	ReleaseTime     time.Time   `json:"releaseTime"`
	Sha1            string      `json:"sha1"`
	ComplianceLevel int         `json:"complianceLevel"`
}

// ReleaseType enum
type ReleaseType string

const (
	ReleaseTypeRelease  ReleaseType = "release"
	ReleaseTypeSnapshot ReleaseType = "snapshot"
	ReleaseTypeOldBeta  ReleaseType = "old_beta"
	ReleaseTypeOldAlpha ReleaseType = "old_alpha"
)

type VersionManifest struct {
	Latest   Latest     `json:"latest"`
	Versions []Versions `json:"versions"`
}

func GetVersionManifest() (VersionManifest, error) {
	var versionManifest VersionManifest
	if err := utils.FetchJSON(http.DefaultClient, VersionManifestV2URL, &versionManifest); err != nil {
		return VersionManifest{}, err
	}
	return versionManifest, nil
}
