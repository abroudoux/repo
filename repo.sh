function repo() {
    repo_ssh_url=$(git config --get remote.origin.url)
    username=$(echo $repo_ssh_url | awk -F 'git@github.com:|/' '{print $2}')
    repo_name=$(echo $repo_ssh_url | awk -F '${username}|/' '{print $2}')
    repo_url="https://github.com/$username/$repo_name"
    os_type=$(uname)

    case "$os_type" in
        "Linux")
            xdg-open "$repo_url" 2>/dev/null || echo "Error: Cannot open the link."
            ;;
        "Darwin")
            open "$repo_url" || echo "Error: Cannot open the link."
            ;;
        "CYGWIN"*|"MINGW"*)
            start "$repo_url" || echo "Error: Cannot open the link."
            ;;
        *)
            echo "Non-supported OS : $os_type"
            ;;
    esac
}