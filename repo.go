package main

import (
	_ "embed"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

//go:embed assets/ascii.txt
var asciiArt string

func main() {
	repositoryUrl, err := getRepositoryUrl()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if len(os.Args) > 1 {
		flagMode(repositoryUrl)
		os.Exit(0)
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
	path := strings.TrimSpace(parts[1])
	path = strings.TrimSuffix(path, ".git")
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

func openRepositoryWithBranch(repositoryUrl string, branch string) error {
	httpsUrl, err := createHttpsUrl(repositoryUrl)
	if err != nil {
		return fmt.Errorf("error opening repository with branch: %v", err)
	}

	httpsUrl = strings.TrimSpace(httpsUrl)
	branch = strings.TrimSpace(branch)
	httpsUrlWithBranch := fmt.Sprintf("%s/tree/%s", httpsUrl, branch)
	err = openHttpsUrlInBrowser(httpsUrlWithBranch)
	if err != nil {
		return fmt.Errorf("error opening repository with branch: %v", err)
	}

	return nil
}

func searchRepository(repositoryName string) error {
	searchUrl := fmt.Sprintf("https://github.com/search?q=%s&type=repositories", repositoryName)
	err := openHttpsUrlInBrowser(searchUrl)
	if err != nil {
		return fmt.Errorf("error searching repository: %v", err)
	}

	return nil
}

func flagMode(repositoryUrl string) {
	flag := os.Args[1]

	if flag == "--branch" || flag == "-b" {
		branch := os.Args[2]
		err := openRepositoryWithBranch(repositoryUrl, branch)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	} else if flag == "--search" || flag == "-s" {
		repositoryName := os.Args[2]
		err := searchRepository(repositoryName)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	} else if flag == "--link" || flag == "-l" {
		err := openHttpsUrlInBrowser(repositoryUrl)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	} else if flag == "--version" || flag == "-v" {
		println(asciiArt)
		println("2.0.0")
	} else if flag == "--help" || flag == "-h" {
		printHelpManual()
	}
}

func printHelpManual() {
	fmt.Println("Usage: repo [options] [command]")
	fmt.Printf("  %-20s %s\n", "repo [--branch | -b]", "Open the current repository with the specified branch")
	fmt.Printf("  %-20s %s\n", "repo [--search | -s]", "Search for a repository")
	fmt.Printf("  %-20s %s\n", "repo [--link | -l]", "Open the repository link in the browser")
	fmt.Printf("  %-20s %s\n", "commit [--help | -h]", "Show this help message")
}