import QtQuick 2.12
import QtQuick.Layouts 1.1
import QtGraphicalEffects 1.12
import QtQuick.Controls 2.12

Rectangle {
  width: 1200
  height: 750
  id:root
  Row {
    spacing: 300
    FileSearch {
      id: fileSearch
    }
    LogArea {
      id: logArea
    }
  }
}