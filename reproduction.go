package main

import (
	"log"
	"os"

	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/getter"
	"helm.sh/helm/v3/pkg/repo"
)

func main() {
	logger := log.Default()

	// Create custom temporary cache folder
	cacheFolder, err := os.MkdirTemp(os.TempDir(), "helm-cache")
	if err != nil {
		logger.Fatalf("Failed to create cache folder: %v", err)
	}
	defer os.RemoveAll(cacheFolder)

	settings := cli.New()
	settings.RepositoryCache = cacheFolder

	logger.Printf("Using custom cache folder: %s", cacheFolder)

	actionConfig := new(action.Configuration)
	actionConfig.Init(settings.RESTClientGetter(), "", "memory", logger.Printf)

	pullClient := action.NewPullWithOpts(action.WithConfig(actionConfig))
	pullClient.Settings = settings
	pullClient.DestDir = cacheFolder
	pullClient.RepoURL = "https://charts.bitnami.com/bitnami"
	pullClient.Version = "19.0.0"

	repoEntry := &repo.Entry{
		Name: "temp-repo",
		URL:  "https://charts.bitnami.com/bitnami",
	}

	helmRepo, err := repo.NewChartRepository(repoEntry, getter.All(settings))
	if err != nil {
		logger.Fatalf("Failed to create repository: %v", err)
	}
	helmRepo.CachePath = settings.RepositoryCache

	logger.Println("Downloading repository index...")
	if _, err := helmRepo.DownloadIndexFile(); err != nil {
		logger.Fatalf("Failed to download index: %v", err)
	}
	logger.Println("Index downloaded successfully to custom cache")

	logger.Println("Attempting to pull chart...")
	_, err = pullClient.Run("nginx")
	if err != nil {
		logger.Printf("Pull failed with error: %v", err)
	} else {
		logger.Println("Pull succeeded (no bugs?)")
	}

	logger.Println("Cache folder contents:")
	logger.Printf("Cache folder: %s", cacheFolder)
	files, err := os.ReadDir(cacheFolder)
	if err != nil {
		logger.Fatalf("Failed to read cache folder: %v", err)
	}
	for _, file := range files {
		logger.Printf("- %s", file.Name())
	}
}

// 2025/07/04 18:02:12 Using custom cache folder: /var/folders/29/jjrd3wv91dn_vr35crz7vswr0000gq/T/helm-cache3468950060
// 2025/07/04 18:02:12 Downloading repository index...
// 2025/07/04 18:02:14 Index downloaded successfully to custom cache
// 2025/07/04 18:02:14 Attempting to pull chart...
// 2025/07/04 18:02:19 Pull succeeded (no bugs?)
// 2025/07/04 18:02:19 Cache folder contents:
// 2025/07/04 18:02:19 Cache folder: /var/folders/29/jjrd3wv91dn_vr35crz7vswr0000gq/T/helm-cache3468950060
// 2025/07/04 18:02:19 - nginx-19.0.0.tgz
// 2025/07/04 18:02:19 - temp-repo-charts.txt
// 2025/07/04 18:02:19 - temp-repo-index.yaml
