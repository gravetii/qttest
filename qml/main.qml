import QtQuick 2.12
import QtQuick.Layouts 1.1
import QtGraphicalEffects 1.12
import QtQuick.Controls 2.12

Rectangle {
  id:root
  Row {
    spacing: 10
    FileSearch {
      id: fileSearch
    }
    LogArea {
      id: logArea
    }
  }
}