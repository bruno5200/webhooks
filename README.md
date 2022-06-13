# webhooks with go

#### Library webhooks allows for easy receiving and parsing of GitHub, Bitbucket and GitLab Webhook Events

### Installation

Use go get.

```shell
go get -u github.com/bruno5200/webhooks/v2
```

Then import the package into your own code.

	import "github.com/bruno5200/webhooks"

##### Examples:
```go
package main

import (
	"fmt"

	"net/http"

	"github.com/bruno5200/webhooks/github"
)

const (
	path = "/webhooks"
)

func main() {
	hook, _ := github.New(github.Options.Secret("[Your_Secret]"))

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
