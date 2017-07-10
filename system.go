package argo

/*// Multicall is not implemented, mainly because it is quite complicated (and the Caller.Call is probably not really adapted)
// Moreover it is weakly type (interface{}) which should not be encouraged to use
// see https://aria2.github.io/manual/en/html/aria2c.html#system.multicall
func (c Client) Multicall() (reply Ok, err error) {
	return Ok(""), errors.New("Not implemented")
}
*/

// ListMethods returns all the available RPC methods .
func (c Client) ListMethods() (reply []string, err error) {
	err = c.Caller.Call("system.listMethods", &reply)
	return
}

// ListNotifications returns all the available RPC notifications
func (c Client) ListNotifications() (reply []string, err error) {
	err = c.Caller.Call("system.listNotifications", &reply)
	return
}
