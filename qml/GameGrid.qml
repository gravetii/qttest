import QtQuick 2.12
import QtQuick.Layouts 1.1
import QtGraphicalEffects 1.12

GridLayout {
	rows: 4
	columns: 4
	rowSpacing: 0;columnSpacing: 0
	anchors.fill: parent
	Layout.fillHeight: true
	Layout.fillWidth: true
	Repeater {
		model: 16
		Rectangle {
			Layout.fillHeight: true
			Layout.fillWidth: true
			color: "transparent"
			Image {
				id: alphaImg
				source: "../images/a.png"
				Layout.fillHeight: true
				Layout.fillWidth: true
				anchors.fill: parent

				MouseArea {
					anchors.fill: parent
					hoverEnabled: true
					cursorShape: Qt.PointingHandCursor
				}
			}
		}
	}
}