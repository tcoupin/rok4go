package backend

import (
	"errors"
	"fmt"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/tcoupin/rok4go/objects"
	"github.com/tcoupin/rok4go/utils/log"
)

const MONGO_COLLECTION = "config"
const GLOBALCONFIG_TYPE = "GlobalConfig"
const NAME_TITLE = "title"
const NAME_KEYWORDS = "keywords"

type MongoDB struct {
	session *mgo.Session
}

func (m *MongoDB) Init(urlstr interface{}) error {
	log.DEBUG("Init MongoDB backend storage: %s", urlstr)
	v, ok := urlstr.(string)
	if !ok {
		return errors.New("Can't parse mongodb url")
	}

	dinfo, err := mgo.ParseURL(v)
	if err != nil {
		return errors.New(fmt.Sprintf("Can't parse mongodb url, %v", err))
	}

	session, err := mgo.DialWithInfo(dinfo)

	m.session = session
	return err
}

func (m *MongoDB) GetGlobalConfig(c *objects.GlobalConfig) error {
	log.DEBUG("Load globalconfig from backend")
	err := m.session.DB("").C(MONGO_COLLECTION).Find(bson.M{"type": GLOBALCONFIG_TYPE}).One(c)
	log.TRACE("%v", c)
	return err
}

func (m *MongoDB) SetGlobalConfig(c *objects.GlobalConfig) error {
	log.DEBUG("Update globalconfig to backend")
	log.TRACE("%v", c)
	_, err := m.session.DB("").C(MONGO_COLLECTION).Upsert(bson.M{"type": GLOBALCONFIG_TYPE}, bson.M{"$set": c})
	return err
}
