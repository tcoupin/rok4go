package backend

import (
	"errors"
	"fmt"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/tcoupin/rok4go/objects"
	"github.com/tcoupin/rok4go/utils/log"
)

const mongoCollection = "config"
const globalConfigType = "GlobalConfig"
const nameTitle = "title"
const nameKeywords = "keywords"

// MongoDB implements backend interface
type MongoDB struct {
	session *mgo.Session
}

// Init the MongoDB bakend connection
func (m *MongoDB) Init(urlstr interface{}) error {
	log.DEBUG("Init MongoDB backend storage: %s", urlstr)
	v, ok := urlstr.(string)
	if !ok {
		return errors.New("Can't parse mongodb url")
	}

	dinfo, err := mgo.ParseURL(v)
	if err != nil {
		return fmt.Errorf("Can't parse mongodb url, %v", err)
	}

	session, err := mgo.DialWithInfo(dinfo)

	m.session = session
	return err
}

// GetGlobalConfig loads GlobalConfig from MongoDB backend
func (m *MongoDB) GetGlobalConfig(c *objects.GlobalConfig) error {
	log.DEBUG("Load globalconfig from backend")
	err := m.session.DB("").C(mongoCollection).Find(bson.M{"type": globalConfigType}).One(c)
	log.TRACE("%v", c)
	return err
}

// SetGlobalConfig save GlobalConfig to MongoDB backend
func (m *MongoDB) SetGlobalConfig(c *objects.GlobalConfig) error {
	log.DEBUG("Update globalconfig to backend")
	log.TRACE("%v", c)
	_, err := m.session.DB("").C(mongoCollection).Upsert(bson.M{"type": globalConfigType}, bson.M{"$set": c})
	return err
}
