

#pragma once

#ifndef GO_MOC_40f1b6_H
#define GO_MOC_40f1b6_H

#include <stdint.h>

#ifdef __cplusplus
class QmlBridge40f1b6;
void QmlBridge40f1b6_QmlBridge40f1b6_QRegisterMetaTypes();
extern "C" {
#endif

struct Moc_PackedString { char* data; long long len; };
struct Moc_PackedList { void* data; long long len; };
void QmlBridge40f1b6_Search(void* ptr, struct Moc_PackedString term);
int QmlBridge40f1b6_QmlBridge40f1b6_QRegisterMetaType();
int QmlBridge40f1b6_QmlBridge40f1b6_QRegisterMetaType2(char* typeName);
int QmlBridge40f1b6_QmlBridge40f1b6_QmlRegisterType();
int QmlBridge40f1b6_QmlBridge40f1b6_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName);
void* QmlBridge40f1b6___children_atList(void* ptr, int i);
void QmlBridge40f1b6___children_setList(void* ptr, void* i);
void* QmlBridge40f1b6___children_newList(void* ptr);
void* QmlBridge40f1b6___dynamicPropertyNames_atList(void* ptr, int i);
void QmlBridge40f1b6___dynamicPropertyNames_setList(void* ptr, void* i);
void* QmlBridge40f1b6___dynamicPropertyNames_newList(void* ptr);
void* QmlBridge40f1b6___findChildren_atList(void* ptr, int i);
void QmlBridge40f1b6___findChildren_setList(void* ptr, void* i);
void* QmlBridge40f1b6___findChildren_newList(void* ptr);
void* QmlBridge40f1b6___findChildren_atList3(void* ptr, int i);
void QmlBridge40f1b6___findChildren_setList3(void* ptr, void* i);
void* QmlBridge40f1b6___findChildren_newList3(void* ptr);
void* QmlBridge40f1b6___qFindChildren_atList2(void* ptr, int i);
void QmlBridge40f1b6___qFindChildren_setList2(void* ptr, void* i);
void* QmlBridge40f1b6___qFindChildren_newList2(void* ptr);
void* QmlBridge40f1b6_NewQmlBridge(void* parent);
void QmlBridge40f1b6_DestroyQmlBridge(void* ptr);
void QmlBridge40f1b6_DestroyQmlBridgeDefault(void* ptr);
void QmlBridge40f1b6_ChildEventDefault(void* ptr, void* event);
void QmlBridge40f1b6_ConnectNotifyDefault(void* ptr, void* sign);
void QmlBridge40f1b6_CustomEventDefault(void* ptr, void* event);
void QmlBridge40f1b6_DeleteLaterDefault(void* ptr);
void QmlBridge40f1b6_DisconnectNotifyDefault(void* ptr, void* sign);
char QmlBridge40f1b6_EventDefault(void* ptr, void* e);
char QmlBridge40f1b6_EventFilterDefault(void* ptr, void* watched, void* event);
void QmlBridge40f1b6_TimerEventDefault(void* ptr, void* event);

#ifdef __cplusplus
}
#endif

#endif