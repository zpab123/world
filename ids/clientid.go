// /////////////////////////////////////////////////////////////////////////////
// ClientID 对象 [代码齐全]

package ids

// ClientID 长度
const CLIENTID_LENGTH = UUID_LENGTH

// ClientID 对象
type ClientID string

// 生成1个新的 ClientID
func GenClientID() ClientID {
	return ClientID(GenUUID())
}

// ClientID 是否 == ""
func (id ClientID) IsNil() bool {
	return id == ""
}
