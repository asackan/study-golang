package koreaPost

import "fmt"

type PostSender struct {
}

func (k *PostSender) Send(parcel string) {
	fmt.Println("우체국 sends", parcel, "parcel")
}
