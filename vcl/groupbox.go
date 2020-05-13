
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

type TGroupBox struct {
    IWinControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与LCL没有太多关系。
    ptr unsafe.Pointer
}

// CN: 创建一个新的对象。
// EN: Create a new object.
func NewGroupBox(owner IComponent) *TGroupBox {
    g := new(TGroupBox)
    g.instance = GroupBox_Create(CheckPtr(owner))
    g.ptr = unsafe.Pointer(g.instance)
    // 不敢启用，因为不知道会发生什么...
    // runtime.SetFinalizer(g, (*TGroupBox).Free)
    return g
}

// CN: 动态转换一个已存在的对象实例。或者使用Obj.As().<目标对象>。
// EN: Dynamically convert an existing object instance. Or use Obj.As().<Target object>.
func AsGroupBox(obj interface{}) *TGroupBox {
    instance, ptr := getInstance(obj)
    if instance == 0 { return nil }
    return &TGroupBox{instance: instance, ptr: ptr}
}

// -------------------------- Deprecated begin --------------------------
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
// Deprecated: use AsGroupBox.
func GroupBoxFromInst(inst uintptr) *TGroupBox {
    return AsGroupBox(inst)
}

// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
// Deprecated: use AsGroupBox.
func GroupBoxFromObj(obj IObject) *TGroupBox {
    return AsGroupBox(obj)
}

// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecured address. Note: Using this function may cause some unclear situations and be used with caution..
// Deprecated: use AsGroupBox.
func GroupBoxFromUnsafePointer(ptr unsafe.Pointer) *TGroupBox {
    return AsGroupBox(ptr)
}

// -------------------------- Deprecated end --------------------------
// CN: 释放对象。
// EN: Free object.
func (g *TGroupBox) Free() {
    if g.instance != 0 {
        GroupBox_Free(g.instance)
        g.instance, g.ptr = 0, nullptr
    }
}

// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (g *TGroupBox) Instance() uintptr {
    return g.instance
}

// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (g *TGroupBox) UnsafeAddr() unsafe.Pointer {
    return g.ptr
}

// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (g *TGroupBox) IsValid() bool {
    return g.instance != 0
}

// CN: 检测当前对象是否继承自目标对象。
// EN: Checks whether the current object is inherited from the target object.
func (g *TGroupBox) Is() TIs {
    return TIs(g.instance)
}

// CN: 动态转换当前对象为目标对象。
// EN: Dynamically convert the current object to the target object.
//func (g *TGroupBox) As() TAs {
//    return TAs(g.instance)
//}

// CN: 获取类信息指针。
// EN: Get class information pointer.
func TGroupBoxClass() TClass {
    return GroupBox_StaticClassType()
}

// CN: 是否可以获得焦点。
// EN: .
func (g *TGroupBox) CanFocus() bool {
    return GroupBox_CanFocus(g.instance)
}

// CN: 返回是否包含指定控件。
// EN: it's contain a specified control.
func (g *TGroupBox) ContainsControl(Control IControl) bool {
    return GroupBox_ContainsControl(g.instance, CheckPtr(Control))
}

// CN: 返回指定坐标及相关属性位置控件。
// EN: Returns the specified coordinate and the relevant attribute position control..
func (g *TGroupBox) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
    return AsControl(GroupBox_ControlAtPos(g.instance, Pos , AllowDisabled , AllowWinControls , AllLevels))
}

// CN: 禁用控件的对齐。
// EN: Disable control alignment.
func (g *TGroupBox) DisableAlign() {
    GroupBox_DisableAlign(g.instance)
}

// CN: 启用控件对齐。
// EN: Enabled control alignment.
func (g *TGroupBox) EnableAlign() {
    GroupBox_EnableAlign(g.instance)
}

// CN: 查找子控件。
// EN: Find sub controls.
func (g *TGroupBox) FindChildControl(ControlName string) *TControl {
    return AsControl(GroupBox_FindChildControl(g.instance, ControlName))
}

func (g *TGroupBox) FlipChildren(AllLevels bool) {
    GroupBox_FlipChildren(g.instance, AllLevels)
}

// CN: 返回是否获取焦点。
// EN: Return to get focus.
func (g *TGroupBox) Focused() bool {
    return GroupBox_Focused(g.instance)
}

// CN: 句柄是否已经分配。
// EN: Is the handle already allocated.
func (g *TGroupBox) HandleAllocated() bool {
    return GroupBox_HandleAllocated(g.instance)
}

