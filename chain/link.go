package chain

import (
	"fmt"
	"github.com/hootuu/domain/scope"
	"github.com/hootuu/utils/errors"
	"github.com/hootuu/utils/strs"
	"regexp"
)

type Code = string

func CodeVerify(code string) *errors.Error {
	if len(code) > 64 {
		return errors.Verify("The length of the code field cannot be greater than 64")
	}
	matched, err := regexp.MatchString("^[a-zA-Z0-9_]+$", code)
	if err != nil || !matched {
		return errors.Verify("Code must be a valid string combination")
	}
	return nil
}

const CreationLinkData = "@"

type CreationLink struct {
	Lead scope.Lead `bson:"lead" json:"lead"`
	Code Code       `bson:"code" json:"code"`
}

func CreationLinkOf(lead scope.Lead, code Code) (*CreationLink, *errors.Error) {
	if err := lead.Verify(); err != nil {
		return nil, errors.Verify("invalid lead:" + err.Error())
	}
	if err := CodeVerify(code); err != nil {
		return nil, err
	}
	return &CreationLink{
		Lead: lead,
		Code: code,
	}, nil
}

func (link *CreationLink) GetChainKey() Key {
	return strs.MD5(fmt.Sprintf("%s.%s", link.Lead.Scope, link.Code))
}

type Link struct {
	Lead scope.Lead `bson:"lead" json:"lead"`
	Code Code       `bson:"code" json:"code"`
	Data Cid        `bson:"data" json:"data"`
}

func LinkOf(lead scope.Lead, code Code, dataCid Cid) (*Link, *errors.Error) {
	if err := lead.Verify(); err != nil {
		return nil, errors.Verify("invalid lead:" + err.Error())
	}
	if err := CodeVerify(code); err != nil {
		return nil, err
	}
	if err := CidVerify(dataCid); err != nil {
		return nil, err
	}
	return &Link{
		Lead: lead,
		Code: code,
		Data: dataCid,
	}, nil
}

func (link *Link) GetCreation() CreationLink {
	return CreationLink{
		Lead: link.Lead,
		Code: link.Code,
	}
}

func (link *Link) GetChainKey() Key {
	return strs.MD5(fmt.Sprintf("%s.%s", link.Lead.Scope, link.Code))
}
