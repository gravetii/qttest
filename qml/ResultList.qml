import QtQuick 2.12
import QtQuick.Layouts 1.1
import QtGraphicalEffects 1.12
import QtQuick.Controls 1.4
import QtQuick.Controls.Styles 1.4

Rectangle {
    id: rootRect
    TableView {
        anchors.fill: parent
        rowDelegate: Rectangle {
            color: "#D3D3D3"
            height: 30
            MouseArea {
                anchors.fill: parent
            }
        }

        itemDelegate: Rectangle {
            width: 100
            height: 50
            border.color: "#000000"
            border.width: 1
            Text {
                id: textItem
                anchors.verticalCenter: parent.verticalCenter
                anchors.horizontalCenter: parent.horizontalCenter
                text: styleData.value
            }
        }

        headerDelegate: Rectangle {
            height: textItem.implicitHeight * 1.2
            width: textItem.implicitWidth
            color: "lightsteelblue"
            Text {
                id: textItem
                anchors.verticalCenter: parent.verticalCenter
                anchors.horizontalCenter: parent.horizontalCenter
                text: styleData.value
                elide: Text.ElideRight
                renderType: Text.NativeRendering
            }
        }
        
        TableViewColumn {
            role: "title"
            title: "Title"
            width: rootRect.width/2
            movable: false
            resizable: false
        }
        TableViewColumn {
            role: "author"
            title: "Author"
            width: rootRect.width/2
            movable: false
            resizable: false
        }

        model: libraryModel
    }

    ListModel {
    id: libraryModel
    ListElement {
        title: "A Masterpiece"
        author: "Gabriel"
    }
    ListElement {
        title: "Brilliance"
        author: "Jens"
    }
    ListElement {
        title: "Outstanding"
        author: "Frederik"
    }
}
}
