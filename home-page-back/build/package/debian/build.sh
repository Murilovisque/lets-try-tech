BASE_DIR="/tmp/home-page-back"
DEB_FILE="/tmp/home-page-back.deb"
DEB_TARGET_DIR="/home-page-back/build/package/debian/target"

currentDir=$(pwd)
rm -fv ${DEB_TARGET_DIR}/*
cd /home-page-back/cmd/home-page-back
rm -v home-page-back
go build
if [[ ! -x home-page-back ]]; then
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
mkdir -pv "${BASE_DIR}/etc/cron.daily"
mkdir -pv "${BASE_DIR}/var/log/home-page-back/archive"
mkdir -pv "${BASE_DIR}/opt/ltt/home-page-back"
mkdir -pv "${BASE_DIR}/etc/init.d"

cp -v /home-page-back/build/package/debian/home-page-back.sh "${BASE_DIR}/etc/init.d/home-page-back"
cp -v /home-page-back/cmd/home-page-back/home-page-back "${BASE_DIR}/opt/ltt/home-page-back"
cp -v /home-page-back/configs/cron/home-page-back "${BASE_DIR}/etc/cron.daily"

chmod +x ${BASE_DIR}/etc/init.d/home-page-back

dpkg-deb --build "${BASE_DIR}"
rm -rfv "${BASE_DIR}"
mkdir -p ${DEB_TARGET_DIR}
mv ${DEB_FILE} ${DEB_TARGET_DIR}
echo "Debian file builded: '${DEB_TARGET_DIR}/home-page-back.deb'"
cd ${currentDir}