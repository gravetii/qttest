import QtQuick 2.12
import QtQuick.Layouts 1.1
import QtGraphicalEffects 1.12
import QtQuick.Controls 1.4
import QtQuick.Controls.Styles 1.4

Rectangle {
    TableView {
        anchors.fill: parent
        rowDelegate: Rectangle {
            color: styleData.row % 2 == 0 ? "#D3D3D3" : "#FFFFFF"
            height: 20
        }

        itemDelegate: Rectangle {
            implicitWidth: 100
            implicitHeight: 50
            border.color: "#000000"
            border.width: 1
            Text {
                anchors.verticalCenter: parent.verticalCenter
                anchors.horizontalCenter: parent.horizontalCenter
                text: styleData.value
            }
        }

        headerDelegate: Rectangle {
            implicitWidth: 100
            implicitHeight: 40
            border.color: "#000000"
            border.width: 1
            Text {
                anchors.verticalCenter: parent.verticalCenter
                anchors.horizontalCenter: parent.horizontalCenter
                text: styleData.value
                font.bold: true
            }
        }
        
        TableViewColumn {
            role: "title"
            title: "Title"
            // width: 100
        }
        TableViewColumn {
            role: "author"
            title: "Author"
            // width: 200
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
