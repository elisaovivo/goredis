package reply

type UnknowErrReply struct {
}

var unknownErrBytes = []byte("-Err unknown\r\n")

func (u UnknowErrReply) Error() string {
	return "Err unknown"
}

func (u UnknowErrReply) ToBytes() []byte {
	return unknownErrBytes
}

type ArgNumErrReply struct {
	Cmd string
}

func (r *ArgNumErrReply) Error() string {
	return "-ERR wrong number of arguments for `" + r.Cmd + "` command\r\n"
}

func (r *ArgNumErrReply) ToBytes() []byte {
	return []byte("-ERR wrong number of arguments for `" + r.Cmd + "` command\r\n")
}

func MakeArgNumErrReply(cmd string) *ArgNumErrReply {
	return &ArgNumErrReply{
		Cmd: cmd,
	}
}
