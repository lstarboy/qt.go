source ~/triline/shell/android-ndk-env.sh


export CGO_ENABLED=1
export GOOS=android
export GOARCH=arm
export GOARM=7
export CGO_LDFLAGS="$CGO_LDFLAGS -L/androidsys/lib"
export PKG_CONFIG_PATH=/androidsys/lib/pkgconfig

set -x
# rm -rf ~/oss/pkg/android_arm
go install -p 1 -v  -pkgdir ~/oss/pkg/android_arm ./qtrt ./qtqt
go install -p 1 -v  -pkgdir ~/oss/pkg/android_arm ./qtrt ./qtcore ./qtgui ./qtwidgets
go install -p 1 -v  -pkgdir ~/oss/pkg/android_arm ./qtrt ./qtnetwork ./qtqml ./qtquick
