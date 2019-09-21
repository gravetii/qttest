import QtQuick 2.12
import QtQuick.Layouts 1.1
import QtGraphicalEffects 1.12
import QtQuick.Controls 1.4
import QtQuick.Controls.Styles 1.4

TabView {
    tabPosition: Qt.BottomEdge
    style: TabViewStyle {
        frameOverlap: 1
        tabsAlignment: Qt.AlignHCenter
        tab: Rectangle {
            color: styleData.selected ? "steelblue" : "lightsteelblue"
            border.color: "steelblue"
            implicitWidth: Math.max(text.width + 4, 80)
            implicitHeight: 20
            radius: 2
            Text {
                id: text
                anchors.centerIn: parent
                text: styleData.title
                color: styleData.selected ? "white" : "black"
            }
        }
        frame: Rectangle {
            color: "steelblue"
        }
    }
    Tab {
        id: redTab
        title: "Red"
        Rectangle {
            width: redTab.width
            height: redTab.height
            color: "red"
        }
    }
    Tab {
        id: greenTab
        title: "Green"
        Rectangle {
            width: greenTab.width
            height: greenTab.height
            color: "green"
        }
    }
    Tab {
        id: blueTab
        title: "Blue"
        Rectangle {
            width: blueTab.width
            height: blueTab.height
            color: "blue"
        }
    }
}