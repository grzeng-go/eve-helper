package model

import "time"

type (
	// 联盟
	Alliance struct {
		CreatorCorporationId  int       `json:"creator_corporation_id"`  // 创建公司ID
		CreatorId             int       `json:"creator_id"`              // 创建人ID
		DateFounded           time.Time `json:"date_founded"`            // 创建时间
		ExecutorCorporationId int       `json:"executor_corporation_id"` // 执行者公司ID
		Name                  string    `json:"name"`                    // 联盟名称
		Ticker                string    `json:"ticker"`                  // 股票代码
	}

	// 军团（公司）
	Corporation struct {
		AllianceId    int       `json:"alliance_id"`		// 联盟ID
		CeoId         int       `json:"ceo_id"`				// ceoID
		CreatorId     int       `json:"creator_id"`			// 创建者ID
		DateFounded   time.Time `json:"date_founded"`		// 创建时间
		Description   string    `json:"description"`		// 军团描述
		HomeStationId int       `json:"home_station_id"`	// 驻地ID
		MemberCount   int       `json:"member_count"`		// 成员数量
		Name          string    `json:"name"`				// 军团名称
		Shares        int       `json:"shares"`				// 股份
		TaxRate       float64   `json:"tax_rate"`			// 军团税率
		Ticker        string    `json:"ticker"`				// 股票代码
		Url           string    `json:"url"`				// url
		WarEligible   bool      `json:"war_eligible"`		// 是否有战争资格
	}
)
