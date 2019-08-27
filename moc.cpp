

#define protected public
#define private public

#include "moc.h"
#include "_cgo_export.h"

#include <QByteArray>
#include <QChildEvent>
#include <QEvent>
#include <QGraphicsObject>
#include <QGraphicsWidget>
#include <QLayout>
#include <QMetaMethod>
#include <QObject>
#include <QOffscreenSurface>
#include <QPaintDeviceWindow>
#include <QPdfWriter>
#include <QQuickItem>
#include <QString>
#include <QTimerEvent>
#include <QWidget>
#include <QWindow>

#ifdef QT_QML_LIB
	#include <QQmlEngine>
#endif


class QmlBridge40f1b6: public QObject
{
Q_OBJECT
public:
	QmlBridge40f1b6(QObject *parent = Q_NULLPTR) : QObject(parent) {qRegisterMetaType<quintptr>("quintptr");QmlBridge40f1b6_QmlBridge40f1b6_QRegisterMetaType();QmlBridge40f1b6_QmlBridge40f1b6_QRegisterMetaTypes();callbackQmlBridge40f1b6_Constructor(this);};
	 ~QmlBridge40f1b6() { callbackQmlBridge40f1b6_DestroyQmlBridge(this); };
	void childEvent(QChildEvent * event) { callbackQmlBridge40f1b6_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQmlBridge40f1b6_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQmlBridge40f1b6_CustomEvent(this, event); };
	void deleteLater() { callbackQmlBridge40f1b6_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQmlBridge40f1b6_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQmlBridge40f1b6_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQmlBridge40f1b6_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQmlBridge40f1b6_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQmlBridge40f1b6_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQmlBridge40f1b6_TimerEvent(this, event); };
signals:
public slots:
	void search(QString term) { QByteArray t2a4c3b = term.toUtf8(); Moc_PackedString termPacked = { const_cast<char*>(t2a4c3b.prepend("WHITESPACE").constData()+10), t2a4c3b.size()-10 };callbackQmlBridge40f1b6_Search(this, termPacked); };
private:
};

Q_DECLARE_METATYPE(QmlBridge40f1b6*)


void QmlBridge40f1b6_QmlBridge40f1b6_QRegisterMetaTypes() {
}

void QmlBridge40f1b6_Search(void* ptr, struct Moc_PackedString term)
{
	QMetaObject::invokeMethod(static_cast<QmlBridge40f1b6*>(ptr), "search", Q_ARG(QString, QString::fromUtf8(term.data, term.len)));
}

int QmlBridge40f1b6_QmlBridge40f1b6_QRegisterMetaType()
{
	return qRegisterMetaType<QmlBridge40f1b6*>();
}

int QmlBridge40f1b6_QmlBridge40f1b6_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<QmlBridge40f1b6*>(const_cast<const char*>(typeName));
}

int QmlBridge40f1b6_QmlBridge40f1b6_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<QmlBridge40f1b6>();
#else
	return 0;
#endif
}

int QmlBridge40f1b6_QmlBridge40f1b6_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<QmlBridge40f1b6>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

void* QmlBridge40f1b6___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QmlBridge40f1b6___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QmlBridge40f1b6___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* QmlBridge40f1b6___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QmlBridge40f1b6___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QmlBridge40f1b6___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QmlBridge40f1b6___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QmlBridge40f1b6___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QmlBridge40f1b6___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QmlBridge40f1b6___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QmlBridge40f1b6___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QmlBridge40f1b6___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QmlBridge40f1b6___qFindChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QmlBridge40f1b6___qFindChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QmlBridge40f1b6___qFindChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QmlBridge40f1b6_NewQmlBridge(void* parent)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new QmlBridge40f1b6(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new QmlBridge40f1b6(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new QmlBridge40f1b6(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new QmlBridge40f1b6(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new QmlBridge40f1b6(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new QmlBridge40f1b6(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new QmlBridge40f1b6(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new QmlBridge40f1b6(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new QmlBridge40f1b6(static_cast<QWindow*>(parent));
	} else {
		return new QmlBridge40f1b6(static_cast<QObject*>(parent));
	}
}

void QmlBridge40f1b6_DestroyQmlBridge(void* ptr)
{
	static_cast<QmlBridge40f1b6*>(ptr)->~QmlBridge40f1b6();
}

void QmlBridge40f1b6_DestroyQmlBridgeDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void QmlBridge40f1b6_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QmlBridge40f1b6*>(ptr)->QObject::childEvent(static_cast<QChildEvent*>(event));
}

void QmlBridge40f1b6_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QmlBridge40f1b6*>(ptr)->QObject::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QmlBridge40f1b6_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QmlBridge40f1b6*>(ptr)->QObject::customEvent(static_cast<QEvent*>(event));
}

void QmlBridge40f1b6_DeleteLaterDefault(void* ptr)
{
	static_cast<QmlBridge40f1b6*>(ptr)->QObject::deleteLater();
}

void QmlBridge40f1b6_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QmlBridge40f1b6*>(ptr)->QObject::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QmlBridge40f1b6_EventDefault(void* ptr, void* e)
{
	return static_cast<QmlBridge40f1b6*>(ptr)->QObject::event(static_cast<QEvent*>(e));
}

char QmlBridge40f1b6_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QmlBridge40f1b6*>(ptr)->QObject::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QmlBridge40f1b6_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QmlBridge40f1b6*>(ptr)->QObject::timerEvent(static_cast<QTimerEvent*>(event));
}

#include "moc_moc.h"
