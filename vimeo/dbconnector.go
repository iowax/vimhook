package vimeo

import (
	mgo "gopkg.in/mgo.v2"
)

//DBConnector handles the connections to the mongodb database
type DBConnector struct {
	session *mgo.Session
	usedDB  string
}

//NewDBConnector provides a DBConnector object after creating a mongodb session.
func NewDBConnector(url string) (*DBConnector, error) {
	session, err := mgo.Dial(url)

	if err != nil {
		return &DBConnector{session: session, usedDB: defaultDatabase}, nil
	}
	return nil, err
}

func (c DBConnector) setUsedDB(name string) {
	c.usedDB = name
}

//Insert a key value pair into the database
func (c DBConnector) Insert(collection string, docs ...interface{}) (bool, error) {
	col := c.session.DB(c.usedDB).C(collection)
	err := col.Insert(docs)
	if err != nil {
		return true, nil
	}
	return false, err
}
