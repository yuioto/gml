package downloader

import (
	"crypto/sha1"
	"fmt"
	"io"
	"os"

	"github.com/cavaliergopher/grab/v3"
	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"
)

type DownloadTask struct {
	URL  string
	Path string
	Sha1 string
	Size int
}

type Downloader interface {
	Enqueue(task DownloadTask)
	Wait() error
}

func DownloadAndCheckSHA1(url, filepath, expectedSHA1 string, client *grab.Client) error {
	// check file SHA1
	if _, err := os.Stat(filepath); err == nil {
		ok, err := verifySHA1(filepath, expectedSHA1)
		if err != nil {
			return fmt.Errorf("failed to verify existing file: %w", err)
		}
		if ok {
			log.Debug().Str("Path", filepath).Msg("File exists and SHA1 matches, skipping download")
			return nil
		}
		log.Warn().Str("Path", filepath).Msg("File exists but SHA1 mismatch, redownloading")
	}

	// download
	req, err := grab.NewRequest(filepath, url)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	resp := client.Do(req)
	<-resp.Done

	if err := resp.Err(); err != nil {
		return fmt.Errorf("download failed: %w", err)
	}

	// download and verify SHA1
	ok, err := verifySHA1(filepath, expectedSHA1)
	if err != nil {
		return fmt.Errorf("failed to verify downloaded file: %w", err)
	}
	if !ok {
		return fmt.Errorf("SHA1 mismatch after download")
	}

	log.Debug().Str("Path", filepath).Msg("Download completed and SHA1 verified")
	return nil
}

func verifySHA1(path, expected string) (bool, error) {
	f, err := os.Open(path)
	if err != nil {
		return false, err
	}
	defer f.Close()

	h := sha1.New()
	if _, err := io.Copy(h, f); err != nil {
		return false, err
	}

	sum := fmt.Sprintf("%x", h.Sum(nil))
	return sum == expected, nil
}

func DownloadFromDownloadTask(downloadTask DownloadTask, client *grab.Client) error {
	if err := DownloadAndCheckSHA1(downloadTask.URL, downloadTask.Path, downloadTask.Sha1, client); err != nil {
		return err
	}
	return nil
}

func DownloadAll(downloadTaskAll []DownloadTask) error {
	var g errgroup.Group
	client := grab.NewClient()

	for _, downloadTask := range downloadTaskAll {
		g.Go(func() error {
			log.Trace().
				Str("URL", downloadTask.URL).
				Str("Path", downloadTask.Path).
				Str("Sha1", downloadTask.Sha1).
				Msg("Download Start")

			return DownloadFromDownloadTask(downloadTask, client)
		})
	}

	return g.Wait()
}
