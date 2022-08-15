package def

import "go.mongodb.org/mongo-driver/bson/primitive"

// 垃圾桶设备信息
type TrashCanInfo struct {
	DeviceId       int    `json:"deviceId" gorm:"type:int"`               // 设备id
	DeviceName     string `json:"deviceName" gorm:"type:varchar(50)"`     // 设备名称
	DeviceType     int    `json:"deviceType" gorm:"type:int"`             // 设备类型id
	DeviceTypeName string `json:"deviceTypeName" gorm:"type:varchar(50)"` // 设备类型名称
	UnitName       string `json:"unitName" gorm:"type:varchar(20)"`       // 所在小区
	CreateTime     int64  `json:"createTime" gorm:"type:int"`             // 创建时间戳,ms
	MainBoardId    string `json:"mainBoardId" gorm:"type:varchar(50)"`    // 主板编码
	DevCode        string `json:"devCode" gorm:"type:varchar(50)"`        // 主板编码
	GaoDe          string `json:"gaode" gorm:"type:varchar(50)"`          // 高德GPS（以此为主）
	Gps            string `json:"gps" gorm:"type:varchar(50)"`            // GPS（原生定位，有偏差）
	Address        string `json:"address" gorm:"type:varchar(50)"`        // 设备地址
}

// 垃圾桶设备状态，半小时一次
type TrashCanState struct {
	DeviceId    int    `json:"device_id" gorm:"type:int"`           // 设备id
	DeviceNo    int    `json:"device_no" gorm:"type:int"`           // 内桶序号
	Type        int    `json:"type" gorm:"type:int"`                // 内桶类型编码
	DevCode     string `json:"devCode" gorm:"type:varchar(50)"`     // 主板编码
	Capacity    int    `json:"capacity" gorm:"type:int"`            // 当前容量
	CreateTime  int64  `json:"createTime" gorm:"type:int"`          // 上报时间
	Temperature int    `json:"temperature" gorm:"type:int"`         // 当前温度
	Weight      int    `json:"weight" gorm:"type:int"`              // 当前重量
	WaterLine   int    `json:"water_line" gorm:"type:int"`          // 水位（仅限特定款式）
	Remark      int    `json:"remark" gorm:"type:int"`              // 备注
	MainBoardId string `json:"mainBoardId" gorm:"type:varchar(50)"` // 主板编码
}

// 设备故障，半小时一次
type TrashCanFault struct {
	Id          primitive.ObjectID `bson:"_id,omitempty"`                       // 主键id
	CreateTime  int64              `json:"createTime" gorm:"type:int"`          // 上报时间
	DevCode     string             `json:"devCode" gorm:"type:varchar(50)"`     // 主板编码
	GarbageType string             `json:"garbageType" gorm:"type:varchar(50)"` // 内桶类型名称
	WarnType    int                `json:"warnType" gorm:"type:int"`            // 异常编码
	WarnName    string             `json:"warnName" gorm:"type:varchar(50)"`    // 异常原因
	DeviceNo    int                `json:"device_no" gorm:"type:int"`           // 内桶序号
	Type        int                `json:"type" gorm:"type:int"`                // 内桶类型编码
	DeviceId    int                `json:"deviceId" gorm:"type:int"`            // 设备id
	MainBoardId string             `json:"mainBoardId" gorm:"type:varchar(50)"` // 主板编码
	Handle      bool               `gorm:"type:bool"`                           // 是否处理，处理设置为true，默认未处理为false
}

// 居民积分，每天一次
type TrashCanScore struct {
	Id           int    `json:"id" gorm:"type:int"`                    // 居民id
	Score        int    `json:"score" gorm:"type:int"`                 // 用户当前账户积分
	Mobile       string `json:"mobile" gorm:"type:varchar(50)"`        // 用户手机号
	Username     string `json:"username" gorm:"type:varchar(50)"`      // 用户名
	Nick         string `json:"nick" gorm:"type:varchar(50)"`          // 用户昵称
	UnitName     string `json:"unitName" gorm:"type:varchar(50)"`      // 所在小区
	SerialNumber string `json:"serial_number" gorm:"type:varchar(50)"` // 随机卡号
	Address      string `json:"address" gorm:"type:varchar(50)"`       // 地址
}

