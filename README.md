# github-profile-card
[![Gitpod Ready-to-Code](https://img.shields.io/badge/Gitpod-Ready--to--Code-blue?logo=gitpod)](https://gitpod.io/#https://github.com/kena0ki/github-profile-card) 

Note: This repository is kind of like user profile version of [gh-card](https://github.com/nwtgck/gh-card).  
This is source code repository for web API which serves github profile card.  
The web API creates cards on each request, currently with no cache, so cards displayed on the client is always up to date.  
## Usage
### Endpoint
`https://gpc.znoo.xyz/api/github/:username`
### Query parameters
Currently, following parameters are available.  
 * width: width of the card
 * height: height of the card
### Sample
The sample page is [here](https://gpc.znoo.xyz/).  
All you need is a normal URL, so this can also be used for Markdown, like this:  
[![kena0ki](https://gpc.znoo.xyz/api/github/kena0ki.svg)](https://github.com/kena0ki)
## Development
* Gitpod: Click the above Gitpod badge  
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
