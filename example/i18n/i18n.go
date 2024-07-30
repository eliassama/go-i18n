// Package i18n Code generated by go generate; DO NOT EDIT. Please adjust your TOML file to regenerate.
package i18n

import (
	"bytes"
	i18ncore "github.com/eliassama/go-i18n/core"
	"golang.org/x/text/language"
	"strings"
	"text/template"
)

var Msg = struct {
	NetworkAuthenticationRequiredMsg *i18ncore.Bundle
	TestMsg                          *i18ncore.Bundle
}{
	NetworkAuthenticationRequiredMsg: &i18ncore.Bundle{
		Status:   100511,
		Code:     511,
		Plural:   map[string]string{},
		Singular: map[string]string{},
		Msg: map[string]string{
			"en-us": "Network authentication required",
			"zh-cn": "需要网络身份验证",
		},
	},
	TestMsg: &i18ncore.Bundle{
		Status: 100555,
		Code:   555,
		Plural: map[string]string{
			"en-us": "This is a plural test message, suitable for cases where the data quantity is greater than 1. Name: {{.Name}}, Age: {{.Age}}, Home Address: {{.Phone}}",
			"zh-cn": "这是一条复数测试消息，适用于数据量大于1的情况。姓名：{{.Name}}，年龄：{{.Age}}，家庭住址：{{.Phone}}",
		},
		Singular: map[string]string{
			"en-us": "This is a singular test message, suitable for the case where the data quantity is 1. Name: {{.Name}}, Age: {{.Age}}, Home Address: {{.Phone}}",
			"zh-cn": "这是一条单数测试消息，适用于数据量为1的情况。姓名：{{.Name}}，年龄：{{.Age}}，家庭住址：{{.Phone}}",
		},
		Msg: map[string]string{
			"en-us": "This is a test message. Name: {{.Name}}, Age: {{.Age}}, Home Address: {{.Phone}}",
			"zh-cn": "这是一条测试消息。姓名：{{.Name}}，年龄：{{.Age}}，家庭住址：{{.Phone}}",
		},
	},
}

func SendNetworkAuthenticationRequiredMsg(langCode string) (int, int, string) {
	bundle := Msg.NetworkAuthenticationRequiredMsg
	langCode = convertLangCode(langCode)

	var msg string
	msg = bundle.Msg[langCode]
	if msg == "" {
		msg = bundle.Msg["zh-cn"]
	}

	return bundle.Code, bundle.Status, msg
}

// TestMsgParams is the parameter struct for the SendTestMsg method
type TestMsgParams struct {
	Name  string
	Age   string
	Phone string
}

func SendTestMsg(langCode string, count int, params TestMsgParams) (int, int, string, error) {
	bundle := Msg.TestMsg
	langCode = convertLangCode(langCode)

	var msg string
	if count > 1 {
		msg = bundle.Plural[langCode]
		if msg == "" {
			msg = bundle.Plural["zh-cn"]
		}
	} else {
		msg = bundle.Singular[langCode]
		if msg == "" {
			msg = bundle.Singular["zh-cn"]
		}
		if msg == "" {
			msg = bundle.Msg[langCode]
			if msg == "" {
				msg = bundle.Msg["zh-cn"]
			}
		}
	}

	tmpl, err := template.New("message").Parse(msg)
	if err != nil {
		return 0, 0, "", err
	}
	var buf bytes.Buffer
	if err = tmpl.Execute(&buf, params); err != nil {
		return 0, 0, "", err
	}
	return bundle.Code, bundle.Status, buf.String(), nil
}

func convertLangCode(langCode string) string {
	if langTag, err := language.Parse(langCode); err != nil {
		return ""
	} else {
		return strings.ToLower(langTag.String())
	}
}