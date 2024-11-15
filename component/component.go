/**
 * @Author: dingQingHui
 * @Description:
 * @File: component
 * @Version: 1.0.0
 * @Date: 2024/9/20 17:53
 */

package component

type (
	IComponent interface {
		Name() string
		IComponentLifecycle
	}
	IComponentLifecycle interface {
		Init()
		Stop()
		SetParent(parent IComponent)
		GetParent() IComponent
		AddComponent(...IComponent)
		RangeChildren(func(IComponent) bool)
	}
)

var _ IComponent = &BuiltinComponent{}

type BuiltinComponent struct {
	parent   IComponent
	children []IComponent
}

func (b *BuiltinComponent) Name() string {
	return ""
}

func (b *BuiltinComponent) Init() {
	for _, child := range b.children {
		child.Init()
	}
}

func (b *BuiltinComponent) Stop() {
	for _, child := range b.children {
		child.Stop()
	}
}

func (b *BuiltinComponent) AddComponent(children ...IComponent) {
	for _, child := range children {
		child.SetParent(b)
	}
	b.children = append(b.children, children...)
}

func (b *BuiltinComponent) SetParent(parent IComponent) {
	b.parent = parent
}

func (b *BuiltinComponent) GetParent() IComponent {
	return b.parent
}

func (b *BuiltinComponent) RangeChildren(f func(IComponent) bool) {
	for _, child := range b.children {
		if !f(child) {
			return
		}
	}
}
