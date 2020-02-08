# github-profile-card
[![Gitpod Ready-to-Code](https://img.shields.io/badge/Gitpod-Ready--to--Code-blue?logo=gitpod)](https://gitpod.io/#https://github.com/kena0ki/github-profile-card) 

This repository is kind of like user profile version of [gh-card](https://github.com/nwtgck/gh-card).  
The URL for image triggers the server to create a new SVG image, so the card displayed is up to date.  
## How to use
Just put a link on your web site. The sample page is [here](https://gpc.znoo.xyz/).  
All you need is a normal URL, so this can not only be used for HTML but for Markdown, like this:  
[![kena0ki](https://gpc.znoo.xyz/api/github/kena0ki.svg)](https://github.com/kena0ki)
### Query parameters
Currently, following parameters are available   
 * width: width of the card
 * height: height of the card
## Development
* Gitpod: Click the above badge  
* Vim:  
```bash
make dev
```
* VS Code: Open this repository by VS Code (Remote Development Extension is needed)  
## This repository is inspired by
* [gh-card](https://github.com/nwtgck/gh-card)
* [GitHub Link Card Creator](https://github.com/po3rin/github_link_creator)
* [Unofficial GitHub Cards](https://github.com/lepture/github-cards)
## License
MIT  