// CN: 插入一个控件。
// EN: Insert a control.
func (g *TGroupBox) InsertControl(AControl IControl) {
    GroupBox_InsertControl(g.instance, CheckPtr(AControl))
}

// CN: 要求重绘。
// EN: Redraw.
func (g *TGroupBox) Invalidate() {
    GroupBox_Invalidate(g.instance)
}

// CN: 移除一个控件。
// EN: Remove a control.
func (g *TGroupBox) RemoveControl(AControl IControl) {
    GroupBox_RemoveControl(g.instance, CheckPtr(AControl))
}

// CN: 重新对齐。
// EN: Realign.
func (g *TGroupBox) Realign() {
    GroupBox_Realign(g.instance)
}

// CN: 重绘。
// EN: Repaint.
func (g *TGroupBox) Repaint() {
    GroupBox_Repaint(g.instance)
}

// CN: 按比例缩放。
// EN: Scale by.
func (g *TGroupBox) ScaleBy(M int32, D int32) {
    GroupBox_ScaleBy(g.instance, M , D)
}

// CN: 滚动至指定位置。
// EN: Scroll by.
func (g *TGroupBox) ScrollBy(DeltaX int32, DeltaY int32) {
    GroupBox_ScrollBy(g.instance, DeltaX , DeltaY)
}

// CN: 设置组件边界。
// EN: Set component boundaries.
func (g *TGroupBox) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    GroupBox_SetBounds(g.instance, ALeft , ATop , AWidth , AHeight)
}

// CN: 设置控件焦点。
// EN: Set control focus.
func (g *TGroupBox) SetFocus() {
    GroupBox_SetFocus(g.instance)
}

// CN: 控件更新。
// EN: Update.
func (g *TGroupBox) Update() {
    GroupBox_Update(g.instance)
}

// CN: 将控件置于最前。
// EN: Bring the control to the front.
func (g *TGroupBox) BringToFront() {
    GroupBox_BringToFront(g.instance)
}

// CN: 将客户端坐标转为绝对的屏幕坐标。
// EN: Convert client coordinates to absolute screen coordinates.
func (g *TGroupBox) ClientToScreen(Point TPoint) TPoint {
    return GroupBox_ClientToScreen(g.instance, Point)
}

// CN: 将客户端坐标转为父容器坐标。
// EN: Convert client coordinates to parent container coordinates.
func (g *TGroupBox) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return GroupBox_ClientToParent(g.instance, Point , CheckPtr(AParent))
}

// CN: 是否在拖拽中。
// EN: Is it in the middle of dragging.
func (g *TGroupBox) Dragging() bool {
    return GroupBox_Dragging(g.instance)
}

// CN: 是否有父容器。
// EN: Is there a parent container.
func (g *TGroupBox) HasParent() bool {
    return GroupBox_HasParent(g.instance)
}

// CN: 隐藏控件。
// EN: Hidden control.
func (g *TGroupBox) Hide() {
    GroupBox_Hide(g.instance)
}

// CN: 发送一个消息。
// EN: Send a message.
func (g *TGroupBox) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return GroupBox_Perform(g.instance, Msg , WParam , LParam)
}

// CN: 刷新控件。
// EN: Refresh control.
func (g *TGroupBox) Refresh() {
    GroupBox_Refresh(g.instance)
}

// CN: 将屏幕坐标转为客户端坐标。
// EN: Convert screen coordinates to client coordinates.
func (g *TGroupBox) ScreenToClient(Point TPoint) TPoint {
    return GroupBox_ScreenToClient(g.instance, Point)
}

// CN: 将父容器坐标转为客户端坐标。
// EN: Convert parent container coordinates to client coordinates.
func (g *TGroupBox) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return GroupBox_ParentToClient(g.instance, Point , CheckPtr(AParent))
}

// CN: 控件至于最后面。
// EN: The control is placed at the end.
func (g *TGroupBox) SendToBack() {
    GroupBox_SendToBack(g.instance)
}

// CN: 显示控件。
// EN: Show control.
func (g *TGroupBox) Show() {
    GroupBox_Show(g.instance)
}

// CN: 获取控件的字符，如果有。
// EN: Get the characters of the control, if any.
func (g *TGroupBox) GetTextBuf(Buffer *string, BufSize int32) int32 {
    return GroupBox_GetTextBuf(g.instance, Buffer , BufSize)
}

// CN: 获取控件的字符长，如果有。
// EN: Get the character length of the control, if any.
func (g *TGroupBox) GetTextLen() int32 {
    return GroupBox_GetTextLen(g.instance)
}

