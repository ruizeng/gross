package gross

import (
	"encoding/json"
	"strings"
)

// CVSEvent is a general struct to describe a cvs event
type CVSEvent struct {
	name   string
	repo   string
	branch string
	extra  string
}

// CVSHandler is a handler interface to handle cvs activity like push, tag, etc.
type CVSHandler interface {
	// repo: the repo name
	// branch: the branch name
	HandlePush(repo string, branch string)
	// repo: the repo name
	// branch: the branch name
	// extra: the tag desc string, usually looks like "xx.xx.xx"
	HandleTag(repo string, tag string)
}

type HTTPParser func(request string) (CVSEvent, error)

// WebHook is a web api adapter to servce web hook for some CVS's API like github.
type WebHook struct {
	Name    string     // name of the hook
	Addr    string     // adress of HTTP server
	Parser  HTTPParser // a web hook http request body parser
	Handler CVSHandler // a CVS handler to handle cvs events
}

// create a new web hook with the given name and listen address.
func NewWebHook(name string, addr string, handler CVSHandler) (*WebHook, error) {
	return nil, nil
}

// gitlab http request parser
func GitlabHTTPParser(request string) (CVSEvent, error) {
	var f interface{}
	err := json.Unmarshal([]byte(request), &f)
	if err != nil {
		return CVSEvent{}, err
	}
	ev := CVSEvent{}
	m := f.(map[string]interface{})
	repository := m["repository"].(map[string]interface{})
	ev.repo = repository["name"].(string)
	ref := m["ref"].(string)
	refs := strings.Split(ref, "/")
	if strings.Contains(ref, "tag") {
		ev.name = "tag"
		ev.extra = refs[len(refs)-1]
	} else {
		ev.name = "push"
		ev.branch = refs[len(refs)-1]
	}

	return ev, nil
}

// github http request parser
// TODO
