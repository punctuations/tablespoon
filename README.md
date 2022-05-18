<h1 align="center">🥄 Tablespoon</h1>
<p align="center">A cli tool to generate semantic commit messages.</p>

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

Tablespoon is a simple cli tool that can generate semantic commit messages based on comments in order to enforce best practices!

## Documentation
All documentation can be found through the tablespoon site here: [docs.tbsp.coffee](https://docs.tbsp.coffee)

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
- Writes [semantic](https://gist.github.com/joshbuchea/6f47e86d2510bce28f8e7f42ae84c716) commit messages for you:
  ```bash
  <type>(<desc>): <summary>
  ```
- Generate just the message **or** generate the message and immediately commit it!
- Customize commentID via the config file!

## TODO
- [ ] Try again for fig x tablespoon for autocompletion
