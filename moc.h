

#pragma once

#ifndef GO_MOC_40f1b6_H
#define GO_MOC_40f1b6_H

#include <stdint.h>

#ifdef __cplusplus
class CtxObject40f1b6;
void CtxObject40f1b6_CtxObject40f1b6_QRegisterMetaTypes();
extern "C" {
#endif

struct Moc_PackedString { char* data; long long len; };
struct Moc_PackedList { void* data; long long len; };
void CtxObject40f1b6_ConnectClicked(void* ptr);
void CtxObject40f1b6_DisconnectClicked(void* ptr);
void CtxObject40f1b6_Clicked(void* ptr);
struct Moc_PackedString CtxObject40f1b6_SomeString(void* ptr);
struct Moc_PackedString CtxObject40f1b6_SomeStringDefault(void* ptr);
void CtxObject40f1b6_SetSomeString(void* ptr, struct Moc_PackedString someString);
void CtxObject40f1b6_SetSomeStringDefault(void* ptr, struct Moc_PackedString someString);
void CtxObject40f1b6_ConnectSomeStringChanged(void* ptr);
void CtxObject40f1b6_DisconnectSomeStringChanged(void* ptr);
void CtxObject40f1b6_SomeStringChanged(void* ptr, struct Moc_PackedString someString);
int CtxObject40f1b6_CtxObject40f1b6_QRegisterMetaType();
int CtxObject40f1b6_CtxObject40f1b6_QRegisterMetaType2(char* typeName);
int CtxObject40f1b6_CtxObject40f1b6_QmlRegisterType();
int CtxObject40f1b6_CtxObject40f1b6_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName);
void* CtxObject40f1b6___children_atList(void* ptr, int i);
void CtxObject40f1b6___children_setList(void* ptr, void* i);
void* CtxObject40f1b6___children_newList(void* ptr);
void* CtxObject40f1b6___dynamicPropertyNames_atList(void* ptr, int i);
void CtxObject40f1b6___dynamicPropertyNames_setList(void* ptr, void* i);
void* CtxObject40f1b6___dynamicPropertyNames_newList(void* ptr);
void* CtxObject40f1b6___findChildren_atList(void* ptr, int i);
void CtxObject40f1b6___findChildren_setList(void* ptr, void* i);
void* CtxObject40f1b6___findChildren_newList(void* ptr);
void* CtxObject40f1b6___findChildren_atList3(void* ptr, int i);
void CtxObject40f1b6___findChildren_setList3(void* ptr, void* i);
void* CtxObject40f1b6___findChildren_newList3(void* ptr);
void* CtxObject40f1b6___qFindChildren_atList2(void* ptr, int i);
void CtxObject40f1b6___qFindChildren_setList2(void* ptr, void* i);
void* CtxObject40f1b6___qFindChildren_newList2(void* ptr);
void* CtxObject40f1b6_NewCtxObject(void* parent);
void CtxObject40f1b6_DestroyCtxObject(void* ptr);
void CtxObject40f1b6_DestroyCtxObjectDefault(void* ptr);
void CtxObject40f1b6_ChildEventDefault(void* ptr, void* event);
void CtxObject40f1b6_ConnectNotifyDefault(void* ptr, void* sign);
void CtxObject40f1b6_CustomEventDefault(void* ptr, void* event);
void CtxObject40f1b6_DeleteLaterDefault(void* ptr);
void CtxObject40f1b6_DisconnectNotifyDefault(void* ptr, void* sign);
char CtxObject40f1b6_EventDefault(void* ptr, void* e);
char CtxObject40f1b6_EventFilterDefault(void* ptr, void* watched, void* event);
void CtxObject40f1b6_TimerEventDefault(void* ptr, void* event);

#ifdef __cplusplus
}
#endif

#endif