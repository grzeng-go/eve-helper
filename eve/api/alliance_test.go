package api

import (
	"context"
	"github.com/grzeng-go/eve-helper/eve/model"
	"reflect"
	"testing"
	"time"
)

var bg = context.Background()

func TestGetAlliances(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{name: "test", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAlliances(context.Background())
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAlliances() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("GetAlliances() got = %v", got)
		})
	}
}

func TestGetAllianceInfo(t *testing.T) {
	type args struct {
		ctx context.Context
		id  int
	}
	tests := []struct {
		name    string
		args    args
		want    *model.Alliance
		wantErr bool
	}{
		{name: "test", args: struct {
			ctx context.Context
			id  int
		}{ctx: bg, id: 99000008}, want: &model.Alliance{
			CreatorCorporationId:  908942008,
			CreatorId:             844310979,
			DateFounded:           time.Unix(1340869200, 0).UTC(),
			ExecutorCorporationId: 908942008,
			Name:                  "蔷薇商业共同体",
			Ticker:                "R.E.C",
		}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAllianceInfo(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllianceInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllianceInfo() got = %v, want %v", got, tt.want)
			} else {
				t.Logf("GetAllianceInfo() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAllianceCorporations(t *testing.T) {
	type args struct {
		ctx context.Context
		id  int
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{name: "test", args: struct {
			ctx context.Context
			id  int
		}{ctx: bg, id: 99001406}, want: []int{
			98090055,
			98090057,
			98090058,
			98090060,
		}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAllianceCorporations(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllianceCorporations() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllianceCorporations() got = %v, want %v", got, tt.want)
			} else {
				t.Logf("GetAllianceInfo() got = %v, want %v", got, tt.want)
			}
		})
	}
}
