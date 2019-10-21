import QtQuick 2.12
import QtQuick.Layouts 1.1
import QtGraphicalEffects 1.12
import QtQuick.Controls 1.4
import QtQuick.Controls.Styles 1.4

Rectangle {
    id: rootRect
    TableView {
        id: resultListTableView
        anchors.fill: parent
        rowDelegate: Rectangle {
            color: "#D3D3D3"
            height: 30
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
            width: resultListTableView.viewport.width * 0.3
            movable: false
            resizable: false
        }
        TableViewColumn {
            role: "type"
            title: "Type"
            width: resultListTableView.viewport.width * 0.2
            movable: false
            resizable: false
        }
        TableViewColumn {
            role: "size"
            title: "Size"
            width: resultListTableView.viewport.width * 0.2
            movable: false
            resizable: false
        }
        TableViewColumn {
            role: "path"
            title: "Path"
            width: resultListTableView.viewport.width * 0.3
            movable: false
            resizable: false
        }

        model: ResultListDataModel {}
        onDoubleClicked: {
            const element = model.get(row)
            console.log("Downloading file ", element.file, "of size", element.size)
        }
    }
}
