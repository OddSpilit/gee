package gee

import (
	"strings"
)

type node struct {
	pattern  string  // 最终路由
	part     string  // 当前匹配到的字符
	children []*node // 子节点
	isWild   bool    // 是否精准匹配
}

/**
获取匹配到的子节点
*/
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

/**
获取所有能匹配到的节点信息
*/

func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if part == child.part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

/**
插入操作
*/
func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		n.pattern = pattern
		return
	}
	part := parts[height]
	child := n.matchChild(part)
	if child == nil {
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
		n.children = append(n.children, child)
	}
	child.insert(pattern, parts, height+1)
}

/**
搜索
*/
func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}

	part := parts[height]
	children := n.matchChildren(part)

	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}

	return nil
}
