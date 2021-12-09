<h1 align="center">🥄 Tablespoon</h1>
<h5 align="center">EXPERIMENTAL PREVIEW</h5> 
<p align="center">A simple generator for semantic git messages.</p>

<p align="center">

<a style="text-decoration: none" href="https://github.com/punctuations/tablespoon/releases">
<img src="https://img.shields.io/github/v/release/punctuations/tablespoon?style=flat-square" alt="Latest Release">
</a>

<a style="text-decoration: none" href="https://github.com/punctuations/tablespoon/releases">
<img src="https://img.shields.io/github/downloads/punctuations/tablespoon/total.svg?style=flat-square" alt="Downloads">
</a>

<a style="text-decoration: none" href="https://github.com/punctuations/tablespoon/stargazers">
<img src="https://img.shields.io/github/stars/punctuations/tablespoon.svg?style=flat-square" alt="Stars">
</a>

<a style="text-decoration: none" href="https://github.com/punctuations/tablespoon/fork">
<img src="https://img.shields.io/github/forks/punctuations/tablespoon.svg?style=flat-square" alt="Forks">
</a>

<a style="text-decoration: none" href="https://github.com/punctuations/tablespoon/issues">
<img src="https://img.shields.io/github/issues/punctuations/tablespoon.svg?style=flat-square" alt="Issues">
</a>

<a style="text-decoration: none" href="https://opensource.org/licenses/MIT">
<img src="https://img.shields.io/badge/License-MIT-yellow.svg?style=flat-square" alt="License: MIT">
</a>

<br/>

<a style="text-decoration: none" href="https://github.com/{{ .ProjectPath }}/releases">
<img src="https://img.shields.io/badge/platform-windows%20%7C%20macos%20%7C%20linux-informational?style=for-the-badge" alt="Downloads">
</a>

<br/>

</p>

<br/>

<p align="center">
<strong><a href="#installation">Installation</a></strong>
|
<strong><a href="#CONTRIBUTING">Contributing</a></strong>
</p>

<br/>

Tablespoon is a simple generator which can generate semantic git commit messages of not only the highlights of the change, but everything.

## Installation

Run the following command in a terminal and you're ready to go!

**Windows**
```powershell
iwr instl.sh/punctuations/tablespoon/windows | iex 
```

**macOS**
```bash
curl -sSL instl.sh/punctuations/tablespoon/macos | sudo bash   
```

**Linux**
```bash
curl -sSL instl.sh/punctuations/tablespoon/linux | sudo bash  
```

----

### Functionality
- Writes semantic git messages for you:
  ```bash
  <type>(<desc>): <summary>
  ```
- Can generate them lone or generate them and create commit
- all semantic-ness of this are based off of [the gist by joshbuchea](https://gist.github.com/joshbuchea/6f47e86d2510bce28f8e7f42ae84c716).
- be able to select which file you wish to be showcased by using `-s` flag then the file name.
- Will list all other changes in long description.
- Allow for tablespoon.config.json file and allow for customization
- Allow for --no-comment flag to manually input short desc

## TODO
- Prompt for framework to determine importance of files? (For whichever goes in the short summary)
- when message is generated have a few to select from (similar to [thefuck](https://github.com/nvbn/thefuck))
- aliases that are recommended: `tbl`, `tbs`, or `tbsp`

## OPTIONS
- [ ] [Javascript](https://github.com/njzydark/project-template-cli)
- [x] [Go](https://github.com/pterm/cli-template)
- [ ] [Rust](https://github.com/rust-cli/cli-template)
