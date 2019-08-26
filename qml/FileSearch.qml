import QtQuick 2.12
import QtQuick.Layouts 1.1
import QtGraphicalEffects 1.12
import QtQuick.Controls 2.3

Rectangle {
    width: 200
    height: 200

    Row {
        x: 10
        y: 10
        spacing: 25
        TextField {
            width: 300
            height: 50
            id: searchBox
            placeholderText: "Search..."
        }
        Button {
            width: 100
            height: 45
            text: "Click!"
        }
    }
}

