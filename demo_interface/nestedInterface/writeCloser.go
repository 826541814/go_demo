package main

import "io"

type device struct{
}

func (d *device) Close() error{
	return nil
}

func main() {
	var c io.Closer = new(device)
	c.Close()
}
