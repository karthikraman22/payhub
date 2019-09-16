package iso8583

type Decoder struct {
	HeaderLength uint32
}

func (dec *Decoder) Decode(data []byte) ([]byte, error) {
	return nil, nil
}
