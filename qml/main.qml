import QtQuick 2.12
import QtQuick.Layouts 1.1
import QtGraphicalEffects 1.12
import QtQuick.Controls 2.12

Rectangle {
    id:root
    TextArea {
        id: logArea
        anchors.fill: parent
		readOnly: true
		textMargin: 5, 5, 5, 5
		text: "Welcome to diztl!"
		selectByMouse: true
		color: "#000000"
		background: null
    }
}