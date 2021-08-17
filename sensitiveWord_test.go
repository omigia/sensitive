package sensitive

import (
	"testing"
)

func TestInit(t *testing.T)  {
	root := Init([]string{
		"78",
		"123456789",
		"345",
		"1234",
		"123",
		"111",
	})
	if len(root.Children) != 3 {
		t.Fatal("根结点子结点数错误")
	}

	// 测试12345678
	if _, ok := root.Children['1']; !ok {
		t.Fatal("结点错误: 1")
	}
	p := root.Children['1']
	if _, ok := p.Children['2']; !ok {
		t.Fatal("结点错误: 12")
	}
	p = p.Children['2']
	if _, ok := p.Children['3']; !ok {
		t.Fatal("结点错误: 123")
	}
	if p.Fail != root {
		t.Fatal("失败结点错误: 12")
	}
	p = p.Children['3']
	if _, ok := p.Children['4']; !ok || !p.Reach {
		t.Fatal("结点错误: 1234")
	}
	p = p.Children['4']
	if _, ok := p.Children['5']; !ok || !p.Reach {
		t.Fatal("结点错误: 12345")
	}
	if p.Fail == nil || p.Fail.Data != '4' {
		t.Fatal("失败结点错误: 1234")
	}
	p = p.Children['5']
	if _, ok := p.Children['6']; !ok || !p.Reach {
		t.Fatal("结点错误: 123456")
	}
	if p.Fail == nil || p.Fail.Data != '5' || !p.Fail.Reach || len(p.Fail.Children) != 0 {
		t.Fatal("失败结点错误: 12345")
	}
	p = p.Children['6']
	if _, ok := p.Children['7']; !ok || p.Fail!=root {
		t.Fatal("结点错误: 1234567")
	}

	// 测试111
	p = root.Children['1']
	if _, ok := p.Children['1']; !ok {
		t.Fatal("结点错误: 11")
	}
	p = p.Children['1']
	if _, ok := p.Children['1']; !ok || p.Reach || p.Fail.Data!='1' || !p.Children['1'].Reach || p.Children['1'].Fail.Data!='1' {
		t.Fatal("结点错误: 111")
	}

	// 测试345
	if _, ok := root.Children['3']; !ok {
		t.Fatal("结点错误: 3")
	}
	p = root.Children['3']
	if _, ok := p.Children['4']; !ok {
		t.Fatal("结点错误: 34")
	}
	p = p.Children['4']
	if _, ok := p.Children['5']; !ok || p.Reach || !p.Children['5'].Reach {
		t.Fatal("结点错误: 345")
	}

	// 测试78
	if _, ok := root.Children['7']; !ok {
		t.Fatal("结点错误: 7")
	}
	p = root.Children['7']
	if _, ok := p.Children['8']; !ok {
		t.Fatal("结点错误: 78")
	}
}

func TestMatch(t *testing.T) {
	root := Init([]string{
		"78",
		"123456789",
		"345",
		"111",
		"好啊",
	})
	if !Match(root, "78") {
		t.Fatal("未找到: 78")
	}
	if Match(root, "1234") {
		t.Fatal("不应该找到: 1234")
	}
	if !Match(root, "12345") {
		t.Fatal("未找到: 345")
	}
	if !Match(root, "12345678") {
		t.Fatal("未找到: 345或78")
	}
	if !Match(root, "34567") {
		t.Fatal("未找到: 345")
	}
	if !Match(root, "45678") {
		t.Fatal("未找到: 78")
	}
	if !Match(root, "你好啊") {
		t.Fatal("未找到: 好啊")
	}
	if Match(root, "你好") {
		t.Fatal("不应该找到: 你好")
	}
}

func TestMatch2(t *testing.T) {
	root := Init([]string{
		"123456789",
		"345",
	})
	if !Match(root, "12345678") {
		t.Fatal("未找到: 345或78")
	}
}