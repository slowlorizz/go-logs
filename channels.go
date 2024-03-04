package logs

type MessageChannel (chan *Message)

//----------------------------------------------------------------------------------------------//

type Channels struct {
	StdOut (chan *string)
	Debug  MessageChannel
	Info   MessageChannel
	Warn   MessageChannel
	Error  MessageChannel
	Fatal  MessageChannel
}

func (c *Channels) init() {
	c.StdOut = make(chan *string)  // Channel for normal Prints
	c.Debug = make(MessageChannel) // Channel for debug-Messages
	c.Info = make(MessageChannel)  // Channel for Info-Messages
	c.Warn = make(MessageChannel)  // Channel for Warning Messages
	c.Error = make(MessageChannel) // Channel for Error-Messages
	c.Fatal = make(MessageChannel) // Channel for Fatal-Error-Messages
}