// CN: 设置控件字符，如果有。
// EN: Set control characters, if any.
func (g *TGroupBox) SetTextBuf(Buffer string) {
    GroupBox_SetTextBuf(g.instance, Buffer)
}

// CN: 查找指定名称的组件。
// EN: Find the component with the specified name.
func (g *TGroupBox) FindComponent(AName string) *TComponent {
    return AsComponent(GroupBox_FindComponent(g.instance, AName))
}

// CN: 获取类名路径。
// EN: Get the class name path.
func (g *TGroupBox) GetNamePath() string {
    return GroupBox_GetNamePath(g.instance)
}

// CN: 复制一个对象，如果对象实现了此方法的话。
// EN: Copy an object, if the object implements this method.
func (g *TGroupBox) Assign(Source IObject) {
    GroupBox_Assign(g.instance, CheckPtr(Source))
}

// CN: 获取类的类型信息。
// EN: Get class type information.
func (g *TGroupBox) ClassType() TClass {
    return GroupBox_ClassType(g.instance)
}

// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (g *TGroupBox) ClassName() string {
    return GroupBox_ClassName(g.instance)
}

// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (g *TGroupBox) InstanceSize() int32 {
    return GroupBox_InstanceSize(g.instance)
}

// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (g *TGroupBox) InheritsFrom(AClass TClass) bool {
    return GroupBox_InheritsFrom(g.instance, AClass)
}

// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (g *TGroupBox) Equals(Obj IObject) bool {
    return GroupBox_Equals(g.instance, CheckPtr(Obj))
}

// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (g *TGroupBox) GetHashCode() int32 {
    return GroupBox_GetHashCode(g.instance)
}

// CN: 文本类信息。
// EN: Text information.
func (g *TGroupBox) ToString() string {
    return GroupBox_ToString(g.instance)
}

func (g *TGroupBox) AnchorToNeighbour(ASide TAnchorKind, ASpace int32, ASibling IControl) {
    GroupBox_AnchorToNeighbour(g.instance, ASide , ASpace , CheckPtr(ASibling))
}

func (g *TGroupBox) AnchorParallel(ASide TAnchorKind, ASpace int32, ASibling IControl) {
    GroupBox_AnchorParallel(g.instance, ASide , ASpace , CheckPtr(ASibling))
}

// CN: 置于指定控件的横向中心。
// EN: .
func (g *TGroupBox) AnchorHorizontalCenterTo(ASibling IControl) {
    GroupBox_AnchorHorizontalCenterTo(g.instance, CheckPtr(ASibling))
}

// CN: 置于指定控件的纵向中心。
// EN: .
func (g *TGroupBox) AnchorVerticalCenterTo(ASibling IControl) {
    GroupBox_AnchorVerticalCenterTo(g.instance, CheckPtr(ASibling))
}

func (g *TGroupBox) AnchorAsAlign(ATheAlign TAlign, ASpace int32) {
    GroupBox_AnchorAsAlign(g.instance, ATheAlign , ASpace)
}

func (g *TGroupBox) AnchorClient(ASpace int32) {
    GroupBox_AnchorClient(g.instance, ASpace)
}

// CN: 获取控件自动调整。
// EN: Get Control automatically adjusts.
func (g *TGroupBox) Align() TAlign {
    return GroupBox_GetAlign(g.instance)
}

// CN: 设置控件自动调整。
// EN: Set Control automatically adjusts.
func (g *TGroupBox) SetAlign(value TAlign) {
    GroupBox_SetAlign(g.instance, value)
}

// CN: 获取四个角位置的锚点。
// EN: .
func (g *TGroupBox) Anchors() TAnchors {
    return GroupBox_GetAnchors(g.instance)
}

// CN: 设置四个角位置的锚点。
// EN: .
func (g *TGroupBox) SetAnchors(value TAnchors) {
    GroupBox_SetAnchors(g.instance, value)
}

func (g *TGroupBox) BiDiMode() TBiDiMode {
    return GroupBox_GetBiDiMode(g.instance)
}

func (g *TGroupBox) SetBiDiMode(value TBiDiMode) {
    GroupBox_SetBiDiMode(g.instance, value)
}

// CN: 获取控件标题。
// EN: Get the control title.
func (g *TGroupBox) Caption() string {
    return GroupBox_GetCaption(g.instance)
}

// CN: 设置控件标题。
// EN: Set the control title.
func (g *TGroupBox) SetCaption(value string) {
    GroupBox_SetCaption(g.instance, value)
}

