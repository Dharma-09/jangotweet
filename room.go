package main

import "net"

type room struct{
	name string
	members  map[net.ddr]*client
}

func(r *room) broadcaste(sender *client,msg string){
	for addr, m := range r.members{
		if addr != sender.conn.RemoteAddr(){
			m.nsg(msg)
		}
	}
}