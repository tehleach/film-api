package db

import "gopkg.in/mgo.v2"

//Conn is database connection
type Conn struct {
	Session *mgo.Session
}

//NewConn gets a new Conn object
func NewConn(host string) (*Conn, error) {
	session, err := mgo.Dial(host)
	conn := Conn{session}
	return &conn, err
}
