package api

import (
	"context"
	"fmt"
	"github.com/grzeng-go/eve-helper/eve/model"
	"github.com/guonaihong/gout"
)

// 获取所有联盟
func GetAlliances(ctx context.Context) ([]int, error) {
	var ids []int
	err := gout.GET(merge(baseUrl, "/alliances/")).
		WithContext(ctx).
		SetHeader(header).
		BindJSON(&ids).
		Do()

	return ids, err
}

// 获取联盟信息
func GetAllianceInfo(ctx context.Context, id int) (*model.Alliance, error) {
	var alliance *model.Alliance
	err := gout.GET(merge(baseUrl, "/alliances/", fmt.Sprint(id))).
		WithContext(ctx).
		SetHeader(header).
		BindJSON(&alliance).
		Do()

	return alliance, err
}

// 获取联盟中所有的公司
func GetAllianceCorporations(ctx context.Context, id int) ([]int, error) {
	var ids []int
	err := gout.GET(merge(baseUrl, "/alliances/", fmt.Sprint(id), "/corporations/")).
		WithContext(ctx).
		SetHeader(header).
		BindJSON(&ids).
		Do()

	return ids, err
}