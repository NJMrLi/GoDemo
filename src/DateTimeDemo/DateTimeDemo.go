package DateTimeDemo

import (
	"time"
	"fmt"
)

func DateTimeBasic() time.Time{
	now:=time.Now()

	fmt.Printf("%v\n",now)

	year:=now.Year()
	month:=now.Month()
	day:=now.Day()
	hour:=now.Hour()
	minute:=now.Minute()
	send:=now.Second()

	fmt.Printf("%02d/%02d/%02d %02d:%02d:%02d\n",year,month,day,hour,minute,send)

	return now
}

func GetDateTimeStamp(datetime time.Time) int64{
	now:=datetime.Unix()
	fmt.Printf("TimeStamp: %v\n",now)
	return now
}

func GetDateTime(timeStamp int64){
	now:=time.Unix(timeStamp,0)
	fmt.Printf("DateTime: %v\n",now)
}

func AddDay(src time.Time) time.Time{
	//添加一天两小时
	src = src.AddDate(0,0,1).Add(time.Hour * 2)
	return src
}


func SimpleTicker(){
	//间隔两秒，会像Channel中写入一个数据
	ticker := time.Tick(time.Second * 2)

	for i := range ticker{
		fmt.Printf("%v\n",i)
		simpleTask()
	}
}

func simpleTask(){
	fmt.Println("Task Start")
}