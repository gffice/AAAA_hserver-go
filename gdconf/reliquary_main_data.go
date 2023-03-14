package gdconf

import (
	"fmt"

	"hk4e/pkg/logger"

	"github.com/jszwec/csvutil"
)

// ReliquaryMainData 圣遗物主属性配置表
type ReliquaryMainData struct {
	MainPropId      int32 `csv:"MainPropId"`                // 主属性ID
	MainPropDepotId int32 `csv:"MainPropDepotId,omitempty"` // 主属性库ID
	PropType        int32 `csv:"PropType,omitempty"`        // 属性类别
	RandomWeight    int32 `csv:"RandomWeight,omitempty"`    // 随机权重
}

func (g *GameDataConfig) loadReliquaryMainData() {
	g.ReliquaryMainDataMap = make(map[int32]map[int32]*ReliquaryMainData)
	data := g.readCsvFileData("ReliquaryMainData.csv")
	var reliquaryMainDataList []*ReliquaryMainData
	err := csvutil.Unmarshal(data, &reliquaryMainDataList)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	for _, reliquaryMainData := range reliquaryMainDataList {
		// 通过主属性库ID找到
		_, ok := g.ReliquaryMainDataMap[reliquaryMainData.MainPropDepotId]
		if !ok {
			g.ReliquaryMainDataMap[reliquaryMainData.MainPropDepotId] = make(map[int32]*ReliquaryMainData)
		}
		// list -> map
		g.ReliquaryMainDataMap[reliquaryMainData.MainPropDepotId][reliquaryMainData.MainPropId] = reliquaryMainData
	}
	logger.Info("ReliquaryMainData count: %v", len(g.ReliquaryMainDataMap))
}

func GetReliquaryMainDataByDepotIdAndPropId(mainPropDepotId int32, mainPropId int32) *ReliquaryMainData {
	value, exist := CONF.ReliquaryMainDataMap[mainPropDepotId]
	if !exist {
		return nil
	}
	return value[mainPropId]
}

func GetReliquaryMainDataMap() map[int32]map[int32]*ReliquaryMainData {
	return CONF.ReliquaryMainDataMap
}