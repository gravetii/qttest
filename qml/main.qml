import QtQuick 2.12
import QtQuick.Layouts 1.1
import QtGraphicalEffects 1.12

Rectangle {
	id: root
	Image {
		id: backImg
		source: "../images/motley.jpg"
		fillMode: Image.Stretch
		anchors.fill: parent
	}
	GameGrid {
		id: gameGrid
	}
	Blend {
		anchors.fill: gameGrid
		source: gameGrid
		foregroundSource: backImg
		mode: "multiply"
	}
}
