package gross

// the Gross Object
type Gross struct {
	hook *WebHook // webhook to serve, ie. "gitlab", "github"
	// rule    BuildRule     // the automatic building rule
	// handler TargetHandler // the handler to handle built targets
	// watcher *Watcher    // fs watcher to watch target folder change
}

// create a new Gross with config
// webhook : webhook to serve, ie. "gitlab", "github"
// rule: the automatic building rule
// folder: the watch folder where the build container put the target
// handler: the handler to handle built targets
func New(hookname string, rule BuildRuler, folder string, handler TargetHandler) (*Gross, error) {
	// w := &Watcher{folder}
	h, err := NewWebHook(hookname, nil)
	if err != nil {
		return nil, err
	}

	return &Gross{
		hook: h,
		// Rule:    rule,
		// Handler: handler,
		// watcher: w,
		// deliverer: d,
	}, nil
}

// begin the service.
// the function will not return utils an error occured
func (gross *Gross) Serve() error {
	return nil
}
