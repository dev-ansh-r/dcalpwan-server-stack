package component

func (c *Component) initACME() error {
	if c.config.TLS.Source != "acme" && !c.config.TLS.ACME.Enable {
		return nil
	}
	var err error
	c.acme, err = c.config.TLS.ACME.Initialize()
	if err != nil {
		return err
	}
	c.web.Prefix("/.well-known/acme-challenge/").Handler(c.acme.HTTPHandler(nil))
	return nil
}