// 投递记录，每天一次
type TrashCanDelivery struct {
	Id             primitive.ObjectID `bson:"_id,omitempty"`                          // 主键id
	DeviceId       int                `json:"deviceId" gorm:"type:int"`               // 设备id
	DeviceName     string             `json:"deviceName" gorm:"type:varchar(50)"`     // 设备名称
	Weight         int                `json:"weight" gorm:"type:int"`                 // 投递重量
	Type           int                `json:"type" gorm:"type:int"`                   // 垃圾类型编码
	CustomerId     int                `json:"customerId" gorm:"type:int"`             // 用户id
	Score          int                `json:"score" gorm:"type:int"`                  // 获取积分
	UnitName       string             `json:"unitName" gorm:"type:varchar(50)"`       // 小区名称
	DeviceNo       int                `json:"device_no" gorm:"type:int"`              // 内桶序号
	GarbageType    string             `json:"garbageType" gorm:"type:varchar(50)"`    // 垃圾类型名称
	CreateTime     int64              `json:"createTime" gorm:"type:int"`             // 上报时间
	CheckStatus    int                `json:"checkStatus" gorm:"type:int"`            // 巡检状态（0：未巡检、1：已通过、2：不通过、3：已清运）
	MainBoardId    string             `json:"mainBoardId" gorm:"type:varchar(50)"`    // 主板编码
	DevCode        string             `json:"devCode" gorm:"type:varchar(50)"`        // 主板编码
	CustomerMobile string             `json:"customerMobile" gorm:"type:varchar(50)"` // 用户手机号
}

// 兑换记录，每天一次
type TrashCanExchange struct {
	Id          int    `gorm:"primaryKey"`                          // 主键id
	DeviceId    int    `json:"deviceId" gorm:"type:int"`            // 设备id
	DeviceName  string `json:"deviceName" gorm:"type:varchar(50)"`  // 设备名称
	DevCode     string `json:"devCode" gorm:"type:varchar(50)"`     // 主板编码
	Type        int    `json:"type" gorm:"type:int"`                // 0=领取、1=兑换
	CustomerId  int    `json:"customerId" gorm:"type:int"`          // 用户id（手机号）
	Score       int    `json:"score" gorm:"type:int"`               // 兑换所需积分
	UnitName    string `json:"unitName" gorm:"type:varchar(50)"`    // 小区名称
	GoodsName   string `json:"goodsName" gorm:"type:varchar(50)"`   // 商品名称
	CreateTime  int64  `json:"createTime" gorm:"type:int"`          // 兑换时间
	Mobile      string `json:"mobile" gorm:"type:varchar(50)"`      // 用户手机号
	MainBoardId string `json:"mainBoardId" gorm:"type:varchar(50)"` // 主板编码
}

type Unit struct {
	UnitName string `json:"unit_name" gorm:"type:varchar(50)"` // 小区名称
	UnitCode string `json:"unitCode" gorm:"type:varchar(50)"`  // 小区编号
	AreaName string `json:"areaName" gorm:"type:varchar(50)"`  // 小区所在区域
	Phone    string `json:"phone" gorm:"type:varchar(50)"`     // 电话
	UserName string `json:"userName" gorm:"type:varchar(50)"`  // 用户名
	UnitType string `json:"unitType" gorm:"type:varchar(50)"`  // 小区类型
}

type UnitList struct {
	Units []Unit `json:"unit"`
}

// 小区数据，一星期一次
type UnitData struct {
	ErrCode int      `json:"errcode"`
	ErrMsg  string   `json:"errmsg"`
	Data    UnitList `json:"data"`
}

// 集中回收数据
type Recycle struct {
	CustomerMobile string `json:"customerMobile" gorm:"type:varchar(50)"` // 手机号
	DeviceId       int    `json:"deviceId" gorm:"type:int"`               // 设备id
	Score          int    `json:"score" gorm:"type:int"`                  // 积分
	GarbageType    string `json:"garbageType" gorm:"type:varchar(50)"`    // 垃圾类型
	CreateTime     int64  `json:"createTime" gorm:"type:int"`             // 时间
	Status         int    `json:"status" gorm:"type:int"`                 // 状态，1通过，2不通过
	UnitName       string `json:"unitName" gorm:"type:varchar(50)"`       // 小区名称
	Weight         int    `json:"weight" gorm:"type:int"`                 // 重量
	CustomerId     int    `json:"customerId" gorm:"type:int"`             // 居民id
	DeviceName     string `json:"deviceName" gorm:"type:varchar(50)"`     // 安卓称名称
}
