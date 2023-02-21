// Code generated by Validator v0.1.4. DO NOT EDIT.

package interact

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
func (p *FavoriteActionRequest) IsValid() error {
	return nil
}
func (p *FavoriteActionResponse) IsValid() error {
	return nil
}
func (p *FavoriteListRequest) IsValid() error {
	return nil
}
func (p *FavoriteListResponse) IsValid() error {
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
func (p *CommentActionRequest) IsValid() error {
	return nil
}
func (p *CommentActionResponse) IsValid() error {
	if p.Comment != nil {
		if err := p.Comment.IsValid(); err != nil {
			return fmt.Errorf("filed Comment not valid, %w", err)
		}
	}
	return nil
}
func (p *CommentListRequest) IsValid() error {
	return nil
}
func (p *CommentListResponse) IsValid() error {
	return nil
}
func (p *Comment) IsValid() error {
	if p.User != nil {
		if err := p.User.IsValid(); err != nil {
			return fmt.Errorf("filed User not valid, %w", err)
		}
	}
	return nil
}
