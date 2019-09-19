#include <QGuiApplication>
#include <QQuickView>
#include <QQmlEngine>
#include <QQmlComponent>
#include <QQuickItem>
#include <QObject>
#include <QQuickWindow>

#include "_cgo_export.h"

extern "C" {

  void start() {
    int argc = 0;
    QGuiApplication app(argc, nullptr);

    QQmlEngine engine;
    QQmlComponent component(&engine, QUrl(QString(ServerAddr()) + "?hello=foo"));
    while (component.isLoading()) {
      QCoreApplication::processEvents();
    }
    if (!component.isReady()) {
      qWarning() << qPrintable(component.errorString());
      return;
    }
    QQuickItem *item = qobject_cast<QQuickItem*>(component.create());
    if (!item && component.isError()) {
      qWarning() << component.errors();
      return;
    } 

    QQuickView view(&engine, nullptr);
    view.setTitle("foo");
    view.setContent(QUrl(), &component, item);
    view.show();

    app.exec();
  }

}