// CN: 获取颜色。
// EN: Get color.
func (g *TGroupBox) Color() TColor {
    return GroupBox_GetColor(g.instance)
}

// CN: 设置颜色。
// EN: Set color.
func (g *TGroupBox) SetColor(value TColor) {
    GroupBox_SetColor(g.instance, value)
}

// CN: 获取约束控件大小。
// EN: .
func (g *TGroupBox) Constraints() *TSizeConstraints {
    return AsSizeConstraints(GroupBox_GetConstraints(g.instance))
}

// CN: 设置约束控件大小。
// EN: .
func (g *TGroupBox) SetConstraints(value *TSizeConstraints) {
    GroupBox_SetConstraints(g.instance, CheckPtr(value))
}

// CN: 获取停靠站点。
// EN: Get Docking site.
func (g *TGroupBox) DockSite() bool {
    return GroupBox_GetDockSite(g.instance)
}

// CN: 设置停靠站点。
// EN: Set Docking site.
func (g *TGroupBox) SetDockSite(value bool) {
    GroupBox_SetDockSite(g.instance, value)
}

// CN: 获取设置控件双缓冲。
// EN: Get Set control double buffering.
func (g *TGroupBox) DoubleBuffered() bool {
    return GroupBox_GetDoubleBuffered(g.instance)
}

// CN: 设置设置控件双缓冲。
// EN: Set Set control double buffering.
func (g *TGroupBox) SetDoubleBuffered(value bool) {
    GroupBox_SetDoubleBuffered(g.instance, value)
}

// CN: 获取设置控件拖拽时的光标。
// EN: Get Set the cursor when the control is dragged.
func (g *TGroupBox) DragCursor() TCursor {
    return GroupBox_GetDragCursor(g.instance)
}

// CN: 设置设置控件拖拽时的光标。
// EN: Set Set the cursor when the control is dragged.
func (g *TGroupBox) SetDragCursor(value TCursor) {
    GroupBox_SetDragCursor(g.instance, value)
}

// CN: 获取拖拽方式。
// EN: Get Drag and drop.
func (g *TGroupBox) DragKind() TDragKind {
    return GroupBox_GetDragKind(g.instance)
}

// CN: 设置拖拽方式。
// EN: Set Drag and drop.
func (g *TGroupBox) SetDragKind(value TDragKind) {
    GroupBox_SetDragKind(g.instance, value)
}

// CN: 获取拖拽模式。
// EN: Get Drag mode.
func (g *TGroupBox) DragMode() TDragMode {
    return GroupBox_GetDragMode(g.instance)
}

// CN: 设置拖拽模式。
// EN: Set Drag mode.
func (g *TGroupBox) SetDragMode(value TDragMode) {
    GroupBox_SetDragMode(g.instance, value)
}

// CN: 获取控件启用。
// EN: Get the control enabled.
func (g *TGroupBox) Enabled() bool {
    return GroupBox_GetEnabled(g.instance)
}

// CN: 设置控件启用。
// EN: Set the control enabled.
func (g *TGroupBox) SetEnabled(value bool) {
    GroupBox_SetEnabled(g.instance, value)
}

// CN: 获取字体。
// EN: Get Font.
func (g *TGroupBox) Font() *TFont {
    return AsFont(GroupBox_GetFont(g.instance))
}

// CN: 设置字体。
// EN: Set Font.
func (g *TGroupBox) SetFont(value *TFont) {
    GroupBox_SetFont(g.instance, CheckPtr(value))
}

// CN: 获取父容器颜色。
// EN: Get parent color.
func (g *TGroupBox) ParentColor() bool {
    return GroupBox_GetParentColor(g.instance)
}

// CN: 设置父容器颜色。
// EN: Set parent color.
func (g *TGroupBox) SetParentColor(value bool) {
    GroupBox_SetParentColor(g.instance, value)
}

// CN: 获取父容器双缓冲。
// EN: Get Parent container double buffering.
func (g *TGroupBox) ParentDoubleBuffered() bool {
    return GroupBox_GetParentDoubleBuffered(g.instance)
}

// CN: 设置父容器双缓冲。
// EN: Set Parent container double buffering.
func (g *TGroupBox) SetParentDoubleBuffered(value bool) {
    GroupBox_SetParentDoubleBuffered(g.instance, value)
}

// CN: 获取父容器字体。
// EN: Get Parent container font.
func (g *TGroupBox) ParentFont() bool {
    return GroupBox_GetParentFont(g.instance)
}

// CN: 设置父容器字体。
// EN: Set Parent container font.
func (g *TGroupBox) SetParentFont(value bool) {
    GroupBox_SetParentFont(g.instance, value)
}

