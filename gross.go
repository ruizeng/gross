package gross

// the Gross Object
type Gross struct {
	hook *WebHook // webhook to serve, ie. "gitlab", "github"
	// rule    BuildRule     // the automatic building rule
	// handler TargetHandler // the handler to handle built targets
	// watcher *Watcher    // fs watcher to watch target folder change
}

// the Gross init config options
// webhook : webhook to serve, ie. "gitlab", "github"
// addr: the listening http service address
// rule: the automatic building rule
// folder: the watch folder where the build container put the target
// handler: the handler to handle built targets
type Config struct {
	HookName string
	Addr     string
	Rule     BuildRuler
	Folder   string
	Handler  TargetHandler
}

// create a new Gross with config
func New(conf Config) (*Gross, error) {
	// w := &Watcher{folder}
	h, err := NewWebHook(conf.HookName, conf.Addr, nil)
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
