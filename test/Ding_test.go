package test

import (
	"go_gin/tools"
	"testing"
)

type SimpleRequest struct {
	Text SimpleRequestContent `json:"text"`
	Msgtype string `json:"msgtype"`
}

type SimpleRequestContent struct {
	Content string `json:"content"`
}
var access_token = "a84a49cee61270e8dbebe5ec6f7cd4b1160e02f3685bfdfee9536ea8cb16b812"
var secert = "SEC54866c8f6f0f0ff0a8c2f04156337f751eea8e848a8c7c0246146b99ec8a8abc"
var api = "https://oapi.dingtalk.com/robot/send?"
func TestDing(t *testing.T)  {
	tools.Send()
}