// CN: 获取以父容器的ShowHint属性为准。
// EN: .
func (g *TGroupBox) ParentShowHint() bool {
    return GroupBox_GetParentShowHint(g.instance)
}

// CN: 设置以父容器的ShowHint属性为准。
// EN: .
func (g *TGroupBox) SetParentShowHint(value bool) {
    GroupBox_SetParentShowHint(g.instance, value)
}

// CN: 获取右键菜单。
// EN: Get Right click menu.
func (g *TGroupBox) PopupMenu() *TPopupMenu {
    return AsPopupMenu(GroupBox_GetPopupMenu(g.instance))
}

// CN: 设置右键菜单。
// EN: Set Right click menu.
func (g *TGroupBox) SetPopupMenu(value IComponent) {
    GroupBox_SetPopupMenu(g.instance, CheckPtr(value))
}

// CN: 获取显示鼠标悬停提示。
// EN: Get Show mouseover tips.
func (g *TGroupBox) ShowHint() bool {
    return GroupBox_GetShowHint(g.instance)
}

// CN: 设置显示鼠标悬停提示。
// EN: Set Show mouseover tips.
func (g *TGroupBox) SetShowHint(value bool) {
    GroupBox_SetShowHint(g.instance, value)
}

// CN: 获取Tab切换顺序序号。
// EN: Get Tab switching sequence number.
func (g *TGroupBox) TabOrder() TTabOrder {
    return GroupBox_GetTabOrder(g.instance)
}

// CN: 设置Tab切换顺序序号。
// EN: Set Tab switching sequence number.
func (g *TGroupBox) SetTabOrder(value TTabOrder) {
    GroupBox_SetTabOrder(g.instance, value)
}

// CN: 获取Tab可停留。
// EN: Get Tab can stay.
func (g *TGroupBox) TabStop() bool {
    return GroupBox_GetTabStop(g.instance)
}

// CN: 设置Tab可停留。
// EN: Set Tab can stay.
func (g *TGroupBox) SetTabStop(value bool) {
    GroupBox_SetTabStop(g.instance, value)
}

// CN: 获取控件可视。
// EN: Get the control visible.
func (g *TGroupBox) Visible() bool {
    return GroupBox_GetVisible(g.instance)
}

// CN: 设置控件可视。
// EN: Set the control visible.
func (g *TGroupBox) SetVisible(value bool) {
    GroupBox_SetVisible(g.instance, value)
}

func (g *TGroupBox) SetOnAlignPosition(fn TAlignPositionEvent) {
    GroupBox_SetOnAlignPosition(g.instance, fn)
}

// CN: 设置控件单击事件。
// EN: Set control click event.
func (g *TGroupBox) SetOnClick(fn TNotifyEvent) {
    GroupBox_SetOnClick(g.instance, fn)
}

// CN: 设置上下文弹出事件，一般是右键时弹出。
// EN: Set Context popup event, usually pop up when right click.
func (g *TGroupBox) SetOnContextPopup(fn TContextPopupEvent) {
    GroupBox_SetOnContextPopup(g.instance, fn)
}

// CN: 设置双击事件。
// EN: .
func (g *TGroupBox) SetOnDblClick(fn TNotifyEvent) {
    GroupBox_SetOnDblClick(g.instance, fn)
}

// CN: 设置拖拽下落事件。
// EN: Set Drag and drop event.
func (g *TGroupBox) SetOnDragDrop(fn TDragDropEvent) {
    GroupBox_SetOnDragDrop(g.instance, fn)
}

func (g *TGroupBox) SetOnDockDrop(fn TDockDropEvent) {
    GroupBox_SetOnDockDrop(g.instance, fn)
}

// CN: 设置拖拽完成事件。
// EN: Set Drag and drop completion event.
func (g *TGroupBox) SetOnDragOver(fn TDragOverEvent) {
    GroupBox_SetOnDragOver(g.instance, fn)
}

// CN: 设置停靠结束事件。
// EN: Set Dock end event.
func (g *TGroupBox) SetOnEndDock(fn TEndDragEvent) {
    GroupBox_SetOnEndDock(g.instance, fn)
}

// CN: 设置拖拽结束。
// EN: Set End of drag.
func (g *TGroupBox) SetOnEndDrag(fn TEndDragEvent) {
    GroupBox_SetOnEndDrag(g.instance, fn)
}

