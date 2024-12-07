# repo

📁 Browse Github from commands

Version : 1.0.0

## 🚀 Installation

### Via Homebrew

Wip 🚧

### Manual

You can copy the function in your shell's RC file. Alternatively, You can create a separate Bash script file and copy `pm.sh` into it. You'll need to load it at the beginning of your shell RC file (e.g., `.bashrc`, `.zshrc`, etc.).

```bash
source path/to/your/script.sh
```

Don't forget to resource your shell RC file:

```bash
source ~/.zshrc
```

## 💻 Usage

`repo` allows you to browser Github Repositories from commands

```bash
# Opens a new browser window with your GitHub repository
repo
```

Multiple tags are existing

```bash
repo -u abroudoux # Searches for a user
repo -r repo # Searches for a repository
repo -b main # Open the current repository with the specified branch
```

You can search for a repository by using the `--link` / `-l` flag.

```bash
repo https://github.com/abroudoux/repo
```

## 🧑‍🤝‍🧑 Contributing

To contribute, fork the repository and open a pull request with the details of your changes.

Create a branch with a [conventionnal name](https://tilburgsciencehub.com/building-blocks/collaborate-and-share-your-work/use-github/naming-git-branches/).

- fix: `bugfix/the-bug-fixed`
- features: `feature/the-amazing-feature`
- test: `test/the-famous-test`
- hotfix `hotfix/oh-my-god-bro`
- wip `wip/the-work-name-in-progress`

## 📌 Roadmap

- [ ] Refactor code
- [ ] Open repository in active browser window

## 📑 License

This project is under MIT license. For more information, please see the file [LICENSE](./LICENSE).
