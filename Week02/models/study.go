package models

type Study struct{
	Id     int    `models:"int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '编号'"`
	Title   string `models:"varchar(200) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL DEFAULT '' COMMENT '标题'"`
	Description string `models:"varchar(2000) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL DEFAULT '' COMMENT '详细信息'"`
}