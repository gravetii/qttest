import QtQuick 2.12
import QtQuick.Layouts 1.1
import QtGraphicalEffects 1.12
import QtQuick.Controls 1.4

Rectangle {
    width: parent.width
    height: 800
    TableView {
        anchors.horizontalCenter: parent.horizontalCenter
        id: downloadsTable
        width: parent.width
        TableViewColumn {
            title: "File"
            width: 200
            horizontalAlignment: Text.AlignHCenter
        }
        TableViewColumn {
            title: "Status"
            width: 200
            horizontalAlignment: Text.AlignHCenter
        }
        TableViewColumn {
            title: "Progress"
            width: 200
            horizontalAlignment: Text.AlignHCenter
        }
        TableViewColumn {
            title: "Size"
            width: 200
            horizontalAlignment: Text.AlignHCenter
        }
        TableViewColumn {
            title: "Type"
            width: 200
            horizontalAlignment: Text.AlignHCenter
        }
        TableViewColumn {
            title: "Path"
            width: 200
            horizontalAlignment: Text.AlignHCenter
        }

    }
}