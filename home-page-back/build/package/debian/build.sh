BASE_DIR="/tmp/home-page-back"
DEB_FILE="/tmp/home-page-back.deb"
DEB_TARGET_DIR="/home-page-back/build/package/debian/target"


currentDit=$(pwd)
rm -fv ${DEB_TARGET_DIR}/*
cd /home-page-back/cmd/home-page
rm -v home-page
go build
if [[ ! -x home-page ]]; then
    echo "go buil failed"
    exit 1
fi

if [[ -d "${BASE_DIR}" ]]; then
    rm -rfv "${BASE_DIR}"
fi

mkdir -pv "${BASE_DIR}/DEBIAN"
cp -v /home-page-back/build/package/debian/control "${BASE_DIR}/DEBIAN"

mkdir -pv "${BASE_DIR}/lib/systemd/system"
cp -v /home-page-back/build/package/debian/home-page-back.service "${BASE_DIR}/lib/systemd/system"

mkdir -pv "${BASE_DIR}/etc/home-page-back"
mkdir -pv "${BASE_DIR}/etc/logrotate.d"
mkdir -pv "${BASE_DIR}/var/log/home-page-back/archive"
mkdir -pv "${BASE_DIR}/opt/ltt/home-page-back/dbs"
cp -v home-page "${BASE_DIR}/opt/ltt/home-page-back"
cp -v /home-page-back/configs/logrotate/home-page-back "${BASE_DIR}/etc/logrotate.d"


dpkg-deb --build "${BASE_DIR}"
rm -rfv "${BASE_DIR}"
mkdir -p ${DEB_TARGET_DIR}
mv ${DEB_FILE} ${DEB_TARGET_DIR}
echo "Debian file builded: '${DEB_TARGET_DIR}/home-page-back.deb'"
cd ${currentDit}