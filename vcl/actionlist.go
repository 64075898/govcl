
//----------------------------------------
// The code is automatically generated by the GenlibLcl tool.
// Copyright © ying32. All Rights Reserved.
// 
// Licensed under Apache License 2.0
//
//----------------------------------------


package vcl


import (
    . "github.com/ying32/govcl/vcl/api"
    . "github.com/ying32/govcl/vcl/types"
    "unsafe"
)

type TActionList struct {
    IComponent
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
    ptr unsafe.Pointer
}

// CN: 创建一个新的对象。
// EN: Create a new object.
func NewActionList(owner IComponent) *TActionList {
    a := new(TActionList)
    a.instance = ActionList_Create(CheckPtr(owner))
    a.ptr = unsafe.Pointer(a.instance)
    // 不敢启用，因为不知道会发生什么...
    // runtime.SetFinalizer(a, (*TActionList).Free)
    return a
}

// CN: 动态转换一个已存在的对象实例。或者使用Obj.As().<目标对象>。
// EN: Dynamically convert an existing object instance. Or use Obj.As().<Target object>.
func AsActionList(obj interface{}) *TActionList {
    instance, ptr := getInstance(obj)
    if instance == 0 { return nil }
    return &TActionList{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
// Deprecated: use AsActionList.
func ActionListFromInst(inst uintptr) *TActionList {
    return AsActionList(inst)
}

// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
// Deprecated: use AsActionList.
func ActionListFromObj(obj IObject) *TActionList {
    return AsActionList(obj)
}

// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecured address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsActionList.
func ActionListFromUnsafePointer(ptr unsafe.Pointer) *TActionList {
    return AsActionList(ptr)
}

// -------------------------- Deprecated end --------------------------
// CN: 释放对象。
// EN: Free object.
func (a *TActionList) Free() {
    if a.instance != 0 {
        ActionList_Free(a.instance)
        a.instance, a.ptr = 0, nullptr
    }
}

// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (a *TActionList) Instance() uintptr {
    return a.instance
}

// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (a *TActionList) UnsafeAddr() unsafe.Pointer {
    return a.ptr
}

// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (a *TActionList) IsValid() bool {
    return a.instance != 0
}

// CN: 检测当前对象是否继承自目标对象。
// EN: Checks whether the current object is inherited from the target object.
func (a *TActionList) Is() TIs {
    return TIs(a.instance)
}

// CN: 动态转换当前对象为目标对象。
// EN: Dynamically convert the current object to the target object.
//func (a *TActionList) As() TAs {
//    return TAs(a.instance)
//}

// CN: 获取类信息指针。
// EN: Get class information pointer.
func TActionListClass() TClass {
    return ActionList_StaticClassType()
}

// CN: 查找指定名称的组件。
// EN: Find the component with the specified name.
func (a *TActionList) FindComponent(AName string) *TComponent {
    return AsComponent(ActionList_FindComponent(a.instance, AName))
}

// CN: 获取类名路径。
// EN: Get the class name path.
func (a *TActionList) GetNamePath() string {
    return ActionList_GetNamePath(a.instance)
}

// CN: 是否有父容器。
// EN: Is there a parent container.
func (a *TActionList) HasParent() bool {
    return ActionList_HasParent(a.instance)
}

// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (a *TActionList) Assign(Source IObject) {
    ActionList_Assign(a.instance, CheckPtr(Source))
}

// CN: 获取类的类型信息。
// EN: Get class type information.
func (a *TActionList) ClassType() TClass {
    return ActionList_ClassType(a.instance)
}

// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (a *TActionList) ClassName() string {
    return ActionList_ClassName(a.instance)
}

// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (a *TActionList) InstanceSize() int32 {
    return ActionList_InstanceSize(a.instance)
}

// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (a *TActionList) InheritsFrom(AClass TClass) bool {
    return ActionList_InheritsFrom(a.instance, AClass)
}

// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (a *TActionList) Equals(Obj IObject) bool {
    return ActionList_Equals(a.instance, CheckPtr(Obj))
}

// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (a *TActionList) GetHashCode() int32 {
    return ActionList_GetHashCode(a.instance)
}

// CN: 文本类信息。
// EN: Text information.
func (a *TActionList) ToString() string {
    return ActionList_ToString(a.instance)
}

// CN: 获取图标索引列表对象。
// EN: .
func (a *TActionList) Images() *TImageList {
    return AsImageList(ActionList_GetImages(a.instance))
}

// CN: 设置图标索引列表对象。
// EN: .
func (a *TActionList) SetImages(value IComponent) {
    ActionList_SetImages(a.instance, CheckPtr(value))
}

func (a *TActionList) State() TActionListState {
    return ActionList_GetState(a.instance)
}

func (a *TActionList) SetState(value TActionListState) {
    ActionList_SetState(a.instance, value)
}

// CN: 设置改变事件。
// EN: Set changed event.
func (a *TActionList) SetOnChange(fn TNotifyEvent) {
    ActionList_SetOnChange(a.instance, fn)
}

// CN: 获取组件总数。
// EN: Get the total number of components.
func (a *TActionList) ComponentCount() int32 {
    return ActionList_GetComponentCount(a.instance)
}

// CN: 获取组件索引。
// EN: Get component index.
func (a *TActionList) ComponentIndex() int32 {
    return ActionList_GetComponentIndex(a.instance)
}

// CN: 设置组件索引。
// EN: Set component index.
func (a *TActionList) SetComponentIndex(value int32) {
    ActionList_SetComponentIndex(a.instance, value)
}

// CN: 获取组件所有者。
// EN: Get component owner.
func (a *TActionList) Owner() *TComponent {
    return AsComponent(ActionList_GetOwner(a.instance))
}

// CN: 获取组件名称。
// EN: Get the component name.
func (a *TActionList) Name() string {
    return ActionList_GetName(a.instance)
}

// CN: 设置组件名称。
// EN: Set the component name.
func (a *TActionList) SetName(value string) {
    ActionList_SetName(a.instance, value)
}

// CN: 获取对象标记。
// EN: Get the control tag.
func (a *TActionList) Tag() int {
    return ActionList_GetTag(a.instance)
}

// CN: 设置对象标记。
// EN: Set the control tag.
func (a *TActionList) SetTag(value int) {
    ActionList_SetTag(a.instance, value)
}

// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (a *TActionList) Components(AIndex int32) *TComponent {
    return AsComponent(ActionList_GetComponents(a.instance, AIndex))
}

