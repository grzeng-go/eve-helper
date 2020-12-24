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

	// 项目类别
	Category struct {
		CategoryId int    `json:"category_id"` // 项目类别ID
		Groups     []int  `json:"groups"`      //	项目组
		Name       string `json:"name"`        // 项目类别名称
		Published  bool   `json:"published"`   // 是否发布
	}

	// 星座
	Constellation struct {
		ConstellationId int      `json:"constellation_id"`
		Name            string   `json:"name"`
		Position        Position `json:"position"`  // 星座坐标
		RegionId        int      `json:"region_id"` // 星座所属区域
		Systems         []int    `json:"systems"`   // 太阳系ID
	}

	// 派系
	Faction struct {
		CorporationId        int    `json:"corporation_id"`         // 星座ID
		Description          string `json:"description"`            // 描述
		FactionId            int    `json:"faction_id"`             // 派系ID
		IsUnique             bool   `json:"is_unique"`              // 是否唯一
		Name                 string `json:"name"`                   // 名称
		SizeFactor           int    `json:"size_factor"`            // 因子  ？
		MilitiaCorporationId int    `json:"militia_corporation_id"` // 民兵公司ID ？
		SolarSystemId        int    `json:"solar_system_id"`        // 太阳系ID
		StationCount         int    `json:"station_count"`          // 空间站个数 ？
		StationSystemCount   int    `json:"station_system_count"`   // 空间站系统个数 ？
	}

	Graphic struct {
		CollisionFile string `json:"collision_file"`
		GraphicFile   string `json:"graphic_file"`
		GraphicId     int    `json:"graphic_id"`
		IconFolder    string `json:"icon_folder"`
		SofDna        string `json:"sof_dna"`
		SofFationName string `json:"sof_fation_name"`
		SofHullName   string `json:"sof_hull_name"`
		SofRaceName   string `json:"sof_race_name"`
	}

	// 组
	Group struct {
		CategoryId int `json:"category_id"`
		GroupId    int `json:"group_id"`
		Name       int `json:"name"`
		Published  int `json:"published"`
		Types      int `json:"types"`
	}

	// int `json:""`
)
