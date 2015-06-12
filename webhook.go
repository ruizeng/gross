package gross

// CVSHandler is a handler interface to handle cvs activity like push, tag, etc.
type CVSHandler interface {
	// event: name of the event, current suport "tag", "commit"
	// repo: the repo address
	// extra: extra info of the trigger event, may be ""
	HandleEvent(event string, repo string, extra string)
}

// WebHook is a web api adapter to servce web hook for some CVS's API like github.
type WebHook struct {
	Name    string     // name of the hook
	Handler CVSHandler // a CVS handler to handle cvs events
}

// create a new web hook with the given name.
func NewWebHook(name string, handler CVSHandler) (*WebHook, error) {
	return nil, nil
}
