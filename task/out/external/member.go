package external

import (
	"fmt"
	"todopoint/common/netutils/networking"
)

func IsMemberExist(memId int) (*networking.ExternalInfo, error) {
	url := fmt.Sprintf("http://localhost:3000/members/%d", memId)
	info, err := networking.RequestGet(url)
	if err != nil {
		return nil, err
	}
	return info, nil
}
