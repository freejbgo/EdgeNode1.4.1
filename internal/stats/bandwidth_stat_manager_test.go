// Copyright 2022 Liuxiangchao iwind.liu@gmail.com. All rights reserved. Official site: https://goedge.cn .

package stats_test

import (
	"github.com/TeaOSLab/EdgeNode/internal/stats"
	"testing"
	"time"
)

func TestBandwidthStatManager_Add(t *testing.T) {
	var manager = stats.NewBandwidthStatManager()
	manager.AddBandwidth(1, 1, 10, 10)
	manager.AddBandwidth(1, 1, 10, 10)
	manager.AddBandwidth(1, 1, 10, 10)
	time.Sleep(1 * time.Second)
	manager.AddBandwidth(1, 1, 85, 85)
	time.Sleep(1 * time.Second)
	manager.AddBandwidth(1, 1, 25, 25)
	manager.AddBandwidth(1, 1, 75, 75)
	manager.Inspect()
}

func TestBandwidthStatManager_Loop(t *testing.T) {
	var manager = stats.NewBandwidthStatManager()
	manager.AddBandwidth(1, 1, 10, 10)
	manager.AddBandwidth(1, 1, 10, 10)
	manager.AddBandwidth(1, 1, 10, 10)
	err := manager.Loop()
	if err != nil {
		t.Fatal(err)
	}
}
