package downloader

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/yuioto/gml/core/search/vanilla"
)

const MinecraftAssetsObjectsURL = "https://resources.download.minecraft.net/"

func DownloadVanilla(versionId string) error {
	versionManifest, err := vanilla.GetVersionManifest()
	if err != nil {
		return err
	}

	version, err := vanilla.GetVersion(versionManifest, versionId)
	if err != nil {
		return err
	}

	assetsIndex, err := vanilla.GetAssetIndex(version)
	if err != nil {
		return err
	}

	downloadTask := TasksFromVanilla(versionManifest, version, assetsIndex, versionId)

	err = DownloadAll(downloadTask)
	if err != nil {
		log.Error().Err(err).Msg("Download faild")
		return err
	}

	return nil
}

func TasksFromVanilla(versionManifest vanilla.VersionManifest, vanilla vanilla.Version, assetsIndex vanilla.AssetsIndex, version string) []DownloadTask {
	libraries := make([]DownloadTask, 0, len(vanilla.Libraries))

	// Version
	versionInfo, err := makeTaskVanillaVersion(version)
	if err != nil {
		return []DownloadTask{}
	}
	tasks := append(libraries, versionInfo)

	// Client
	tasks = append(tasks, makeTaskVanillaClient(vanilla.Downloads, version))

	// Server
	tasks = append(tasks, makeTaskVanillaServer(vanilla.Downloads, version))

	// Assets index
	tasks = append(tasks, DownloadTask{
		URL:  vanilla.AssetIndex.URL,
		Path: filepath.Join("assets", "indexes", vanilla.AssetIndex.ID+".json"),
		Sha1: vanilla.AssetIndex.Sha1,
		Size: vanilla.AssetIndex.Size,
	})

	// Logging config
	tasks = append(tasks, DownloadTask{
		URL:  vanilla.Logging.Client.File.URL,
		Path: filepath.Join("assets", "log_configs", vanilla.Logging.Client.File.ID+".json"),
		Sha1: vanilla.Logging.Client.File.Sha1,
		Size: vanilla.Logging.Client.File.Size,
	})

	// Assets Objects
	assetsObjectTasks := makeTaskVanillaAssetsObjects(assetsIndex, version)
	tasks = append(tasks, assetsObjectTasks...)

	// Libraries
	librariesTask := makeTaskVanillaLibrary(vanilla.Libraries)
	tasks = append(tasks, librariesTask...)

	return tasks
}

// Version
func makeTaskVanillaVersion(versionId string) (DownloadTask, error) {
	versionManifest, err := vanilla.GetVersionManifest()
	if err != nil {
		return DownloadTask{}, err
	}
	for _, version := range versionManifest.Versions {
		if version.ID == versionId {
			return DownloadTask{
				URL:  version.URL,
				Path: filepath.Join("versions", versionId, versionId+".json"),
				Sha1: version.Sha1,
			}, nil
		}
	}
	log.Error().Str("version", versionId).Msg("not fount")
	return DownloadTask{}, fmt.Errorf("not fount version")
}

// Client
func makeTaskVanillaClient(downloads vanilla.Downloads, version string) DownloadTask {
	return DownloadTask{
		URL:  downloads.Client.URL,
		Path: filepath.Join("versions", version, version+".jar"),
		Sha1: downloads.Client.Sha1,
		Size: downloads.Client.Size,
	}
}

// Server
func makeTaskVanillaServer(downloads vanilla.Downloads, version string) DownloadTask {
	return DownloadTask{
		URL:  downloads.Server.URL,
		Path: filepath.Join("versions", version, "server-"+version+".jar"),
		Sha1: downloads.Server.Sha1,
		Size: downloads.Server.Size,
	}
}

// AssetsObject
func makeTaskVanillaAssetsObject(sha1 string, size int) DownloadTask {
	return DownloadTask{
		URL:  MinecraftAssetsObjectsURL + sha1[:2] + "/" + sha1,
		Path: filepath.Join("assets", "objects", sha1[:2], sha1),
		Sha1: sha1,
		Size: size,
	}
}

func makeTaskVanillaAssetsObjects(assetIndex vanilla.AssetsIndex, versionId string) []DownloadTask {
	var downloadTask []DownloadTask

	for path, object := range assetIndex.Objects {
		log.Trace().Str("FilePath", path).Str("Hash", object.Hash).Int("Size", object.Size).Msg("Object")
		downloadTask = append(downloadTask, makeTaskVanillaAssetsObject(object.Hash, object.Size))
	}
	return downloadTask
}

// Library
var osMap = map[string]string{
	"darwin":  "osx",
	"linux":   "linux",
	"windows": "windows",
}
var osArch = map[string]string{
	"386":   "x86",
	"amd64": "x86_64",
	"arm64": "aarch64",
}

func LibraryAllowed(library vanilla.Library, os string, arch string) bool {
	allowed := true

	if mappedOS, ok := osMap[os]; ok {
		os = mappedOS
	}
	if mappedArch, ok := osArch[arch]; ok {
		arch = mappedArch
	}

	for _, rule := range library.Rules {
		match := rule.OS.Name == "" || rule.OS.Name == os
		switch rule.Action {
		case "allow":
			if !match {
				return false
			}
		case "disallow":
			if match {
				return false
			}
		}
	}

	if strings.Contains(library.Name, "natives") {
		if strings.Contains(library.Name, "x86") && arch != "x86" {
			allowed = false
		}
		if strings.Contains(library.Name, "arm64") && arch != "arm64" {
			allowed = false
		}
		if strings.Contains(library.Name, "x86_64") && arch != "x86_64" {
			allowed = false
		}
	}
	return allowed
}

func makeTaskVanillaLibrary(libraries []vanilla.Library) []DownloadTask {
	var tasks []DownloadTask
	for _, library := range libraries {
		log.Trace().Str("name", library.Name).Msg("checking Library")
		artifact := library.Downloads.Artifact
		if !LibraryAllowed(library, runtime.GOOS, runtime.GOARCH) {
			continue
		}

		log.Trace().Str("name", library.Name).Msg("added library")
		tasks = append(tasks, DownloadTask{
			URL:  artifact.URL,
			Path: filepath.Join("libraries", artifact.Path),
			Sha1: artifact.Sha1,
			Size: artifact.Size,
		})
	}
	return tasks
}
