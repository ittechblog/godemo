package main

import (
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"igen.com/bee/thrift-hbase/hbase"
	"net"
	"os"
	"reflect"
	"time"
)

const (
	HOST       = "test40"
	PORT       = "9090"
	TESTRECORD = 10
)

func main() {
	startTime := currentTimeMillis()
	logformatstr_ := "----%s\n"
	logformatstr := "----%s 用时:%d-%d=%d毫秒\n\n"
	logformattitle := "建立连接"
	rowkey := "01_100101896_1499237802_08"
	temptable := "analysis:device_history_index"


	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	transport, err := thrift.NewTSocket(net.JoinHostPort(HOST, PORT))
	if err != nil {
		fmt.Fprintln(os.Stderr, "error resolving address:", err)
		os.Exit(1)
	}

	client := hbase.NewTHBaseServiceClientFactory(transport, protocolFactory)
	if err := transport.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to "+HOST+":"+PORT, " ", err)
		os.Exit(1)
	}
	tmpendTime := currentTimeMillis()
	fmt.Printf(logformatstr, logformattitle, tmpendTime, startTime, (tmpendTime - startTime))

	//--------------Exists
	//logformattitle = "调用Exists方法"
	//fmt.Printf(logformatstr_, logformattitle)
	//tmpstartTime := currentTimeMillis()
	////
	//isexists, err := client.Exists([]byte(temptable), &hbase.TGet{Row: []byte(rowkey)})
	//fmt.Printf("rowkey{%s} in table{%s} Exists:%t\t", rowkey, temptable, isexists)
	//if err != nil {
	//	fmt.Printf("Exists err:%s\n", err)
	//}
	//fmt.Println("")
	//tmpendTime = currentTimeMillis()
	//fmt.Printf(logformatstr, logformattitle, tmpendTime, tmpstartTime, (tmpendTime - tmpstartTime))

	////--------------put
	//logformattitle = "调用Put方法写数据"
	//fmt.Printf(logformatstr_, logformattitle)
	//tmpstartTime = currentTimeMillis()
	//cvarr := []*hbase.TColumnValue{
	//	&hbase.TColumnValue{
	//		Family:    []byte("name"),
	//		Qualifier: []byte("idoall.org"),
	//		Value:     []byte("welcome idoall.org")}}
	//temptput := hbase.TPut{Row: []byte(rowkey), ColumnValues: cvarr}
	//err = client.Put([]byte(temptable), &temptput)
	//if err != nil {
	//	fmt.Printf("Put err:%s\n", err)
	//} else {
	//	fmt.Println("Put done")
	//}
	//tmpendTime = currentTimeMillis()
	//fmt.Printf(logformatstr, logformattitle, tmpendTime, tmpstartTime, (tmpendTime - tmpstartTime))
	//
	////--------------put2
	//logformattitle = "调用Put2方法写数据"
	//fmt.Printf(logformatstr_, logformattitle)
	//tmpstartTime = currentTimeMillis()
	//cvarr2 := []*hbase.TColumnValue{
	//	&hbase.TColumnValue{
	//		Family:    []byte("name"),
	//		Qualifier: []byte("idoall.org"),
	//		Value:     []byte("welcome idoall.org")}}
	//temptput2 := hbase.TPut{Row: []byte("2"), ColumnValues: cvarr2}
	//err = client.Put([]byte(temptable), &temptput2)
	//if err != nil {
	//	fmt.Printf("Put err:%s\n", err)
	//} else {
	//	fmt.Println("Put done")
	//}
	//tmpendTime = currentTimeMillis()
	//fmt.Printf(logformatstr, logformattitle, tmpendTime, tmpstartTime, (tmpendTime - tmpstartTime))

	//------------Get---------------
	logformattitle = "调用Get方法获取新增加的数据"
	fmt.Printf(logformatstr_, logformattitle)
	tmpstartTime := currentTimeMillis()
	//
	result, err := client.Get([]byte(temptable), &hbase.TGet{Row: []byte(rowkey)})
	if err != nil {
		fmt.Printf("Get err:%s\n", err)

	} else {
		fmt.Println("Rowkey:" + string(result.Row))
		for _, cv := range result.ColumnValues {
			printscruct(cv)
		}
	}
	tmpendTime = currentTimeMillis()
	fmt.Printf(logformatstr, logformattitle, tmpendTime, tmpstartTime, (tmpendTime - tmpstartTime))

	////--------------put update
	//logformattitle = "调用Put update方法'修改'数据"
	//fmt.Printf(logformatstr_, logformattitle)
	//tmpstartTime = currentTimeMillis()
	//cvarr = []*hbase.TColumnValue{
	//	&hbase.TColumnValue{
	//		Family:    []byte("name"),
	//		Qualifier: []byte("idoall.org"),
	//		Value:     []byte("welcome idoall.org---update")}}
	//temptput = hbase.TPut{Row: []byte(rowkey), ColumnValues: cvarr}
	//err = client.Put([]byte(temptable), &temptput)
	//if err != nil {
	//	fmt.Printf("Put update err:%s\n", err)
	//} else {
	//	fmt.Println("Put update done")
	//}
	//tmpendTime = currentTimeMillis()
	//fmt.Printf(logformatstr, logformattitle, tmpendTime, tmpstartTime, (tmpendTime - tmpstartTime))
	//
	////------------Get  update---------------
	//logformattitle = "调用Get方法获取'修改'后的数据"
	//fmt.Printf(logformatstr_, logformattitle)
	//tmpstartTime = currentTimeMillis()
	////
	//result, err = (client.Get([]byte(temptable), &hbase.TGet{Row: []byte(rowkey)}))
	//if err != nil {
	//	fmt.Printf("Get update err:%s\n", err)
	//} else {
	//	fmt.Println("update Rowkey:" + string(result.Row))
	//	for _, cv := range result.ColumnValues {
	//		printscruct(cv)
	//	}
	//}
	//tmpendTime = currentTimeMillis()
	//fmt.Printf(logformatstr, logformattitle, tmpendTime, tmpstartTime, (tmpendTime - tmpstartTime))
	//
	////------------DeleteSingle------------
	//logformattitle = "调用DeleteSingle方法删除一条数据"
	//fmt.Printf(logformatstr_, logformattitle)
	//tmpstartTime = currentTimeMillis()
	//tdelete := hbase.TDelete{Row: []byte(rowkey)}
	//err = client.DeleteSingle([]byte(temptable), &tdelete)
	//if err != nil {
	//	fmt.Printf("DeleteSingle err:%s\n", err)
	//} else {
	//	fmt.Printf("DeleteSingel done\n")
	//}
	//
	//tmpendTime = currentTimeMillis()
	//fmt.Printf(logformatstr, logformattitle, tmpendTime, tmpstartTime, (tmpendTime - tmpstartTime))
	//
	////-------------PutMultiple----------------
	//logformattitle = "调用PutMultiple方法添加" + strconv.Itoa(TESTRECORD) + "条数据"
	//fmt.Printf(logformatstr_, logformattitle)
	//tmpstartTime = currentTimeMillis()
	//
	//var tputArr []*hbase.TPut
	//for i := 0; i < TESTRECORD; i++ {
	//	putrowkey := strconv.Itoa(i)
	//
	//	tputArr = append(tputArr, &hbase.TPut{
	//		Row: []byte(putrowkey),
	//		ColumnValues: []*hbase.TColumnValue{
	//			&hbase.TColumnValue{
	//				Family:    []byte("name"),
	//				Qualifier: []byte("idoall.org"),
	//				Value:     []byte(time.Now().String())}}})
	//}
	//
	//err = client.PutMultiple([]byte(temptable), tputArr)
	//if err != nil {
	//	fmt.Printf("PutMultiple err:%s\n", err)
	//} else {
	//	fmt.Printf("PutMultiple done\n")
	//}
	//
	//tmpendTime = currentTimeMillis()
	//fmt.Printf(logformatstr, logformattitle, tmpendTime, tmpstartTime, (tmpendTime - tmpstartTime))
	//
	////------------------GetMultiple-----------------------------
	//logformattitle = "调用GetMultiple方法获取" + strconv.Itoa(TESTRECORD) + "数据"
	//fmt.Printf(logformatstr_, logformattitle)
	//tmpstartTime = currentTimeMillis()
	////
	//var tgets []*hbase.TGet
	//for i := 0; i < TESTRECORD; i++ {
	//	putrowkey := strconv.Itoa(i)
	//	tgets = append(tgets, &hbase.TGet{
	//		Row: []byte(putrowkey)})
	//}
	//results, err := client.GetMultiple([]byte(temptable), tgets)
	//if err != nil {
	//	fmt.Printf("GetMultiple err:%s", err)
	//} else {
	//	fmt.Printf("GetMultiple Count:%d\n", len(results))
	//	for _, k := range results {
	//		fmt.Println("Rowkey:" + string(k.Row))
	//		for _, cv := range k.ColumnValues {
	//			printscruct(cv)
	//		}
	//	}
	//}
	//tmpendTime = currentTimeMillis()
	//fmt.Printf(logformatstr, logformattitle, tmpendTime, tmpstartTime, (tmpendTime - tmpstartTime))
	////-------------------TMutation
	////TMutation包含一个TGet一个TPut，就不做测试了
	////可以和MutateRow结合使用
	////
	//
	////-------------------OpenScanner
	//logformattitle = "调用OpenScanner方法"
	//fmt.Printf(logformatstr_, logformattitle)
	//tmpstartTime = currentTimeMillis()
	//startrow := make([]byte, 4)
	//binary.LittleEndian.PutUint32(startrow, 1)
	//stoprow := make([]byte, 4)
	//binary.LittleEndian.PutUint32(stoprow, 10)
	//
	//scanresultnum, err := client.OpenScanner([]byte(temptable), &hbase.TScan{
	//	StartRow: startrow,
	//	StopRow:  stoprow,
	//	//     FilterString: []byte("RowFilter(=, 'regexstring:00[1-3]00')"),
	//	//     FilterString: []byte("PrefixFilter('1407658495588-')"),
	//	Columns: []*hbase.TColumn{
	//		&hbase.TColumn{
	//			Family:    []byte("name"),
	//			Qualifier: []byte("idoall.org")}}})
	//if err != nil {
	//	fmt.Printf("OpenScanner err:%s\n", err)
	//} else {
	//	fmt.Printf("OpenScanner %d done\n", scanresultnum)
	//
	//	scanresult, err := client.GetScannerRows(scanresultnum, 100)
	//	if err != nil {
	//		fmt.Printf("GetScannerRows err:%s\n", err)
	//	} else {
	//		fmt.Printf("GetScannerRows %d done\n", len(scanresult))
	//		for _, k := range scanresult {
	//			fmt.Println("scan Rowkey:" + string(k.Row))
	//			for _, cv := range k.ColumnValues {
	//				printscruct(cv)
	//			}
	//		}
	//	}
	//}
	//tmpendTime = currentTimeMillis()
	//fmt.Printf(logformatstr, logformattitle, tmpendTime, tmpstartTime, (tmpendTime - tmpstartTime))
	//
	////--closescanner
	//logformattitle = "调用CloseScanner方法"
	//fmt.Printf(logformatstr_, logformattitle)
	//tmpstartTime = currentTimeMillis()
	//err = client.CloseScanner(scanresultnum)
	//if err != nil {
	//	fmt.Printf("CloseScanner err:%s\n", err)
	//}
	//tmpendTime = currentTimeMillis()
	//fmt.Printf(logformatstr, logformattitle, tmpendTime, tmpstartTime, (tmpendTime - tmpstartTime))
	//
	////-------------------GetScannerResults
	//
	//logformattitle = "调用GetScannerResults方法"
	//fmt.Printf(logformatstr_, logformattitle)
	//tmpstartTime = currentTimeMillis() //
	//gsr, err := client.GetScannerResults([]byte(temptable), &hbase.TScan{
	//	StartRow: startrow,
	//	StopRow:  stoprow,
	//	//     FilterString: []byte("RowFilter(=, 'regexstring:00[1-3]00')"),
	//	//     FilterString: []byte("PrefixFilter('1407658495588-')"),
	//	Columns: []*hbase.TColumn{
	//		&hbase.TColumn{
	//			Family:    []byte("name"),
	//			Qualifier: []byte("idoall.org")}}}, 100)
	//if err != nil {
	//	fmt.Printf("GetScannerResults err:%s\n", err)
	//} else {
	//	fmt.Printf("GetScannerResults %d done\n", len(gsr))
	//	for _, k := range gsr {
	//		fmt.Println("scan Rowkey:" + string(k.Row))
	//		for _, cv := range k.ColumnValues {
	//			printscruct(cv)
	//		}
	//	}
	//}
	//
	//tmpendTime = currentTimeMillis()
	//fmt.Printf(logformatstr, logformattitle, tmpendTime, tmpstartTime, (tmpendTime - tmpstartTime))
	//
	////---------------DeleteMultiple--------------
	//logformattitle = "调用DeleteMultiple方法删除" + strconv.Itoa(TESTRECORD) + "数据"
	//fmt.Printf(logformatstr_, logformattitle)
	//tmpstartTime = currentTimeMillis()
	//var tdelArr []*hbase.TDelete
	//for i := 0; i < TESTRECORD; i++ {
	//	putrowkey := strconv.Itoa(i)
	//	tdelArr = append(tdelArr, &hbase.TDelete{
	//		Row: []byte(putrowkey)})
	//}
	//r, err := client.DeleteMultiple([]byte(temptable), tdelArr)
	//if err != nil {
	//	fmt.Printf("DeleteMultiple err:%s\n", err)
	//} else {
	//	fmt.Printf("DeleteMultiple %d done\n", TESTRECORD)
	//	fmt.Println(r)
	//}
	//tmpendTime = currentTimeMillis()
	//fmt.Printf(logformatstr, logformattitle, tmpendTime, tmpstartTime, (tmpendTime - tmpstartTime))
	//
	//endTime := currentTimeMillis()
	//fmt.Printf("\nGolang调用总计用时:%d-%d=%d毫秒\n", endTime, startTime, (endTime - startTime))
}

func printscruct(cv interface{}) {
	switch reflect.ValueOf(cv).Interface().(type) {
	case *hbase.TColumnValue:
		s := reflect.ValueOf(cv).Elem()
		typeOfT := s.Type()
		//获取Thrift2中struct的field
		for i := 0; i < s.NumField(); i++ {
			f := s.Field(i)
			fileldformatstr := "\t%d: %s(%s)= %v\n"
			switch f.Interface().(type) {
			case []uint8:
				fmt.Printf(fileldformatstr, i, typeOfT.Field(i).Name, f.Type(), string(f.Interface().([]uint8)))
			case *int64:
				var tempint64 int64
				if f.Interface().(*int64) == nil {
					tempint64 = 0
				} else {
					tempint64 = *f.Interface().(*int64)
				}
				fmt.Printf(fileldformatstr, i, typeOfT.Field(i).Name, f.Type(), tempint64)
			default:
				fmt.Printf("I don't know")
			}
		}
	default:
		fmt.Printf("I don't know")
		fmt.Print(reflect.ValueOf(cv))
	}

}

func currentTimeMillis() int64 {
	return time.Now().UnixNano() / 1000000
}