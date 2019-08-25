import QtQuick 2.12
import QtQuick.Layouts 1.1
import QtGraphicalEffects 1.12
import QtQuick.Controls 2.3

Rectangle {
    id:root
    TextArea {
        id: logArea
        placeholderText: qsTr("Welcome to diztl!")
        anchors.fill: parent
    }
}