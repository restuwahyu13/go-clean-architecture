package pkg

import (
	"time"

	"github.com/sirupsen/logrus"
)

func Logrus(Type string, Msg any, Args ...any) {
	format := false
	logrus.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: time.RFC3339,
	})

	if Args != nil {
		format = true
	}

	switch Type {

	case "info":
		if format {
			logrus.Infof(Msg.(string), Args...)
		} else {
			logrus.Info(Msg)
		}

		break

	case "error":
		if format {
			logrus.Errorf(Msg.(string), Args...)
		} else {
			logrus.Error(Msg)
		}

		break

	case "print":
		if format {
			logrus.Printf(Msg.(string), Args...)
		} else {
			logrus.Print(Msg)
		}

		break

	case "fatal":
		if format {
			logrus.Fatalf(Msg.(string), Args...)
		} else {
			logrus.Fatal(Msg)
		}

		break

	case "debug":
		if format {
			logrus.Debugf(Msg.(string), Args...)
		} else {
			logrus.Debug(Msg)
		}

		break

	case "panic":
		if format {
			logrus.Panicf(Msg.(string), Args...)
		} else {
			logrus.Panic(Msg)
		}

		break

	default:
		logrus.Println(Msg)

		break
	}
}
