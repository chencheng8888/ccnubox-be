// Code generated by protoc-gen-go-errors. DO NOT EDIT.

package feedv1

import (
	fmt "fmt"
	errors "github.com/go-kratos/kratos/v2/errors"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
const _ = errors.SupportPackageIsVersion1

func IsTokenAlreadyExist(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_TOKEN_ALREADY_EXIST.String() && e.Code == 501
}

func ErrorTokenAlreadyExist(format string, args ...interface{}) *errors.Error {
	return errors.New(501, ErrorReason_TOKEN_ALREADY_EXIST.String(), fmt.Sprintf(format, args...))
}

func IsUserNotFound(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_USER_NOT_FOUND.String() && e.Code == 502
}

func ErrorUserNotFound(format string, args ...interface{}) *errors.Error {
	return errors.New(502, ErrorReason_USER_NOT_FOUND.String(), fmt.Sprintf(format, args...))
}

func IsGetFeedEventError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_GET_FEED_EVENT_ERROR.String() && e.Code == 503
}

func ErrorGetFeedEventError(format string, args ...interface{}) *errors.Error {
	return errors.New(503, ErrorReason_GET_FEED_EVENT_ERROR.String(), fmt.Sprintf(format, args...))
}

func IsClearFeedEventError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_CLEAR_FEED_EVENT_ERROR.String() && e.Code == 504
}

func ErrorClearFeedEventError(format string, args ...interface{}) *errors.Error {
	return errors.New(504, ErrorReason_CLEAR_FEED_EVENT_ERROR.String(), fmt.Sprintf(format, args...))
}

func IsPublicFeedEventError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_PUBLIC_FEED_EVENT_ERROR.String() && e.Code == 505
}

func ErrorPublicFeedEventError(format string, args ...interface{}) *errors.Error {
	return errors.New(505, ErrorReason_PUBLIC_FEED_EVENT_ERROR.String(), fmt.Sprintf(format, args...))
}

func IsFindConfigOrTokenError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_FIND_CONFIG_OR_TOKEN_ERROR.String() && e.Code == 506
}

func ErrorFindConfigOrTokenError(format string, args ...interface{}) *errors.Error {
	return errors.New(506, ErrorReason_FIND_CONFIG_OR_TOKEN_ERROR.String(), fmt.Sprintf(format, args...))
}

func IsChangeConfigOrTokenError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_CHANGE_CONFIG_OR_TOKEN_ERROR.String() && e.Code == 507
}

func ErrorChangeConfigOrTokenError(format string, args ...interface{}) *errors.Error {
	return errors.New(507, ErrorReason_CHANGE_CONFIG_OR_TOKEN_ERROR.String(), fmt.Sprintf(format, args...))
}

func IsRemoveConfigOrTokenError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_REMOVE_CONFIG_OR_TOKEN_ERROR.String() && e.Code == 508
}

func ErrorRemoveConfigOrTokenError(format string, args ...interface{}) *errors.Error {
	return errors.New(508, ErrorReason_REMOVE_CONFIG_OR_TOKEN_ERROR.String(), fmt.Sprintf(format, args...))
}

func IsGetMuxiFeedError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_GET_MUXI_FEED_ERROR.String() && e.Code == 509
}

func ErrorGetMuxiFeedError(format string, args ...interface{}) *errors.Error {
	return errors.New(509, ErrorReason_GET_MUXI_FEED_ERROR.String(), fmt.Sprintf(format, args...))
}

func IsInsertMuxiFeedError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_INSERT_MUXI_FEED_ERROR.String() && e.Code == 510
}

func ErrorInsertMuxiFeedError(format string, args ...interface{}) *errors.Error {
	return errors.New(510, ErrorReason_INSERT_MUXI_FEED_ERROR.String(), fmt.Sprintf(format, args...))
}

func IsRemoveMuxiFeedError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_REMOVE_MUXI_FEED_ERROR.String() && e.Code == 511
}

func ErrorRemoveMuxiFeedError(format string, args ...interface{}) *errors.Error {
	return errors.New(511, ErrorReason_REMOVE_MUXI_FEED_ERROR.String(), fmt.Sprintf(format, args...))
}