// CN: 设置焦点进入。
// EN: Set Focus entry.
func (g *TGroupBox) SetOnEnter(fn TNotifyEvent) {
    GroupBox_SetOnEnter(g.instance, fn)
}

// CN: 设置焦点退出。
// EN: Set Focus exit.
func (g *TGroupBox) SetOnExit(fn TNotifyEvent) {
    GroupBox_SetOnExit(g.instance, fn)
}

func (g *TGroupBox) SetOnGetSiteInfo(fn TGetSiteInfoEvent) {
    GroupBox_SetOnGetSiteInfo(g.instance, fn)
}

// CN: 设置鼠标按下事件。
// EN: Set Mouse down event.
func (g *TGroupBox) SetOnMouseDown(fn TMouseEvent) {
    GroupBox_SetOnMouseDown(g.instance, fn)
}

// CN: 设置鼠标进入事件。
// EN: Set Mouse entry event.
func (g *TGroupBox) SetOnMouseEnter(fn TNotifyEvent) {
    GroupBox_SetOnMouseEnter(g.instance, fn)
}

// CN: 设置鼠标离开事件。
// EN: Set Mouse leave event.
func (g *TGroupBox) SetOnMouseLeave(fn TNotifyEvent) {
    GroupBox_SetOnMouseLeave(g.instance, fn)
}

// CN: 设置鼠标移动事件。
// EN: .
func (g *TGroupBox) SetOnMouseMove(fn TMouseMoveEvent) {
    GroupBox_SetOnMouseMove(g.instance, fn)
}

// CN: 设置鼠标抬起事件。
// EN: Set Mouse lift event.
func (g *TGroupBox) SetOnMouseUp(fn TMouseEvent) {
    GroupBox_SetOnMouseUp(g.instance, fn)
}

// CN: 设置启动停靠。
// EN: .
func (g *TGroupBox) SetOnStartDock(fn TStartDockEvent) {
    GroupBox_SetOnStartDock(g.instance, fn)
}

func (g *TGroupBox) SetOnUnDock(fn TUnDockEvent) {
    GroupBox_SetOnUnDock(g.instance, fn)
}

// CN: 获取依靠客户端总数。
// EN: .
func (g *TGroupBox) DockClientCount() int32 {
    return GroupBox_GetDockClientCount(g.instance)
}

// CN: 获取鼠标是否在客户端，仅VCL有效。
// EN: Get Whether the mouse is on the client, only VCL is valid.
func (g *TGroupBox) MouseInClient() bool {
    return GroupBox_GetMouseInClient(g.instance)
}

// CN: 获取当前停靠的可视总数。
// EN: Get The total number of visible calls currently docked.
func (g *TGroupBox) VisibleDockClientCount() int32 {
    return GroupBox_GetVisibleDockClientCount(g.instance)
}

// CN: 获取画刷对象。
// EN: Get Brush.
func (g *TGroupBox) Brush() *TBrush {
    return AsBrush(GroupBox_GetBrush(g.instance))
}

// CN: 获取子控件数。
// EN: Get Number of child controls.
func (g *TGroupBox) ControlCount() int32 {
    return GroupBox_GetControlCount(g.instance)
}

// CN: 获取控件句柄。
// EN: Get Control handle.
func (g *TGroupBox) Handle() HWND {
    return GroupBox_GetHandle(g.instance)
}

// CN: 获取父容器句柄。
// EN: Get Parent container handle.
func (g *TGroupBox) ParentWindow() HWND {
    return GroupBox_GetParentWindow(g.instance)
}

// CN: 设置父容器句柄。
// EN: Set Parent container handle.
func (g *TGroupBox) SetParentWindow(value HWND) {
    GroupBox_SetParentWindow(g.instance, value)
}

// CN: 获取使用停靠管理。
// EN: .
func (g *TGroupBox) UseDockManager() bool {
    return GroupBox_GetUseDockManager(g.instance)
}

// CN: 设置使用停靠管理。
// EN: .
func (g *TGroupBox) SetUseDockManager(value bool) {
    GroupBox_SetUseDockManager(g.instance, value)
}

func (g *TGroupBox) Action() *TAction {
    return AsAction(GroupBox_GetAction(g.instance))
}

func (g *TGroupBox) SetAction(value IComponent) {
    GroupBox_SetAction(g.instance, CheckPtr(value))
}

func (g *TGroupBox) BoundsRect() TRect {
    return GroupBox_GetBoundsRect(g.instance)
}

func (g *TGroupBox) SetBoundsRect(value TRect) {
    GroupBox_SetBoundsRect(g.instance, value)
}

