package api

type (
	// 祖先
	Ancestry struct {
		BloodlineId      int    `json:"bloodline_id"`      // 血统ID
		Description      string `json:"description"`       // 描述
		Id               int    `json:"id"`                // 祖先ID
		Name             string `json:"name"`              // 名称
		ShortDescription string `json:"short_description"` // 简介
	}

	// 小行星带信息
	AsteroidBelt struct {
		name      string   `json:"name"`      // 名称
		position  Position `json:"position"`  // 坐标
		system_id int      `json:"system_id"` // 所属太阳系ID
	}

	// 坐标
	Position struct {
		X int `json:"x"`
		Y int `json:"y"`
		Z int `json:"z"`
	}

	// 血统
	Bloodline struct {
		BloodlineId   int    `json:"bloodline_id"`   // 血统ID
		CorporationId int    `json:"corporation_id"` // 公司ID
		RaceId        int    `json:"race_id"`        // 种族ID
		ShipTypeId    int    `json:"ship_type_id"`   // 舰船类型ID
		Name          string `json:"name"`           // 名称
		Description   string `json:"description"`    // 描述
		Intelligence  int    `json:"intelligence"`   // 情报
		Memory        int    `json:"memory"`         // 记忆
		Perception    int    `json:"perception"`     // 感知
		Charisma      int    `json:"charisma"`       // 魅力
		Willpower     int    `json:"willpower"`      // 意志力
	}
)
