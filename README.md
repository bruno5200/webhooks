# webhooks with go

#### Library webhooks allows for easy receiving and parsing of GitHub, Bitbucket and GitLab Webhook Events

#### Features:

* Parses the entire payload, not just a few fields.
* Fields + Schema directly lines up with webhook posted json

#### Notes:

* Currently only accepting json payloads.

### Installation


Use go get.


```shell
go get -u github.com/bruno5200/webhooks/v2
```

Then import the package into your own code.

	import "github.com/bruno5200/webhooks/v2"


##### Examples:
```go
package main

import (
	"fmt"

	"net/http"

	"github.com/bruno5200/webhooks/v2/github"
)

const (
	path = "/webhooks"
)

func main() {
	hook, _ := github.New(github.Options.Secret("MyGitHubSuperSecretSecrect...?"))

	http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		payload, err := hook.Parse(r, github.ReleaseEvent, github.PullRequestEvent)
		if err != nil {
			if err == github.ErrEventNotFound {
				// ok event wasn;t one of the ones asked to be parsed
			}
		}
		switch payload.(type) {

		case github.ReleasePayload:
			release := payload.(github.ReleasePayload)
			// Do whatever you want from here...
			fmt.Printf("%+v", release)

		case github.PullRequestPayload:
			pullRequest := payload.(github.PullRequestPayload)
			// Do whatever you want from here...
			fmt.Printf("%+v", pullRequest)
		}
	})
	http.ListenAndServe(":3000", nil)
}
```

##### License: MIT
