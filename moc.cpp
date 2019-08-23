

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


class CtxObject40f1b6: public QObject
{
Q_OBJECT
Q_PROPERTY(QString someString READ someString WRITE setSomeString NOTIFY someStringChanged)
public:
	CtxObject40f1b6(QObject *parent = Q_NULLPTR) : QObject(parent) {qRegisterMetaType<quintptr>("quintptr");CtxObject40f1b6_CtxObject40f1b6_QRegisterMetaType();CtxObject40f1b6_CtxObject40f1b6_QRegisterMetaTypes();callbackCtxObject40f1b6_Constructor(this);};
	void Signal_Clicked() { callbackCtxObject40f1b6_Clicked(this); };
	QString someString() { return ({ Moc_PackedString tempVal = callbackCtxObject40f1b6_SomeString(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setSomeString(QString someString) { QByteArray te0f92a = someString.toUtf8(); Moc_PackedString someStringPacked = { const_cast<char*>(te0f92a.prepend("WHITESPACE").constData()+10), te0f92a.size()-10 };callbackCtxObject40f1b6_SetSomeString(this, someStringPacked); };
	void Signal_SomeStringChanged(QString someString) { QByteArray te0f92a = someString.toUtf8(); Moc_PackedString someStringPacked = { const_cast<char*>(te0f92a.prepend("WHITESPACE").constData()+10), te0f92a.size()-10 };callbackCtxObject40f1b6_SomeStringChanged(this, someStringPacked); };
	 ~CtxObject40f1b6() { callbackCtxObject40f1b6_DestroyCtxObject(this); };
	void childEvent(QChildEvent * event) { callbackCtxObject40f1b6_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackCtxObject40f1b6_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackCtxObject40f1b6_CustomEvent(this, event); };
	void deleteLater() { callbackCtxObject40f1b6_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackCtxObject40f1b6_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackCtxObject40f1b6_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackCtxObject40f1b6_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackCtxObject40f1b6_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackCtxObject40f1b6_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackCtxObject40f1b6_TimerEvent(this, event); };
	QString someStringDefault() { return _someString; };
	void setSomeStringDefault(QString p) { if (p != _someString) { _someString = p; someStringChanged(_someString); } };
signals:
	void clicked();
	void someStringChanged(QString someString);
public slots:
private:
	QString _someString;
};

Q_DECLARE_METATYPE(CtxObject40f1b6*)


void CtxObject40f1b6_CtxObject40f1b6_QRegisterMetaTypes() {
	qRegisterMetaType<QString>();
}

void CtxObject40f1b6_ConnectClicked(void* ptr)
{
	QObject::connect(static_cast<CtxObject40f1b6*>(ptr), static_cast<void (CtxObject40f1b6::*)()>(&CtxObject40f1b6::clicked), static_cast<CtxObject40f1b6*>(ptr), static_cast<void (CtxObject40f1b6::*)()>(&CtxObject40f1b6::Signal_Clicked));
}

void CtxObject40f1b6_DisconnectClicked(void* ptr)
{
	QObject::disconnect(static_cast<CtxObject40f1b6*>(ptr), static_cast<void (CtxObject40f1b6::*)()>(&CtxObject40f1b6::clicked), static_cast<CtxObject40f1b6*>(ptr), static_cast<void (CtxObject40f1b6::*)()>(&CtxObject40f1b6::Signal_Clicked));
}

void CtxObject40f1b6_Clicked(void* ptr)
{
	static_cast<CtxObject40f1b6*>(ptr)->clicked();
}

struct Moc_PackedString CtxObject40f1b6_SomeString(void* ptr)
{
	return ({ QByteArray tbc551a = static_cast<CtxObject40f1b6*>(ptr)->someString().toUtf8(); Moc_PackedString { const_cast<char*>(tbc551a.prepend("WHITESPACE").constData()+10), tbc551a.size()-10 }; });
}

struct Moc_PackedString CtxObject40f1b6_SomeStringDefault(void* ptr)
{
	return ({ QByteArray tf2f6e2 = static_cast<CtxObject40f1b6*>(ptr)->someStringDefault().toUtf8(); Moc_PackedString { const_cast<char*>(tf2f6e2.prepend("WHITESPACE").constData()+10), tf2f6e2.size()-10 }; });
}

void CtxObject40f1b6_SetSomeString(void* ptr, struct Moc_PackedString someString)
{
	static_cast<CtxObject40f1b6*>(ptr)->setSomeString(QString::fromUtf8(someString.data, someString.len));
}

void CtxObject40f1b6_SetSomeStringDefault(void* ptr, struct Moc_PackedString someString)
{
	static_cast<CtxObject40f1b6*>(ptr)->setSomeStringDefault(QString::fromUtf8(someString.data, someString.len));
}

void CtxObject40f1b6_ConnectSomeStringChanged(void* ptr)
{
	QObject::connect(static_cast<CtxObject40f1b6*>(ptr), static_cast<void (CtxObject40f1b6::*)(QString)>(&CtxObject40f1b6::someStringChanged), static_cast<CtxObject40f1b6*>(ptr), static_cast<void (CtxObject40f1b6::*)(QString)>(&CtxObject40f1b6::Signal_SomeStringChanged));
}

void CtxObject40f1b6_DisconnectSomeStringChanged(void* ptr)
{
	QObject::disconnect(static_cast<CtxObject40f1b6*>(ptr), static_cast<void (CtxObject40f1b6::*)(QString)>(&CtxObject40f1b6::someStringChanged), static_cast<CtxObject40f1b6*>(ptr), static_cast<void (CtxObject40f1b6::*)(QString)>(&CtxObject40f1b6::Signal_SomeStringChanged));
}

void CtxObject40f1b6_SomeStringChanged(void* ptr, struct Moc_PackedString someString)
{
	static_cast<CtxObject40f1b6*>(ptr)->someStringChanged(QString::fromUtf8(someString.data, someString.len));
}

int CtxObject40f1b6_CtxObject40f1b6_QRegisterMetaType()
{
	return qRegisterMetaType<CtxObject40f1b6*>();
}

int CtxObject40f1b6_CtxObject40f1b6_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<CtxObject40f1b6*>(const_cast<const char*>(typeName));
}

int CtxObject40f1b6_CtxObject40f1b6_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<CtxObject40f1b6>();
#else
	return 0;
#endif
}

int CtxObject40f1b6_CtxObject40f1b6_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<CtxObject40f1b6>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

void* CtxObject40f1b6___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void CtxObject40f1b6___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* CtxObject40f1b6___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* CtxObject40f1b6___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void CtxObject40f1b6___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* CtxObject40f1b6___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* CtxObject40f1b6___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void CtxObject40f1b6___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* CtxObject40f1b6___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* CtxObject40f1b6___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void CtxObject40f1b6___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* CtxObject40f1b6___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* CtxObject40f1b6___qFindChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void CtxObject40f1b6___qFindChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* CtxObject40f1b6___qFindChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* CtxObject40f1b6_NewCtxObject(void* parent)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new CtxObject40f1b6(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new CtxObject40f1b6(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new CtxObject40f1b6(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new CtxObject40f1b6(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new CtxObject40f1b6(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new CtxObject40f1b6(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new CtxObject40f1b6(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new CtxObject40f1b6(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new CtxObject40f1b6(static_cast<QWindow*>(parent));
	} else {
		return new CtxObject40f1b6(static_cast<QObject*>(parent));
	}
}

void CtxObject40f1b6_DestroyCtxObject(void* ptr)
{
	static_cast<CtxObject40f1b6*>(ptr)->~CtxObject40f1b6();
}

void CtxObject40f1b6_DestroyCtxObjectDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void CtxObject40f1b6_ChildEventDefault(void* ptr, void* event)
{
	static_cast<CtxObject40f1b6*>(ptr)->QObject::childEvent(static_cast<QChildEvent*>(event));
}

void CtxObject40f1b6_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<CtxObject40f1b6*>(ptr)->QObject::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void CtxObject40f1b6_CustomEventDefault(void* ptr, void* event)
{
	static_cast<CtxObject40f1b6*>(ptr)->QObject::customEvent(static_cast<QEvent*>(event));
}

void CtxObject40f1b6_DeleteLaterDefault(void* ptr)
{
	static_cast<CtxObject40f1b6*>(ptr)->QObject::deleteLater();
}

void CtxObject40f1b6_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<CtxObject40f1b6*>(ptr)->QObject::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char CtxObject40f1b6_EventDefault(void* ptr, void* e)
{
	return static_cast<CtxObject40f1b6*>(ptr)->QObject::event(static_cast<QEvent*>(e));
}

char CtxObject40f1b6_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<CtxObject40f1b6*>(ptr)->QObject::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void CtxObject40f1b6_TimerEventDefault(void* ptr, void* event)
{
	static_cast<CtxObject40f1b6*>(ptr)->QObject::timerEvent(static_cast<QTimerEvent*>(event));
}

#include "moc_moc.h"
