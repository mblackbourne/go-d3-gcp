package plugin

import (
	"github.com/ipfs/go-ipfs/plugin"
	"github.com/ipfs/go-ipfs/repo/fsrepo"
)

var Plugins = []plugin.Plugin{
	&GCPPlugin{},
}

type GCPPlugin struct{}

func (gsUtilDs GCPPlugin) Name() string {
	return "GSUTIL-datastore-plugin"
}

func (gsUtilDs GCPPlugin) Version() string {
	return "0.0.1"
}

func (gsUtilDs GCPPlugin) Init(env *plugin.Environment) error {
	return nil
}

func (gsUtilDs GCPPlugin) DatastoreTypeName() string {
	return "gsUtilDs"
}

func (gsUtilDs GCPPlugin) DatastoreConfigParser() fsrepo.ConfigFromMap {
	return nil
}