// CN: 获取客户区高度。
// EN: Get client height.
func (g *TGroupBox) ClientHeight() int32 {
    return GroupBox_GetClientHeight(g.instance)
}

// CN: 设置客户区高度。
// EN: Set client height.
func (g *TGroupBox) SetClientHeight(value int32) {
    GroupBox_SetClientHeight(g.instance, value)
}

func (g *TGroupBox) ClientOrigin() TPoint {
    return GroupBox_GetClientOrigin(g.instance)
}

// CN: 获取客户区矩形。
// EN: Get client rectangle.
func (g *TGroupBox) ClientRect() TRect {
    return GroupBox_GetClientRect(g.instance)
}

// CN: 获取客户区宽度。
// EN: Get client width.
func (g *TGroupBox) ClientWidth() int32 {
    return GroupBox_GetClientWidth(g.instance)
}

// CN: 设置客户区宽度。
// EN: Set client width.
func (g *TGroupBox) SetClientWidth(value int32) {
    GroupBox_SetClientWidth(g.instance, value)
}

// CN: 获取控件状态。
// EN: Get control state.
func (g *TGroupBox) ControlState() TControlState {
    return GroupBox_GetControlState(g.instance)
}

// CN: 设置控件状态。
// EN: Set control state.
func (g *TGroupBox) SetControlState(value TControlState) {
    GroupBox_SetControlState(g.instance, value)
}

// CN: 获取控件样式。
// EN: Get control style.
func (g *TGroupBox) ControlStyle() TControlStyle {
    return GroupBox_GetControlStyle(g.instance)
}

// CN: 设置控件样式。
// EN: Set control style.
func (g *TGroupBox) SetControlStyle(value TControlStyle) {
    GroupBox_SetControlStyle(g.instance, value)
}

func (g *TGroupBox) Floating() bool {
    return GroupBox_GetFloating(g.instance)
}

// CN: 获取控件父容器。
// EN: Get control parent container.
func (g *TGroupBox) Parent() *TWinControl {
    return AsWinControl(GroupBox_GetParent(g.instance))
}

// CN: 设置控件父容器。
// EN: Set control parent container.
func (g *TGroupBox) SetParent(value IWinControl) {
    GroupBox_SetParent(g.instance, CheckPtr(value))
}

// CN: 获取左边位置。
// EN: Get Left position.
func (g *TGroupBox) Left() int32 {
    return GroupBox_GetLeft(g.instance)
}

// CN: 设置左边位置。
// EN: Set Left position.
func (g *TGroupBox) SetLeft(value int32) {
    GroupBox_SetLeft(g.instance, value)
}

// CN: 获取顶边位置。
// EN: Get Top position.
func (g *TGroupBox) Top() int32 {
    return GroupBox_GetTop(g.instance)
}

// CN: 设置顶边位置。
// EN: Set Top position.
func (g *TGroupBox) SetTop(value int32) {
    GroupBox_SetTop(g.instance, value)
}

// CN: 获取宽度。
// EN: Get width.
func (g *TGroupBox) Width() int32 {
    return GroupBox_GetWidth(g.instance)
}

// CN: 设置宽度。
// EN: Set width.
func (g *TGroupBox) SetWidth(value int32) {
    GroupBox_SetWidth(g.instance, value)
}

// CN: 获取高度。
// EN: Get height.
func (g *TGroupBox) Height() int32 {
    return GroupBox_GetHeight(g.instance)
}

// CN: 设置高度。
// EN: Set height.
func (g *TGroupBox) SetHeight(value int32) {
    GroupBox_SetHeight(g.instance, value)
}

// CN: 获取控件光标。
// EN: Get control cursor.
func (g *TGroupBox) Cursor() TCursor {
    return GroupBox_GetCursor(g.instance)
}

// CN: 设置控件光标。
// EN: Set control cursor.
func (g *TGroupBox) SetCursor(value TCursor) {
    GroupBox_SetCursor(g.instance, value)
}

// CN: 获取组件鼠标悬停提示。
// EN: Get component mouse hints.
func (g *TGroupBox) Hint() string {
    return GroupBox_GetHint(g.instance)
}

// CN: 设置组件鼠标悬停提示。
// EN: Set component mouse hints.
func (g *TGroupBox) SetHint(value string) {
    GroupBox_SetHint(g.instance, value)
}

// CN: 获取组件总数。
// EN: Get the total number of components.
func (g *TGroupBox) ComponentCount() int32 {
    return GroupBox_GetComponentCount(g.instance)
}

