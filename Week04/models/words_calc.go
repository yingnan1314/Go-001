package models

type WordsCalc struct {
	Id        int    `models:"int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '编号'" json:"id"`
	WordId    int    `models:"int(10) unsigned NOT NULL DEFAULT '0'" json:"word_id"`
	UserId    int    `models:"int(10) unsigned NOT NULL DEFAULT '0'" json:"user_id"`
	Active    int    `models:"tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态：0 默认，1 认识，不再出现，2 不认识'" json:"-"`
	Click     int    `models:"int(10) unsigned NOT NULL DEFAULT '0' COMMENT '计数器：点击次数'" json:"click"`
	CurrClick int    `models:"int(10) unsigned NOT NULL DEFAULT '0' COMMENT '当前题库点击数'" json:"curr_click"`
	CalcType  int    `models:"tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '题库类型：0首次记忆库，1常规库，2测试库，3复习库，4深度检测1，5深度检测2，6深度检测3'" json:"calc_type"`
	CreatedAt string `models:"timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间'" json:"created_at"`
	UpdatedAt string `models:"timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间'" json:"updated_at"`
}
