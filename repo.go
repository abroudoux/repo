package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	repositoryUrl, err := getRepositoryUrl()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if (isHttpsUrl(repositoryUrl)) {
		err := openHttpsUrlInBrowser(repositoryUrl)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	} else {
		httpsUrl, err := createHttpsUrl(repositoryUrl)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		err = openHttpsUrlInBrowser(httpsUrl)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

func getRepositoryUrl() (string, error) {
	err := isInGitRepository()
	if err != nil {
		return "", err
	}

	cmd := exec.Command("git", "config", "--get", "remote.origin.url")
	repositoryUrl, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("error getting repo url: %v", err)
	}

	return string(repositoryUrl), nil
}

func isInGitRepository() error {
	err := isGitInstalled()
	if err != nil {
		return fmt.Errorf("error checking if in git repository: %v", err)
	}

	cmd := exec.Command("git", "rev-parse", "--is-inside-work-tree")
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("error checking if in git repository: %v", err)
	}

	return nil
}

func isGitInstalled() error {
	cmd := exec.Command("git", "version")
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("git is not installed: %v", err)
	}

	return nil
}

func isHttpsUrl(repositoryUrl string) bool {
	return strings.HasPrefix(repositoryUrl, "https://github.com")
}

func createHttpsUrl(repositoryUrl string) (string, error) {
	if !strings.HasPrefix(repositoryUrl, "git@") {
		return "", fmt.Errorf("invalid SSH URL: %s", repositoryUrl)
	}

	parts := strings.SplitN(repositoryUrl, ":", 2)
	if len(parts) != 2 {
		return "", fmt.Errorf("invalid SSH URL format: %s", repositoryUrl)
	}

	host := strings.TrimPrefix(parts[0], "git@")
	path := strings.TrimSuffix(parts[1], ".git")

	httpsURL := fmt.Sprintf("https://%s/%s", host, path)
	return httpsURL, nil
}

func openHttpsUrlInBrowser(httpsUrl string) error {
	osType, err := determineOs()
	if err != nil {
		return fmt.Errorf("error opening browser: %v", err)
	}

	switch osType {
		case "Linux\n":
			err := openLinkLinux(httpsUrl)
			if err != nil {
				return fmt.Errorf("error opening browser: %v", err)
			}
		case "Darwin\n":
			err := openLinkDarwin(httpsUrl)
			if err != nil {
				return fmt.Errorf("error opening browser: %v", err)
			}
		default:
			return fmt.Errorf("unsupported OS: %s", osType)
	}

	return nil
}

func determineOs() (string, error) {
	cmd := exec.Command("uname")
	osType, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("error determining OS: %v", err)
	}

	return string(osType), nil
}

func openLinkLinux(url string) error {
	cmd := exec.Command("xdg-open", url)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error opening browser: %v", err)
	}

	return nil
}

func openLinkDarwin(url string) error {
	cmd := exec.Command("open", url)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error opening browser: %v", err)
	}

	return nil
}
