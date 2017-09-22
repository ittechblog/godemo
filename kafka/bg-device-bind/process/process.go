package process

import (
	"github.com/tidwall/gjson"
	"strings"
	"igen.com/bee/kafka/bg-device-bind/config"
	"github.com/astaxie/beego/logs"
	"time"
	"strconv"
)

var (
	count int
)

//处理更新数据库
func Process(jsonMsg string, processChan chan int) {
	//json str 转map
	dataType := gjson.Get(jsonMsg, "zi").String()
	commandType := gjson.Get(jsonMsg, "a").String()
	sn := gjson.Get(jsonMsg, "g").String()
	collectSn := gjson.Get(jsonMsg, "za").String()
	sensor := gjson.Get(jsonMsg, "b").String()
	ipPort := gjson.Get(jsonMsg, "zh").String()
	deviceId := gjson.Get(jsonMsg, "zf").String()
	deviceType := gjson.Get(jsonMsg, "a0").String()
	//接收数据的时间戳(秒) UTC
	uploadDate := gjson.Get(jsonMsg, "zd").String()
	uploadDate = ConvertUploadTime(uploadDate)
	var datetime = time.Now()
	datetime.Format("2006-01-02 15:04:05")
	//tx,_  := config.Db.Begin()
	switch dataType {
	//信息帧
	case "01":
		if (strings.EqualFold(commandType, "11") || strings.EqualFold(commandType, "21") ) {
			if len(deviceId) > 0 {
				//更新device.info
				//logs.Info("信息帧->%s更新device.info", deviceId)
				stmt, err := config.Db.Prepare("update device set info = ?,upload_date=?,update_date=? where device_id=?")
				//logs.Info("update device set info = %s,upload_date=%s,update_date=%s where device_id=%s\n", jsonMsg, uploadDate, datetime, deviceId)
				if (err != nil) {
					LogErr(err)
				}
				_, err = stmt.Exec(jsonMsg,datetime, datetime, deviceId)
				if (err != nil) {
					LogErr(err)
				}
				defer stmt.Close();
			}
		} else {
			//更新datalogger.extend
			//logs.Info("信息帧->%s更新datalogger.extend", collectSn)
			firmwareVersion := gjson.Get(jsonMsg, "h").String()
			dataLoggerSql := ""
			if len(firmwareVersion) > 0{
				dataLoggerSql = "update datalogger set command_type=?,extend = ?,upload_date=?,update_date=?,firmware_developer_version=? where sn =?"
			}else{
				dataLoggerSql = "update datalogger set command_type=?,extend = ?,upload_date=?,update_date=? where sn =?"
			}
			stmt, err := config.Db.Prepare(dataLoggerSql)
			//logs.Info("update datalogger set command_type=%s,extend = %s,upload_date=%s,update_date=%s where sn =%s\n", deviceType, jsonMsg, uploadDate, datetime, collectSn)
			if (err != nil) {
				LogErr(err)
			}
			if len(firmwareVersion) > 0{
				_, err = stmt.Exec(commandType, jsonMsg, datetime, datetime, firmwareVersion,collectSn)
			}else{
				_, err = stmt.Exec(commandType, jsonMsg, datetime, datetime, collectSn)
			}
			if (err != nil) {
				LogErr(err)
			}
			defer stmt.Close();
		}
	//数据帧
	case "02":
		//如果有数据就更新device.data，否则新增
		if len(deviceId) > 0{
			//logs.Info("数据帧->%s更新device.data->更新device_id", deviceId)
			rows, queryErr := config.Db.Query("select count(*) from device where device_id=" + deviceId)
			if (queryErr != nil) {
				LogErr(queryErr)
			}
			for rows.Next() {
				err := rows.Scan(&count)
				if err != nil {
					//log.Fatal(err)
				}
			}
			defer rows.Close()

			addr := gjson.Get(jsonMsg, "f").String()
			if (count > 0) {
				stmt, err := config.Db.Prepare("update device set sn = ?,datalogger_sn=?,sensor=?,type=?,ip=?,upload_date=?,update_date=?,addr=? where device_id=?")
				//logs.Info("update device set sn = %s,datalogger_sn=%s,sensor=%s,type=%s,ip=%s,data = %s,upload_date=%s,update_date=%s where device_id=%s\n", sn, collectSn, sensor, deviceType, ipPort, jsonMsg, uploadDate, datetime, deviceId)
				if (err != nil) {
					LogErr(err)
				}
				_, err = stmt.Exec(sn, collectSn, sensor, deviceType, ipPort, datetime, datetime,addr, deviceId)
				if (err != nil) {
					LogErr(err)
				}
				defer stmt.Close();
			} else {
				//新增device数据
				stmt, err := config.Db.Prepare("insert into device(device_id, sn,datalogger_sn,sensor,type,ip,upload_date,update_date,addr) values(?,?,?,?,?,?,?,?,?)")
				//logs.Info("insert into device(device_id, sn,datalogger_sn,sensor,type,ip,data,upload_date,update_date) values(%s,%s,%s,%s,%s,%s,%s,%s,%s)\n",deviceId, sn, collectSn, sensor, deviceType, ipPort, jsonMsg, uploadDate, datetime)
				if (err != nil) {
					LogErr(err)
				}
				_, err = stmt.Exec(deviceId, sn, collectSn, sensor, deviceType, ipPort, datetime, datetime,addr)
				if (err != nil) {
					LogErr(err)
				}
				defer stmt.Close();
			}

			//更新device_id
			idStmt, error := config.Db.Prepare("update device_id set type = ?,datalogger_sn=? where type is null and device_id=?")
			//logs.Info("update device_id set type = %s,datalogger_sn=%s where type is null and device_id=%s\n", deviceType, collectSn, deviceId)
			if (error != nil) {
				LogErr(error)
			}
			_, error = idStmt.Exec(deviceType, collectSn, deviceId)
			if (error != nil) {
				LogErr(error)
			}
			defer idStmt.Close();
		}
	default:
	}
	//tx.Commit()
	processChan <- 1
}

//打印error
func LogErr(error error) {
	logs.Info("error-----------------------------")
	logs.Info(error)
}

//转换为时间格式
func ConvertUploadTime(uploadDate string) string {
	uploadDateInt, _ := strconv.ParseInt(uploadDate, 10, 64)
	tm := time.Unix(uploadDateInt, 0)
	return tm.Format("2006-01-02 15:04:05")
}