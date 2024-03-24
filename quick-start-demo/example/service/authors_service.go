package service

import "go-code-generate/quick-start-demo/example/service/vo"
import "go-code-generate/quick-start-demo/example/service/to"

func AuthorsService(vo *vo.AuthorsVO) *AuthorsServiceHelper {
	return &AuthorsServiceHelper{vo: vo}
}

type AuthorsServiceHelper struct {
	vo *vo.AuthorsVO
	to *to.AuthorsTO
}

func (u *AuthorsServiceHelper) DoXXX() (*to.AuthorsTO, error) {

	//将打包好的to返回
	return u.to, nil
}