// CN: 获取组件索引。
// EN: Get component index.
func (g *TGroupBox) ComponentIndex() int32 {
    return GroupBox_GetComponentIndex(g.instance)
}

// CN: 设置组件索引。
// EN: Set component index.
func (g *TGroupBox) SetComponentIndex(value int32) {
    GroupBox_SetComponentIndex(g.instance, value)
}

// CN: 获取组件所有者。
// EN: Get component owner.
func (g *TGroupBox) Owner() *TComponent {
    return AsComponent(GroupBox_GetOwner(g.instance))
}

// CN: 获取组件名称。
// EN: Get the component name.
func (g *TGroupBox) Name() string {
    return GroupBox_GetName(g.instance)
}

// CN: 设置组件名称。
// EN: Set the component name.
func (g *TGroupBox) SetName(value string) {
    GroupBox_SetName(g.instance, value)
}

// CN: 获取对象标记。
// EN: Get the control tag.
func (g *TGroupBox) Tag() int {
    return GroupBox_GetTag(g.instance)
}

// CN: 设置对象标记。
// EN: Set the control tag.
func (g *TGroupBox) SetTag(value int) {
    GroupBox_SetTag(g.instance, value)
}

// CN: 获取左边锚点。
// EN: .
func (g *TGroupBox) AnchorSideLeft() *TAnchorSide {
    return AsAnchorSide(GroupBox_GetAnchorSideLeft(g.instance))
}

// CN: 设置左边锚点。
// EN: .
func (g *TGroupBox) SetAnchorSideLeft(value *TAnchorSide) {
    GroupBox_SetAnchorSideLeft(g.instance, CheckPtr(value))
}

// CN: 获取顶边锚点。
// EN: .
func (g *TGroupBox) AnchorSideTop() *TAnchorSide {
    return AsAnchorSide(GroupBox_GetAnchorSideTop(g.instance))
}

// CN: 设置顶边锚点。
// EN: .
func (g *TGroupBox) SetAnchorSideTop(value *TAnchorSide) {
    GroupBox_SetAnchorSideTop(g.instance, CheckPtr(value))
}

// CN: 获取右边锚点。
// EN: .
func (g *TGroupBox) AnchorSideRight() *TAnchorSide {
    return AsAnchorSide(GroupBox_GetAnchorSideRight(g.instance))
}

// CN: 设置右边锚点。
// EN: .
func (g *TGroupBox) SetAnchorSideRight(value *TAnchorSide) {
    GroupBox_SetAnchorSideRight(g.instance, CheckPtr(value))
}

// CN: 获取底边锚点。
// EN: .
func (g *TGroupBox) AnchorSideBottom() *TAnchorSide {
    return AsAnchorSide(GroupBox_GetAnchorSideBottom(g.instance))
}

// CN: 设置底边锚点。
// EN: .
func (g *TGroupBox) SetAnchorSideBottom(value *TAnchorSide) {
    GroupBox_SetAnchorSideBottom(g.instance, CheckPtr(value))
}

func (g *TGroupBox) ChildSizing() *TControlChildSizing {
    return AsControlChildSizing(GroupBox_GetChildSizing(g.instance))
}

func (g *TGroupBox) SetChildSizing(value *TControlChildSizing) {
    GroupBox_SetChildSizing(g.instance, CheckPtr(value))
}

// CN: 获取边框间距。
// EN: .
func (g *TGroupBox) BorderSpacing() *TControlBorderSpacing {
    return AsControlBorderSpacing(GroupBox_GetBorderSpacing(g.instance))
}

// CN: 设置边框间距。
// EN: .
func (g *TGroupBox) SetBorderSpacing(value *TControlBorderSpacing) {
    GroupBox_SetBorderSpacing(g.instance, CheckPtr(value))
}

// CN: 获取指定索引停靠客户端。
// EN: .
func (g *TGroupBox) DockClients(Index int32) *TControl {
    return AsControl(GroupBox_GetDockClients(g.instance, Index))
}

// CN: 获取指定索引子控件。
// EN: .
func (g *TGroupBox) Controls(Index int32) *TControl {
    return AsControl(GroupBox_GetControls(g.instance, Index))
}

// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (g *TGroupBox) Components(AIndex int32) *TComponent {
    return AsComponent(GroupBox_GetComponents(g.instance, AIndex))
}

// CN: 获取锚侧面。
// EN: .
func (g *TGroupBox) AnchorSide(AKind TAnchorKind) *TAnchorSide {
    return AsAnchorSide(GroupBox_GetAnchorSide(g.instance, AKind))
}

