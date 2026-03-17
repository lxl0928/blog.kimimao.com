package models

// Response 统一响应结构
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// PageResult 分页结果
type PageResult struct {
	List  interface{} `json:"list"`
	Total int64       `json:"total"`
	Page  int         `json:"page"`
	Size  int         `json:"size"`
}

// StatsResult 统计结果
type StatsResult struct {
	TotalCount    int64       `json:"total_count"`
	TodayNewCount int64       `json:"today_new_count"`
	TodayActive   int64       `json:"today_active_count,omitempty"`
	RankList      interface{} `json:"rank_list,omitempty"`
}
