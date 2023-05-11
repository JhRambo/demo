package utils

import (
	pb "demo/grpc/proto/logs"
	"demo/grpc/tools"
	"fmt"
	"strconv"
	"time"
)

// 日志统计 原生SQL
func Logs(actionType pb.ActionType, uid int64, eid int64, spaceId int64, devId string) error {
	ym := time.Now().Format("2006-01") //当前年月
	sql := fmt.Sprintf("insert into `t_logs_%s` (`action_type`,`uid`,`eid`,`space_id`,`dev_id`) values(?,?,?,?,?)", ym)
	err := tools.DB.Exec(sql, actionType, uid, eid, spaceId, devId)
	return err.Error
}

// 创建表，每个月创建一张表，每月执行一次即可
func CreateTable() (string, error) {
	ym := ""
	year := time.Now().Format("2006")
	month := time.Now().Format("1")
	y, _ := strconv.Atoi(year)
	m, _ := strconv.Atoi(month)
	if m == 12 { //次年第一个月
		y += 1
		ym = fmt.Sprintf("%d-%02d", y, 1)
	} else { //当年下一个月
		m += 1
		ym = fmt.Sprintf("%d-%02d", y, m)
	}
	table := fmt.Sprintf("t_logs_%s", ym)
	sql := "CREATE TABLE `" + table + "` (" +
		"`id` int NOT NULL AUTO_INCREMENT," +
		"`action_type` int NOT NULL," +
		"`uid` bigint DEFAULT '0'," +
		"`eid` bigint DEFAULT '0'," +
		"`space_id` bigint NOT NULL DEFAULT '0'," +
		"`dev_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT ''," +
		"`create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP," +
		"`bak1` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL," +
		"PRIMARY KEY (`id`)" +
		") ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci"
	err := tools.DB.Exec(sql)
	return table, err.Error
}
