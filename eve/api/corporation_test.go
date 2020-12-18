package api

import (
	"context"
	"github.com/grzeng-go/eve-helper/eve/model"
	"reflect"
	"testing"
	"time"
)

func TestGetCorporationInfo(t *testing.T) {
	type args struct {
		ctx context.Context
		id  int
	}
	tests := []struct {
		name    string
		args    args
		want    *model.Corporation
		wantErr bool
	}{
		{name: "test", args: struct {
			ctx context.Context
			id  int
		}{ctx: context.Background(), id: 98021437}, want: &model.Corporation{
			AllianceId:    403524217,
			CeoId:         410614068,
			CreatorId:     102149490,
			DateFounded:   time.Unix(1369565996, 0).UTC(),
			Description:   "<font size=\"12\" color=\"#bfffffff\"></font><font size=\"16\" color=\"#fffaac58\"><b>只</font><font size=\"1\" color=\"#fff7d358\"> </font><font size=\"16\" color=\"#fff7d358\">有</font><font size=\"1\" color=\"#fff4fa58\"> </font><font size=\"16\" color=\"#fff4fa58\">强</font><font size=\"1\" color=\"#ffacfa58\"> </font><font size=\"16\" color=\"#ffacfa58\">大</font><font size=\"1\" color=\"#ff82fa58\"> </font><font size=\"16\" color=\"#ff82fa58\">的</font><font size=\"1\" color=\"#ff58fa58\"> </font><font size=\"16\" color=\"#ff58fa58\">人</font><font size=\"1\" color=\"#ff58fa82\"> </font><font size=\"16\" color=\"#ff58fa82\">民</font><font size=\"1\" color=\"#ff58faac\"> </font><font size=\"16\" color=\"#ff58faac\">，</font><font size=\"1\" color=\"#ff58fad0\"> </font><font size=\"16\" color=\"#ff58fad0\">才</font><font size=\"1\" color=\"#ff58faf4\"> </font><font size=\"16\" color=\"#ff58faf4\">是</font><font size=\"1\" color=\"#ff58d3f7\"> </font><font size=\"16\" color=\"#ff58d3f7\">自</font><font size=\"1\" color=\"#ff58acfa\"> </font><font size=\"16\" color=\"#ff58acfa\">由</font><font size=\"1\" color=\"#ff5882fa\"> </font><font size=\"16\" color=\"#ff5882fa\">的</font><font size=\"1\" color=\"#ff5858fa\"> </font><font size=\"16\" color=\"#ff5858fa\">人</font><font size=\"1\" color=\"#ff8258fa\"> </font><font size=\"16\" color=\"#ff8258fa\">民</font><font size=\"1\" color=\"#ffac58fa\"> </font><font size=\"16\" color=\"#ffac58fa\">。</b></font>",
			HomeStationId: 60014944,
			MemberCount:   15,
			Name:          "PIBC执行军团",
			Shares:        10000000,
			TaxRate:       0,
			Ticker:        "PIBC执",
			Url:           "https://www.mango-eve.online/",
			WarEligible:   true,
		}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetCorporationInfo(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCorporationInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCorporationInfo() got = %v, want %v", got, tt.want)
			} else {
				t.Logf("GetCorporationInfo() got = %v", got)
			}
		})
	}
}

func TestGetNpcCorps(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "test", args: struct{ ctx context.Context }{ctx: context.Background()}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetNpcCorps(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetNpcCorps() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("GetNpcCorps() got = %v", got)
		})
	}
}
