function repo() {
    repo_url=$(git config --get remote.origin.url)
    request_url=""

    if ! command -v git &> /dev/null; then
        echo "Error: Git is not installed."
        return 1
    fi

    if [[ $repo_url == "https://"* ]]; then
        $request_url=$repo_url
    else
        username=$(echo $repo_url | awk -F 'git@github.com:|/' '{print $2}')
        repo_name=$(echo $repo_url | awk -F '${username}|/' '{print $2}' | sed 's/\.git$//')
        request_url="https://github.com/$username/$repo_name"
    fi

    while [[ $# -gt 0 ]]; do
        case "$1" in
            --branch|-b)
                if [[ -n "$2" ]]; then
                    request_url="$request_url/tree/$2"
                    return
                else
                    echo "Error: --branch requires a value." >&2
                    return 1
                fi
                ;;
            --repository|-r)
                if [[ -n "$2" ]]; then
                    request_url="https://github.com/search?q=$2&type=repositories"
                    return
                else
                    echo "Error: --repository requires a value." >&2
                    return 1
                fi
                ;;
            --user|-u)
                if [[ -n "$2" ]]; then
                    request_url="https://github.com/search?q=$2&type=users"
                    return
                else
                    echo "Error: --user requires a value." >&2
                    return 1
                fi
                ;;
            --link|-l)
                request_url="https://github.com/$1"
                return
                ;;
            --version|-v)
                version=$(jq -r '.version' ./package.json) 
                echo "$version"
                return
                ;;
            --help|-h)
                echo "Usage: repo [options] [command]"
                echo "Options:"
                echo "  --branch, -b          Open the current repository with the specified branch"
                echo "  --repository, -r      Search for a repository"
                echo "  --user, -u            Search for a user"
                echo "  --link, -l            Open the repository link in the browser"
                echo "  --help, -h            Show this help message"
                return
                ;;
            *)
                echo "Unknown option: $1" >&2
                echo "Use --help or -h for usage information."
                return 1
                ;;
        esac
        shift
    done

    os_type=$(uname)

    case "$os_type" in
        "Linux")
            xdg-open "$request_url" 2>/dev/null || echo "Error: Cannot open the link."
            ;;
        "Darwin")
            open "$request_url" || echo "Error: Cannot open the link."
            ;;
        "CYGWIN"*|"MINGW"*)
            start "$request_url" || echo "Error: Cannot open the link."
            ;;
        *)
            echo "Non-supported OS : $os_type"
            ;;
    esac
}