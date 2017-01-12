package main 

import (
	"ethos/syscall"
	"ethos/ethos"
	"log"
	"ethos/efmt"
)

func main () {
	me := syscall.GetUser()
	path := "/user/" + me + "/myDir/"
	fd, status := ethos.OpenDirectoryPath(path)
	if status != syscall.StatusOk {
		log.Fatalf ("Error opening %v: %v\n", path, status)
	}
	
	timeA := syscall.GetTime();
	efmt.Println("Time A is",timeA)
	timeB := syscall.GetTime();
	efmt.Println("Time B is",timeB)
	efmt.Println("The time difference between Time B & Time A is",timeB - timeA, "nanoseconds")


	t1 := Date{day: 15,month: 6,year: 2016,hours: 6,minutes: 15,seconds: 15}
	t1.Write(fd)
	t1.WriteVar(path +"date")
	//efmt.Println("The start date is","(Day)",t1.day,"(Month)",t1.month,"(Year)",t1.year,"(Hours)",t1.hours,"(Minutes)",t1.minutes,"(Seconds)",t1.seconds)
	
	
	t2:= Date{day: 1,month: 1,year: 2016,hours: 0,minutes: 0,seconds: 0}
	efmt.Println("The end date is","(Day)",t2.day,"(Month)",t2.month,"(Year)",t2.year,"(Hours)",t2.hours,"(Minutes)",t2.minutes,"(Seconds)",t2.seconds)
	

	//new := dayarray{31,28,31}
	//efmt.Println("Days in February",new[1])
	
	var startsec = t1.year*365 + (t1.month-1)*30 + t1.day
	startsec = startsec*86400 + t1.hours*3600 + t1.minutes*60 + t1.seconds
	
	//efmt.Println(startsec)
	
	var endsec = (t2.year-1970)*365 + (t2.day-1)
	//endsec = endsec*86400 + t2.hours*3600 + t2.minutes*60 + t2.seconds

	var i int64 = 1	
	for i<t2.month {
	if(i==1) {
		endsec=endsec+31}
	if(i==2) {
		endsec=endsec+28}
	if(i==3) {
		endsec=endsec+31}
	if(i==4) {
		endsec=endsec+30}
	if(i==5) {
		endsec=endsec+31}
	if(i==6) {
		endsec=endsec+30}
	if(i==7) {
		endsec=endsec+31}
	if(i==8) {
		endsec=endsec+31}
	if(i==9) {
		endsec=endsec+30}
	if(i==10) {
		endsec=endsec+31}
	if(i==11) {
		endsec=endsec+30}
	i++
	}
	

	var j int64 = 1970
	for j<=t2.year {
	if j % 4 != 0 {
		endsec = endsec
	} else if j % 400 == 0 {
		endsec = endsec+1
		efmt.Println("Following leap year counted in calculation",j)
	} else if j % 100 == 0 {
		endsec = endsec
	} else {
		if !(j==t2.year && t2.month<3){
		endsec = endsec+1
		efmt.Println("Following leap year counted in calculation",j)
		}
	}
	j++
	}

	endsec = endsec*86400 + t2.hours*3600 + t2.minutes*60 + t2.seconds	
	//efmt.Println(endsec)
	
	var diff = int64(timeB/1000000000) - endsec + 101639
	//var diff = endsec - startsec
	efmt.Println("The time difference between Time B & End date is",diff,"Seconds")
	var diffdays = diff/86400 
	var diffmonths = diffdays/30
	if(diffmonths < 1 && diffdays > 0){		
		efmt.Println("The time difference is", diffdays,"Days")
	}
	if(diffmonths >=1){
		efmt.Println("The time difference is",diffdays,"Days")
	}

}
