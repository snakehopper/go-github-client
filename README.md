go-github-client: Github v3 API Client for Go
=============================================

**githib.com/alcacoop/go-github-client** is a package to simplify interaction
with the Github Rest API.

Quickstart Examples
-------------------

### Create an authenticated Github API Client using Basic Authentication method

```
import (
  ghclient "github.com/alcacoop/go-github-client/client"
)

...

ghc, _ := ghclient.NewGithubClient("testuser", "password", AUTH_USER_PASSWORD)

```  

### Get authenticated user info 

```
import (
  ghusers "github.com/alcacoop/go-github-client/users"
)

usersc = ghusers.NewUsers(ghc)
res, err := usersc.GetUserInfo(*userinfo)
jr, err := res.Json()
name := jr.GetString("login")

```

### Paginate Public Gists 

```
import (
  ghgist "github.com/alcacoop/go-github-client/gists"
)

gistsc = ghgists.NewUsers(ghc)
res, err := gistsc.GetPublicGists()
jr, err := res.Json()

res2, err := res.NextPage()
jr2, err := res.Json()
```

More Info
-------------

### API Reference
* [Base Client](http://gopkgdoc.appspot.com/pkg/github.com/alcacoop/go-github-client/client)
* [Users](http://gopkgdoc.appspot.com/pkg/github.com/alcacoop/go-github-client/users)
* [Gists](http://gopkgdoc.appspot.com/pkg/github.com/alcacoop/go-github-client/gists)

### Examples
* [Usage Examples](http:///github.com/alcacoop/go-github-client/examples)

### Dev Updates

We've extracted and refactored this package from [cocode.io](http://cocode.io), 
our real-time web collaborative code editor integrated with Github Gists, and we're enhancing our github integrations so we'll add more feature to this package,
to keep you updated on future developments, please feel free to follow us or 
give us some feedback on:

* [@cocodeio (twitter)](https://twitter.com/#!/cocodeio)
* [+cocodeio (google+)](https://plus.google.com/110953439702828767840/posts)