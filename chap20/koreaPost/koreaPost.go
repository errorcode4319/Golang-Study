package koreaPost

import "fmt"

type PostSender struct {
}

func (k *PostSender) Send(parcel string) {
	fmt.Printf("KoreadPost Send %v", parcel)
}
