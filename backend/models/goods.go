package models

// GoodsStatus 货物状态
type GoodsStatus string

const (
	StatusInStock   GoodsStatus = "in_stock"  // 在库
	StatusOutStock  GoodsStatus = "out_stock" // 已出库
	StatusMortgaged GoodsStatus = "mortgaged" // 抵押中
)

// GoodsType 货物类型
type GoodsType string

const (
	TypeShrimp    GoodsType = "shrimp"    // 冷冻虾类
	TypeFish      GoodsType = "fish"      // 冷冻鱼类
	TypeShellfish GoodsType = "shellfish" // 冷冻贝类
	TypeCrab      GoodsType = "crab"      // 冷冻蟹类
)

// Goods 货物信息
type Goods struct {
	ID            string  `json:"goodsId"`       // 批次编号
	DeviceID      string  `json:"deviceId"`      // 设备标识
	Name          string  `json:"name"`          // 海鲜名称
	Type          string  `json:"type"`          // 类型
	Specification string  `json:"specification"` // 等级
	Size          string  `json:"size"`          // 包装
	Weight        float64 `json:"weight"`        // 重量(吨)
	Temperature   float64 `json:"temperature"`   // 温度(°C)
	Location      string  `json:"location"`      // 存储位置
	InTime        string  `json:"inTime"`        // 入库时间
	Status        string  `json:"status"`        // 状态
}

// GoodsFilter 货物查询过滤条件
type GoodsFilter struct {
	GoodsID string      `json:"goodsId"`
	Type    GoodsType   `json:"type"`
	Status  GoodsStatus `json:"status"`
}

// GoodsListResponse 货物列表响应
type GoodsListResponse struct {
	Total int64   `json:"total"`
	Items []Goods `json:"items"`
}
