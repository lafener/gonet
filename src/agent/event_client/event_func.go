package event_client

import (
	"log"
	"fmt"
)

import (
	event "event/protos"
	"misc/packet"
)

func Ping() bool {
	defer _event_err()
	req := event.INT{}
	req.F_v = 1
	ret := _call(packet.Pack(event.Code["ping_req"], req, nil))
	reader := packet.Reader(ret)
	tbl, _ := event.PKT_INT(reader)
	fmt.Println(tbl)
	if tbl.F_v != req.F_v {
		return false
	}

	return true
}

func Add(oid uint32, user_id int32, timeout int64) uint32 {
	defer _event_err()
	req := event.ADD_REQ{}
	req.F_oid = oid
	req.F_user_id = user_id
	req.F_timeout = timeout
	ret := _call(packet.Pack(event.Code["add_req"], req, nil))
	reader := packet.Reader(ret)
	tbl, _ := event.PKT_INT(reader)
	return tbl.F_v
}

func _event_err() {
	if x := recover(); x != nil {
		log.Println(x)
	}
}
