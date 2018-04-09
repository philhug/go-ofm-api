package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	ofmdb "github.com/philhug/go-ofm-api"
	"github.com/philhug/go-ofm-api/gen/restapi/operations/nodes"
	"log"
	"github.com/hashicorp/golang-lru"
	"github.com/philhug/go-ofm-api/gen/models"
	"strings"
)

func NewNodesGetNodeHandler(rt *ofmdb.Runtime) nodes.GetNodeHandler {
	cache, _ := lru.New(1024*1024*10)
	return &nodeGetNodeHandler{rt: rt, cache: cache}
}

type nodeGetNodeHandler struct {
	rt *ofmdb.Runtime
	cache *lru.Cache
}

func (d *nodeGetNodeHandler) Handle(params nodes.GetNodeParams, more interface{}) middleware.Responder {
	nodemodel := d.rt.ServiceInstances()

	if v, ok := d.cache.Get(params.ID); ok {
		node := v.((models.Node))
		return nodes.NewGetNodeOK().WithPayload(&node)
	}

	srvmodel := d.rt.Services()

	srvid, err := srvmodel.LoadServiceByName(params.Db)
	if err != nil {
		// TODO
		return nodes.NewSearchNodeOK()
	}

	node, err := nodemodel.LoadNode(srvid, params.ID)
	if err != nil {
		// TODO
		log.Fatal(err)
	}
	d.cache.Add(params.ID, *node)
	return nodes.NewGetNodeOK().WithPayload(node)
}


func NewNodesGetMultipleNodesHandler(rt *ofmdb.Runtime) nodes.GetMultipleNodesHandler {
	return &nodeGetMultipleNodesHandler{rt: rt}
}

type nodeGetMultipleNodesHandler struct {
	rt *ofmdb.Runtime
}

func (d *nodeGetMultipleNodesHandler) Handle(params nodes.GetMultipleNodesParams, more interface{}) middleware.Responder {
	srvmodel := d.rt.Services()

	srvid, err := srvmodel.LoadServiceByName(params.Db)
	if err != nil {
		// TODO
		return nodes.NewSearchNodeOK()
	}

	nodemodel := d.rt.ServiceInstances()

	var nodeList models.NodeList
	n := strings.Split(params.Nodes, ",")
	for _, x := range n {
		node, err := nodemodel.LoadNode(srvid, x)
		if err != nil {
			// TODO
			log.Fatal(err)
		}
		nodeList.Items = append(nodeList.Items, node)
	}
	return nodes.NewGetMultipleNodesOK().WithPayload(&nodeList)
}

func NewNodeSearchNodeHandler(rt *ofmdb.Runtime) nodes.SearchNodeHandler {
	return &searchNodeHandler{rt: rt}
}

type searchNodeHandler struct {
	rt *ofmdb.Runtime
}

func (d *searchNodeHandler) Handle(params nodes.SearchNodeParams, more interface{}) middleware.Responder {
	srvmodel := d.rt.Services()

	srvid, err := srvmodel.LoadServiceByName(params.Db)
	if err != nil {
		// TODO
		return nodes.NewSearchNodeOK()
	}

	nodemodel := d.rt.ServiceInstances()

	node, err := nodemodel.QueryNode(srvid, params.Query, params.Deleted, params.Bbox, params.Fulltext)
	if err != nil {
		// TODO
		log.Fatal(err)
	}
	return nodes.NewSearchNodeOK().WithPayload(node)
}

