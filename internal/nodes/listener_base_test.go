// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package nodes

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/nodeconfigs"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs"
	"github.com/iwind/TeaGo/types"
	"testing"
	"time"
)

func TestBaseListener_FindServer(t *testing.T) {
	sharedNodeConfig = &nodeconfigs.NodeConfig{}

	var listener = &BaseListener{namedServers: map[string]*NamedServer{}}
	listener.Group = &serverconfigs.ServerAddressGroup{}
	for i := 0; i < 1_000_000; i++ {
		var server = &serverconfigs.ServerConfig{
			IsOn: true,
			Name: types.String(i) + ".hello.com",
			ServerNames: []*serverconfigs.ServerNameConfig{
				{Name: types.String(i) + ".hello.com"},
			},
		}
		_ = server.Init()
		listener.Group.Servers = append(listener.Group.Servers, server)
	}

	var before = time.Now()
	defer func() {
		t.Log(time.Since(before).Seconds()*1000, "ms")
	}()

	t.Log(listener.findNamedServerMatched("855555.hello.com"))
}
