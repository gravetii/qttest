import QtQuick 2.6
import QtQuick.Layouts 1.1

GridLayout {
	anchors.fill: parent
	rows: 4;
	columns: 4;
	rowSpacing: 0; columnSpacing: 0;
	Layout.fillHeight: true;
    Layout.fillWidth: true;
	// spacing: 0;
    Repeater {
        model: 16
        Rectangle {
			Layout.fillHeight: true;
			Layout.fillWidth: true;
            Image {
                source: "../images/a.png";
				Layout.fillHeight: true;
				Layout.fillWidth: true;
				anchors.fill: parent
                MouseArea {
                    anchors.fill: parent;
                    hoverEnabled: true;
					cursorShape: Qt.PointingHandCursor;
                }
            }
        }
    }
}