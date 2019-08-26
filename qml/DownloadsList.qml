import QtQuick 2.12
import QtQuick.Layouts 1.1
import QtGraphicalEffects 1.12
import QtQuick.Controls 1.4

Rectangle {
    height: 500
    TableView {
        id: downloadsTable
        TableViewColumn {
            title: "File"
            width: 200
        }
        TableViewColumn {
            title: "Status"
            width: 200
        }
        TableViewColumn {
            title: "Progress"
            width: 200
        }
        TableViewColumn {
            title: "Size"
            width: 200
        }
        TableViewColumn {
            title: "Type"
            width: 200
        }
        TableViewColumn {
            title: "Path"
            width: 200
        }

    }
}