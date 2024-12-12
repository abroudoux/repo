# repo

📁 Browse Github from commands

Version : 2.0.0

## 🚀 Installation

### Via Homebrew

Wip 🚧

### Manual

You can paste the binary in your `bin` directory (e.g., on mac it's `/usr/bin/local`). \
Don't forget to grant execution permissions to the binary.

```bash
chmox +x repo
```

## 💻 Usage

`repo` allows you to browser Github Repositories from commands.

```bash
# Opens a new browser window with your current GitHub repository
repo
```

Use the `--branch` / `-b` flag to open the current Gihtub repository with the specified branch.

```bash
repo --branch develop
```

You can search for a repository by using the `--link` / `-l` flag.

```bash
repo https://github.com/abroudoux/repo
```

If you don't have the URl, use the `--search` / `-s` flag to make a research.

```bash
repo --search repo
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

- [x] Refactor code in Go
- [ ] Open repository in active browser window
- [ ] Add more option to browser repositories from commands
- [ ] Open specified menus of the specified repository (e.g., Pull Requests, Issues...)
- [ ] Window supports

## 📑 License

This project is under MIT license. For more information, please see the file [LICENSE](./LICENSE).
