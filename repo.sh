function repo() {
    repo_url=$(git config --get remote.origin.url)
    request_url=""

    if [[ $repo_url == "https://"* ]]; then
        $request_url=$repo_url
    else
        username=$(echo $repo_url | awk -F 'git@github.com:|/' '{print $2}')
        repo_name=$(echo $repo_url | awk -F '${username}|/' '{print $2}' | sed 's/\.git$//')
        request_url="https://github.com/$username/$repo_name"
    fi

    if [[ "$1" == "-b" && -n "$2" ]]; then
        branch_name=$2
        request_url="$request_url/tree/$branch_name"
    fi

    if [[ "$1" == "-u" && -n "$2" ]]; then
        username=$2
        request_url="https://github.com/search?q=${username}&type=users";
    fi

    if [[ "$1" == "-r" && -n "$2" ]]; then
        repository_name=$2
        request_url="https://github.com/search?q=${repository_name}&type=repositories";
    fi

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