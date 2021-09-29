package timer

import (
	"encoding/json"
	"fmt"
	"log"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

var event = `
{
    "appid": "NMDPTRIAL_peter_freshman_nuance_com_20191015T195942962874",
    "data": {
        "clientData": {},
        "dataContentType": "application/x-nuance-dlg-nii-logs.v1+json",
        "events": [
            {
                "name": "transition",
                "value": {
                    "from": {
                        "component": "HI",
                        "name": "Question Router",
                        "type": "questionrouter",
                        "uuid": "4346224e-6337-4839-bd0f-4ab034a81a79"
                    },
                    "to": {
                        "component": "HI",
                        "name": "Hey",
                        "type": "playprompt",
                        "uuid": "1b4a8496-29a4-4662-be2a-1208057f525e"
                    }
                }
            }
        ],
        "locale": "en-US",
        "requestid": "d3056a02-c357-4a9a-8d34-c74b78d98592",
        "seqid": "28",
        "sessionid": "3a0ed103-5d0c-4104-bb7c-a194764f39ea",
        "traceid": "4090041a4f716f74"
    },
    "datacontenttype": "application/json",
    "id": "0a6dd533-31af-4b38-bb94-69021613b33f",
    "key": "{\"service\":\"DLGaaS\"}",
    "offset": 406,
    "partition": 0,
    "partitionKey": "{\"service\": \"DLGaaS\", \"id\": \"0a6dd533-31af-4b38-bb94-69021613b33f\"}",
    "service": "DLGaaS",
    "source": "NIIEventLogger",
    "specversion": "1.0",
    "timestamp": "2021-08-09T18:51:21.971Z",
    "topic": "NMDPTRIAL_peter_freshman_nuance_com_20191015T195942962874",
    "type": "NIIEventLog",
    "value": {
        "appid": "NMDPTRIAL_peter_freshman_nuance_com_20191015T195942962874",
        "data": {
            "clientData": {},
            "dataContentType": "application/x-nuance-dlg-nii-logs.v1+json",
            "events": [
                {
                    "name": "transition",
                    "value": {
                        "from": {
                            "component": "HI",
                            "name": "Question Router",
                            "type": "questionrouter",
                            "uuid": "4346224e-6337-4839-bd0f-4ab034a81a79"
                        },
                        "to": {
                            "component": "HI",
                            "name": "Hey",
                            "type": "playprompt",
                            "uuid": "1b4a8496-29a4-4662-be2a-1208057f525e"
                        }
                    }
                }
            ],
            "locale": "en-US",
            "requestid": "d3056a02-c357-4a9a-8d34-c74b78d98592",
            "seqid": "28",
            "sessionid": "3a0ed103-5d0c-4104-bb7c-a194764f39ea",
            "traceid": "4090041a4f716f74"
        },
        "datacontenttype": "application/json",
        "id": "0a6dd533-31af-4b38-bb94-69021613b33f",
        "partitionKey": "{\"service\": \"DLGaaS\", \"id\": \"0a6dd533-31af-4b38-bb94-69021613b33f\"}",
        "service": "DLGaaS",
        "source": "NIIEventLogger",
        "specversion": "1.0",
        "timestamp": "2021-08-09T18:51:21.971Z",
        "type": "NIIEventLog"
    }
}
`

type Records struct {
	val     int
	mu      sync.Mutex
	Records jsonObj
}

type (
	jsonObj = map[string]interface{}
	jsonArr = []interface{}
)

func DoTest(LogFlushTimeoutSec int) {

	records := &Records{
		Records: make(jsonObj),
		val:     0,
	}
	records.Records["records"] = make(jsonArr, 0)
	var wg sync.WaitGroup
	wg.Add(1)

	done := make(chan int, 1)
	now := time.Tick(time.Duration(LogFlushTimeoutSec) * time.Nanosecond)

	count := 1000
	// Go Routine processing cache
	go func(done chan int, now <-chan time.Time, records *Records, wg *sync.WaitGroup, counter int) {
		for {
			select {
			case <-done:
				fmt.Println("Received Done")
				wg.Done()
				return
			case <-now:
				records.processCache(done, counter)
			}
		}

	}(done, now, records, &wg, count)

	var obj jsonObj
	err := json.Unmarshal([]byte(event), &obj)
	_obj, err := json.Marshal(obj)
	fmt.Printf("size in bytes: %v \n", len(_obj))
	if err != nil {
		log.Fatalf("%v", err)
	}

	// Go routine adding val
	for i := 0; i < count; i++ {
		wg.Add(1)
		go records.generateLoad(done, &wg, obj, count)
	}
	wg.Wait()
	fmt.Println(records.val)
}

func (r *Records) generateLoad(done chan int, wg *sync.WaitGroup, obj jsonObj, counter int) {
	defer wg.Done()
	r.mu.Lock()
	defer r.mu.Unlock()
	r.Records["records"] = append(r.Records["records"].(jsonArr), obj)
	_records, err := json.Marshal(r.Records["records"].(jsonArr))
	if err != nil {
		panic(err)
	}
	size := len([]byte(_records))
	fmt.Printf("increased size in bytes in total: %v \n", size)
	if size > 1*1e+3 {
		r.processCache(done, counter)
	}
	// fmt.Printf("Goroutine: %v, Increments: %v \n", goid(), len(r.Records["records"].(jsonArr)))

}

func (r *Records) processCache(done chan int, counter int) {
	r.mu.Lock()
	defer r.mu.Unlock()
	count := len(r.Records["records"].(jsonArr))
	if r.Records != nil && len(r.Records) > 0 {
		if count > 0 {
			r.val += count
			fmt.Printf("Goroutine: %v , Purging : %v, Total: %v \n", goid(), len(r.Records["records"].(jsonArr)), r.val)
			_records, err := json.Marshal(r.Records["records"].(jsonArr))
			if err != nil {
				panic(err)
			}
			fmt.Printf("Purged size in bytes in total: %v \n", len([]byte(_records)))
			if r.val == counter {
				done <- 1
			}
			r.Records["records"] = make(jsonArr, 0)
		}
	} else {
		fmt.Println("records empty")
	}

}
func goid() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}
