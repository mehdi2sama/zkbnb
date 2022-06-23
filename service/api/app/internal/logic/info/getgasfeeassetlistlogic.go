package info

import (
	"context"

	table "github.com/zecrey-labs/zecrey-legend/common/model/assetInfo"
	"github.com/zecrey-labs/zecrey-legend/service/api/app/internal/repo/l2asset"
	"github.com/zecrey-labs/zecrey-legend/service/api/app/internal/svc"
	"github.com/zecrey-labs/zecrey-legend/service/api/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGasFeeAssetListLogic struct {
	logx.Logger
	ctx     context.Context
	svcCtx  *svc.ServiceContext
	l2asset l2asset.L2asset
}

func NewGetGasFeeAssetListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGasFeeAssetListLogic {
	return &GetGasFeeAssetListLogic{
		Logger:  logx.WithContext(ctx),
		ctx:     ctx,
		svcCtx:  svcCtx,
		l2asset: l2asset.New(svcCtx),
	}
}

func (l *GetGasFeeAssetListLogic) GetGasFeeAssetList(req *types.ReqGetGasFeeAssetList) (*types.RespGetGasFeeAssetList, error) {
	assets, err := l.l2asset.GetL2AssetsList()
	if err != nil {
		logx.Errorf("[GetL2AssetsList] err:%v", err)
		return nil, err
	}
	resp := &types.RespGetGasFeeAssetList{
		Assets: make([]types.AssetInfo, 0),
	}
	for _, asset := range assets {
		if asset.IsGasAsset != table.IsGasAsset {
			continue
		}
		resp.Assets = append(resp.Assets, types.AssetInfo{
			AssetId:       asset.AssetId,
			AssetName:     asset.AssetName,
			AssetDecimals: asset.Decimals,
			AssetSymbol:   asset.AssetSymbol,
			AssetAddress:  asset.L1Address,
			IsGasAsset:    asset.IsGasAsset,
		})
	}
	return resp, nil
}