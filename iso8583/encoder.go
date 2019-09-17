package iso8583

type Encoder struct {
	HeaderLength uint32
}

func (enc *Encoder) Encode(data []byte) ([]byte, error) {
	return nil, nil
}
