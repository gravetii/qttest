

#define protected public
#define private public

#include "moc.h"
#include "_cgo_export.h"

#include <QByteArray>
#include <QChildEvent>
#include <QDragEnterEvent>
#include <QDragLeaveEvent>
#include <QDragMoveEvent>
#include <QDropEvent>
#include <QEvent>
#include <QFocusEvent>
#include <QHoverEvent>
#include <QInputMethodEvent>
#include <QKeyEvent>
#include <QMap>
#include <QMetaMethod>
#include <QMouseEvent>
#include <QObject>
#include <QPointF>
#include <QQuickItem>
#include <QQuickWindow>
#include <QRectF>
#include <QSGTextureProvider>
#include <QString>
#include <QTimerEvent>
#include <QTouchEvent>
#include <QVariant>
#include <QVector>
#include <QWheelEvent>

#ifdef QT_QML_LIB
	#include <QQmlEngine>
#endif


typedef QMap<QString, QVariant> type424d06;
class ItemTemplate40f1b6: public QQuickItem
{
Q_OBJECT
Q_PROPERTY(QString someString READ someString WRITE setSomeString NOTIFY someStringChanged)
public:
	ItemTemplate40f1b6(QQuickItem *parent = Q_NULLPTR) : QQuickItem(parent) {qRegisterMetaType<quintptr>("quintptr");ItemTemplate40f1b6_ItemTemplate40f1b6_QRegisterMetaType();ItemTemplate40f1b6_ItemTemplate40f1b6_QRegisterMetaTypes();callbackItemTemplate40f1b6_Constructor(this);};
	void Signal_SendBool(bool v0, QList<bool> v1) { callbackItemTemplate40f1b6_SendBool(this, v0, ({ QList<bool>* tmpValue = new QList<bool>(v1); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	void Signal_SendInt(qint32 v0, QList<qint32> v1) { callbackItemTemplate40f1b6_SendInt(this, v0, ({ QList<qint32>* tmpValue = new QList<qint32>(v1); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	void Signal_SendFloat(float v0) { callbackItemTemplate40f1b6_SendFloat(this, v0); };
	void Signal_SendDouble(qreal v0, QList<qreal> v1) { callbackItemTemplate40f1b6_SendDouble(this, v0, ({ QList<qreal>* tmpValue = new QList<qreal>(v1); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	void Signal_SendString(QString v0, QStringList v1) { QByteArray tea1dd7 = v0.toUtf8(); Moc_PackedString v0Packed = { const_cast<char*>(tea1dd7.prepend("WHITESPACE").constData()+10), tea1dd7.size()-10 };QByteArray t5a6df7 = v1.join("¡¦!").toUtf8(); Moc_PackedString v1Packed = { const_cast<char*>(t5a6df7.prepend("WHITESPACE").constData()+10), t5a6df7.size()-10 };callbackItemTemplate40f1b6_SendString(this, v0Packed, v1Packed); };
	void Signal_SendError(QString v0, QList<QString> v1) { QByteArray tea1dd7 = v0.toUtf8(); Moc_PackedString v0Packed = { const_cast<char*>(tea1dd7.prepend("WHITESPACE").constData()+10), tea1dd7.size()-10 };callbackItemTemplate40f1b6_SendError(this, v0Packed, ({ QList<QString>* tmpValue = new QList<QString>(v1); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	void Signal_SendVariantListMap(QVariant v0, QList<QVariant> v1, type424d06 v2) { callbackItemTemplate40f1b6_SendVariantListMap(this, new QVariant(v0), ({ QList<QVariant>* tmpValue = new QList<QVariant>(v1); Moc_PackedList { tmpValue, tmpValue->size() }; }), ({ QMap<QString, QVariant>* tmpValue = new QMap<QString, QVariant>(v2); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	void Signal_SendItemTemplate(ItemTemplate40f1b6* v0, QList<ItemTemplate40f1b6*> v1) { callbackItemTemplate40f1b6_SendItemTemplate(this, v0, ({ QList<ItemTemplate40f1b6*>* tmpValue = new QList<ItemTemplate40f1b6*>(v1); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	void Signal_SendItem(QQuickItem* v0, QList<QQuickItem*> v1) { callbackItemTemplate40f1b6_SendItem(this, v0, ({ QList<QQuickItem*>* tmpValue = new QList<QQuickItem*>(v1); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	void Signal_SendObject(QObject* v0, QList<QObject*> v1) { callbackItemTemplate40f1b6_SendObject(this, v0, ({ QList<QObject*>* tmpValue = new QList<QObject*>(v1); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	QString someString() { return ({ Moc_PackedString tempVal = callbackItemTemplate40f1b6_SomeString(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setSomeString(QString someString) { QByteArray te0f92a = someString.toUtf8(); Moc_PackedString someStringPacked = { const_cast<char*>(te0f92a.prepend("WHITESPACE").constData()+10), te0f92a.size()-10 };callbackItemTemplate40f1b6_SetSomeString(this, someStringPacked); };
	void Signal_SomeStringChanged(QString someString) { QByteArray te0f92a = someString.toUtf8(); Moc_PackedString someStringPacked = { const_cast<char*>(te0f92a.prepend("WHITESPACE").constData()+10), te0f92a.size()-10 };callbackItemTemplate40f1b6_SomeStringChanged(this, someStringPacked); };
	 ~ItemTemplate40f1b6() { callbackItemTemplate40f1b6_DestroyItemTemplate(this); };
	void Signal_ActiveFocusChanged(bool vbo) { callbackItemTemplate40f1b6_ActiveFocusChanged(this, vbo); };
	void Signal_ActiveFocusOnTabChanged(bool vbo) { callbackItemTemplate40f1b6_ActiveFocusOnTabChanged(this, vbo); };
	void Signal_AntialiasingChanged(bool vbo) { callbackItemTemplate40f1b6_AntialiasingChanged(this, vbo); };
	void Signal_BaselineOffsetChanged(qreal vqr) { callbackItemTemplate40f1b6_BaselineOffsetChanged(this, vqr); };
	bool childMouseEventFilter(QQuickItem * item, QEvent * event) { return callbackItemTemplate40f1b6_ChildMouseEventFilter(this, item, event) != 0; };
	void Signal_ChildrenRectChanged(const QRectF & vqr) { callbackItemTemplate40f1b6_ChildrenRectChanged(this, const_cast<QRectF*>(&vqr)); };
	void classBegin() { callbackItemTemplate40f1b6_ClassBegin(this); };
	void Signal_ClipChanged(bool vbo) { callbackItemTemplate40f1b6_ClipChanged(this, vbo); };
	void componentComplete() { callbackItemTemplate40f1b6_ComponentComplete(this); };
	void Signal_ContainmentMaskChanged() { callbackItemTemplate40f1b6_ContainmentMaskChanged(this); };
	bool contains(const QPointF & point) const { return callbackItemTemplate40f1b6_Contains(const_cast<void*>(static_cast<const void*>(this)), const_cast<QPointF*>(&point)) != 0; };
	void dragEnterEvent(QDragEnterEvent * event) { callbackItemTemplate40f1b6_DragEnterEvent(this, event); };
	void dragLeaveEvent(QDragLeaveEvent * event) { callbackItemTemplate40f1b6_DragLeaveEvent(this, event); };
	void dragMoveEvent(QDragMoveEvent * event) { callbackItemTemplate40f1b6_DragMoveEvent(this, event); };
	void dropEvent(QDropEvent * event) { callbackItemTemplate40f1b6_DropEvent(this, event); };
	void Signal_EnabledChanged() { callbackItemTemplate40f1b6_EnabledChanged(this); };
	bool event(QEvent * ev) { return callbackItemTemplate40f1b6_Event(this, ev) != 0; };
	void Signal_FocusChanged(bool vbo) { callbackItemTemplate40f1b6_FocusChanged(this, vbo); };
	void focusInEvent(QFocusEvent * vqf) { callbackItemTemplate40f1b6_FocusInEvent(this, vqf); };
	void focusOutEvent(QFocusEvent * vqf) { callbackItemTemplate40f1b6_FocusOutEvent(this, vqf); };
	void geometryChanged(const QRectF & newGeometry, const QRectF & oldGeometry) { callbackItemTemplate40f1b6_GeometryChanged(this, const_cast<QRectF*>(&newGeometry), const_cast<QRectF*>(&oldGeometry)); };
	void Signal_HeightChanged() { callbackItemTemplate40f1b6_HeightChanged(this); };
	void hoverEnterEvent(QHoverEvent * event) { callbackItemTemplate40f1b6_HoverEnterEvent(this, event); };
	void hoverLeaveEvent(QHoverEvent * event) { callbackItemTemplate40f1b6_HoverLeaveEvent(this, event); };
	void hoverMoveEvent(QHoverEvent * event) { callbackItemTemplate40f1b6_HoverMoveEvent(this, event); };
	void Signal_ImplicitHeightChanged() { callbackItemTemplate40f1b6_ImplicitHeightChanged(this); };
	void Signal_ImplicitWidthChanged() { callbackItemTemplate40f1b6_ImplicitWidthChanged(this); };
	void inputMethodEvent(QInputMethodEvent * event) { callbackItemTemplate40f1b6_InputMethodEvent(this, event); };
	QVariant inputMethodQuery(Qt::InputMethodQuery query) const { return *static_cast<QVariant*>(callbackItemTemplate40f1b6_InputMethodQuery(const_cast<void*>(static_cast<const void*>(this)), query)); };
	bool isTextureProvider() const { return callbackItemTemplate40f1b6_IsTextureProvider(const_cast<void*>(static_cast<const void*>(this))) != 0; };
	void keyPressEvent(QKeyEvent * event) { callbackItemTemplate40f1b6_KeyPressEvent(this, event); };
	void keyReleaseEvent(QKeyEvent * event) { callbackItemTemplate40f1b6_KeyReleaseEvent(this, event); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackItemTemplate40f1b6_MouseDoubleClickEvent(this, event); };
	void mouseMoveEvent(QMouseEvent * event) { callbackItemTemplate40f1b6_MouseMoveEvent(this, event); };
	void mousePressEvent(QMouseEvent * event) { callbackItemTemplate40f1b6_MousePressEvent(this, event); };
	void mouseReleaseEvent(QMouseEvent * event) { callbackItemTemplate40f1b6_MouseReleaseEvent(this, event); };
	void mouseUngrabEvent() { callbackItemTemplate40f1b6_MouseUngrabEvent(this); };
	void Signal_OpacityChanged() { callbackItemTemplate40f1b6_OpacityChanged(this); };
	void Signal_ParentChanged(QQuickItem * vqq) { callbackItemTemplate40f1b6_ParentChanged(this, vqq); };
	void releaseResources() { callbackItemTemplate40f1b6_ReleaseResources(this); };
	void Signal_RotationChanged() { callbackItemTemplate40f1b6_RotationChanged(this); };
	void Signal_ScaleChanged() { callbackItemTemplate40f1b6_ScaleChanged(this); };
	void Signal_SmoothChanged(bool vbo) { callbackItemTemplate40f1b6_SmoothChanged(this, vbo); };
	void Signal_StateChanged(const QString & vqs) { QByteArray tda39a3 = vqs.toUtf8(); Moc_PackedString vqsPacked = { const_cast<char*>(tda39a3.prepend("WHITESPACE").constData()+10), tda39a3.size()-10 };callbackItemTemplate40f1b6_StateChanged(this, vqsPacked); };
	QSGTextureProvider * textureProvider() const { return static_cast<QSGTextureProvider*>(callbackItemTemplate40f1b6_TextureProvider(const_cast<void*>(static_cast<const void*>(this)))); };
	void touchEvent(QTouchEvent * event) { callbackItemTemplate40f1b6_TouchEvent(this, event); };
	void touchUngrabEvent() { callbackItemTemplate40f1b6_TouchUngrabEvent(this); };
	void Signal_TransformOriginChanged(QQuickItem::TransformOrigin vqq) { callbackItemTemplate40f1b6_TransformOriginChanged(this, vqq); };
	void update() { callbackItemTemplate40f1b6_Update(this); };
	void updatePolish() { callbackItemTemplate40f1b6_UpdatePolish(this); };
	void Signal_VisibleChanged() { callbackItemTemplate40f1b6_VisibleChanged(this); };
	void wheelEvent(QWheelEvent * event) { callbackItemTemplate40f1b6_WheelEvent(this, event); };
	void Signal_WidthChanged() { callbackItemTemplate40f1b6_WidthChanged(this); };
	void Signal_WindowChanged(QQuickWindow * window) { callbackItemTemplate40f1b6_WindowChanged(this, window); };
	void Signal_XChanged() { callbackItemTemplate40f1b6_XChanged(this); };
	void Signal_YChanged() { callbackItemTemplate40f1b6_YChanged(this); };
	void Signal_ZChanged() { callbackItemTemplate40f1b6_ZChanged(this); };
	void childEvent(QChildEvent * event) { callbackItemTemplate40f1b6_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackItemTemplate40f1b6_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackItemTemplate40f1b6_CustomEvent(this, event); };
	void deleteLater() { callbackItemTemplate40f1b6_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackItemTemplate40f1b6_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackItemTemplate40f1b6_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackItemTemplate40f1b6_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackItemTemplate40f1b6_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackItemTemplate40f1b6_TimerEvent(this, event); };
	QString someStringDefault() { return _someString; };
	void setSomeStringDefault(QString p) { if (p != _someString) { _someString = p; someStringChanged(_someString); } };
signals:
	void sendBool(bool v0, QList<bool> v1);
	void sendInt(qint32 v0, QList<qint32> v1);
	void sendFloat(float v0);
	void sendDouble(qreal v0, QList<qreal> v1);
	void sendString(QString v0, QStringList v1);
	void sendError(QString v0, QList<QString> v1);
	void sendVariantListMap(QVariant v0, QList<QVariant> v1, type424d06 v2);
	void sendItemTemplate(ItemTemplate40f1b6* v0, QList<ItemTemplate40f1b6*> v1);
	void sendItem(QQuickItem* v0, QList<QQuickItem*> v1);
	void sendObject(QObject* v0, QList<QObject*> v1);
	void someStringChanged(QString someString);
public slots:
private:
	QString _someString;
};

Q_DECLARE_METATYPE(ItemTemplate40f1b6*)


void ItemTemplate40f1b6_ItemTemplate40f1b6_QRegisterMetaTypes() {
	qRegisterMetaType<type424d06>("type424d06");
	qRegisterMetaType<QList<QObject*>>("QList<ItemTemplate40f1b6*>");
	qRegisterMetaType<QList<QObject*>>("QList<QQuickItem*>");
	qRegisterMetaType<QString>();
}

void ItemTemplate40f1b6_ConnectSendBool(void* ptr)
{
	QObject::connect(static_cast<ItemTemplate40f1b6*>(ptr), static_cast<void (ItemTemplate40f1b6::*)(bool, QList<bool>)>(&ItemTemplate40f1b6::sendBool), static_cast<ItemTemplate40f1b6*>(ptr), static_cast<void (ItemTemplate40f1b6::*)(bool, QList<bool>)>(&ItemTemplate40f1b6::Signal_SendBool));
}

void ItemTemplate40f1b6_DisconnectSendBool(void* ptr)
{
	QObject::disconnect(static_cast<ItemTemplate40f1b6*>(ptr), static_cast<void (ItemTemplate40f1b6::*)(bool, QList<bool>)>(&ItemTemplate40f1b6::sendBool), static_cast<ItemTemplate40f1b6*>(ptr), static_cast<void (ItemTemplate40f1b6::*)(bool, QList<bool>)>(&ItemTemplate40f1b6::Signal_SendBool));
}

void ItemTemplate40f1b6_SendBool(void* ptr, char v0, void* v1)
{
	static_cast<ItemTemplate40f1b6*>(ptr)->sendBool(v0 != 0, ({ QList<bool>* tmpP = static_cast<QList<bool>*>(v1); QList<bool> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

void ItemTemplate40f1b6_ConnectSendInt(void* ptr)
{
	QObject::connect(static_cast<ItemTemplate40f1b6*>(ptr), static_cast<void (ItemTemplate40f1b6::*)(qint32, QList<qint32>)>(&ItemTemplate40f1b6::sendInt), static_cast<ItemTemplate40f1b6*>(ptr), static_cast<void (ItemTemplate40f1b6::*)(qint32, QList<qint32>)>(&ItemTemplate40f1b6::Signal_SendInt));
}

void ItemTemplate40f1b6_DisconnectSendInt(void* ptr)
{
	QObject::disconnect(static_cast<ItemTemplate40f1b6*>(ptr), static_cast<void (ItemTemplate40f1b6::*)(qint32, QList<qint32>)>(&ItemTemplate40f1b6::sendInt), static_cast<ItemTemplate40f1b6*>(ptr), static_cast<void (ItemTemplate40f1b6::*)(qint32, QList<qint32>)>(&ItemTemplate40f1b6::Signal_SendInt));
}

void ItemTemplate40f1b6_SendInt(void* ptr, int v0, void* v1)
{
	static_cast<ItemTemplate40f1b6*>(ptr)->sendInt(v0, ({ QList<qint32>* tmpP = static_cast<QList<qint32>*>(v1); QList<qint32> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

void ItemTemplate40f1b6_ConnectSendFloat(void* ptr)
{
	QObject::connect(static_cast<ItemTemplate40f1b6*>(ptr), static_cast<void (ItemTemplate40f1b6::*)(float)>(&ItemTemplate40f1b6::sendFloat), static_cast<ItemTemplate40f1b6*>(ptr), static_cast<void (ItemTemplate40f1b6::*)(float)>(&ItemTemplate40f1b6::Signal_SendFloat));
}

void ItemTemplate40f1b6_DisconnectSendFloat(void* ptr)
{
	QObject::disconnect(static_cast<ItemTemplate40f1b6*>(ptr), static_cast<void (ItemTemplate40f1b6::*)(float)>(&ItemTemplate40f1b6::sendFloat), static_cast<ItemTemplate40f1b6*>(ptr), static_cast<void (ItemTemplate40f1b6::*)(float)>(&ItemTemplate40f1b6::Signal_SendFloat));
}

void ItemTemplate40f1b6_SendFloat(void* ptr, float v0)
{
	static_cast<ItemTemplate40f1b6*>(ptr)->sendFloat(v0);
}

void ItemTemplate40f1b6_ConnectSendDouble(void* ptr)
{
	QObject::connect(static_cast<ItemTemplate40f1b6*>(ptr), static_cast<void (ItemTemplate40f1b6::*)(qreal, QList<qreal>)>(&ItemTemplate40f1b6::sendDouble), static_cast<ItemTemplate40f1b6*>(ptr), static_cast<void (ItemTemplate40f1b6::*)(qreal, QList<qreal>)>(&ItemTemplate40f1b6::Signal_SendDouble));
}

void ItemTemplate40f1b6_DisconnectSendDouble(void* ptr)
{
	QObject::disconnect(static_cast<ItemTemplate40f1b6*>(ptr), static_cast<void (ItemTemplate40f1b6::*)(qreal, QList<qreal>)>(&ItemTemplate40f1b6::sendDouble), static_cast<ItemTemplate40f1b6*>(ptr), static_cast<void (ItemTemplate40f1b6::*)(qreal, QList<qreal>)>(&ItemTemplate40f1b6::Signal_SendDouble));
}

void ItemTemplate40f1b6_SendDouble(void* ptr, double v0, void* v1)
{
	static_cast<ItemTemplate40f1b6*>(ptr)->sendDouble(v0, ({ QList<qreal>* tmpP = static_cast<QList<qreal>*>(v1); QList<qreal> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

void ItemTemplate40f1b6_ConnectSendString(void* ptr)
{
	QObject::connect(static_cast<ItemTemplate40f1b6*>(ptr), static_cast<void (ItemTemplate40f1b6::*)(QString, QStringList)>(&ItemTemplate40f1b6::sendString), static_cast<ItemTemplate40f1b6*>(ptr), static_cast<void (ItemTemplate40f1b6::*)(QString, QStringList)>(&ItemTemplate40f1b6::Signal_SendString));
}

void ItemTemplate40f1b6_DisconnectSendString(void* ptr)
{
	QObject::disconnect(static_cast<ItemTemplate40f1b6*>(ptr), static_cast<void (ItemTemplate40f1b6::*)(QString, QStringList)>(&ItemTemplate40f1b6::sendString), static_cast<ItemTemplate40f1b6*>(ptr), static_cast<void (ItemTemplate40f1b6::*)(QString, QStringList)>(&ItemTemplate40f1b6::Signal_SendString));
}

void ItemTemplate40f1b6_SendString(void* ptr, struct Moc_PackedString v0, struct Moc_PackedString v1)
{
	static_cast<ItemTemplate40f1b6*>(ptr)->sendString(QString::fromUtf8(v0.data, v0.len), QString::fromUtf8(v1.data, v1.len).split("¡¦!", QString::SkipEmptyParts));
}

void ItemTemplate40f1b6_ConnectSendError(void* ptr)
{
	QObject::connect(static_cast<ItemTemplate40f1b6*>(ptr), static_cast<void (ItemTemplate40f1b6::*)(QString, QList<QString>)>(&ItemTemplate40f1b6::sendError), static_cast<ItemTemplate40f1b6*>(ptr), static_cast<void (ItemTemplate40f1b6::*)(QString, QList<QString>)>(&ItemTemplate40f1b6::Signal_SendError));
}

void ItemTemplate40f1b6_DisconnectSendError(void* ptr)
{
	QObject::disconnect(static_cast<ItemTemplate40f1b6*>(ptr), static_cast<void (ItemTemplate40f1b6::*)(QString, QList<QString>)>(&ItemTemplate40f1b6::sendError), static_cast<ItemTemplate40f1b6*>(ptr), static_cast<void (ItemTemplate40f1b6::*)(QString, QList<QString>)>(&ItemTemplate40f1b6::Signal_SendError));
}

void ItemTemplate40f1b6_SendError(void* ptr, struct Moc_PackedString v0, void* v1)
{
	static_cast<ItemTemplate40f1b6*>(ptr)->sendError(QString::fromUtf8(v0.data, v0.len), ({ QList<QString>* tmpP = static_cast<QList<QString>*>(v1); QList<QString> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

void ItemTemplate40f1b6_ConnectSendVariantListMap(void* ptr)
{
	QObject::connect(static_cast<ItemTemplate40f1b6*>(ptr), static_cast<void (ItemTemplate40f1b6::*)(QVariant, QList<QVariant>, QMap<QString, QVariant>)>(&ItemTemplate40f1b6::sendVariantListMap), static_cast<ItemTemplate40f1b6*>(ptr), static_cast<void (ItemTemplate40f1b6::*)(QVariant, QList<QVariant>, QMap<QString, QVariant>)>(&ItemTemplate40f1b6::Signal_SendVariantListMap));
}

void ItemTemplate40f1b6_DisconnectSendVariantListMap(void* ptr)
{
	QObject::disconnect(static_cast<ItemTemplate40f1b6*>(ptr), static_cast<void (ItemTemplate40f1b6::*)(QVariant, QList<QVariant>, QMap<QString, QVariant>)>(&ItemTemplate40f1b6::sendVariantListMap), static_cast<ItemTemplate40f1b6*>(ptr), static_cast<void (ItemTemplate40f1b6::*)(QVariant, QList<QVariant>, QMap<QString, QVariant>)>(&ItemTemplate40f1b6::Signal_SendVariantListMap));
}

void ItemTemplate40f1b6_SendVariantListMap(void* ptr, void* v0, void* v1, void* v2)
{
	static_cast<ItemTemplate40f1b6*>(ptr)->sendVariantListMap(*static_cast<QVariant*>(v0), ({ QList<QVariant>* tmpP = static_cast<QList<QVariant>*>(v1); QList<QVariant> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }), ({ QMap<QString, QVariant>* tmpP = static_cast<QMap<QString, QVariant>*>(v2); QMap<QString, QVariant> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }));
}

void ItemTemplate40f1b6_ConnectSendItemTemplate(void* ptr)
{
	QObject::connect(static_cast<ItemTemplate40f1b6*>(ptr), static_cast<void (ItemTemplate40f1b6::*)(ItemTemplate40f1b6*, QList<ItemTemplate40f1b6*>)>(&ItemTemplate40f1b6::sendItemTemplate), static_cast<ItemTemplate40f1b6*>(ptr), static_cast<void (ItemTemplate40f1b6::*)(ItemTemplate40f1b6*, QList<ItemTemplate40f1b6*>)>(&ItemTemplate40f1b6::Signal_SendItemTemplate));
}

void ItemTemplate40f1b6_DisconnectSendItemTemplate(void* ptr)
{
	QObject::disconnect(static_cast<ItemTemplate40f1b6*>(ptr), static_cast<void (ItemTemplate40f1b6::*)(ItemTemplate40f1b6*, QList<ItemTemplate40f1b6*>)>(&ItemTemplate40f1b6::sendItemTemplate), static_cast<ItemTemplate40f1b6*>(ptr), static_cast<void (ItemTemplate40f1b6::*)(ItemTemplate40f1b6*, QList<ItemTemplate40f1b6*>)>(&ItemTemplate40f1b6::Signal_SendItemTemplate));
}

void ItemTemplate40f1b6_SendItemTemplate(void* ptr, void* v0, void* v1)
{
	static_cast<ItemTemplate40f1b6*>(ptr)->sendItemTemplate(static_cast<ItemTemplate40f1b6*>(v0), ({ QList<ItemTemplate40f1b6*>* tmpP = static_cast<QList<ItemTemplate40f1b6*>*>(v1); QList<ItemTemplate40f1b6*> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

void ItemTemplate40f1b6_ConnectSendItem(void* ptr)
{
	QObject::connect(static_cast<ItemTemplate40f1b6*>(ptr), static_cast<void (ItemTemplate40f1b6::*)(QQuickItem*, QList<QQuickItem*>)>(&ItemTemplate40f1b6::sendItem), static_cast<ItemTemplate40f1b6*>(ptr), static_cast<void (ItemTemplate40f1b6::*)(QQuickItem*, QList<QQuickItem*>)>(&ItemTemplate40f1b6::Signal_SendItem));
}

void ItemTemplate40f1b6_DisconnectSendItem(void* ptr)
{
	QObject::disconnect(static_cast<ItemTemplate40f1b6*>(ptr), static_cast<void (ItemTemplate40f1b6::*)(QQuickItem*, QList<QQuickItem*>)>(&ItemTemplate40f1b6::sendItem), static_cast<ItemTemplate40f1b6*>(ptr), static_cast<void (ItemTemplate40f1b6::*)(QQuickItem*, QList<QQuickItem*>)>(&ItemTemplate40f1b6::Signal_SendItem));
}

void ItemTemplate40f1b6_SendItem(void* ptr, void* v0, void* v1)
{
	static_cast<ItemTemplate40f1b6*>(ptr)->sendItem(static_cast<QQuickItem*>(v0), ({ QList<QQuickItem*>* tmpP = static_cast<QList<QQuickItem*>*>(v1); QList<QQuickItem*> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

void ItemTemplate40f1b6_ConnectSendObject(void* ptr)
{
	QObject::connect(static_cast<ItemTemplate40f1b6*>(ptr), static_cast<void (ItemTemplate40f1b6::*)(QObject*, QList<QObject*>)>(&ItemTemplate40f1b6::sendObject), static_cast<ItemTemplate40f1b6*>(ptr), static_cast<void (ItemTemplate40f1b6::*)(QObject*, QList<QObject*>)>(&ItemTemplate40f1b6::Signal_SendObject));
}

void ItemTemplate40f1b6_DisconnectSendObject(void* ptr)
{
	QObject::disconnect(static_cast<ItemTemplate40f1b6*>(ptr), static_cast<void (ItemTemplate40f1b6::*)(QObject*, QList<QObject*>)>(&ItemTemplate40f1b6::sendObject), static_cast<ItemTemplate40f1b6*>(ptr), static_cast<void (ItemTemplate40f1b6::*)(QObject*, QList<QObject*>)>(&ItemTemplate40f1b6::Signal_SendObject));
}

void ItemTemplate40f1b6_SendObject(void* ptr, void* v0, void* v1)
{
	static_cast<ItemTemplate40f1b6*>(ptr)->sendObject(static_cast<QObject*>(v0), ({ QList<QObject*>* tmpP = static_cast<QList<QObject*>*>(v1); QList<QObject*> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

struct Moc_PackedString ItemTemplate40f1b6_SomeString(void* ptr)
{
	return ({ QByteArray tb26607 = static_cast<ItemTemplate40f1b6*>(ptr)->someString().toUtf8(); Moc_PackedString { const_cast<char*>(tb26607.prepend("WHITESPACE").constData()+10), tb26607.size()-10 }; });
}

struct Moc_PackedString ItemTemplate40f1b6_SomeStringDefault(void* ptr)
{
	return ({ QByteArray t207da5 = static_cast<ItemTemplate40f1b6*>(ptr)->someStringDefault().toUtf8(); Moc_PackedString { const_cast<char*>(t207da5.prepend("WHITESPACE").constData()+10), t207da5.size()-10 }; });
}

void ItemTemplate40f1b6_SetSomeString(void* ptr, struct Moc_PackedString someString)
{
	static_cast<ItemTemplate40f1b6*>(ptr)->setSomeString(QString::fromUtf8(someString.data, someString.len));
}

void ItemTemplate40f1b6_SetSomeStringDefault(void* ptr, struct Moc_PackedString someString)
{
	static_cast<ItemTemplate40f1b6*>(ptr)->setSomeStringDefault(QString::fromUtf8(someString.data, someString.len));
}

void ItemTemplate40f1b6_ConnectSomeStringChanged(void* ptr)
{
	QObject::connect(static_cast<ItemTemplate40f1b6*>(ptr), static_cast<void (ItemTemplate40f1b6::*)(QString)>(&ItemTemplate40f1b6::someStringChanged), static_cast<ItemTemplate40f1b6*>(ptr), static_cast<void (ItemTemplate40f1b6::*)(QString)>(&ItemTemplate40f1b6::Signal_SomeStringChanged));
}

void ItemTemplate40f1b6_DisconnectSomeStringChanged(void* ptr)
{
	QObject::disconnect(static_cast<ItemTemplate40f1b6*>(ptr), static_cast<void (ItemTemplate40f1b6::*)(QString)>(&ItemTemplate40f1b6::someStringChanged), static_cast<ItemTemplate40f1b6*>(ptr), static_cast<void (ItemTemplate40f1b6::*)(QString)>(&ItemTemplate40f1b6::Signal_SomeStringChanged));
}

void ItemTemplate40f1b6_SomeStringChanged(void* ptr, struct Moc_PackedString someString)
{
	static_cast<ItemTemplate40f1b6*>(ptr)->someStringChanged(QString::fromUtf8(someString.data, someString.len));
}

int ItemTemplate40f1b6_ItemTemplate40f1b6_QRegisterMetaType()
{
	return qRegisterMetaType<ItemTemplate40f1b6*>();
}

int ItemTemplate40f1b6_ItemTemplate40f1b6_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<ItemTemplate40f1b6*>(const_cast<const char*>(typeName));
}

int ItemTemplate40f1b6_ItemTemplate40f1b6_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<ItemTemplate40f1b6>();
#else
	return 0;
#endif
}

int ItemTemplate40f1b6_ItemTemplate40f1b6_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<ItemTemplate40f1b6>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

void* ItemTemplate40f1b6___childItems_atList(void* ptr, int i)
{
	return ({QQuickItem * tmp = static_cast<QList<QQuickItem *>*>(ptr)->at(i); if (i == static_cast<QList<QQuickItem *>*>(ptr)->size()-1) { static_cast<QList<QQuickItem *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void ItemTemplate40f1b6___childItems_setList(void* ptr, void* i)
{
	static_cast<QList<QQuickItem *>*>(ptr)->append(static_cast<QQuickItem*>(i));
}

void* ItemTemplate40f1b6___childItems_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QQuickItem *>();
}

int ItemTemplate40f1b6___grabTouchPoints_ids_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QVector<int>*>(ptr)->at(i); if (i == static_cast<QVector<int>*>(ptr)->size()-1) { static_cast<QVector<int>*>(ptr)->~QVector(); free(ptr); }; tmp; });
}

void ItemTemplate40f1b6___grabTouchPoints_ids_setList(void* ptr, int i)
{
	static_cast<QVector<int>*>(ptr)->append(i);
}

void* ItemTemplate40f1b6___grabTouchPoints_ids_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<int>();
}

void* ItemTemplate40f1b6___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void ItemTemplate40f1b6___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* ItemTemplate40f1b6___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* ItemTemplate40f1b6___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void ItemTemplate40f1b6___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* ItemTemplate40f1b6___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* ItemTemplate40f1b6___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void ItemTemplate40f1b6___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* ItemTemplate40f1b6___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* ItemTemplate40f1b6___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void ItemTemplate40f1b6___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* ItemTemplate40f1b6___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* ItemTemplate40f1b6___qFindChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void ItemTemplate40f1b6___qFindChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* ItemTemplate40f1b6___qFindChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

char ItemTemplate40f1b6___sendBool_v1_atList(void* ptr, int i)
{
	return ({bool tmp = static_cast<QList<bool>*>(ptr)->at(i); if (i == static_cast<QList<bool>*>(ptr)->size()-1) { static_cast<QList<bool>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void ItemTemplate40f1b6___sendBool_v1_setList(void* ptr, char i)
{
	static_cast<QList<bool>*>(ptr)->append(i != 0);
}

void* ItemTemplate40f1b6___sendBool_v1_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<bool>();
}

int ItemTemplate40f1b6___sendInt_v1_atList(void* ptr, int i)
{
	return ({qint32 tmp = static_cast<QList<qint32>*>(ptr)->at(i); if (i == static_cast<QList<qint32>*>(ptr)->size()-1) { static_cast<QList<qint32>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void ItemTemplate40f1b6___sendInt_v1_setList(void* ptr, int i)
{
	static_cast<QList<qint32>*>(ptr)->append(i);
}

void* ItemTemplate40f1b6___sendInt_v1_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<qint32>();
}

double ItemTemplate40f1b6___sendDouble_v1_atList(void* ptr, int i)
{
	return ({qreal tmp = static_cast<QList<qreal>*>(ptr)->at(i); if (i == static_cast<QList<qreal>*>(ptr)->size()-1) { static_cast<QList<qreal>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void ItemTemplate40f1b6___sendDouble_v1_setList(void* ptr, double i)
{
	static_cast<QList<qreal>*>(ptr)->append(i);
}

void* ItemTemplate40f1b6___sendDouble_v1_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<qreal>();
}

struct Moc_PackedString ItemTemplate40f1b6___sendError_v1_atList(void* ptr, int i)
{
	return ({ QByteArray t94aa5e = ({QString tmp = static_cast<QList<QString>*>(ptr)->at(i); if (i == static_cast<QList<QString>*>(ptr)->size()-1) { static_cast<QList<QString>*>(ptr)->~QList(); free(ptr); }; tmp; }).toUtf8(); Moc_PackedString { const_cast<char*>(t94aa5e.prepend("WHITESPACE").constData()+10), t94aa5e.size()-10 }; });
}

void ItemTemplate40f1b6___sendError_v1_setList(void* ptr, struct Moc_PackedString i)
{
	static_cast<QList<QString>*>(ptr)->append(QString::fromUtf8(i.data, i.len));
}

void* ItemTemplate40f1b6___sendError_v1_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QString>();
}

void* ItemTemplate40f1b6___sendVariantListMap_v1_atList(void* ptr, int i)
{
	return new QVariant(({QVariant tmp = static_cast<QList<QVariant>*>(ptr)->at(i); if (i == static_cast<QList<QVariant>*>(ptr)->size()-1) { static_cast<QList<QVariant>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void ItemTemplate40f1b6___sendVariantListMap_v1_setList(void* ptr, void* i)
{
	static_cast<QList<QVariant>*>(ptr)->append(*static_cast<QVariant*>(i));
}

void* ItemTemplate40f1b6___sendVariantListMap_v1_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QVariant>();
}

void* ItemTemplate40f1b6___sendVariantListMap_v2_atList(void* ptr, struct Moc_PackedString v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<QString, QVariant>*>(ptr)->value(QString::fromUtf8(v.data, v.len)); if (i == static_cast<QMap<QString, QVariant>*>(ptr)->size()-1) { static_cast<QMap<QString, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void ItemTemplate40f1b6___sendVariantListMap_v2_setList(void* ptr, struct Moc_PackedString key, void* i)
{
	static_cast<QMap<QString, QVariant>*>(ptr)->insert(QString::fromUtf8(key.data, key.len), *static_cast<QVariant*>(i));
}

void* ItemTemplate40f1b6___sendVariantListMap_v2_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<QString, QVariant>();
}

struct Moc_PackedList ItemTemplate40f1b6___sendVariantListMap_v2_keyList(void* ptr)
{
	return ({ QList<QString>* tmpValue = new QList<QString>(static_cast<QMap<QString, QVariant>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* ItemTemplate40f1b6___sendItemTemplate_v1_atList(void* ptr, int i)
{
	return ({ItemTemplate40f1b6* tmp = static_cast<QList<ItemTemplate40f1b6*>*>(ptr)->at(i); if (i == static_cast<QList<ItemTemplate40f1b6*>*>(ptr)->size()-1) { static_cast<QList<ItemTemplate40f1b6*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void ItemTemplate40f1b6___sendItemTemplate_v1_setList(void* ptr, void* i)
{
	static_cast<QList<ItemTemplate40f1b6*>*>(ptr)->append(static_cast<ItemTemplate40f1b6*>(i));
}

void* ItemTemplate40f1b6___sendItemTemplate_v1_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<ItemTemplate40f1b6*>();
}

void* ItemTemplate40f1b6___sendItem_v1_atList(void* ptr, int i)
{
	return ({QQuickItem* tmp = static_cast<QList<QQuickItem*>*>(ptr)->at(i); if (i == static_cast<QList<QQuickItem*>*>(ptr)->size()-1) { static_cast<QList<QQuickItem*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void ItemTemplate40f1b6___sendItem_v1_setList(void* ptr, void* i)
{
	static_cast<QList<QQuickItem*>*>(ptr)->append(static_cast<QQuickItem*>(i));
}

void* ItemTemplate40f1b6___sendItem_v1_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QQuickItem*>();
}

void* ItemTemplate40f1b6___sendObject_v1_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void ItemTemplate40f1b6___sendObject_v1_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* ItemTemplate40f1b6___sendObject_v1_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

struct Moc_PackedString ItemTemplate40f1b6_____sendVariantListMap_v2_keyList_atList(void* ptr, int i)
{
	return ({ QByteArray t94aa5e = ({QString tmp = static_cast<QList<QString>*>(ptr)->at(i); if (i == static_cast<QList<QString>*>(ptr)->size()-1) { static_cast<QList<QString>*>(ptr)->~QList(); free(ptr); }; tmp; }).toUtf8(); Moc_PackedString { const_cast<char*>(t94aa5e.prepend("WHITESPACE").constData()+10), t94aa5e.size()-10 }; });
}

void ItemTemplate40f1b6_____sendVariantListMap_v2_keyList_setList(void* ptr, struct Moc_PackedString i)
{
	static_cast<QList<QString>*>(ptr)->append(QString::fromUtf8(i.data, i.len));
}

void* ItemTemplate40f1b6_____sendVariantListMap_v2_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QString>();
}

void* ItemTemplate40f1b6_NewItemTemplate(void* parent)
{
		return new ItemTemplate40f1b6(static_cast<QQuickItem*>(parent));
}

void ItemTemplate40f1b6_DestroyItemTemplate(void* ptr)
{
	static_cast<ItemTemplate40f1b6*>(ptr)->~ItemTemplate40f1b6();
}

void ItemTemplate40f1b6_DestroyItemTemplateDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

char ItemTemplate40f1b6_ChildMouseEventFilterDefault(void* ptr, void* item, void* event)
{
	return static_cast<ItemTemplate40f1b6*>(ptr)->QQuickItem::childMouseEventFilter(static_cast<QQuickItem*>(item), static_cast<QEvent*>(event));
}

void ItemTemplate40f1b6_ClassBeginDefault(void* ptr)
{
	static_cast<ItemTemplate40f1b6*>(ptr)->QQuickItem::classBegin();
}

void ItemTemplate40f1b6_ComponentCompleteDefault(void* ptr)
{
	static_cast<ItemTemplate40f1b6*>(ptr)->QQuickItem::componentComplete();
}

char ItemTemplate40f1b6_ContainsDefault(void* ptr, void* point)
{
	return static_cast<ItemTemplate40f1b6*>(ptr)->QQuickItem::contains(*static_cast<QPointF*>(point));
}

void ItemTemplate40f1b6_DragEnterEventDefault(void* ptr, void* event)
{
	static_cast<ItemTemplate40f1b6*>(ptr)->QQuickItem::dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void ItemTemplate40f1b6_DragLeaveEventDefault(void* ptr, void* event)
{
	static_cast<ItemTemplate40f1b6*>(ptr)->QQuickItem::dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
}

void ItemTemplate40f1b6_DragMoveEventDefault(void* ptr, void* event)
{
	static_cast<ItemTemplate40f1b6*>(ptr)->QQuickItem::dragMoveEvent(static_cast<QDragMoveEvent*>(event));
}

void ItemTemplate40f1b6_DropEventDefault(void* ptr, void* event)
{
	static_cast<ItemTemplate40f1b6*>(ptr)->QQuickItem::dropEvent(static_cast<QDropEvent*>(event));
}

char ItemTemplate40f1b6_EventDefault(void* ptr, void* ev)
{
	return static_cast<ItemTemplate40f1b6*>(ptr)->QQuickItem::event(static_cast<QEvent*>(ev));
}

void ItemTemplate40f1b6_FocusInEventDefault(void* ptr, void* vqf)
{
	static_cast<ItemTemplate40f1b6*>(ptr)->QQuickItem::focusInEvent(static_cast<QFocusEvent*>(vqf));
}

void ItemTemplate40f1b6_FocusOutEventDefault(void* ptr, void* vqf)
{
	static_cast<ItemTemplate40f1b6*>(ptr)->QQuickItem::focusOutEvent(static_cast<QFocusEvent*>(vqf));
}

void ItemTemplate40f1b6_GeometryChangedDefault(void* ptr, void* newGeometry, void* oldGeometry)
{
	static_cast<ItemTemplate40f1b6*>(ptr)->QQuickItem::geometryChanged(*static_cast<QRectF*>(newGeometry), *static_cast<QRectF*>(oldGeometry));
}

void ItemTemplate40f1b6_HoverEnterEventDefault(void* ptr, void* event)
{
	static_cast<ItemTemplate40f1b6*>(ptr)->QQuickItem::hoverEnterEvent(static_cast<QHoverEvent*>(event));
}

void ItemTemplate40f1b6_HoverLeaveEventDefault(void* ptr, void* event)
{
	static_cast<ItemTemplate40f1b6*>(ptr)->QQuickItem::hoverLeaveEvent(static_cast<QHoverEvent*>(event));
}

void ItemTemplate40f1b6_HoverMoveEventDefault(void* ptr, void* event)
{
	static_cast<ItemTemplate40f1b6*>(ptr)->QQuickItem::hoverMoveEvent(static_cast<QHoverEvent*>(event));
}

void ItemTemplate40f1b6_InputMethodEventDefault(void* ptr, void* event)
{
	static_cast<ItemTemplate40f1b6*>(ptr)->QQuickItem::inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void* ItemTemplate40f1b6_InputMethodQueryDefault(void* ptr, long long query)
{
	return new QVariant(static_cast<ItemTemplate40f1b6*>(ptr)->QQuickItem::inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

char ItemTemplate40f1b6_IsTextureProviderDefault(void* ptr)
{
	return static_cast<ItemTemplate40f1b6*>(ptr)->QQuickItem::isTextureProvider();
}

void ItemTemplate40f1b6_KeyPressEventDefault(void* ptr, void* event)
{
	static_cast<ItemTemplate40f1b6*>(ptr)->QQuickItem::keyPressEvent(static_cast<QKeyEvent*>(event));
}

void ItemTemplate40f1b6_KeyReleaseEventDefault(void* ptr, void* event)
{
	static_cast<ItemTemplate40f1b6*>(ptr)->QQuickItem::keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void ItemTemplate40f1b6_MouseDoubleClickEventDefault(void* ptr, void* event)
{
	static_cast<ItemTemplate40f1b6*>(ptr)->QQuickItem::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void ItemTemplate40f1b6_MouseMoveEventDefault(void* ptr, void* event)
{
	static_cast<ItemTemplate40f1b6*>(ptr)->QQuickItem::mouseMoveEvent(static_cast<QMouseEvent*>(event));
}

void ItemTemplate40f1b6_MousePressEventDefault(void* ptr, void* event)
{
	static_cast<ItemTemplate40f1b6*>(ptr)->QQuickItem::mousePressEvent(static_cast<QMouseEvent*>(event));
}

void ItemTemplate40f1b6_MouseReleaseEventDefault(void* ptr, void* event)
{
	static_cast<ItemTemplate40f1b6*>(ptr)->QQuickItem::mouseReleaseEvent(static_cast<QMouseEvent*>(event));
}

void ItemTemplate40f1b6_MouseUngrabEventDefault(void* ptr)
{
	static_cast<ItemTemplate40f1b6*>(ptr)->QQuickItem::mouseUngrabEvent();
}

void ItemTemplate40f1b6_ReleaseResourcesDefault(void* ptr)
{
	static_cast<ItemTemplate40f1b6*>(ptr)->QQuickItem::releaseResources();
}

void* ItemTemplate40f1b6_TextureProviderDefault(void* ptr)
{
	return static_cast<ItemTemplate40f1b6*>(ptr)->QQuickItem::textureProvider();
}

void ItemTemplate40f1b6_TouchEventDefault(void* ptr, void* event)
{
	static_cast<ItemTemplate40f1b6*>(ptr)->QQuickItem::touchEvent(static_cast<QTouchEvent*>(event));
}

void ItemTemplate40f1b6_TouchUngrabEventDefault(void* ptr)
{
	static_cast<ItemTemplate40f1b6*>(ptr)->QQuickItem::touchUngrabEvent();
}

void ItemTemplate40f1b6_UpdateDefault(void* ptr)
{
	static_cast<ItemTemplate40f1b6*>(ptr)->QQuickItem::update();
}

void ItemTemplate40f1b6_UpdatePolishDefault(void* ptr)
{
	static_cast<ItemTemplate40f1b6*>(ptr)->QQuickItem::updatePolish();
}

void ItemTemplate40f1b6_WheelEventDefault(void* ptr, void* event)
{
	static_cast<ItemTemplate40f1b6*>(ptr)->QQuickItem::wheelEvent(static_cast<QWheelEvent*>(event));
}

void ItemTemplate40f1b6_ChildEventDefault(void* ptr, void* event)
{
	static_cast<ItemTemplate40f1b6*>(ptr)->QQuickItem::childEvent(static_cast<QChildEvent*>(event));
}

void ItemTemplate40f1b6_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<ItemTemplate40f1b6*>(ptr)->QQuickItem::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void ItemTemplate40f1b6_CustomEventDefault(void* ptr, void* event)
{
	static_cast<ItemTemplate40f1b6*>(ptr)->QQuickItem::customEvent(static_cast<QEvent*>(event));
}

void ItemTemplate40f1b6_DeleteLaterDefault(void* ptr)
{
	static_cast<ItemTemplate40f1b6*>(ptr)->QQuickItem::deleteLater();
}

void ItemTemplate40f1b6_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<ItemTemplate40f1b6*>(ptr)->QQuickItem::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char ItemTemplate40f1b6_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<ItemTemplate40f1b6*>(ptr)->QQuickItem::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void ItemTemplate40f1b6_TimerEventDefault(void* ptr, void* event)
{
	static_cast<ItemTemplate40f1b6*>(ptr)->QQuickItem::timerEvent(static_cast<QTimerEvent*>(event));
}

#include "moc_moc.h"
