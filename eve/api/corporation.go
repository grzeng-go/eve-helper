package api

import (
	"context"
	"fmt"
	"github.com/grzeng-go/eve-helper/eve/model"
	"github.com/guonaihong/gout"
)

// 获取军团（公司）信息
func GetCorporationInfo(ctx context.Context, id int) (*model.Corporation, error) {
	var corp *model.Corporation
	err := gout.GET(merge(baseUrl, "/corporations/", fmt.Sprint(id))).
		WithContext(ctx).
		SetHeader(header).
		BindJSON(&corp).
		Do()

	return corp, err
}

// 获取NPC军团信息
func GetNpcCorps(ctx context.Context) ([]int, error) {
	var ids []int
	err := gout.GET(merge(baseUrl, "/corporations/npccorps/")).
		WithContext(ctx).
		SetHeader(header).
		BindJSON(&ids).
		Do()

	return ids, err
}