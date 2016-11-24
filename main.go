package main;

import "time"
import "flag"
import "fmt"
import "os"
import "math"
import "math/rand"
//import "sync/atomic"

type Tester struct {
	Tstamp time.Time 
	Counter	int64
}

func newTester() *Tester {
	return &Tester{Tstamp: time.Now(), Counter:0}
}


var (
	_Tester *Tester
	_Pid=os.Getpid()
	_IntervalSeconds time.Duration
	_Exit=false
)



func loopA() {

	for {
		a:=_Tester
		//Tried a sort of computing
		math.Atan(math.Tan(rand.Float64()))	
		a.Counter++
	//	a.Counter = atomic.AddInt64(&a.Counter,1)
		if (_Exit) {
			return
		}
	
	}
}


func syncStart() {
	fmt.Printf("%d:%d %d\t%d\t%s\t\n", time.Now().Hour(), time.Now().Minute(), time.Now().Second(),_Pid,"STARTING");
	vNext:= time.Now().Truncate(_IntervalSeconds).Add(_IntervalSeconds)
	time.Sleep(vNext.Sub(time.Now()))
	
}


func main() {


	vFlagIntervalS:=flag.Int("intervals", 5, "Interval in seconds between samples")
	vFlagDurationS:=flag.Int("durations", 600, "Duration of test in seconds")
	flag.Parse()


	_IntervalSeconds=time.Second * time.Duration(*vFlagIntervalS)

	syncStart()

	vChannelEnd:=time.After(time.Second* time.Duration(*vFlagDurationS))
	_Tester = newTester()

	go loopA()

	MAINLOOP:
	for {

		select { 
			case <- vChannelEnd:
				break MAINLOOP
	 		case <- time.After(_IntervalSeconds):
				vTmp:= _Tester
				_Tester = newTester()
				fmt.Printf("%d:%d %d\t%d\t%s\t%d\n", vTmp.Tstamp.Hour(), vTmp.Tstamp.Minute(), vTmp.Tstamp.Second(),_Pid,"SAMPLE", vTmp.Counter)
		}
	
	}

	fmt.Printf("%d:%d %d\t%d\t%s\t\n", time.Now().Hour(), time.Now().Minute(), time.Now().Second(),_Pid,"FINISHED");
	_Exit=true
	time.Sleep(time.Millisecond * 100)
}
