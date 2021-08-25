package interfaces

import "log"

//Interface
type Logic interface {

	//Declare Function Process
	Process(data string) string
}

//LogicProvider that implements interfaces
type LogicProvider struct {}

// LogicProvider's method that implememtns the interface's function
func (lp LogicProvider) Process(data string) string {
	log.Printf("Logic Provider's Process method: implicitly implemented Logic's interface")
	return data
}


// Client struct that embeds Logic type
type Client struct{
	L Logic
}

// NewClient that creates Client instance with LogicProvider that implements Logic
func NewClient() *Client {
	l := &LogicProvider{}
	log.Printf("NewClient L: %T",l)
	return &Client{L: *l}
}

func(c Client) Program(data string) string{
	// get data from somewhere
	return c.L.Process(data)
}