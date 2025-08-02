package vanilla

import (
	"fmt"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/yuioto/gml/internal/utils"
)

// Arguments
type Arguments struct {
	Game []any `json:"game"`
	Jvm  []any `json:"jvm"`
}

// AssetIndex
type AssetIndex struct {
	ID        string `json:"id"`
	Sha1      string `json:"sha1"`
	Size      int    `json:"size"`
	TotalSize int    `json:"totalSize"`
	URL       string `json:"url"`
}

// Downloads
type Downloads struct {
	Client         Client         `json:"client"`
	ClientMappings ClientMappings `json:"client_mappings"`
	Server         Server         `json:"server"`
	ServerMappings ServerMappings `json:"server_mappings"`
}

type DownloadsMetadata struct {
	Sha1 string `json:"sha1"`
	Size int    `json:"size"`
	URL  string `json:"url"`
}

type Client DownloadsMetadata
type ClientMappings DownloadsMetadata
type Server DownloadsMetadata
type ServerMappings DownloadsMetadata

type JavaVersion struct {
	Component    string `json:"component"`
	MajorVersion int    `json:"majorVersion"`
}

// Libraries

type Library struct {
	Downloads LibraryDownloads `json:"downloads"`
	Name      string           `json:"name"`
	Rules     []LibraryRules   `json:"rules"`
}

type LibraryDownloads struct {
	Artifact Artifact `json:"artifact"`
}

type Artifact struct {
	Path string `json:"path"`
	Sha1 string `json:"sha1"`
	Size int    `json:"size"`
	URL  string `json:"url"`
}

type LibraryRules struct {
	Action string `json:"action"`
	OS     OS     `json:"os"`
}

type OS struct {
	Name string `json:"name"`
}

// Logging
type Logging struct {
	Client LoggingClient `json:"client"`
}

type LoggingClient struct {
	Argument string `json:"argument"`
	File     File   `json:"file"`
}

type File struct {
	ID   string `json:"id"`
	Sha1 string `json:"sha1"`
	Size int    `json:"size"`
	URL  string `json:"url"`
}

/* version VersionManifest.Versions.VersionAll.URL.Version */
type Version struct {
	Arguments              Arguments   `json:"arguments"`
	AssetIndex             AssetIndex  `json:"assetIndex"`
	Assets                 string      `json:"assets"`
	ComplianceLevel        int         `json:"complianceLevel"`
	Downloads              Downloads   `json:"downloads"`
	Id                     string      `json:"id"`
	JavaVersion            JavaVersion `json:"javaVersion"`
	Libraries              []Library   `json:"libraries"`
	Logging                Logging     `json:"logging"`
	MainClass              string      `json:"mainClass"`
	MinimumLauncherVersion int         `json:"minimumLauncherVersion"`
	ReleaseTime            time.Time   `json:"releaseTime"`
	Time                   time.Time   `json:"time"`
	Type                   ReleaseType `json:"type"`
}

func getVersionFromURL(url string) (Version, error) {
	var versionInfo Version
	if err := utils.FetchJSON(http.DefaultClient, url, &versionInfo); err != nil {
		return versionInfo, err
	}
	return versionInfo, nil
}

func GetVersion(versionManifest VersionManifest, versionId string) (Version, error) {
	for _, version := range versionManifest.Versions {
		if version.ID == versionId {
			return getVersionFromURL(version.URL)
		}
	}
	log.Error().Str("version", versionId).Msg("not fount")
	return Version{}, fmt.Errorf("not fount version")
}
