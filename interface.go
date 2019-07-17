package gogm

import dsl "github.com/mindstand/go-cypherdsl"

type IEdge interface {
	GetStartNode() interface{}
	SetStartNode(v interface{}) error

	GetEndNode() interface{}
	SetEndNode(v interface{}) error
}

//inspiration from -- https://github.com/neo4j/neo4j-ogm/blob/master/core/src/main/java/org/neo4j/ogm/session/Session.java

//session object for ogm interactions
type ISession interface {
	//transaction functions
	ITransaction

	//load single object
	Load(respObj interface{}, id string) error

	//load object with depth
	LoadDepth(respObj interface{}, id string, depth int) error

	//load with depth and filter
	LoadDepthFilter(respObj interface{}, id string, depth int, filter *dsl.ConditionBuilder) error

	//load with depth, filter and pagination
	LoadDepthFilterPagination(respObj interface{}, id string, depth int, filter dsl.ConditionOperator, pagination *Pagination) error

	//load slice of something
	LoadAll(respObj interface{}) error

	//load all of depth
	LoadAllDepth(respObj interface{}, depth int) error

	//load all of type with depth and filter
	LoadAllDepthFilter(respObj interface{}, depth int, filter *dsl.ConditionBuilder) error

	//load all with depth, filter and pagination
	LoadAllDepthFilterPagination(respObj interface{}, depth int, filter *dsl.ConditionBuilder, pagination *Pagination) error

	//save object
	Save(saveObj interface{}) error

	//save object with depth
	SaveDepth(saveObj interface{}, depth int) error

	//delete
	Delete(deleteObj interface{}) error

	//specific query, only respond with single object
	Query(query string, properties map[string]interface{}, respObj interface{}) error

	//specify query, return slice of info
	QuerySlice(query string, properties map[string]interface{}, respObj interface{}) error

	//delete everything, this will literally delete everything
	PurgeDatabase() error
}

type ITransaction interface {
	Begin() error
	Rollback() error
	Commit() error
}