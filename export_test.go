package snmpgo

import (
	"time"
)

var StripHexPrefix = stripHexPrefix
var ToHexStr = toHexStr
var Retry = retry
var GenRequestId = genRequestId
var GenSalt32 = genSalt32
var GenSalt64 = genSalt64
var GenMessageId = genMessageId
var NewNotInTimeWindowError = func() error { return &notInTimeWindowError{&MessageError{}} }

// For snmpgo testing
var NewSNMPEngine = newSNMPEngine

func ArgsValidate(args *SNMPArguments) error { return args.validate() }
func CheckPdu(engine *snmpEngine, pdu Pdu, args *SNMPArguments) error {
	return engine.checkPdu(pdu, args)
}

// For message testing
var NewMessage = NewMessage
var UnmarshalMessage = UnmarshalMessage
var NewMessageWithPdu = NewMessageWithPdu
var NewMessageProcessing = NewMessageProcessing

func ToMessageV1(msg Message) *messageV1 { return msg.(*messageV1) }
func ToMessageV3(msg Message) *messageV3 { return msg.(*messageV3) }
func ToUsm(sec Security) *usm            { return sec.(*usm) }

// For Security testing
var NewSecurity = newSecurity
var PasswordToKey = passwordToKey
var EncryptDES = encryptDES
var EncryptAES = encryptAES
var DecryptDES = decryptDES
var DecryptAES = decryptAES
var NewSecurityMap = NewSecurityMap

func NewCommunity() *community { return &community{} }
func NewUsm() *usm             { return &usm{} }

// For server
func ListeningUDPAddress(s *TrapServer) string {
	for i := 0; i < 12; i++ {
		if conn := s.transport.(*packetTransport).conn; conn != nil {
			return conn.LocalAddr().String()
		}
		// XXX Wait until a connection is available, but this code is a kludge
		time.Sleep(time.Millisecond * time.Duration(1<<uint(i)))
	}
	return ""
}
