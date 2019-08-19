import QtQuick 2.3
import QtQuick.Layouts 1.3
import QtQuick.Controls 1.3

Rectangle {
	id: root
	Image {
		source: "../images/motley.jpg"
		fillMode: Image.Stretch
		anchors.fill: parent
	}
	GameGrid {
	}
}
