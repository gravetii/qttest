import QtQuick 2.12
import QtQuick.Layouts 1.1
import QtGraphicalEffects 1.12
import QtQuick.Controls 1.4
import QtQuick.Controls.Styles 1.4

Rectangle {
    id: rootRect
    width: parent.width
    height: 800
    TableView {
        id: downloadsListTableView
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
                text: styleData.value
                anchors.fill: parent
                verticalAlignment: Text.AlignVCenter
                horizontalAlignment: Text.AlignHCenter
                elide: Text.ElideRight
            }
        }

        headerDelegate: Rectangle {
            height: textItem.implicitHeight * 1.2
            width: textItem.implicitWidth
            color: "lightsteelblue"
            Text {
                id: textItem
                anchors.centerIn: parent
                text: styleData.value
                elide: Text.ElideRight
            }
        }

        TableViewColumn {
            role: "file"
            title: "File"
            width: downloadsListTableView.viewport.width/6
            movable: false
            resizable: false
        }
        TableViewColumn {
            role: "status"
            title: "Status"
            width: downloadsListTableView.viewport.width/6
            movable: false
            resizable: false
        }
        TableViewColumn {
            role: "progress"
            title: "Progress"
            width: downloadsListTableView.viewport.width/6
            movable: false
            resizable: false
        }
        TableViewColumn {
            role: "size"
            title: "Size"
            width: downloadsListTableView.viewport.width/6
            movable: false
            resizable: false
        }
        TableViewColumn {
            role: "type"
            title: "Type"
            width: downloadsListTableView.viewport.width/6
            movable: false
            resizable: false
        }
        TableViewColumn {
            role: "path"
            title: "Path"
            width: downloadsListTableView.viewport.width/6
            movable: false
            resizable: false
        }
    }
}