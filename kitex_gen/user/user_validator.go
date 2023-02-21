// Code generated by Validator v0.1.4. DO NOT EDIT.

package user

import (
	"bytes"
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"time"
)

// unused protection
var (
	_ = fmt.Formatter(nil)
	_ = (*bytes.Buffer)(nil)
	_ = (*strings.Builder)(nil)
	_ = reflect.Type(nil)
	_ = (*regexp.Regexp)(nil)
	_ = time.Nanosecond
)

func (p *BaseResp) IsValid() error {
	return nil
}
func (p *UserLogin) IsValid() error {
	return nil
}
func (p *User) IsValid() error {
	return nil
}
func (p *Video) IsValid() error {
	if p.Author != nil {
		if err := p.Author.IsValid(); err != nil {
			return fmt.Errorf("filed Author not valid, %w", err)
		}
	}
	return nil
}
func (p *LoginUserRequest) IsValid() error {
	if len(p.Username) < int(1) {
		return fmt.Errorf("field Username min_len rule failed, current value: %d", len(p.Username))
	}
	if len(p.Password) < int(1) {
		return fmt.Errorf("field Password min_len rule failed, current value: %d", len(p.Password))
	}
	return nil
}
func (p *LoginUserResponse) IsValid() error {
	return nil
}
func (p *LogoutUserRequest) IsValid() error {
	return nil
}
func (p *LogoutUserResponse) IsValid() error {
	return nil
}
func (p *RegisterUserRequest) IsValid() error {
	return nil
}
func (p *RegisterUserResponse) IsValid() error {
	return nil
}
func (p *UserInfoRequest) IsValid() error {
	return nil
}
func (p *UserInfoResponse) IsValid() error {
	if p.User != nil {
		if err := p.User.IsValid(); err != nil {
			return fmt.Errorf("filed User not valid, %w", err)
		}
	}
	return nil
}
func (p *PublishActionRequest) IsValid() error {
	return nil
}
func (p *PublishActionResponse) IsValid() error {
	return nil
}
func (p *PublishListRequest) IsValid() error {
	return nil
}
func (p *PublishListResponse) IsValid() error {
	return nil
}
<<<<<<< HEAD
func (p *DouyinFeedRequest) IsValid() error {
	return nil
}
func (p *DouyinFeedResponse) IsValid() error {
	return nil
}
func (p *VideoIdRequest) IsValid() error {
	return nil
}
=======
>>>>>>> origin/guo
