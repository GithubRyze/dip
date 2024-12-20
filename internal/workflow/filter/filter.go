package filter

import (
	"dip/internal/errors"
	"time"
)

type Filter[T any] interface {
	DoFilter(parameter T) error
}

type StatusFilter struct{}

func (filter StatusFilter) DoFilter(opened bool) error {
	if !opened {
		return errors.NotOpenedError
	}
	return nil
}

type TimeFilter struct {
	StartTime time.Time
	EndTime   time.Time
}

func (filter TimeFilter) DoFilter(t time.Time) error {
	if t.Before(filter.StartTime) {
		return errors.NotInEffectiveError
	}
	if t.After(filter.EndTime) {
		return errors.ExpirationError
	}
	return nil
}

type IpBlackListFilter struct {
	IpList []string
}

func (filter IpBlackListFilter) DoFilter(ip string) error {
	for _, value := range filter.IpList {
		if ip == value {
			return errors.IpForbiddenError
		}
	}
	return nil
}

type IpAllowListFilter struct {
	IpList []string
}

func (filter IpAllowListFilter) DoFilter(ip string) error {
	for _, value := range filter.IpList {
		if ip == value {
			return nil
		}
	}
	return errors.IpNotAllowedError
}

// TODO
type TokenFilter struct {
}

func (filter TokenFilter) DoFilter(token string) error {
	return nil
}

// TODO
type MsgFormatFilter struct {
}

func (filter MsgFormatFilter) DoFilter(msg string) error {
	return nil
}

// TODO
type MsgTransferFilter struct {
}

func (filter MsgTransferFilter) DoFilter(msg string) error {
	return nil
}

type FilterExecute struct {
	FilterType string
	Executor   Filter[any]
}
