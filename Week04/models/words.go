package models

type Words struct {
	Id               int    `models:"int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '编号'" json:"id"`
	Word             string `models:"varchar(200) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL DEFAULT '' COMMENT '单词拼写'" json:"word"`
	Meaning          string `models:"varchar(2000) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL DEFAULT '' COMMENT '中文意思'" json:"meaning"`
	PhoneticSymbolEn string `models:"varchar(80) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL DEFAULT '' COMMENT '音标-英音'" json:"phonetic_symbol_en"`
	PhoneticSymbolUn string `models:"varchar(80) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL DEFAULT '' COMMENT '音标-美音'" json:"phonetic_symbol_un"`
	Partofspeech     string `models:"char(4) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL DEFAULT '' COMMENT '词性'" json:"partofspeesh"`
	Grade            int    `models:"tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '年级'" json:"grade"`
	Type             int    `models:"tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '单词类别，1大纲，2超纲，3阅读，4听力，5写作，6词组'" json:"type"`
	CreatedAt        string `models:"timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间'" json:"created_at"`
	UpdatedAt        string `models:"timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间'" json:"updatedd_at"`
}
